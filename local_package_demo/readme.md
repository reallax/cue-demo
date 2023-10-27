**此目录测试：本地package相互导入**
* 定义了cue module：example.com
* 定义package y，路径：local_package_demo/b/b.cue
* 定义package z，路径：local_package_demo/c/c.cue
* 将packge y导入c.cue，并使用：`import b "example.com/b:y"`
* 测试命令：`cd c; cue eval c.cue`