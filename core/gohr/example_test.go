package gohr

import (
	"fmt"
	spiderman2 "github.com/nanguohuakai/spiderman"
	"github.com/nanguohuakai/spiderman/dto"
)

func ExampleClient_GetRecordJournal() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	var input = dto.RecordJournalUserInput{
		AppID:          "1233455",
		ViewerWorkcode: "076533",
		TargetWorkcode: "076533",
		Type:           "lesson",
		Size:           10,
	}
	s, _ := p.GetRecordJournal(input)
	fmt.Println(s)
}


func ExampleClient_GetTalentOkrUserList() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)
	s, _ := p.GetTalentOkrUserList("123455","0755555")
	fmt.Println(s)
}

func ExampleClient_GetHonorUserInfo() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	var input = dto.HonorUserInfoInput{
		AppID:    "12212",
		Workcode: "033333",
		GroupId:  "1111",
	}

	s, _ := p.GetHonorUserInfo(input)
	fmt.Println(s)
}

func ExampleClient_GetYachWeekReport() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	var input = dto.WeekReportUserInput{
		AppID:          "1111",
		ViewerWorkcode: "076533",
		TargetWorkcode: "076533",
		Size:           0,
	}

	s, _ := p.GetYachWeekReport(input)
	fmt.Println(s)
}

func ExampleClient_GetReportListFrom360() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	s, _ := p.GetReportListFrom360("076533", 1, 10)
	fmt.Println(s)
}

func ExampleClient_GetJobChannel() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	s, _ := p.GetJobChannel(dto.JobChannelInput{})
	fmt.Println(s)
}

func ExampleClient_GetJobc() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	s, _ := p.GetJobc(dto.JobcInput{})
	fmt.Println(s)
}

func ExampleClient_GetJobSeq() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	s, _ := p.GetJobSeq(dto.JobSeqInput{})
	fmt.Println(s)
}

func ExampleClient_GetJobSub() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	s, _ := p.GetJobSub(dto.JobSubInput{})
	fmt.Println(s)
}

func ExampleClient_GetQuitList() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	s, _ := p.GetQuitList(dto.QuitOrTalInput{})
	fmt.Println(s)
}

func ExampleClient_GetFullToTal() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	s, _ := p.GetFullToTal(dto.QuitOrTalInput{})
	fmt.Println(s)
}

func ExampleClient_ScheduleCreate() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	s, _ := p.ScheduleCreate(dto.ScheduleCreateOrUpdateInput{})
	fmt.Println(s)
}

func ExampleClient_ScheduleUpdate() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	s, _ := p.ScheduleUpdate(dto.ScheduleCreateOrUpdateInput{})
	fmt.Println(s)
}

func ExampleClient_ScheduleCancel() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	s, _ := p.ScheduleCancel(dto.ScheduleCancelInput{})
	fmt.Println(s)
}

func ExampleClient_ScheduleInvite() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	s, _ := p.ScheduleInvite(dto.ScheduleInviteInput{})
	fmt.Println(s)
}

func ExampleClient_ScheduleList24() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	s, _ := p.ScheduleList24(dto.ScheduleLisInput{})
	fmt.Println(s)
}

func ExampleClient_GetSecretDeptByUser() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	s, _ := p.GetSecretDeptByUser(dto.SecretDeptInput{})
	fmt.Println(s)
}

func ExampleClient_RecentSecret() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	s, _ := p.RecentSecret(dto.SecretInput{})
	fmt.Println(s)
}

func ExampleClient_BeforeSecret() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	s, _ := p.BeforeSecret(dto.SecretInput{})
	fmt.Println(s)
}

func ExampleClient_DecodeSecret() {
	conf := dto.AppConf{
		ServiceName: "TEST",
		Token:       "TEST",
	}
	spiderman ,_ :=spiderman2.NewSpiderman(conf)

	gohrConf := dto.GoHrConf{
		BaseUri:     "http://127.0.0.1",

	}
	p,_ := spiderman.GoHr(gohrConf)

	var input = []map[string]string{}
	s, _ := p.DecodeSecret(input)
	fmt.Println(s)
}