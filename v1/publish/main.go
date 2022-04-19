package publish

import (
	"net/http"
	"upload/v1/controller"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

var bucketname string
var AWSRegion string
var filepath string

//HealthCheck function
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "200",
	})
}

//UploadImage function
func UploadImage(c *gin.Context) {
	sess := c.MustGet("sess").(*session.Session)
	uploader := s3manager.NewUploader(sess)
	bucketname = c.Request.Header.Get("bucketname")
	file, header, err := c.Request.FormFile("filename")
	filename := header.Filename

	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketname),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    err.Error(),
			"uploader": up,
		})
		return
	}
	filepath = "https://" + bucketname + "." + "s3-" + AWSRegion + ".amazonaws.com/" + filename
	c.JSON(http.StatusOK, gin.H{
		"filepath": filepath,
	})
}

//StartServer function
func StartServer() {
	sess := controller.ConnectAws()
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("sess", sess)
		c.Next()
	})
	router.POST("/v1/upload", UploadImage)
	router.GET("/health", HealthCheck)
	_ = router.Run(":8080")
}
