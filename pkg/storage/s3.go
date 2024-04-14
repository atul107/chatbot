package storage

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/chatbot/pkg/config"
	"github.com/chatbot/pkg/logger"
)

// S3Client...
type S3Client struct {
	s3     *s3.S3
	config config.AWS
	logger logger.Interface
}

// NewS3Client ...
func NewS3Client(awsConfig config.AWS, logger logger.Interface) (*S3Client, error) {
	s3Config := &aws.Config{
		Region:      aws.String(awsConfig.Region),
		Credentials: credentials.NewStaticCredentials(awsConfig.AccessKeyID, awsConfig.AccessKeySecret, ""),
	}
	s3 := s3.New(session.Must(session.NewSession(s3Config)))
	store := &S3Client{
		s3:     s3,
		config: awsConfig,
		logger: logger,
	}
	return store, nil
}

// Upload ...
func (s *S3Client) Upload(file *multipart.FileHeader, userId string) (string, error) {
	// Open the file to be uploaded
	fileReader, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileReader.Close()

	// Set the file type based on the extension
	contentType := "application/octet-stream"
	switch ext := filepath.Ext(file.Filename); ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	}

	objectName := userId + "-" + file.Filename
	// Upload the file to S3
	_, err = s.s3.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(s.config.BucketName),
		Key:         aws.String(objectName),
		Body:        fileReader,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}

	// Get the URL of the uploaded file
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.config.BucketName, s.config.Region, objectName)
	return url, nil
}

// Retrieve downloads a file from S3 using the provided unique identifier (S3 object key).
func (c *S3Client) Retrieve(key string) ([]byte, error) {
	resp, err := c.s3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(c.config.BucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve file from S3: %w", err)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		return nil, fmt.Errorf("failed to read file body from S3 response: %w", err)
	}

	return buf.Bytes(), nil
}
