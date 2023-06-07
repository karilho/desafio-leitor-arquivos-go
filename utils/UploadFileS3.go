package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"path/filepath"
)

const (
	AWS_S3_BUCKET  = "buckettest"
	AWS_S3_REGION  = "us-east-1"
	ACESSKEY       = "localstack"
	SECRETACESSKEY = "localstack"
)

func UploadFile(bucketName, filePath string, region string) error {
	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(region),
		Credentials:      credentials.NewStaticCredentials(ACESSKEY, SECRETACESSKEY, ""),
		Endpoint:         aws.String("http://localhost:4566"),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}

	// Create an S3 service client
	svc := s3.New(sess)

	// Open the file you want to upload
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Get the file size
	fileInfo, _ := file.Stat()
	var _ int64 = fileInfo.Size()

	filename := filepath.Base(filePath)

	// Create an S3 upload input object
	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   file,
		ACL:    aws.String("public-read"), // Set the ACL to public-read if you want the uploaded file to be publicly accessible
	}

	// Perform the upload
	_, err = svc.PutObject(input)
	if err != nil {
		return fmt.Errorf("failed to upload file: %v", err)
	}

	return nil
}
