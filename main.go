package main

import (
	"context"
	"time"

	"github.com/demonCoder95/ready-wait-controller/pkg/utils"
	log "github.com/sirupsen/logrus"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

func main() {
	clientset, err := utils.CreateLocalClient()
	if err != nil {
		log.Fatalf("error creating k8s client: %v", err)
	}

	factory := informers.NewFilteredSharedInformerFactory(clientset, 0, "kube-system", nil)
	deploymentInformer := factory.Apps().V1().Deployments().Informer()

	deploymentInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			log.Printf("Deployment added: %s", utils.AsDeployment(obj).Name)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			log.Printf("Deployment updated: %s", utils.AsDeployment(oldObj).Name)
		},
		DeleteFunc: func(obj interface{}) {
			log.Printf("Deployment deleted: %s", utils.AsDeployment(obj).Name)
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	deploymentInformer.Run(ctx.Done())

	log.Infof("Exiting...")
}
