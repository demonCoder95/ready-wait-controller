package utils

import (
	"reflect"

	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
)

// This package provides utility functions for working with Kubernetes objects

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
