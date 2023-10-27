package test

import (
	abcdef "reallx.edu/edsn/internaal/argo"
)

abcdef.edsn_argo_template,edsn_inputs: abcdef.edsn_argo_template.edsn_inputs & {
	name: "argo_template_full_demo",
	inputs: {
		parameters: [
			{
				name:  "input_file"
				value: "/Users/reallx/GoProjects/cue-demo/input_file"
			},
			{
				name: "traffic_group"
				enum: [ "blue", "green", "gray"]
				default: "blue"
			},
		]
	},
	outputs: {
		parameters: [
			{
				name:  "output_file"
				value: "/Users/reallx/GoProjects/cue-demo/output_file"
			},
		]
	},
	script: {
		name:  "gen-random-int"
		image: "python:alpine3.6"
		command: [ "python"]
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
		source: """
			import random
			i = random.randint(1, 100)
			print(i)
			"""
	}
}
