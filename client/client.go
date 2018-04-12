package client

import (
	"context"
	"github.com/ericchiang/k8s"
	appsv1 "github.com/ericchiang/k8s/apis/apps/v1"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
	"github.com/ericchiang/k8s/util/intstr"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-k8s-helper/config"
	"strconv"
)

// Client represents a headr-k8s-client that is responsible for create/delete a caddy server container in the cluster.
type Client interface {
	CreateCaddyService(siteID uint) error
	DeleteCaddyService(siteID uint) error
}

type k8sclient struct {
	client *k8s.Client
	logger log.Logger
}

func (c k8sclient) CreateCaddyService(siteID uint) error {
	siteIDs := strconv.Itoa(int(siteID))
	// create deployment
	var (
		name      = "siteid-" + siteIDs + "-service"
		namespace = "user"
		labels    = map[string]string{
			"app": name,
		}
		replicas        int32 = 1
		volumeName            = "data"
		mountPath             = "/www"
		image                 = "seagullbird/headr-caddy:1.0.0"
		imagePullPolicy       = "IfNotPresent"
		hostPath              = "/home/docker/data/sites/" + siteIDs + "/public"
		nfsPvcName            = "nfs"
	)

	var volumeSource corev1.VolumeSource
	switch config.Dev {
	case "true":
		volumeSource = corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: &hostPath,
			},
		}
	case "false":
		volumeSource = corev1.VolumeSource{
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: &nfsPvcName,
			},
		}
	}

	dp := &appsv1.Deployment{
		Metadata: &metav1.ObjectMeta{
			Name:      &name,
			Namespace: &namespace,
			Labels:    labels,
		},
		Spec: &appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: &corev1.PodTemplateSpec{
				Metadata: &metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: &corev1.PodSpec{
					Volumes: []*corev1.Volume{
						{
							Name:         &volumeName,
							VolumeSource: &volumeSource,
						},
					},
					Containers: []*corev1.Container{
						{
							Name:            &name,
							Image:           &image,
							ImagePullPolicy: &imagePullPolicy,
							VolumeMounts: []*corev1.VolumeMount{
								{
									Name:      &volumeName,
									MountPath: &mountPath,
								},
							},
						},
					},
				},
			},
		},
	}
	if err := c.client.Create(context.TODO(), dp); err != nil {
		return err
	}

	// create service
	var (
		svcType    string
		svcProto         = "TCP"
		port       int32 = 2018
		targetPort int32 = 2015
	)

	switch config.Dev {
	case "true":
		svcType = "NodePort"
	case "false":
		svcType = "LoadBalancer"
	}

	svc := &corev1.Service{
		Metadata: &metav1.ObjectMeta{
			Name:      &name,
			Namespace: &namespace,
			Labels:    labels,
		},
		Spec: &corev1.ServiceSpec{
			Selector: labels,
			Type:     &svcType,
			Ports: []*corev1.ServicePort{
				{
					Protocol: &svcProto,
					Port:     &port,
					TargetPort: &intstr.IntOrString{
						IntVal: &targetPort,
					},
				},
			},
		},
	}
	c.logger.Log("svc", svc.String())
	return c.client.Create(context.TODO(), svc)
}

func (c k8sclient) DeleteCaddyService(siteID uint) error {
	// delete deployment
	name := "siteid-" + strconv.Itoa(int(siteID)) + "-service"

	var dp appsv1.Deployment
	if err := c.client.Get(context.TODO(), "user", name, &dp); err != nil {
		c.logger.Log("error_desc", "failed to get deployment resource", "error", err)
		return err
	}
	if err := c.client.Delete(context.TODO(), &dp); err != nil {
		c.logger.Log("error_desc", "failed to delete deployment resource", "error", err)
		return err
	}
	// delete service
	var svc corev1.Service
	if err := c.client.Get(context.TODO(), "user", name, &svc); err != nil {
		c.logger.Log("error_desc", "failed to get service resource", "error", err)
		return err
	}
	if err := c.client.Delete(context.TODO(), &svc); err != nil {
		c.logger.Log("error_desc", "failed to delete service resource", "error", err)
		return err
	}
	return nil
}

// NewClient returns a Client instance with given logger.
func NewClient(logger log.Logger) (Client, error) {
	client, err := k8s.NewInClusterClient()
	if err != nil {
		return nil, err
	}

	return k8sclient{
		client: client,
		logger: logger,
	}, nil
}
