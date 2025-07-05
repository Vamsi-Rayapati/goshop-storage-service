package presign

type UploadSignedUrlRequest struct {
	Key         string `json:"key" validate:"required"`
	FileName    string `json:"file_name" validate:"required"`
	ContentType string `json:"content_type"  validate:"required"`
}

type UploadSignedUrlResponse struct {
	UploadUrl string `json:"upload_url"`
	FilePath  string `json:"path"`
	PublicUrl string `json:"public_url"`
}

type DownloadSignedUrlRequest struct {
	Key      string `json:"key" validate:"required"`
	FilePath string `json:"path" validate:"required"`
}

type DownloadSignedUrlResponse struct {
	DownloadUrl string `json:"download_url"`
}
