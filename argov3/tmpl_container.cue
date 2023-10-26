package argo

#Resources: {
        requests: {
                cpu:    string
                memory: string
        }
        limits: {
                cpu:    string
                memory: string
        }
}

argo_template_container: {
    edsn_inputs: {
        image: string
        command: [...string]
        args: [...string]
        resources: #Resources
    }
    edsn_contents: {
        argo_template.edsn_contents
    }
}

argo_template: edsn_inputs: argo_template.edsn_inputs & {
    name: argo_template_container.edsn_inputs.name
    container: {
        name: argo_template_container.edsn_inputs.name
        image: argo_template_container.edsn_inputs.image
        args: argo_template_container.edsn_inputs.args
        command: argo_template_container.edsn_inputs.command
        resources: {
            requests: {
                "cpu": argo_template_container.edsn_inputs.resources.requests.cpu
                "memory": argo_template_container.edsn_inputs.resources.requests.memory
            }
            limits: {
                "cpu": argo_template_container.edsn_inputs.resources.limits.cpu
                "memory": argo_template_container.edsn_inputs.resources.limits.memory
            }
        }
    }
}