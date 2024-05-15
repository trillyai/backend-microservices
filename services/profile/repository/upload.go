package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/trillyai/backend-microservices/core/auth"
	"github.com/trillyai/backend-microservices/services/profile/shared"
)

// UploadProfileImage implements contracts.Repository.
func (r repository) UploadProfileImage(ctx context.Context, request shared.UploadProfileImageRequest) (shared.UploadProfileImageResponse, error) {
	claims := ctx.Value("user").(*auth.Claims)
	if claims.UserName == "" {
		return shared.UploadProfileImageResponse{}, errors.New("context not found")
	}

	const maxProfileImageSize int64 = 5 * 1024 * 1024 // 5 megabayt
	if request.File.Size > maxProfileImageSize {
		return shared.UploadProfileImageResponse{}, errors.New("too big image")
	}

	// Get the file extension
	fileExt := getFileExtension(request.File.Filename)

	// Supported image extensions
	supportedExts := map[string]bool{"jpg": true, "jpeg": true, "png": true}

	// Return error if the file extension is not supported
	if !supportedExts[fileExt] {
		return shared.UploadProfileImageResponse{}, errors.New("unsupported image format")
	}

	// Open the file
	file, err := request.File.Open()
	if err != nil {
		return shared.UploadProfileImageResponse{}, err
	}
	defer file.Close()

	// Set AWS session with provided credentials
	sess, err := createAWSSession()
	if err != nil {
		return shared.UploadProfileImageResponse{}, err
	}

	// Create S3 service
	svc := s3.New(sess)

	// Upload the profile image
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("YOUR_BUCKET_NAME"),
		Key:    aws.String("profile-images/" + claims.UserName + "." + fileExt),
		Body:   file,
	})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.UploadProfileImageResponse{}, err
	}

	return shared.UploadProfileImageResponse{}, nil
}

// createAWSSession creates and returns a new AWS session
func createAWSSession() (*session.Session, error) {
	accessKey := "YOUR_ACCESS_KEY"
	secretKey := "YOUR_SECRET_KEY"
	region := "YOUR_REGION"

	// Create and return AWS session with provided credentials
	return session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
}

// getFileExtension returns the extension of the file
func getFileExtension(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}
