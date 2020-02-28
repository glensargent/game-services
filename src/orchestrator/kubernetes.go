package main

import (
	"encoding/json"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type cluster struct {
	config    *rest.Config
	clientset *kubernetes.Clientset
}

func newCluster() cluster {
	return cluster{}
}

func (c *cluster) init() {
	// creates the in-cluster config
	var err error
	c.config, err = rest.InClusterConfig()
	if err != nil {
		// no point exiting if it can't connect to the cluster
		panic(err.Error())
	}

	// creates the clientset
	c.clientset, err = kubernetes.NewForConfig(c.config)
	if err != nil {
		// no point exiting if it can't connect to the cluster
		panic(err.Error())
	}
}

func (c cluster) makePod(yaml []byte) error {
	var p v1.Pod
	err := json.Unmarshal([]byte(yaml), &p)
	if err != nil {
		return err
	}
	pod, poderr := c.clientset.CoreV1().Pods(apiv1.NamespaceDefault).Create(&p)
	if poderr != nil {
		return poderr
	}
	fmt.Printf("Created Pod %q.\n", pod.GetObjectMeta().GetName())
	return nil
}
