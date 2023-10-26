package argo

import (
	argo "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	apiv1 "k8s.io/api/core/v1"
)

argo_template: {
	edsn_inputs: {
		name?:    string
		inputs?:  argo.#Inputs
		outputs?: argo.#Outputs
		nodeSelector?: [string]: string
		affinity?: apiv1.#Affinity
		metadata?: argo.#Metadata
		daemon?:   bool
		steps?: [...argo.#ParallelSteps]
		container?: apiv1.#Container
		containerSet?: [...apiv1.#ContainerSetTemplate]
		script?:   argo.#ScriptTemplate
		resource?: argo.#ResourceTemplate
		dag?:      argo.#DAGTemplate
		suspend?:  argo.#SuspendTemplate
		data?:     argo.#Data
		http?:     argo.#HTTP
		plugin?:   argo.#Plugin
		volumes?: [...apiv1.#Volume]
		initContainers?: argo.#UserContainer
		sidecars?:       argo.#UserContainer
		timeout?:        string
	}
	edsn_contents: {
		if edsn_inputs.name != _|_ {
			"name": edsn_inputs.name
		}
		if edsn_inputs.inputs != _|_ {
			"inputs": edsn_inputs.inputs
		}
		if edsn_inputs.outputs != _|_ {
			"outputs": edsn_inputs.outputs
		}
		if edsn_inputs.nodeSelector != _|_ {
			"nodeSelector": edsn_inputs.nodeSelector
		}
		if edsn_inputs.affinity != _|_ {
			"affinity": edsn_inputs.affinity
		}
		if edsn_inputs.metadata != _|_ {
			"metadata": edsn_inputs.metadata
		}
		if edsn_inputs.daemon != _|_ {
			"daemon": edsn_inputs.daemon
		}
		if edsn_inputs.steps != _|_ {
			"steps": edsn_inputs.steps
		}
		if edsn_inputs.container != _|_ {
			"container": edsn_inputs.container
		}
		if edsn_inputs.containerSet != _|_ {
			"containerSet": edsn_inputs.containerSet
		}
		if edsn_inputs.script != _|_ {
			"script": edsn_inputs.script
		}
		if edsn_inputs.resource != _|_ {
			"resource": edsn_inputs.resource
		}
		if edsn_inputs.dag != _|_ {
			"dag": edsn_inputs.dag
		}
		if edsn_inputs.suspend != _|_ {
			"suspend": edsn_inputs.suspend
		}
		if edsn_inputs.data != _|_ {
			"data": edsn_inputs.data
		}
		if edsn_inputs.http != _|_ {
			"http": edsn_inputs.http
		}
		if edsn_inputs.plugin != _|_ {
			"plugin": edsn_inputs.plugin
		}
		if edsn_inputs.volumes != _|_ {
			"volumes": edsn_inputs.volumes
		}
		if edsn_inputs.initContainers != _|_ {
			"initContainers": edsn_inputs.initContainers
		}
		if edsn_inputs.sidecars != _|_ {
			"sidecars": edsn_inputs.sidecars
		}
		if edsn_inputs.timeout != _|_ {
			"timeout": edsn_inputs.timeout
		}
	}
}
