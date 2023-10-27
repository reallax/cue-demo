package deemo

import (
    "reallx.edu/edsn/internaal/argo"
)

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

demo_tmpl: {
    edsn_inputs: {
        image: string
        command: [...string]
        args: [...string]
        resources: #Resources
    }
    edsn_contents: {
        abc: "sdf"
        argo.edsn_argo_template.edsn_contents
    }
}


argo.edsn_argo_template,edsn_inputs: {
    name: demo_tmpl.edsn_inputs.name
    container: {
        name: demo_tmpl.edsn_inputs.name
        image: demo_tmpl.edsn_inputs.image
        args: demo_tmpl.edsn_inputs.args
        command: demo_tmpl.edsn_inputs.command
        resources: {
            requests: {
                "cpu": demo_tmpl.edsn_inputs.resources.requests.cpu
                "memory": demo_tmpl.edsn_inputs.resources.requests.memory
            }
            limits: {
                "cpu": demo_tmpl.edsn_inputs.resources.limits.cpu
                "memory": demo_tmpl.edsn_inputs.resources.limits.memory
            }
        }
    }
}