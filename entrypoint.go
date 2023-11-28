package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/muhamadrafipamungkas/cloud-function-golang/src/hello"
	"github.com/muhamadrafipamungkas/cloud-function-golang/src/user"
)

func init() {
	functions.HTTP("HelloWorld", hello.HelloWorld)
	functions.HTTP("HelloDynamic", hello.HelloDynamic)
	functions.HTTP("HelloUser", user.HelloUser)
}

func main() {
	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
