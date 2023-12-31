// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1

package v1alpha1

// the workflow's phase
#WorkflowPhase: string // #enumWorkflowPhase

#enumWorkflowPhase:
	#WorkflowUnknown |
	#WorkflowPending |
	#WorkflowRunning |
	#WorkflowSucceeded |
	#WorkflowFailed |
	#WorkflowError

#WorkflowUnknown:   #WorkflowPhase & ""
#WorkflowPending:   #WorkflowPhase & "Pending"
#WorkflowRunning:   #WorkflowPhase & "Running"
#WorkflowSucceeded: #WorkflowPhase & "Succeeded"
#WorkflowFailed:    #WorkflowPhase & "Failed"
#WorkflowError:     #WorkflowPhase & "Error"
