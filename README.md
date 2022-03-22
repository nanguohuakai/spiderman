## 测试用例执行命令，依赖文件需要手动加载 实例：
````
方法1
go.mod文件添加
require (
	spiderman v0.0.1
)

replace (
	spiderman v0.0.1 => github.com/nanguohuakai/spiderman v0.0.1
)

方法2
go test -v spiderman_test.go spiderman.go
````

## 调用方式
### 1、pizza接口调用， 示例：
````
svc := NewSpiderman(conf).Pizza()
res := svc.GetEmployeeInfo()

````