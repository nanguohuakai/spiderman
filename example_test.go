package spiderman

import (
	"fmt"
	"github.com/nanguohuakai/spiderman/dto"
)

//NewSpiderman
func ExampleNewSpiderman() {
	//app conf
	conf := dto.AppConf{
		ServiceName: "test",
		Token:       "example",
	}
	s, _ := NewSpiderman(conf)
	fmt.Println(s)
}

func ExampleSpidermanClient_Alert() {

	conf := dto.AppConf{
		ServiceName: "test",
		Token:       "example",
	}
	spiderman, _ := NewSpiderman(conf)

	alertConf := dto.AlertConf{
		BaseUri: "127.0.0.1",
		Level:   "warn",
		Env:     "test",
	}

	alert, _ := spiderman.Alert(alertConf)
	fmt.Println(alert)
}

func ExampleSpidermanClient_Pizza() {
	conf := dto.AppConf{
		ServiceName: "test",
		Token:       "example",
	}
	spiderman, _ := NewSpiderman(conf)

	pizzaConf := dto.PizzaConf{
		BaseUri: "127.0.0.1",
	}

	pizza, _ := spiderman.Pizza(pizzaConf)
	fmt.Println(pizza)
}

func ExampleSpidermanClient_Sso() {
	conf := dto.AppConf{
		ServiceName: "test",
		Token:       "example",
	}
	spiderman, _ := NewSpiderman(conf)
	ssoConf := dto.SsoConf{
		BaseUri: "127.0.0.1",
		AppId:   "test",
		AppKey:  "test1",
	}
	sso, _ := spiderman.Sso(ssoConf)
	fmt.Println(sso)
}

func ExampleSpidermanClient_Schedule() {
	conf := dto.AppConf{
		ServiceName: "test",
		Token:       "example",
	}
	spiderman, _ := NewSpiderman(conf)

	schedule,_  := spiderman.Schedule(dto.ScheduleConf{
		BaseUri:     "http://127.0.0.1:8081",
		CallbackUri: "http://127.0.0.1:8080",
	})
	fmt.Println(schedule)
}
