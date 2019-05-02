package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/metalmatze/cd/api"
	"github.com/metalmatze/cd/cd"
	"github.com/oklog/run"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const apiURL = "http://localhost:6660"
const namespace = "default"

func main() {
	kubeconfig := flag.String("kubeconfig", "", "Path to kubeconfig")
	agentName := flag.String("name", "", "Name for this agent")
	flag.Parse()

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		level.Error(logger).Log(
			"msg", "failed to read Kubernetes config",
			"err", err,
		)
		os.Exit(2)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		level.Error(logger).Log(
			"msg", "failed to create Kubernetes client",
			"err", err,
		)
		os.Exit(3)
	}

	ctx, cancel := context.WithCancel(context.Background())

	var gr run.Group
	{
		u := &updater{client: client, logger: logger, agentName: *agentName}

		gr.Add(func() error {
			return u.pollLoop(ctx)
		}, func(err error) {
			cancel()
		})
	}
	{
		sig := make(chan os.Signal)

		gr.Add(func() error {
			signal.Notify(sig, os.Interrupt)
			<-sig
			cancel()
			return nil
		}, func(err error) {
			close(sig)
		})
	}

	if err := gr.Run(); err != nil {
		level.Error(logger).Log(
			"msg", "error running",
			"err", err,
		)
		os.Exit(1)
	}
}

type updater struct {
	client    *kubernetes.Clientset
	logger    log.Logger
	agentName string

	currentPipeline cd.Pipeline
}

func (u *updater) pollLoop(ctx context.Context) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	level.Info(u.logger).Log("msg", "starting poll loop")

	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return nil
		case <-ticker.C:
			if err := u.poll(); err != nil {
				level.Warn(u.logger).Log(
					"msg", "failed to poll",
					"err", err,
				)
			}
		}
	}
}

func (u *updater) poll() error {
	p, err := u.pipeline()
	if err != nil {
		return err
	}

	if u.currentPipeline.ID == "" {
		u.currentPipeline = p

		level.Info(u.logger).Log("msg", "unknown pipeline", "pipeline", p.ID)
		if err := u.runSteps(p); err != nil {
			return err
		}
	}
	if u.currentPipeline.ID != p.ID {
		u.currentPipeline = p

		level.Info(u.logger).Log("msg", "updated pipeline", "pipeline", p.ID)
		if err := u.runSteps(p); err != nil {
			return err
		}
		if err := u.runChecks(p); err != nil {
			return err
		}
	}

	return nil
}

func (u *updater) pipeline() (cd.Pipeline, error) {
	var w cd.Pipeline

	resp, err := http.Get(apiURL + api.PipelineCurrent)
	if err != nil {
		return w, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&w)
	if err != nil {
		return w, err
	}

	return w, err
}

func (u *updater) pipelineStatus(status appsv1.DeploymentStatus) error {
	payload, err := json.Marshal(cd.Agent{
		Name:   u.agentName,
		Status: status,
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(apiURL+api.PipelinesStatus, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("no 200 OK, but %s", resp.Status)
	}

	return nil
}

func (u *updater) runSteps(p cd.Pipeline) error {
	for _, s := range p.Steps {
		if err := u.runStep(s); err != nil {
			return err
		}
	}

	return nil
}

func (u *updater) runStep(step cd.Step) error {
	args := []string{"-c"}
	for _, c := range step.Commands {
		args = append(args, c)
	}

	p := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      step.Name,
			Namespace: namespace,
		},
		Spec: corev1.PodSpec{
			ServiceAccountName: "cd",
			Containers: []corev1.Container{{
				Name:            step.Name,
				Image:           step.Image,
				ImagePullPolicy: corev1.PullAlways,
				Command:         []string{"sh"},
				Args:            args,
			}},
			RestartPolicy: corev1.RestartPolicyNever,
		},
	}

	_, err := u.client.CoreV1().Pods(namespace).Create(&p)
	if err != nil {
		return err
	}
	defer func(p *corev1.Pod) {
		_ = u.client.CoreV1().Pods(namespace).Delete(p.Name, nil)
	}(&p)

	watch, err := u.client.CoreV1().Pods(namespace).Watch(metav1.ListOptions{
		LabelSelector: labelsSelector(p.GetLabels()),
		Watch:         true,
	})
	if err != nil {
		return err
	}

	for event := range watch.ResultChan() {
		pod := event.Object.(*corev1.Pod)

		if pod.Status.Phase == corev1.PodSucceeded {
			return nil
		}
		if pod.Status.Phase == corev1.PodFailed {
			return fmt.Errorf("step failed")
		}
	}

	return nil
}

func (u *updater) runChecks(p cd.Pipeline) error {
	for _, c := range p.Checks {
		if err := u.runCheck(c); err != nil {
			return err
		}
	}

	return nil
}

func (u *updater) runCheck(check cd.Check) error {
	p := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      strings.ToLower(check.Name),
			Namespace: namespace,
		},
		Spec: corev1.PodSpec{
			ServiceAccountName: "cd",
			Containers: []corev1.Container{{
				Name:            strings.ToLower(check.Name),
				Image:           check.Image,
				ImagePullPolicy: corev1.PullAlways,
			}},
			RestartPolicy: corev1.RestartPolicyNever,
		},
	}

	pod, err := u.client.CoreV1().Pods(namespace).Create(&p)
	if err != nil {
		return err
	}

	defer func(pod *corev1.Pod) {
		time.Sleep(time.Minute)
		_ = u.client.CoreV1().Pods(namespace).Delete(pod.Name, nil)
	}(pod)

	return nil
}

func labelsSelector(ls map[string]string) string {
	var selectors []string
	for key, value := range ls {
		selectors = append(selectors, key+"="+value)
	}
	return strings.Join(selectors, ",")
}
