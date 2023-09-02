package alyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

// ECS 接口
type OpsOss interface {
	ListBuckets() []string
	ListFiles(bucketName string) []string
	Get(id string) (string, error)
	IsFileExist(bucketName, fileName string) (bool, error)
}

// 实现 ECS 接口的结构体
type MyOpsOss struct {
	Client *oss.Client
}

func (e *MyOpsOss) ListBuckets() []string {
	// 列举当前账号所有地域下的存储空间。
	var buckets []string
	marker := ""
	for {
		lsRes, err := e.Client.ListBuckets(oss.Marker(marker))
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}

		// 默认情况下一次返回100条记录。
		for _, bucket := range lsRes.Buckets {
			fmt.Println("Bucket: ", bucket.Name)
			buckets = append(buckets, bucket.Name)
		}

		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
	return buckets
}

// list files in a bucket
func (e *MyOpsOss) ListFiles(bucketName string) []string {
	// 模拟返回 ECS 列表
	// 列举当前账号所有地域下的存储空间。
	bucket, err := e.Client.Bucket(bucketName)
	if err != nil {
		fmt.Println(err)
	}
	// 列举所有文件。
	marker := ""
	var files []string
	for {
		lsRes, err := bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			fmt.Println(err)
		}
		// 打印列举结果。默认情况下，一次返回100条记录。
		for _, object := range lsRes.Objects {
			fmt.Println("Object Name: ", object.Key)
			files = append(files, object.Key)
		}
		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
	return files
}

func (e *MyOpsOss) IsFileExist(bucketName, fileName string) (bool, error) {
	// 判断文件是否存在。
	// yourObjectName填写不包含Bucket名称在内的Object的完整路径。
	bucket, err := e.Client.Bucket(bucketName)
	if err != nil {
		fmt.Println(err)
	}
	isExist, err := bucket.IsObjectExist(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("Exist:", isExist)
	return isExist, nil
}

// Get 方法的实现
func (e *MyOpsOss) Get(id string) (string, error) {
	// 根据指定 id 模拟获取 ECS 信息
	if id == "ecs-1" {
		return "ECS-1 details", nil
	} else if id == "ecs-2" {
		return "ECS-2 details", nil
	} else if id == "ecs-3" {
		return "ECS-3 details", nil
	} else {
		return "", fmt.Errorf("ECS not found")
	}
}
