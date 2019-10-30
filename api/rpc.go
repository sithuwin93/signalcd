package api

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/signalcd/signalcd/signalcd"
	signalcdproto "github.com/signalcd/signalcd/signalcd/proto"
	"golang.org/x/net/context"
	"golang.org/x/xerrors"
)

// RPC implement the gRPC server connecting it to a SignalDB
type RPC struct {
	DB     SignalDB
	logger log.Logger
}

// NewRPC creates a new gRPC Server
func NewRPC(db SignalDB, logger log.Logger) *RPC {
	return &RPC{
		DB:     db,
		logger: logger,
	}
}

// CurrentDeployment returns the current Deployment
func (r *RPC) CurrentDeployment(ctx context.Context, req *signalcdproto.CurrentDeploymentRequest) (*signalcdproto.CurrentDeploymentResponse, error) {
	deployment, err := r.DB.GetCurrentDeployment()
	if err != nil {
		return nil, xerrors.Errorf("failed to get current deployment: %w", err)
	}

	steps := func(steps1 []signalcd.Step) []*signalcdproto.Step {
		var steps2 []*signalcdproto.Step
		for _, s := range steps1 {
			steps2 = append(steps2, &signalcdproto.Step{
				Name:             s.Name,
				Image:            s.Image,
				ImagePullSecrets: s.ImagePullSecrets,
				Commands:         s.Commands,
			})
		}
		return steps2
	}

	checks := func(checks1 []signalcd.Check) []*signalcdproto.Check {
		var checks2 []*signalcdproto.Check
		for _, c := range checks1 {
			checks2 = append(checks2, &signalcdproto.Check{
				Name:             c.Name,
				Image:            c.Image,
				ImagePullSecrets: c.ImagePullSecrets,
				Duration:         int64(c.Duration.Seconds()),
			})
		}
		return checks2
	}

	return &signalcdproto.CurrentDeploymentResponse{
		CurrentDeployment: &signalcdproto.Deployment{
			Number:  deployment.Number,
			Created: deployment.Created.Unix(),

			Pipeline: &signalcdproto.Pipeline{
				Id:     deployment.Pipeline.ID,
				Name:   deployment.Pipeline.Name,
				Steps:  steps(deployment.Pipeline.Steps),
				Checks: checks(deployment.Pipeline.Checks),
			},
		},
	}, nil
}

// DeploymentStatusSetter sets the phase for a specific Deployment by its number
type DeploymentStatusSetter interface {
	SetDeploymentStatus(context.Context, int64, signalcd.DeploymentPhase) (signalcd.Deployment, error)
}

// SetDeploymentStatus sets the phase for a specific Deployment when receiving a request
func (r *RPC) SetDeploymentStatus(ctx context.Context, req *signalcdproto.SetDeploymentStatusRequest) (*signalcdproto.SetDeploymentStatusResponse, error) {
	var phase signalcd.DeploymentPhase

	switch req.Status.Phase {
	case signalcdproto.DeploymentStatus_UNKNOWN:
		phase = signalcd.Unknown
	case signalcdproto.DeploymentStatus_SUCCESS:
		phase = signalcd.Success
	case signalcdproto.DeploymentStatus_FAILURE:
		phase = signalcd.Failure
	case signalcdproto.DeploymentStatus_PROGRESS:
		phase = signalcd.Progress
	case signalcdproto.DeploymentStatus_PENDING:
		phase = signalcd.Pending
	case signalcdproto.DeploymentStatus_KILLED:
		phase = signalcd.Killed
	}

	// TODO: Use returned Deployment in SetDeploymentStatusResponse
	_, err := r.DB.SetDeploymentStatus(ctx, req.Number, phase)
	if err != nil {
		return nil, xerrors.Errorf("failed to update status: %w", err)
	}

	return &signalcdproto.SetDeploymentStatusResponse{}, nil
}

// StepStatusSetter saves the logs for a Deployment step by its number
type StepStatusSetter interface {
	SetStepStatus(ctx context.Context, deployment int64, step int64, status signalcd.Status) error
}

// StepStatus saves the logs for a specific deployment and step coming from an agent
func (r *RPC) StepStatus(ctx context.Context, req *signalcdproto.StepStatusRequest) (*signalcdproto.StepStatusResponse, error) {
	var phase signalcd.DeploymentPhase

	switch req.GetPhase() {
	case signalcdproto.StepStatusRequest_UNKNOWN:
		phase = signalcd.Unknown
	case signalcdproto.StepStatusRequest_SUCCESS:
		phase = signalcd.Success
	case signalcdproto.StepStatusRequest_FAILURE:
		phase = signalcd.Failure
	case signalcdproto.StepStatusRequest_PROGRESS:
		phase = signalcd.Progress
	case signalcdproto.StepStatusRequest_PENDING:
		phase = signalcd.Pending
	case signalcdproto.StepStatusRequest_KILLED:
		phase = signalcd.Killed
	}

	err := r.DB.SetStepStatus(ctx, req.GetDeployment(), req.GetStep(), signalcd.Status{
		Phase: phase,
		Logs:  req.GetLogs(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to save logs: %w", err)
	}

	return &signalcdproto.StepStatusResponse{}, nil
}
