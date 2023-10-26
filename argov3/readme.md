## 一、引入argo workflow包
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

## 二、依赖argo导入测试Demo
`template.cue`是一个通用的模板，直接引用argo workflow的定义。  
但也有**缺点**，就是将全量参数暴露给用户，增加了配置的复杂度。   
cue文件测试命令：  
`cue eval data.cue template.cue -e argo_template.edsn_contents --out yaml`

## 三、基于通用模板选择性暴露参数
* 构建tmpl_container.cue，引入基础模板template.cue#edsn_contents
* 在tmpl_container.cue中定义希望暴露给用户的参数，并定义template.cue#edsn_inputs，其值来源于tmpl_container.cue#edsn_inputs
* 构建tmpl_container_data.cue，传入tmpl_container.cue#edsn_inputs参数值。  
cue文件测试命令：  
```
cue eval tmpl_container.cue tmpl_container_data.cue template.cue -e argo_template_container.edsn_contents --out yaml
```
*注意*：多文件渲染，多文件需要有共同的package名称。如：package argo
