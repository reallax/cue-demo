argo_template_container: edsn_inputs: argo_template_container.edsn_inputs & {
    name: "name-abc"
	image: "docker/whalesay"
	command: [ "cowsay"]
	args: [ "hello world"]
	resources: {
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
