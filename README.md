# 🖼️ S3 File Upload Microservice (Go)

This is a lightweight microservice written in Go for handling image uploads to AWS S3 using **presigned URLs**.
## 📦 Features
* ✅ Generate S3 presigned URLs for uploading images

* ✅ Secure upload without exposing credentials

* ✅ Built using Gin and AWS SDK v2

* ✅ Easily deployable as a container or microservice

## 🚀 Getting Started
1. Update .env with AWS IAM user credentials
```bash
AWS_API_KEY="XXX"
AWS_SECRET="XXXXXXX"
```
2. Install Dependencies
```bash
go mod tidy
```
3. Run the Service
```bash
go run main.go
```

## 🛠 Built With
- Go
- Gin
- AWS SDK for Go v2


## ✨ Author
Made with ❤️ by Vamsi Rayapati
