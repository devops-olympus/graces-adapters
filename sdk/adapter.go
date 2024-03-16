package sdk

// WorkflowAdapter 定义工作流引擎适配器的接口
type WorkflowAdapter interface {
	StartWorkflow(workflowID, workflowType string, params interface{}) error
	GetWorkflowStatus(workflowID string) (string, error)
	StopWorkflow(workflowID string) error
}
