package presign

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smartbot/storage/pkg/validator"
)

type PresignController struct {
	service PresignService
}

func (ac *PresignController) GetUploadPreSignedUrl(c *gin.Context) {
	var request UploadSignedUrlRequest
	err := validator.ValidateBody(c, &request)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	res, err := ac.service.GetUploadPreSignedUrl(request)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, res)

}

func (ac *PresignController) GetDownloadPreSignedUrl(c *gin.Context) {
	var request DownloadSignedUrlRequest
	err := validator.ValidateBody(c, &request)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	res, err := ac.service.GetDownloadPreSignedUrl(request)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, res)

}
