package container

import (
    argo "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	apiv1 "k8s.io/api/core/v1"
)

tmpl_container: {
    edsn_inputs: {
        name: string
        image: string
        command: [...string]
        args: [...string]
        resources: apiv1.#ResourceRequirements
    }
    edsn_contents: argo.#Template & {
        name: edsn_inputs.name
        container: {
            name: edsn_inputs.name
            image: edsn_inputs.image
            args: edsn_inputs.args
            command: edsn_inputs.command
            resources: edsn_inputs.resources
        }
    }
}