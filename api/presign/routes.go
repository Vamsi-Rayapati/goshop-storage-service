package presign

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup) {
	service := PresignService{}
	controller := PresignController{service: service}
	group.POST("/presign/upload", controller.GetUploadPreSignedUrl)
	group.POST("/presign/download", controller.GetDownloadPreSignedUrl)
}
