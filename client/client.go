package client

import (
	"context"
	"github.com/ericchiang/k8s"
	appsv1 "github.com/ericchiang/k8s/apis/apps/v1"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	extensionsv1beta1 "github.com/ericchiang/k8s/apis/extensions/v1beta1"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
	"github.com/ericchiang/k8s/util/intstr"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-k8s-helper/config"
	"path/filepath"
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
	siteIDstr := strconv.Itoa(int(siteID))
	// create deployment
	var (
		name      = "siteid-" + siteIDstr + "-service"
		namespace = "default"
		labels    = map[string]string{
			"app": name,
		}
		replicas        int32 = 1
		volumeName            = "data"
		mountPath             = "/www"
		serverRootPath        = ""
		image                 = "seagullbird/headr-caddy:2.0.0"
		imagePullPolicy       = "IfNotPresent"
		hostPath              = "/home/docker/data/sites/" + siteIDstr + "/public"
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
		serverRootPath = mountPath
	case "false":
		volumeSource = corev1.VolumeSource{
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: &nfsPvcName,
			},
		}
		serverRootPath = filepath.Join(mountPath, "sites", siteIDstr, "public")
	}

	command := []string{"/bin/parent", "caddy", "--conf", "/etc/Caddyfile", "-root", serverRootPath, "--log", "stdout"}

	env_name := "SITENAME"
	env_val := "/" + siteIDstr
	env := corev1.EnvVar{Name: &env_name, Value: &env_val}

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
							Command:         command,
							Env:             []*corev1.EnvVar{&env},
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
		svcType          = "NodePort"
		svcProto         = "TCP"
		port       int32 = 2018
		targetPort int32 = 2015
	)

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
	if err := c.client.Create(context.TODO(), svc); err != nil {
		return err
	}

	if config.Dev == "true" {
		return nil
	}

	// Add usersites-ingress entry
	var ing extensionsv1beta1.Ingress
	if err := c.client.Get(context.TODO(), "default", "usersites-ingress", &ing); err != nil {
		return err
	}
	backendPath := "/" + siteIDstr
	backend := extensionsv1beta1.IngressBackend{
		ServiceName: &name,
		ServicePort: &intstr.IntOrString{
			IntVal: &port,
		},
	}
	newHTTPPath := extensionsv1beta1.HTTPIngressPath{
		Path:    &backendPath,
		Backend: &backend,
	}
	if ing.Spec.Rules[0].IngressRuleValue.Http == nil {
		ing.Spec.Rules[0].IngressRuleValue.Http = &extensionsv1beta1.HTTPIngressRuleValue{}
		ing.Spec.Rules[0].IngressRuleValue.Http.Paths = []*extensionsv1beta1.HTTPIngressPath{}
	}
	ing.Spec.Rules[0].IngressRuleValue.Http.Paths = append(ing.Spec.Rules[0].IngressRuleValue.Http.Paths, &newHTTPPath)
	return c.client.Update(context.TODO(), &ing)
}

func (c k8sclient) DeleteCaddyService(siteID uint) error {
	// delete deployment
	name := "siteid-" + strconv.Itoa(int(siteID)) + "-service"

	var dp appsv1.Deployment
	if err := c.client.Get(context.TODO(), "default", name, &dp); err != nil {
		c.logger.Log("error_desc", "failed to get deployment resource", "error", err)
		return err
	}
	if err := c.client.Delete(context.TODO(), &dp); err != nil {
		c.logger.Log("error_desc", "failed to delete deployment resource", "error", err)
		return err
	}
	// delete service
	var svc corev1.Service
	if err := c.client.Get(context.TODO(), "default", name, &svc); err != nil {
		c.logger.Log("error_desc", "failed to get service resource", "error", err)
		return err
	}
	if err := c.client.Delete(context.TODO(), &svc); err != nil {
		c.logger.Log("error_desc", "failed to delete service resource", "error", err)
		return err
	}

	if config.Dev == "true" {
		return nil
	}

	// delete usersites-ingress entry
	var ing extensionsv1beta1.Ingress
	if err := c.client.Get(context.TODO(), "default", "usersites-ingress", &ing); err != nil {
		c.logger.Log("error_desc", "failed to get usersites-ingress resource", "error", err)
		return err
	}

	index := 0
	found := false
	for i, v := range ing.Spec.Rules[0].IngressRuleValue.Http.Paths {
		if *v.Backend.ServiceName == name {
			found = true
			index = i
		}
	}
	if found {
		ing.Spec.Rules[0].IngressRuleValue.Http.Paths[index] = ing.Spec.Rules[0].IngressRuleValue.Http.Paths[len(ing.Spec.Rules[0].IngressRuleValue.Http.Paths)-1]
		ing.Spec.Rules[0].IngressRuleValue.Http.Paths = ing.Spec.Rules[0].IngressRuleValue.Http.Paths[:len(ing.Spec.Rules[0].IngressRuleValue.Http.Paths)-1]
		if err := c.client.Update(context.TODO(), &ing); err != nil {
			return err
		}
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
