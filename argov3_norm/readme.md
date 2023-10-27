**// FIXME: 本地package导入与引用报错，未得出期望结果**
此构建，是基于argov3的改进和优化

1. cue mod初始化
`cue mod init reallx.edu/edsn`

2. go mod 初始化
`go mod init mod.edsn`

3. 导入argo workflow依赖
```
go get github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1
cue get go github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1
```

4. 创建基础模板
