## Installation
````
go get github.com/nanguohuakai/spiderman

````

## Usage Example
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