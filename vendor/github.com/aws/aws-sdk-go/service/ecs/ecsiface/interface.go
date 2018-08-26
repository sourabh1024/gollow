// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ecsiface provides an interface to enable mocking the Amazon EC2 Container Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ecsiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ecs"
)

// ECSAPI provides an interface to enable mocking the
// ecs.ECS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon EC2 Container Service.
//    func myFunc(svc ecsiface.ECSAPI) bool {
//        // Make svc.CreateCluster request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := ecs.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockECSClient struct {
//        ecsiface.ECSAPI
//    }
//    func (m *mockECSClient) CreateCluster(input *ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockECSClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ECSAPI interface {
	CreateCluster(*ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error)
	CreateClusterWithContext(aws.Context, *ecs.CreateClusterInput, ...request.Option) (*ecs.CreateClusterOutput, error)
	CreateClusterRequest(*ecs.CreateClusterInput) (*request.Request, *ecs.CreateClusterOutput)

	CreateService(*ecs.CreateServiceInput) (*ecs.CreateServiceOutput, error)
	CreateServiceWithContext(aws.Context, *ecs.CreateServiceInput, ...request.Option) (*ecs.CreateServiceOutput, error)
	CreateServiceRequest(*ecs.CreateServiceInput) (*request.Request, *ecs.CreateServiceOutput)

	DeleteAttributes(*ecs.DeleteAttributesInput) (*ecs.DeleteAttributesOutput, error)
	DeleteAttributesWithContext(aws.Context, *ecs.DeleteAttributesInput, ...request.Option) (*ecs.DeleteAttributesOutput, error)
	DeleteAttributesRequest(*ecs.DeleteAttributesInput) (*request.Request, *ecs.DeleteAttributesOutput)

	DeleteCluster(*ecs.DeleteClusterInput) (*ecs.DeleteClusterOutput, error)
	DeleteClusterWithContext(aws.Context, *ecs.DeleteClusterInput, ...request.Option) (*ecs.DeleteClusterOutput, error)
	DeleteClusterRequest(*ecs.DeleteClusterInput) (*request.Request, *ecs.DeleteClusterOutput)

	DeleteService(*ecs.DeleteServiceInput) (*ecs.DeleteServiceOutput, error)
	DeleteServiceWithContext(aws.Context, *ecs.DeleteServiceInput, ...request.Option) (*ecs.DeleteServiceOutput, error)
	DeleteServiceRequest(*ecs.DeleteServiceInput) (*request.Request, *ecs.DeleteServiceOutput)

	DeregisterContainerInstance(*ecs.DeregisterContainerInstanceInput) (*ecs.DeregisterContainerInstanceOutput, error)
	DeregisterContainerInstanceWithContext(aws.Context, *ecs.DeregisterContainerInstanceInput, ...request.Option) (*ecs.DeregisterContainerInstanceOutput, error)
	DeregisterContainerInstanceRequest(*ecs.DeregisterContainerInstanceInput) (*request.Request, *ecs.DeregisterContainerInstanceOutput)

	DeregisterTaskDefinition(*ecs.DeregisterTaskDefinitionInput) (*ecs.DeregisterTaskDefinitionOutput, error)
	DeregisterTaskDefinitionWithContext(aws.Context, *ecs.DeregisterTaskDefinitionInput, ...request.Option) (*ecs.DeregisterTaskDefinitionOutput, error)
	DeregisterTaskDefinitionRequest(*ecs.DeregisterTaskDefinitionInput) (*request.Request, *ecs.DeregisterTaskDefinitionOutput)

	DescribeClusters(*ecs.DescribeClustersInput) (*ecs.DescribeClustersOutput, error)
	DescribeClustersWithContext(aws.Context, *ecs.DescribeClustersInput, ...request.Option) (*ecs.DescribeClustersOutput, error)
	DescribeClustersRequest(*ecs.DescribeClustersInput) (*request.Request, *ecs.DescribeClustersOutput)

	DescribeContainerInstances(*ecs.DescribeContainerInstancesInput) (*ecs.DescribeContainerInstancesOutput, error)
	DescribeContainerInstancesWithContext(aws.Context, *ecs.DescribeContainerInstancesInput, ...request.Option) (*ecs.DescribeContainerInstancesOutput, error)
	DescribeContainerInstancesRequest(*ecs.DescribeContainerInstancesInput) (*request.Request, *ecs.DescribeContainerInstancesOutput)

	DescribeServices(*ecs.DescribeServicesInput) (*ecs.DescribeServicesOutput, error)
	DescribeServicesWithContext(aws.Context, *ecs.DescribeServicesInput, ...request.Option) (*ecs.DescribeServicesOutput, error)
	DescribeServicesRequest(*ecs.DescribeServicesInput) (*request.Request, *ecs.DescribeServicesOutput)

	DescribeTaskDefinition(*ecs.DescribeTaskDefinitionInput) (*ecs.DescribeTaskDefinitionOutput, error)
	DescribeTaskDefinitionWithContext(aws.Context, *ecs.DescribeTaskDefinitionInput, ...request.Option) (*ecs.DescribeTaskDefinitionOutput, error)
	DescribeTaskDefinitionRequest(*ecs.DescribeTaskDefinitionInput) (*request.Request, *ecs.DescribeTaskDefinitionOutput)

	DescribeTasks(*ecs.DescribeTasksInput) (*ecs.DescribeTasksOutput, error)
	DescribeTasksWithContext(aws.Context, *ecs.DescribeTasksInput, ...request.Option) (*ecs.DescribeTasksOutput, error)
	DescribeTasksRequest(*ecs.DescribeTasksInput) (*request.Request, *ecs.DescribeTasksOutput)

	DiscoverPollEndpoint(*ecs.DiscoverPollEndpointInput) (*ecs.DiscoverPollEndpointOutput, error)
	DiscoverPollEndpointWithContext(aws.Context, *ecs.DiscoverPollEndpointInput, ...request.Option) (*ecs.DiscoverPollEndpointOutput, error)
	DiscoverPollEndpointRequest(*ecs.DiscoverPollEndpointInput) (*request.Request, *ecs.DiscoverPollEndpointOutput)

	ListAttributes(*ecs.ListAttributesInput) (*ecs.ListAttributesOutput, error)
	ListAttributesWithContext(aws.Context, *ecs.ListAttributesInput, ...request.Option) (*ecs.ListAttributesOutput, error)
	ListAttributesRequest(*ecs.ListAttributesInput) (*request.Request, *ecs.ListAttributesOutput)

	ListClusters(*ecs.ListClustersInput) (*ecs.ListClustersOutput, error)
	ListClustersWithContext(aws.Context, *ecs.ListClustersInput, ...request.Option) (*ecs.ListClustersOutput, error)
	ListClustersRequest(*ecs.ListClustersInput) (*request.Request, *ecs.ListClustersOutput)

	ListClustersPages(*ecs.ListClustersInput, func(*ecs.ListClustersOutput, bool) bool) error
	ListClustersPagesWithContext(aws.Context, *ecs.ListClustersInput, func(*ecs.ListClustersOutput, bool) bool, ...request.Option) error

	ListContainerInstances(*ecs.ListContainerInstancesInput) (*ecs.ListContainerInstancesOutput, error)
	ListContainerInstancesWithContext(aws.Context, *ecs.ListContainerInstancesInput, ...request.Option) (*ecs.ListContainerInstancesOutput, error)
	ListContainerInstancesRequest(*ecs.ListContainerInstancesInput) (*request.Request, *ecs.ListContainerInstancesOutput)

	ListContainerInstancesPages(*ecs.ListContainerInstancesInput, func(*ecs.ListContainerInstancesOutput, bool) bool) error
	ListContainerInstancesPagesWithContext(aws.Context, *ecs.ListContainerInstancesInput, func(*ecs.ListContainerInstancesOutput, bool) bool, ...request.Option) error

	ListServices(*ecs.ListServicesInput) (*ecs.ListServicesOutput, error)
	ListServicesWithContext(aws.Context, *ecs.ListServicesInput, ...request.Option) (*ecs.ListServicesOutput, error)
	ListServicesRequest(*ecs.ListServicesInput) (*request.Request, *ecs.ListServicesOutput)

	ListServicesPages(*ecs.ListServicesInput, func(*ecs.ListServicesOutput, bool) bool) error
	ListServicesPagesWithContext(aws.Context, *ecs.ListServicesInput, func(*ecs.ListServicesOutput, bool) bool, ...request.Option) error

	ListTaskDefinitionFamilies(*ecs.ListTaskDefinitionFamiliesInput) (*ecs.ListTaskDefinitionFamiliesOutput, error)
	ListTaskDefinitionFamiliesWithContext(aws.Context, *ecs.ListTaskDefinitionFamiliesInput, ...request.Option) (*ecs.ListTaskDefinitionFamiliesOutput, error)
	ListTaskDefinitionFamiliesRequest(*ecs.ListTaskDefinitionFamiliesInput) (*request.Request, *ecs.ListTaskDefinitionFamiliesOutput)

	ListTaskDefinitionFamiliesPages(*ecs.ListTaskDefinitionFamiliesInput, func(*ecs.ListTaskDefinitionFamiliesOutput, bool) bool) error
	ListTaskDefinitionFamiliesPagesWithContext(aws.Context, *ecs.ListTaskDefinitionFamiliesInput, func(*ecs.ListTaskDefinitionFamiliesOutput, bool) bool, ...request.Option) error

	ListTaskDefinitions(*ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error)
	ListTaskDefinitionsWithContext(aws.Context, *ecs.ListTaskDefinitionsInput, ...request.Option) (*ecs.ListTaskDefinitionsOutput, error)
	ListTaskDefinitionsRequest(*ecs.ListTaskDefinitionsInput) (*request.Request, *ecs.ListTaskDefinitionsOutput)

	ListTaskDefinitionsPages(*ecs.ListTaskDefinitionsInput, func(*ecs.ListTaskDefinitionsOutput, bool) bool) error
	ListTaskDefinitionsPagesWithContext(aws.Context, *ecs.ListTaskDefinitionsInput, func(*ecs.ListTaskDefinitionsOutput, bool) bool, ...request.Option) error

	ListTasks(*ecs.ListTasksInput) (*ecs.ListTasksOutput, error)
	ListTasksWithContext(aws.Context, *ecs.ListTasksInput, ...request.Option) (*ecs.ListTasksOutput, error)
	ListTasksRequest(*ecs.ListTasksInput) (*request.Request, *ecs.ListTasksOutput)

	ListTasksPages(*ecs.ListTasksInput, func(*ecs.ListTasksOutput, bool) bool) error
	ListTasksPagesWithContext(aws.Context, *ecs.ListTasksInput, func(*ecs.ListTasksOutput, bool) bool, ...request.Option) error

	PutAttributes(*ecs.PutAttributesInput) (*ecs.PutAttributesOutput, error)
	PutAttributesWithContext(aws.Context, *ecs.PutAttributesInput, ...request.Option) (*ecs.PutAttributesOutput, error)
	PutAttributesRequest(*ecs.PutAttributesInput) (*request.Request, *ecs.PutAttributesOutput)

	RegisterContainerInstance(*ecs.RegisterContainerInstanceInput) (*ecs.RegisterContainerInstanceOutput, error)
	RegisterContainerInstanceWithContext(aws.Context, *ecs.RegisterContainerInstanceInput, ...request.Option) (*ecs.RegisterContainerInstanceOutput, error)
	RegisterContainerInstanceRequest(*ecs.RegisterContainerInstanceInput) (*request.Request, *ecs.RegisterContainerInstanceOutput)

	RegisterTaskDefinition(*ecs.RegisterTaskDefinitionInput) (*ecs.RegisterTaskDefinitionOutput, error)
	RegisterTaskDefinitionWithContext(aws.Context, *ecs.RegisterTaskDefinitionInput, ...request.Option) (*ecs.RegisterTaskDefinitionOutput, error)
	RegisterTaskDefinitionRequest(*ecs.RegisterTaskDefinitionInput) (*request.Request, *ecs.RegisterTaskDefinitionOutput)

	RunTask(*ecs.RunTaskInput) (*ecs.RunTaskOutput, error)
	RunTaskWithContext(aws.Context, *ecs.RunTaskInput, ...request.Option) (*ecs.RunTaskOutput, error)
	RunTaskRequest(*ecs.RunTaskInput) (*request.Request, *ecs.RunTaskOutput)

	StartTask(*ecs.StartTaskInput) (*ecs.StartTaskOutput, error)
	StartTaskWithContext(aws.Context, *ecs.StartTaskInput, ...request.Option) (*ecs.StartTaskOutput, error)
	StartTaskRequest(*ecs.StartTaskInput) (*request.Request, *ecs.StartTaskOutput)

	StopTask(*ecs.StopTaskInput) (*ecs.StopTaskOutput, error)
	StopTaskWithContext(aws.Context, *ecs.StopTaskInput, ...request.Option) (*ecs.StopTaskOutput, error)
	StopTaskRequest(*ecs.StopTaskInput) (*request.Request, *ecs.StopTaskOutput)

	SubmitContainerStateChange(*ecs.SubmitContainerStateChangeInput) (*ecs.SubmitContainerStateChangeOutput, error)
	SubmitContainerStateChangeWithContext(aws.Context, *ecs.SubmitContainerStateChangeInput, ...request.Option) (*ecs.SubmitContainerStateChangeOutput, error)
	SubmitContainerStateChangeRequest(*ecs.SubmitContainerStateChangeInput) (*request.Request, *ecs.SubmitContainerStateChangeOutput)

	SubmitTaskStateChange(*ecs.SubmitTaskStateChangeInput) (*ecs.SubmitTaskStateChangeOutput, error)
	SubmitTaskStateChangeWithContext(aws.Context, *ecs.SubmitTaskStateChangeInput, ...request.Option) (*ecs.SubmitTaskStateChangeOutput, error)
	SubmitTaskStateChangeRequest(*ecs.SubmitTaskStateChangeInput) (*request.Request, *ecs.SubmitTaskStateChangeOutput)

	UpdateContainerAgent(*ecs.UpdateContainerAgentInput) (*ecs.UpdateContainerAgentOutput, error)
	UpdateContainerAgentWithContext(aws.Context, *ecs.UpdateContainerAgentInput, ...request.Option) (*ecs.UpdateContainerAgentOutput, error)
	UpdateContainerAgentRequest(*ecs.UpdateContainerAgentInput) (*request.Request, *ecs.UpdateContainerAgentOutput)

	UpdateContainerInstancesState(*ecs.UpdateContainerInstancesStateInput) (*ecs.UpdateContainerInstancesStateOutput, error)
	UpdateContainerInstancesStateWithContext(aws.Context, *ecs.UpdateContainerInstancesStateInput, ...request.Option) (*ecs.UpdateContainerInstancesStateOutput, error)
	UpdateContainerInstancesStateRequest(*ecs.UpdateContainerInstancesStateInput) (*request.Request, *ecs.UpdateContainerInstancesStateOutput)

	UpdateService(*ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error)
	UpdateServiceWithContext(aws.Context, *ecs.UpdateServiceInput, ...request.Option) (*ecs.UpdateServiceOutput, error)
	UpdateServiceRequest(*ecs.UpdateServiceInput) (*request.Request, *ecs.UpdateServiceOutput)

	WaitUntilServicesInactive(*ecs.DescribeServicesInput) error
	WaitUntilServicesInactiveWithContext(aws.Context, *ecs.DescribeServicesInput, ...request.WaiterOption) error

	WaitUntilServicesStable(*ecs.DescribeServicesInput) error
	WaitUntilServicesStableWithContext(aws.Context, *ecs.DescribeServicesInput, ...request.WaiterOption) error

	WaitUntilTasksRunning(*ecs.DescribeTasksInput) error
	WaitUntilTasksRunningWithContext(aws.Context, *ecs.DescribeTasksInput, ...request.WaiterOption) error

	WaitUntilTasksStopped(*ecs.DescribeTasksInput) error
	WaitUntilTasksStoppedWithContext(aws.Context, *ecs.DescribeTasksInput, ...request.WaiterOption) error
}

var _ ECSAPI = (*ecs.ECS)(nil)
