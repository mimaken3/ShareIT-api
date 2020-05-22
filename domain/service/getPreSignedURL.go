package service

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetPreSignedURL(iconName string) (preSignedURL string, err error) {
	accessKey := os.Getenv("AWS_S3_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_S3_SECRET_ACCESS_KEY")
	region := "ap-northeast-1"
	bucketName := "share-it-test"

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secretAccessKey, ""),
		Region:      aws.String(region),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		return "", err
	}

	s3Client := s3.New(newSession)

	req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("user-icons/" + iconName),
	})

	// 3日間
	preSignedURL, _, err = req.PresignRequest(24 * time.Hour * 3)

	return
}
