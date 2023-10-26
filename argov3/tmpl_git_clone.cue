tmpl_git_clone: {
	edsn_inputs: {}
	edsn_contents: {
		{
			"name": "git-clone"
			"inputs": {
				"parameters": [
					{
						"name": "repo"
					},
					{
						"name": "branch"
					},
					{
						"name": "workdir"
					},
				]
			}
			"outputs": {
				"artifacts": [
					{
						"name": "repo"
						"path": "{{inputs.parameters.workdir}}"
					},
				]
			}
			"container": {
				"volumeMounts": [
					{
						"mountPath": "{{inputs.parameters.workdir}}"
					},
				]
				"name":       "workdir"
				"image":      "alpine/git:v2.26.2"
				"workingDir": "/work"
				"args": [
					"clone",
					"--depth",
					"1",
					"--branch",
					"{{inputs.parameters.branch}}",
					"--single-branch",
					"{{inputs.parameters.repo}}",
				]
			}
		}
	}
}
