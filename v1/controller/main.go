package controller

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

//ConnectAws function
func ConnectAws() *session.Session {
	var AWSSecretKey string
	var AWSAccessKey string
	var AWSRegion string

	AWSAccessKey = GetEnvWithKey("AWS_ACCESS_KEY_ID")
	AWSSecretKey = GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
	AWSRegion = GetEnvWithKey("REGION")
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(AWSRegion),
			Credentials: credentials.NewStaticCredentials(
				AWSAccessKey,
				AWSSecretKey,
				"",
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}

//GetEnvWithKey : get env value
func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}
