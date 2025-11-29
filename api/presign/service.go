package presign

import (
	"fmt"
	"log"

	"github.com/smartbot/storage/pkg/aws"
	"github.com/smartbot/storage/pkg/config"
	"github.com/smartbot/storage/pkg/errors"
)

type PresignService struct {
}

func (as *PresignService) GetUploadPreSignedUrl(request UploadSignedUrlRequest) (*UploadSignedUrlResponse, *errors.ApiError) {
	bucket, ok := BucketMap[request.Key]

	if !ok {
		return nil, errors.NotFoundError("Key not found")
	}
	url, key, err := aws.GetNewSignedUrl(bucket, request.FileName, request.ContentType)

	if err != nil {
		return nil, err
	}

	publicUrl := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, config.Config.AWSRegion, *key)

	return &UploadSignedUrlResponse{
		UploadUrl: *url,
		FilePath:  *key,
		PublicUrl: publicUrl,
	}, nil

}

//aws.GetSignedUrl("goshop-avatar", user.Avatar)

func (as *PresignService) GetDownloadPreSignedUrl(request DownloadSignedUrlRequest) (*DownloadSignedUrlResponse, *errors.ApiError) {
	bucket, ok := BucketMap[request.Key]

	if !ok {
		return nil, errors.NotFoundError("Key not found")
	}

	url, err := aws.GetSignedUrl(bucket, request.FilePath)
	log.Println("GetDownloadPreSignedUrl", url)

	if err != nil {
		return nil, err
	}

	return &DownloadSignedUrlResponse{
		DownloadUrl: *url,
	}, nil

}
