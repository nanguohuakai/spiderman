## Installation
````
go get github.com/nanguohuakai/spiderman

````

## Usage Example
### sso
````
conf := dto.AppConf{
		ServiceName: "",
		Token:       "",
	}
s,_ :=NewSpiderman(conf)

ssoConf := dto.SsoConf{
		BaseUri:     "",
		AppId:       "",
		AppKey:      "",
	}
sso, _ := s.Sso(ssoConf)
r,err := sso.Verify("")

````

### pizza
````
conf := dto.AppConf{
		ServiceName: "",
		Token:       "",
	}
s,_ :=NewSpiderman(conf)

pizzaConf := dto.SsoConf{
		BaseUri:     "",
	}
pizza, _ := s.pizza(pizzaConf)
var input dto.EmployeeListInput
r,err := pizza.GetEmployeeList(input)

````

