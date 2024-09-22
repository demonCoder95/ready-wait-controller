package main

import (
	"context"
	"time"

	"github.com/demonCoder95/ready-wait-controller/pkg/utils"
	log "github.com/sirupsen/logrus"
)

func main() {
	clientset, err := utils.CreateLocalClient()
	if err != nil {
		log.Fatalf("error creating k8s client: %v", err)
	}

	objType := "StatefulSet"
	informer, err := utils.NewInformerGenerator(clientset, objType, "kube-system").GetInformer()
	if err != nil {
		log.Fatalf("error creating informer for %s: %v", objType, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	informer.Run(ctx.Done())

	log.Infof("Exiting...")
}
