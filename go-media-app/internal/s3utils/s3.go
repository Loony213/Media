package s3utils

import (
	"fmt"
	"time"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	awsRegion   = "us-west-2" 
	bucketName  = "profile_pics" 
)

func GetProfilePhotoURL(userId string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	if err != nil {
		return "", fmt.Errorf("unable to create AWS session: %v", err)
	}


	s3Client := s3.New(sess)


	key := fmt.Sprintf("profile_pics/%s.jpg", userId)

	req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})

	url, err := req.Presign(1 * time.Minute)
	if err != nil {
		return "", fmt.Errorf("unable to sign request: %v", err)
	}

	return url, nil
}
