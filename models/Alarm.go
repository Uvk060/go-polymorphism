package models

import "time"

type Alarm interface {
	setTime(time time.Time)
	setAlarmSound(string2 string)
	PlayAlarm()
}
