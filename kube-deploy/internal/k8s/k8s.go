package k8s

import (
	"context"
	"fmt"
	"io"
	"kube-deploy/internal/logger"
	"log/slog"

	corev1 "k8s.io/api/core/v1"
	k8Err "k8s.io/apimachinery/pkg/api/errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var log = logger.GetLogger()

// Checks if a given namespace exists in the cluster
func NamespaceExists(namespace string) (bool, error) {
	// Use in-cluster configuration
	config, err := rest.InClusterConfig()
	if err != nil {
		return false, fmt.Errorf("failed to create in-cluster config: %w", err)
	}

	// Create Kubernetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return false, fmt.Errorf("failed to create clientset: %w", err)
	}

	// Try to get the namespace
	_, err = clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err != nil {
		if k8Err.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("error getting namespace: %w", err)
	}

	return true, nil
}

func GetPodLogs(namespace, deploymentName string) (map[string]string, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get in-cluster config: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset: %w", err)
	}

	ctx := context.Background()

	// Get the deployment to retrieve its label selector
	deploy, err := clientset.AppsV1().Deployments(namespace).Get(ctx, deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get deployment: %w", err)
	}

	selector := metav1.FormatLabelSelector(&metav1.LabelSelector{
		MatchLabels: deploy.Spec.Selector.MatchLabels,
	})

	pods, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: selector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list pods: %w", err)
	}

	logs := make(map[string]string)
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			logOpts := &corev1.PodLogOptions{
				Container: container.Name,
			}

			req := clientset.CoreV1().Pods(namespace).GetLogs(pod.Name, logOpts)
			stream, err := req.Stream(ctx)
			if err != nil {
				log.Error("failed to get logs", slog.String("pod", pod.Name), slog.String("container", container.Name), slog.Any("error", err))
				continue
			}
			defer stream.Close()

			data, err := io.ReadAll(stream)
			if err != nil {
				log.Error("failed to read log stream", slog.String("pod", pod.Name), slog.String("container", container.Name), slog.Any("error", err))
				continue
			}

			key := fmt.Sprintf("%s/%s", pod.Name, container.Name)
			logs[key] = string(data)
		}
	}

	return logs, nil
}
