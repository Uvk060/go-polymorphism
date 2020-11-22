package models

import "time"

type OwnAlarm struct {
	ID             int
	WelcomeMessage string
	playDate       time.Time
}

func (a *OwnAlarm) setTime() {
	panic("implement me")
}

func (a *OwnAlarm) SetAlarmSound(_ string) {
	panic("implement me")
}

func (a *OwnAlarm) PlayAlarm() {
	panic("implement me")
}
