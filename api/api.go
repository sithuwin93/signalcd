package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-openapi/loads"
	restmiddleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/signalcd/signalcd/api/v1/models"
	"github.com/signalcd/signalcd/api/v1/restapi"
	"github.com/signalcd/signalcd/api/v1/restapi/operations"
	"github.com/signalcd/signalcd/api/v1/restapi/operations/agents"
	"github.com/signalcd/signalcd/api/v1/restapi/operations/deployments"
	"github.com/signalcd/signalcd/api/v1/restapi/operations/pipeline"
	"github.com/signalcd/signalcd/signalcd"
	"golang.org/x/xerrors"
)

// SignalDB is the union of all necessary interfaces for the API
type SignalDB interface {
	DeploymentLister
	DeploymentStatusSetter
	CurrentDeploymentGetter
	CurrentDeploymentSetter
	PipelinesLister
	PipelineCreator
	AgentsLister
	AgentStatusSetter
}

// Events to Deployments that should be sent via SSE (Server Sent Events)
type Events interface {
	SubscribeDeployments(channel chan signalcd.Deployment) signalcd.Subscription
	UnsubscribeDeployments(s signalcd.Subscription)
}

// NewV1 creates a new v1 API
func NewV1(logger log.Logger, db SignalDB, events Events) (*chi.Mux, error) {
	router := chi.NewRouter()

	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return nil, xerrors.Errorf("failed to load embedded swagger file: %w", err.Error())
	}

	api := operations.NewCdAPI(swaggerSpec)

	// Skip the  redoc middleware, only serving the OpenAPI specification and
	// the API itself via RoutesHandler. See:
	// https://github.com/go-swagger/go-swagger/issues/1779
	api.Middleware = func(b restmiddleware.Builder) http.Handler {
		return restmiddleware.Spec("", swaggerSpec.Raw(), api.Context().RoutesHandler(b))
	}

	api.AgentsAgentsHandler = getAgentsHandler(db)
	api.DeploymentsDeploymentsHandler = getDeploymentsHandler(db)
	api.DeploymentsCurrentDeploymentHandler = getCurrentDeploymentHandler(db)
	api.DeploymentsSetCurrentDeploymentHandler = setCurrentDeploymentHandler(db, logger)
	api.PipelinePipelineHandler = getPipelineHandler(db)
	api.PipelinePipelinesHandler = getPipelinesHandler(db)
	api.PipelineCreateHandler = createPipelineHandler(db)

	router.Mount("/", api.Serve(nil))

	router.Get("/api/v1/deployments/events", deploymentEventsHandler(logger, events))

	return router, nil
}

func deploymentEventsHandler(logger log.Logger, events Events) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusMethodNotAllowed)
			return
		}

		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		level.Debug(logger).Log("msg", "streaming deployment http connection just opened")

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		deploymentEvents := make(chan signalcd.Deployment, 8)
		subscription := events.SubscribeDeployments(deploymentEvents)

		defer func() {
			events.UnsubscribeDeployments(subscription)
			level.Debug(logger).Log("msg", "streaming deployment http connection just closed")
		}()

		for {
			select {
			case <-ctx.Done():
				close(deploymentEvents)
				return
			case deployment := <-deploymentEvents:
				model := getModelsDeployment(deployment)
				j, err := json.Marshal(model)
				if err != nil {
					return // TODO
				}

				_, err = fmt.Fprintf(w, "data: %s\n\n", j)
				if err != nil {
					return // TODO
				}

				flusher.Flush()
			}
		}
	}
}

func getModelsPipeline(p signalcd.Pipeline) *models.Pipeline {
	mp := &models.Pipeline{
		ID:      strfmt.UUID(p.ID),
		Name:    p.Name,
		Created: strfmt.DateTime(p.Created),
	}

	for _, s := range p.Steps {
		mp.Steps = append(mp.Steps, &models.Step{
			Name:             &s.Name,
			Image:            &s.Image,
			ImagePullSecrets: s.ImagePullSecrets,
			Commands:         s.Commands,
		})
	}

	for _, c := range p.Checks {
		var env []*models.CheckEnvironmentItems0
		for key, value := range c.Environment {
			env = append(env, &models.CheckEnvironmentItems0{
				Key:   key,
				Value: value,
			})
		}

		mp.Checks = append(mp.Checks, &models.Check{
			Name:             &c.Name,
			Image:            &c.Image,
			ImagePullSecrets: c.ImagePullSecrets,
			Duration:         c.Duration.Seconds(),
			Environment:      env,
		})
	}

	return mp
}

type AgentsLister interface {
	ListAgents() ([]signalcd.Agent, error)
}

func getAgentsHandler(lister AgentsLister) agents.AgentsHandlerFunc {
	return func(params agents.AgentsParams) restmiddleware.Responder {
		as, err := lister.ListAgents()
		if err != nil {
			return agents.NewAgentsInternalServerError()
		}

		var payload []*models.Agent

		for _, a := range as {
			name := a.Name
			ready := a.Ready()
			payload = append(payload, &models.Agent{
				Name:  &name,
				Ready: &ready,
			})
		}

		return agents.NewAgentsOK().WithPayload(payload)
	}
}

// PipelinesLister returns a list of Pipelines
type PipelinesLister interface {
	ListPipelines() ([]signalcd.Pipeline, error)
}

func getPipelinesHandler(lister PipelinesLister) pipeline.PipelinesHandlerFunc {
	return func(params pipeline.PipelinesParams) restmiddleware.Responder {
		var payload []*models.Pipeline

		pipelines, err := lister.ListPipelines()
		if err != nil {
			return pipeline.NewPipelinesInternalServerError()
		}

		for _, p := range pipelines {
			payload = append(payload, getModelsPipeline(p))
		}

		return pipeline.NewPipelinesOK().WithPayload(payload)
	}
}

func getDeploymentStatusPhase(phase signalcd.DeploymentPhase) string {
	switch phase {
	case signalcd.Success:
		return models.DeploymentstatusPhaseSuccess
	case signalcd.Failure:
		return models.DeploymentstatusPhaseFailure
	case signalcd.Progress:
		return models.DeploymentstatusPhaseProgress
	default:
		return models.DeploymentstatusPhaseUnknown
	}
}

// DeploymentLister lists all Deployments
type DeploymentLister interface {
	ListDeployments() ([]signalcd.Deployment, error)
}

func getDeploymentsHandler(lister DeploymentLister) deployments.DeploymentsHandlerFunc {
	return func(params deployments.DeploymentsParams) restmiddleware.Responder {
		var payload []*models.Deployment

		list, err := lister.ListDeployments()
		if err != nil {
			return deployments.NewDeploymentsInternalServerError()
		}

		for _, d := range list {
			payload = append(payload, getModelsDeployment(d))
		}

		return deployments.NewDeploymentsOK().WithPayload(payload)
	}
}

func getModelsDeployment(fd signalcd.Deployment) *models.Deployment {
	return &models.Deployment{
		Number:   &fd.Number,
		Created:  strfmt.DateTime(fd.Created),
		Started:  strfmt.DateTime(fd.Started),
		Finished: strfmt.DateTime(fd.Finished),
		Pipeline: getModelsPipeline(fd.Pipeline),
		Status: &models.Deploymentstatus{
			Phase: getDeploymentStatusPhase(fd.Status.Phase),
		},
	}
}

// CurrentDeploymentGetter gets the current Deployment
type CurrentDeploymentGetter interface {
	GetCurrentDeployment() (signalcd.Deployment, error)
}

func getCurrentDeploymentHandler(getter CurrentDeploymentGetter) deployments.CurrentDeploymentHandlerFunc {
	return func(params deployments.CurrentDeploymentParams) restmiddleware.Responder {
		d, err := getter.GetCurrentDeployment()
		if err != nil {
			return deployments.NewSetCurrentDeploymentInternalServerError()
		}

		return deployments.NewCurrentDeploymentOK().WithPayload(getModelsDeployment(d))
	}
}

// CurrentDeploymentSetter gets a Pipeline and then creates a new Deployments
type CurrentDeploymentSetter interface {
	PipelineGetter
	CreateDeployment(signalcd.Pipeline) (signalcd.Deployment, error)
}

func setCurrentDeploymentHandler(creator CurrentDeploymentSetter, logger log.Logger) deployments.SetCurrentDeploymentHandlerFunc {
	return func(params deployments.SetCurrentDeploymentParams) restmiddleware.Responder {
		p, err := creator.GetPipeline(params.Pipeline)
		if err != nil {
			logger.Log("msg", "failed to get pipeline", "id", params.Pipeline, "err", err)
			return deployments.NewSetCurrentDeploymentInternalServerError()
		}

		d, err := creator.CreateDeployment(p)
		if err != nil {
			logger.Log("msg", "failed to create deployment", "err", err)
			return deployments.NewSetCurrentDeploymentInternalServerError()
		}

		return deployments.NewSetCurrentDeploymentOK().WithPayload(getModelsDeployment(d))
	}
}

// PipelineGetter gets a new Pipeline
type PipelineGetter interface {
	GetPipeline(id string) (signalcd.Pipeline, error)
}

func getPipelineHandler(getter PipelineGetter) pipeline.PipelineHandlerFunc {
	return func(params pipeline.PipelineParams) restmiddleware.Responder {
		p, err := getter.GetPipeline(params.ID)
		if err != nil {
			return pipeline.NewPipelineInternalServerError()
		}
		return pipeline.NewPipelineOK().WithPayload(getModelsPipeline(p))
	}
}

// PipelineCreator creates a new Pipeline
type PipelineCreator interface {
	CreatePipeline(signalcd.Pipeline) (signalcd.Pipeline, error)
}

func createPipelineHandler(creator PipelineCreator) pipeline.CreateHandlerFunc {
	return func(params pipeline.CreateParams) restmiddleware.Responder {
		p, err := creator.CreatePipeline(fromModelPipeline(params.Pipeline))
		if err != nil {
			return pipeline.NewCreateInternalServerError()
		}

		return pipeline.NewCreateOK().WithPayload(getModelsPipeline(p))
	}
}

func fromModelPipeline(m *models.Pipeline) signalcd.Pipeline {
	p := signalcd.Pipeline{
		ID:      m.ID.String(),
		Name:    m.Name,
		Created: time.Time(m.Created),
	}

	for _, c := range m.Checks {
		check := signalcd.Check{
			Name:             *c.Name,
			Image:            *c.Image,
			ImagePullSecrets: c.ImagePullSecrets,
			Duration:         time.Duration(c.Duration) * time.Second,
		}

		for _, env := range c.Environment {
			check.Environment[env.Key] = env.Value
		}

		p.Checks = append(p.Checks, check)
	}

	for _, s := range m.Steps {
		step := signalcd.Step{
			Name:             *s.Name,
			Image:            *s.Image,
			ImagePullSecrets: s.ImagePullSecrets,
			Commands:         s.Commands,
		}

		p.Steps = append(p.Steps, step)
	}

	return p
}
