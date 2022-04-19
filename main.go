package main

import (
	"upload/v1/publish"
)

func main() {
	// Set mode to release
	//gin.SetMode(gin.ReleaseMode)
	publish.StartServer()
}
