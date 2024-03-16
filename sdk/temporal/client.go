package temporal

import (
	"go.temporal.io/sdk/client"
)

// ClientWrapper 提供了对Temporal client.Client的封装，可以添加额外的功能或简化操作
type ClientWrapper struct {
	client client.Client
}

// NewClientWrapper 创建一个新的ClientWrapper实例
func NewClientWrapper(options client.Options) (*ClientWrapper, error) {
	c, err := client.Dial(options)
	if err != nil {
		return nil, err
	}
	return &ClientWrapper{client: c}, nil
}

// GetClient 返回封装的client.Client实例，以便进行工作流操作
func (cw *ClientWrapper) GetClient() client.Client {
	return cw.client
}

// Close 关闭Temporal客户端连接
func (cw *ClientWrapper) Close() {
	cw.client.Close()
}
