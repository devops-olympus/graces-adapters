package main

import (
	gracesAdapters "172.17.162.204/devops-olympus/graces-adapters/sdk/temporal"
	"go.temporal.io/sdk/client"
	"log"
)

func main() {
	// 使用ClientWrapper初始化Temporal客户端
	cw, err := gracesAdapters.NewClientWrapper(client.Options{
		HostPort:  "localhost:7233", // 指定Temporal服务地址
		Namespace: "default",
	})
	if err != nil {
		log.Fatalf("Failed to create Temporal client wrapper: %v", err)
	}
	defer cw.Close() // 确保在程序结束时关闭客户端

	// 使用封装的客户端创建适配器实例
	adapter := gracesAdapters.NewAdapter(cw.GetClient())

	// 使用适配器启动工作流
	workflowID := "exampleWorkflowID"
	workflowType := "ExampleWorkflow"
	err = adapter.StartWorkflow(workflowID, workflowType, nil)
	if err != nil {
		log.Fatalf("Failed to start workflow: %v", err)
	}
	log.Println("Workflow started successfully")

	// 查询工作流状态
	status, err := adapter.GetWorkflowStatus(workflowID)
	if err != nil {
		log.Fatalf("Failed to get workflow status: %v", err)
	}
	log.Printf("Workflow status: %s\n", status)

	// 停止工作流
	err = adapter.StopWorkflow(workflowID)
	if err != nil {
		log.Fatalf("Failed to stop workflow: %v", err)
	}
	log.Println("Workflow stopped successfully")
}
