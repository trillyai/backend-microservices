package repository

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/core/auth"
	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
	"github.com/trillyai/backend-microservices/core/env"
	"github.com/trillyai/backend-microservices/services/profile/shared"
)

func (r repository) UploadProfileImage(ctx context.Context, request shared.UploadProfileImageRequest) (shared.UploadProfileImageResponse, error) {
	claims := ctx.Value("user").(*auth.Claims)
	if claims.UserName == "" {
		return shared.UploadProfileImageResponse{}, errors.New("context not found")
	}

	fileExt, err := validateFile(request.File)
	if err != nil {
		return shared.UploadProfileImageResponse{}, err
	}

	// Upload the file to S3
	fullpath, err := r.uploadToS3(ctx, request.File, claims.UserName, fileExt)
	if err != nil {
		return shared.UploadProfileImageResponse{}, err
	}

	// Update user's profile image URL
	err = updateUserProfileImage(ctx, claims.UserId, fullpath)
	if err != nil {
		return shared.UploadProfileImageResponse{}, err
	}

	return shared.UploadProfileImageResponse{Url: fullpath}, nil
}

func (r repository) uploadToS3(ctx context.Context, file *multipart.FileHeader, userName, fileExt string) (string, error) {
	// Open the file
	fileContent, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileContent.Close()

	// Set AWS session with provided credentials
	sess, err := createAWSSession()
	if err != nil {
		return "", err
	}

	// Create S3 service
	svc := s3.New(sess)
	putObjectKey := env.AwsProfileImageFolderPath + userName + "." + fileExt

	// Upload the profile image
	_, err = svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(env.AwsProfileImageBucketName),
		Key:    aws.String(putObjectKey),
		Body:   fileContent,
	})
	if err != nil {
		r.logger.Error(err.Error())
		return "", err
	}

	return fmt.Sprintf("https://trillyai.s3.amazonaws.com/%s", putObjectKey), nil
}

func updateUserProfileImage(ctx context.Context, userID uuid.UUID, imageURL string) error {
	user := tables.User{}
	user.ProfileImage = imageURL
	_, err := postgres.Update[struct{}, tables.User](ctx, map[string]interface{}{"Id": userID}, user)
	if err != nil {
		return err
	}
	return nil
}

func validateFile(file *multipart.FileHeader) (string, error) {
	const maxProfileImageSize int64 = 5 * 1024 * 1024 // 5 megabayt
	if file.Size > maxProfileImageSize {
		return "", errors.New("too big image")
	}

	// Get the file extension
	fileExt := getFileExtension(file.Filename)

	// Supported image extensions
	supportedExts := map[string]bool{"jpg": true, "jpeg": true, "png": true}

	// Return error if the file extension is not supported
	if !supportedExts[fileExt] {
		return "", errors.New("unsupported image format")
	}

	return fileExt, nil
}

func getFileExtension(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}

func createAWSSession() (*session.Session, error) {
	accessKey := env.AwsProfileImageAccessKey
	secretKey := env.AwsProfileImageSecretAccessKey
	region := env.AwsProfileImageRegion

	// Create and return AWS session with provided credentials
	return session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
}
