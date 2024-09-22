// This package provides utility functions for working with Kubernetes objects
package utils

import (
	"fmt"
	"reflect"

	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

// The Informer Generator pattern
type InformerGenerator struct {
	// The Kubernetes client
	clientSet kubernetes.Interface
	// The type of the object
	objType string
	// The namespace
	namespace string
}

// NewInformGenerator creates a new InformerGenerator
func NewInformerGenerator(cs kubernetes.Interface, objType, ns string) *InformerGenerator {
	return &InformerGenerator{
		clientSet: cs,
		objType:   objType,
		namespace: ns,
	}
}

// GetInformer returns a new shared informer for the specified type
func (g *InformerGenerator) GetInformer() (cache.SharedIndexInformer, error) {

	// Create a new shared informer factory
	factory := informers.NewFilteredSharedInformerFactory(g.clientSet, 0, g.namespace, nil)

	// Create a new informer, so we can update the handler functions
	var informer cache.SharedIndexInformer

	// Create a new informer for the specified type
	switch g.objType {
	case "Deployment":
		informer = factory.Apps().V1().Deployments().Informer()
	case "DaemonSet":
		informer = factory.Apps().V1().DaemonSets().Informer()
	case "StatefulSet":
		informer = factory.Apps().V1().StatefulSets().Informer()
	default:
		log.Errorf("unsupported type: %s", g.objType)
		// TODO: add custom error types?
		return nil, fmt.Errorf("unsupported type: %s", g.objType)
	}

	return g.AddEventHandler(informer)
}

// AddEventHandler adds event handlers to the informer
func (g *InformerGenerator) AddEventHandler(informer cache.SharedIndexInformer) (cache.SharedIndexInformer, error) {
	var err error

	// Add handler functions for events
	_, err = informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			log.Printf("%s added: %s", g.objType, g.GetName(obj))
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			log.Printf("%s updated: %s", g.objType, g.GetName(oldObj))
		},
		DeleteFunc: func(obj interface{}) {
			log.Printf("%s deleted: %s", g.objType, g.GetName(obj))
		},
	})

	return informer, err
}

// GetName returns the name of the object
func (g *InformerGenerator) GetName(obj interface{}) string {
	switch g.objType {
	case "Deployment":
		return AsDeployment(obj).Name
	case "DaemonSet":
		return AsDaemonSet(obj).Name
	case "StatefulSet":
		return AsStatefulSet(obj).Name
	default:
		log.Errorf("unsupported type: %s", g.objType)
		return ""
	}
}

func AsDeployment(obj interface{}) *appsv1.Deployment {
	dep, ok := obj.(*appsv1.Deployment)
	if !ok {
		log.Errorf("expected Deployment, got %s: %v", reflect.TypeOf(obj), obj)
	}
	return dep
}

func AsDaemonSet(obj interface{}) *appsv1.DaemonSet {
	ds, ok := obj.(*appsv1.DaemonSet)
	if !ok {
		log.Errorf("expected DaemonSet, got %s: %v", reflect.TypeOf(obj), obj)
	}
	return ds
}

func AsStatefulSet(obj interface{}) *appsv1.StatefulSet {
	sts, ok := obj.(*appsv1.StatefulSet)
	if !ok {
		log.Errorf("expected StatefulSet, got %s: %v", reflect.TypeOf(obj), obj)
	}
	return sts
}
