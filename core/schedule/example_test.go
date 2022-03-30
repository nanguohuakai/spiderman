package schedule

import (
	spiderman2 "github.com/nanguohuakai/spiderman"
	"github.com/nanguohuakai/spiderman/dto"
)

func ExampleClient_RegisterCorn() {
	conf := dto.AppConf{
		ServiceName: "test",
		Token:       "example",
	}
	spiderman, _ := spiderman2.NewSpiderman(conf)

	schedule,_  := spiderman.Schedule(dto.ScheduleConf{
		BaseUri:     "http://127.0.0.1:8081",
		CallbackUri: "http://127.0.0.1:8080",
	})
	_ = schedule.RegisterCorn(dto.ScheduleRegisterInput{
		ScheduleId:     "test-talent",
		ScheduleName:   "ttttt",
		CallbackParams:  nil,
		Value:          "*/1 * * * *",
	})
}


func ExampleClient_RegisterCornTime() {
	conf := dto.AppConf{
		ServiceName: "test",
		Token:       "example",
	}
	spiderman, _ := spiderman2.NewSpiderman(conf)

	schedule,_  := spiderman.Schedule(dto.ScheduleConf{
		BaseUri:     "http://127.0.0.1:8081",
		CallbackUri: "http://127.0.0.1:8080",
	})
	_ = schedule.RegisterCornTime(dto.ScheduleRegisterInput{
		ScheduleId:     "test-talent",
		ScheduleName:   "ttttt",
		CallbackParams:  nil,
		Value:          "2022-01-01 12:01",
	})
}