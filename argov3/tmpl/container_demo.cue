import (
    argo "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	apiv1 "k8s.io/api/core/v1"
)

template: argo.#Template & {
    name: "name-abc"
    container: {
        name: "name-abc"
        image: "docker/whalesay"
        args: [ "hello world"]
        command: [ "cowsay"]
        resources: apiv1.#ResourceRequirements & {
            requests: {
                cpu:    "250m"
                memory: "64Mi"
            }
            limits: {
                cpu:    "500m"
                memory: "128Mi"
            }
        }
    }
}