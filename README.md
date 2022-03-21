## 测试用例执行命令，依赖文件需要手动加载 实例：
````
go test -v spiderman_test.go spiderman.go
````

## 调用方式
### 1、pizza接口调用， 示例：
````
svc := NewSpiderman(conf).Pizza()
res := svc.GetEmployeeInfo()

````