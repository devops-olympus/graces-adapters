package temporal

import (
	"context"
	"errors"
	"go.temporal.io/api/serviceerror"
	"go.temporal.io/sdk/client"
)

// Adapter 实现 WorkflowAdapter 接口，适配Temporal工作流引擎
type Adapter struct {
	client client.Client // Temporal的客户端实例
}

// NewAdapter 创建新的Temporal适配器实例
func NewAdapter(c client.Client) *Adapter {
	return &Adapter{client: c}
}

// StartWorkflow 实现启动Temporal工作流的逻辑
func (a *Adapter) StartWorkflow(workflowID string, workflowType string, params interface{}) error {
	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "XXX_TASK_QUEUE",
	}
	_, err := a.client.ExecuteWorkflow(context.Background(), options, workflowType, params)
	return err
}

// GetWorkflowStatus 使用客户端和工作流ID查询状态
func (a *Adapter) GetWorkflowStatus(workflowID string) (string, error) {
	// 注意：此处假设a.client是已初始化的Temporal客户端实例
	describeResp, err := a.client.DescribeWorkflowExecution(context.Background(), workflowID, "")
	if err != nil {
		var notFound *serviceerror.NotFound
		if errors.As(err, &notFound) {
			return "NotFound", nil
		}
		return "", err // 返回错误
	}

	status := describeResp.WorkflowExecutionInfo.Status
	return status.String(), nil // 返回工作流状态的字符串表示
	//return enums.WorkflowExecutionStatus_name[int32(status)], nil
}

// StopWorkflow 停止正在运行的工作流
func (a *Adapter) StopWorkflow(workflowID string) error {
	return a.client.TerminateWorkflow(context.Background(), workflowID, "", "Manual termination", nil)
}
