**./tmpl文件夹：用于测试对argo.#Template内容参数的二次过滤**
* 通过在模板中导入argo依赖，我们可以非常方便的生成argo template，甚至可以生成完整的workflow yaml。
但缺点是，把argo template的参数全部暴露给用户，导致用户使用门槛增加，易用性低下
* 通过定义类似container.cue文件，在edsn_inputs中【自定义】暴露部分参数，可以减少用户使用模板时必要输入参数，提升易用性
    * 测试命令：`cue eval container.cue data.cue -e tmpl_container.edsn_contents --out yaml`
* 为了生成一个argo template，最简单的方式，直接编辑数据即可。参考文件：container_demo.cue。但此方法不能完成参数的二次过滤。
    * 测试命令：`cue eval container_demo.cue -e template --out yaml`