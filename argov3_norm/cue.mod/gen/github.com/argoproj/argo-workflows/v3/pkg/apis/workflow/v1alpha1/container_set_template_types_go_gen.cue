// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

#ContainerSetTemplate: {
	containers: [...#ContainerNode] @go(Containers,[]ContainerNode) @protobuf(4,bytes,rep)
	volumeMounts?: [...corev1.#VolumeMount] @go(VolumeMounts,[]corev1.VolumeMount) @protobuf(3,bytes,rep)

	// RetryStrategy describes how to retry a container nodes in the container set if it fails.
	// Nbr of retries(default 0) and sleep duration between retries(default 0s, instant retry) can be set.
	retryStrategy?: null | #ContainerSetRetryStrategy @go(RetryStrategy,*ContainerSetRetryStrategy) @protobuf(5,bytes,opt)
}

#ContainerSetRetryStrategy: {
	// Duration is the time between each retry, examples values are "300ms", "1s" or "5m".
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	duration?: string @go(Duration) @protobuf(1,bytes,opt)

	// Nbr of retries
	retries?: null | intstr.#IntOrString @go(Retries,*intstr.IntOrString) @protobuf(2,bytes,rep)
}

#ContainerNode: {
	corev1.#Container
	dependencies?: [...string] @go(Dependencies,[]string) @protobuf(2,bytes,rep)
}
