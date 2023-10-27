###测试文件
**测试：cue可以在相同module中，导入自定义的package**
* data.cue 导入自定义的package argo。相对路径（相对于module根，而不是相对于本文件）：./internal/argo/template.cue
* data.cue导入：`import abcdef "reallx.edu/edsn/internal/argo"`
    *  相对路径导入格式：`<module identifier>/<relative position of package within module>:<package name>`。[参考](https://cuelang.org/docs/concepts/packages/#import-path)
    * abcdef 是包别名。像golang一样，如果没有冲突，可以省略
    * reallx.edu/edsn 是cue module名称。在argov3_norm/cue.mod/module.cue中查看。
    * internal/argo 是相对于module root的路径
    * 本来应该是`import abcdef "reallx.edu/edsn/internal/argo:argo"`，但因为argo目录名与package名相同，所以`:argo`被省略
* 注意：`abcdef.edsn_argo_template, `(第七行)，后面跟的是逗号，换称分号就报错。不知道为什么，目前没找到文档说明。
* 执行测试命令：`cue eval ./ -e edsn_contents --out yaml`。注意：是`-e edsn_contents`，而不是`-e abcdef.edsn_argo_template.edsn_contents`