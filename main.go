package main

import (
	. "go-polymorphism/models"
)

func main() {
	user := NewUser("Mary", "Moonpax", 10000)
	user.GetSalary()
	alarmExample(OwnAlarm)
}
func alarmExample(alarm OwnAlarm) {
	alarm.PlayAlarm()
	alarm.SetAlarmSound("Gai")
}
