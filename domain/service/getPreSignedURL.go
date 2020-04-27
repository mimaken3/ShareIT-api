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
	privateKey := os.Getenv("AWS_S3_SECRET_ACCESS_KEY")
	region := "ap-northeast-1"
	bucketName := "share-it-test"
	fileName := iconName

	creds := credentials.NewStaticCredentials(accessKey, privateKey, "")
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(region),
	}))
	s3Client := s3.New(sess)

	req, _ := s3Client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("/user-icons/" + fileName),
	})

	preSignedURL, err = req.Presign(15 * time.Minute)

	return
}
