### 1. 创建工作目录
```
mkdir argov3
cd argov3
```

### 2. cue mod初始化
`cue mod init`
作用：初始化cue.mod，以及导入关系路径

### 3. go mod初始化
`go mod init mod.test`
作用：可以用go命令拉取远端go代码，然后被cue转换结构定义

### 4. 导入argo包 [包地址](https://pkg.go.dev/github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1)
```
go get github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1
cue get go github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1
```
作用：
- go get：定义go的依赖；
- cue get：将go lib中的的struct生成Definition
- cue文件生成路径：argov3/cue.mod/gen/github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1/

### 5. cue文件测试
`cue eval {data.cue} {tmpl.cue} -e {wf_container}.edsn_contents --out yaml`
