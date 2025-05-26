package core

import (
	"context"
	"fmt"

	"cloud.google.com/go/iam/apiv1/iampb"
	"cloud.google.com/go/run/apiv2"
	runpb "cloud.google.com/go/run/apiv2/runpb"
)

type PigenCoreGCP struct {
	ProjectID string
	Region    string
}

func (g *PigenCoreGCP) DeployPigenCore() (string, error) {
	ctx := context.Background()
	containerPort := &runpb.ContainerPort{
		ContainerPort:5000,
	}
	ressources := &runpb.ResourceRequirements{
		Limits: map[string]string{
			"cpu":    "1",
			"memory": "1G",
		},
	}
	googleCloudRunV2Container := &runpb.Container{
		Image: "fedimersni/pigen-core:latest",
		Name: "pigen-core",
		Ports: []*runpb.ContainerPort{containerPort},
		Resources: ressources,
	}

	revisionTemplate := &runpb.RevisionTemplate{
		Containers: []*runpb.Container{googleCloudRunV2Container},
	}
	service := &runpb.Service{
		Template: revisionTemplate,
		Ingress: runpb.IngressTraffic_INGRESS_TRAFFIC_ALL,
	}

	createServiceRequest := &runpb.CreateServiceRequest{
		Parent: "projects/"+g.ProjectID+"/locations/" + g.Region,
		Service: service,
		ServiceId:"pigen-core",
	}
	client, err := run.NewServicesClient(ctx)
	if err != nil {
		return "", fmt.Errorf("error waiting for Cloud Run service creation: %v", err)
	}
	op, err := client.CreateService(ctx, createServiceRequest)
	if err != nil {
		return "", fmt.Errorf("failed to create Cloud Run service: %v", err)
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		return "", fmt.Errorf("cloud run creator failed to wait for service creation: %v", err)
	}
	// Get the service name
	serviceName := resp.Name

	// Set up IAM policy to allow unauthenticated invocations
	iamClient, err := run.NewServicesClient(ctx)
	if err != nil {
		return "", fmt.Errorf("error creating IAM client: %v", err)
	}

	// First, get the current IAM policy
	getIamPolicyRequest := &iampb.GetIamPolicyRequest{
		Resource: serviceName,
	}
	policy, err := iamClient.GetIamPolicy(ctx, getIamPolicyRequest)
	if err != nil {
		return "", fmt.Errorf("error getting IAM policy: %v", err)
	}

	// Add the binding for allUsers
	policy.Bindings = append(policy.Bindings, &iampb.Binding{
		Role:    "roles/run.invoker",
		Members: []string{"allUsers"},
	})

	// Set the updated policy
	setIamPolicyRequest := &iampb.SetIamPolicyRequest{
		Resource: serviceName,
		Policy:   policy,
	}
	_, err = iamClient.SetIamPolicy(ctx, setIamPolicyRequest)
	if err != nil {
		return "", fmt.Errorf("error setting IAM policy: %v", err)
	}

	return resp.Uri, nil
}