package temporal_test

import (
	"context"
	gracesAdapters "github.com/devops-olympus/graces-adapters/sdk/temporal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"testing"
)

// mockClient 是Temporal客户端的一个mock实现，用于测试
type mockClient struct {
	mock.Mock
	client.Client
}

func (m *mockClient) ExecuteWorkflow(ctx context.Context, options client.StartWorkflowOptions, workflow interface{}, args ...interface{}) (client.WorkflowRun, error) {
	args = append([]interface{}{ctx, options, workflow}, args...)
	returns := m.Called(args...)
	return returns.Get(0).(client.WorkflowRun), returns.Error(1)
}

func (m *mockClient) DescribeWorkflowExecution(ctx context.Context, request *workflowservice.DescribeWorkflowExecutionRequest) (*workflowservice.DescribeWorkflowExecutionResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*workflowservice.DescribeWorkflowExecutionResponse), args.Error(1)
}

func (m *mockClient) TerminateWorkflow(ctx context.Context, workflowID, runID, reason string, details ...interface{}) error {
	args := m.Called(ctx, workflowID, runID, reason, details)
	return args.Error(0)
}

// TestStartWorkflow 测试启动工作流的逻辑
func TestStartWorkflow(t *testing.T) {
	mc := new(mockClient)
	adapter := gracesAdapters.NewAdapter(mc)
	mc.On("ExecuteWorkflow", mock.Anything, mock.AnythingOfType("client.StartWorkflowOptions"), "TestWorkflow", nil).Return(nil, nil)

	err := adapter.StartWorkflow("test-workflow-id", "TestWorkflow", nil)
	assert.NoError(t, err)
	mc.AssertExpectations(t)
}

// TestGetWorkflowStatus 测试获取工作流状态的逻辑
func TestGetWorkflowStatus(t *testing.T) {
	mc := new(mockClient)
	adapter := gracesAdapters.NewAdapter(mc)
	response := &workflowservice.DescribeWorkflowExecutionResponse{
		//WorkflowExecutionInfo: &workflowservice.WorkflowExecutionInfo{
		//	Status: enums.WORKFLOW_EXECUTION_STATUS_RUNNING,
		//},
	}
	mc.On("DescribeWorkflowExecution", mock.Anything, mock.AnythingOfType("*workflowservice.DescribeWorkflowExecutionRequest")).Return(response, nil)

	status, err := adapter.GetWorkflowStatus("test-workflow-id")
	assert.NoError(t, err)
	assert.Equal(t, enums.WorkflowExecutionStatus_name[int32(enums.WORKFLOW_EXECUTION_STATUS_RUNNING)], status)
	mc.AssertExpectations(t)
}

// TestStopWorkflow 测试停止工作流的逻辑
func TestStopWorkflow(t *testing.T) {
	mc := new(mockClient)
	adapter := gracesAdapters.NewAdapter(mc)
	mc.On("TerminateWorkflow", mock.Anything, "test-workflow-id", "", "Manual termination", nil).Return(nil)

	err := adapter.StopWorkflow("test-workflow-id")
	assert.NoError(t, err)
	mc.AssertExpectations(t)
}
