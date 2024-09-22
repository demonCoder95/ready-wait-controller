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

	informer := utils.NewInformerGenerator(clientset, "StatefulSet", "kube-system").GetInformer()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	informer.Run(ctx.Done())

	log.Infof("Exiting...")
}
