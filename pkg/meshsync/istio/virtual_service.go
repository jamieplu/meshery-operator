package istio

import (
	"log"

	common "github.com/layer5io/meshery-operator/pkg/common"
	client "github.com/layer5io/meshery-operator/pkg/meshsync/istio/client"
	"github.com/myntra/pipeline"
)

// VirtualService will implement step interface for VirtualService
type VirtualService struct {
	pipeline.StepContext
	// clients
	client     *client.IstioClient
	kubeclient *common.KubeClient
}

// Exec - step interface
func (vs *VirtualService) Exec(request *pipeline.Request) *pipeline.Result {
	// it will contain a pipeline to run
	log.Println("Virtual Service Discovery Started")

	// get all namespaces
	namespaces, err := vs.kubeclient.ListNamespace()
	if err != nil {
		return &pipeline.Result{
			Error: err,
		}
	}

	// virtual service for all namespace
	for _, namespace := range namespaces {
		virtualServices, err := vs.client.ListVirtualService(namespace.Name)
		if err != nil {
			return &pipeline.Result{
				Error: err,
			}
		}

		// process virtualServices
		for _, virtualService := range virtualServices {
			log.Printf("Discovered virtual service named %s in namespace %s", virtualService.Name, namespace.Name)
		}
	}

	// no data is feeded to future steps or stages
	return &pipeline.Result{
		Error: nil,
	}
}

// Cancel - step interface
func (vs *VirtualService) Cancel() error {
	vs.Status("cancel step")
	return nil
}
