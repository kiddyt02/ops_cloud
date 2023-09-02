package alyun

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// ECS 接口
type OpsEcs interface {
	ListEcs() []string
}

// 实现 ECS 接口的结构体
type MyOpsEcs struct {
	Client *ecs.Client
}

func (e *MyOpsEcs) ListEcs() []string {
	// 列举当前账号所有地域下的存储空间。
	request := ecs.CreateDescribeInstancesRequest()
	response, err := e.Client.DescribeInstances(request)
	if err != nil {
		fmt.Println("Failed to describe instances:", err)
	}

	for _, instance := range response.Instances.Instance {
		fmt.Printf("Instance ID: %s, Instance Name: %s\n", instance.InstanceId, instance.InstanceName)
	}
	return nil
}
