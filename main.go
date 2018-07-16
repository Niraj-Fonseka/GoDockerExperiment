package main

import (
	"log"
	"time"

	sdk "github.com/gaia-pipeline/gosdk"
	"github.com/gin-gonic/gin"
)

var Jobs = sdk.Jobs{
	sdk.Job{
		Handler:     CreateUser,
		Title:       "Create DB User",
		Description: "Creates a database user with least privileged permissions.",
		Priority:    0,
	},
	sdk.Job{
		Handler:     MigrateDB,
		Title:       "DB Migration",
		Description: "Imports newest test data dump and migrates to newest version.",
		Priority:    10,
	},
	sdk.Job{
		Handler:     CreateNamespace,
		Title:       "Create K8S Namespace",
		Description: "Creates a new Kubernetes namespace for the new test environment.",
		Priority:    20,
	},
	sdk.Job{
		Handler:     CreateDeployment,
		Title:       "Create K8S Deployment",
		Description: "Creates a new Kubernetes deployment for the new test environment.",
		Priority:    30,
	},
	sdk.Job{
		Handler:     CreateService,
		Title:       "Create K8S Service",
		Description: "Creates a new Kubernetes service for the new test environment.",
		Priority:    30,
	},
	sdk.Job{
		Handler:     CreateIngress,
		Title:       "Create K8S Ingress",
		Description: "Creates a new Kubernetes ingress for the new test environment.",
		Priority:    30,
	},
	sdk.Job{
		Handler:     Cleanup,
		Title:       "Clean up",
		Description: "Removes all temporary files.",
		Priority:    40,
	},
}

func main() {

	if err := sdk.Serve(Jobs); err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Health is ok",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func CreateUser() error {
	log.Println("CreateUser has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	log.Println("CreateUser has been finished!")

	return nil
}

func MigrateDB() error {
	log.Println("MigrateDB has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	log.Println("MigrateDB has been finished!")

	return nil
}

func CreateNamespace() error {
	log.Println("CreateNamespace has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	log.Println("CreateNamespace has been finished!")

	return nil
}

func CreateDeployment() error {
	log.Println("CreateDeployment has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	log.Println("CreateDeployment has been finished!")

	return nil
}

func CreateService() error {
	log.Println("CreateService has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	log.Println("CreateService has been finished!")

	return nil
}

func CreateIngress() error {
	log.Println("CreateIngress has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	log.Println("CreateIngress has been finished!")

	return nil
}

func Cleanup() error {
	log.Println("Cleanup has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	log.Println("Cleanup has been finished!")

	return nil
}
