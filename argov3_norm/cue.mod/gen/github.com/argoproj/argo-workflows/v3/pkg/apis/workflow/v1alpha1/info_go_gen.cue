// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1

package v1alpha1

// A link to another app.
// +patchStrategy=merge
// +patchMergeKey=name
#Link: {
	// The name of the link, E.g. "Workflow Logs" or "Pod Logs"
	name: string @go(Name) @protobuf(1,bytes,opt)

	// "workflow", "pod", "pod-logs", "event-source-logs", "sensor-logs", "workflow-list" or "chat"
	scope: string @go(Scope) @protobuf(2,bytes,opt)

	// The URL. Can contain "${metadata.namespace}", "${metadata.name}", "${status.startedAt}", "${status.finishedAt}" or any other element in workflow yaml, e.g. "${workflow.metadata.annotations.userDefinedKey}"
	url: string @go(URL) @protobuf(3,bytes,opt)
}

// Column is a custom column that will be exposed in the Workflow List View.
// +patchStrategy=merge
// +patchMergeKey=name
#Column: {
	// The name of this column, e.g., "Workflow Completed".
	name: string @go(Name) @protobuf(1,bytes,opt)

	// The type of this column, "label" or "annotation".
	type: string @go(Type) @protobuf(2,bytes,opt)

	// The key of the label or annotation, e.g., "workflows.argoproj.io/completed".
	key: string @go(Key) @protobuf(3,bytes,opt)
}
