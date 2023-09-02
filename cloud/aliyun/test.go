package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"net/http"
	ops "ops_cloud/cloud"
)

const (
	EndPoint        = "https://oss-cn-hangzhou.aliyuncs.com"
	AccessKeyId     = "LTAIbevdrr7MYQRd"
	AccessKeySecret = "hk3JYSaPb8K0RXI6LrSZlRWp4lTOMj"
)

var ossClient ops.OpsOss
var ecsClient ops.OpsEcs

func init() {
	ossclient, _ := oss.New(EndPoint, AccessKeyId, AccessKeySecret)

	ossClient = &ops.MyOpsOss{
		Client: ossclient,
	}
	ecsclient, _ := ecs.NewClientWithAccessKey(EndPoint, AccessKeyId, AccessKeySecret)
	ecsClient = &ops.MyOpsEcs{
		Client: ecsclient,
	}
}

func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func ListOssBucketsHandler(c *gin.Context) {
	buckets := ossClient.ListBuckets()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": buckets,
	})
}

func ListEcsInstanceHandler(c *gin.Context) {
	buckets := ecsClient.ListEcs()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": buckets,
	})
}

func FileExistsHandler(c *gin.Context) {
	bucket := c.Request.Header.Get("bucketname")
	filename := c.Request.Header.Get("filename")
	exists, err := ossClient.IsFileExist(bucket, filename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "fail",
			"data": err.Error(),
		})
		return
	}
	if exists {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "file exist",
			"data": true,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "file not exist",
			"data": false,
		})
		return
	}
}
