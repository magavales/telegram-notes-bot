package models

import (
	"fmt"
	"strings"
	"time"
)

type MyTime struct {
	t time.Time
}

func NewMyTime() *MyTime {
	return &MyTime{}
}

func (t *MyTime) Get() time.Time {
	return t.t
}

func (t *MyTime) Set(currentTime string) error {
	temp, err := time.Parse("January _2 15:04:05", parseTime(currentTime))
	t.t = temp

	return err
}

func parseTime(currentTime string) string {
	strs := strings.Split(currentTime, " ")
	month := strs[0]
	day := strs[1]
	cloks := strings.Split(strs[2], ":")
	hour := cloks[0]
	minute := cloks[1]
	second := ""
	if len(cloks) > 2 {
		second = cloks[2]
	} else {
		second = "00"
	}

	month = getMonth(month)

	newTime := fmt.Sprintf("%s %s %s:%s:%s", month, day, hour, minute, second)

	return newTime
}

func getMonth(month string) string {
	switch month {
	case "Январь":
		return time.January.String()
	case "Января":
		return time.January.String()
	case "Февраль":
		return time.February.String()
	case "Февраля":
		return time.February.String()
	case "Март":
		return time.March.String()
	case "Марта":
		return time.March.String()
	case "Апрель":
		return time.April.String()
	case "Апреля":
		return time.April.String()
	case "Май":
		return time.May.String()
	case "Мая":
		return time.May.String()
	case "Июнь":
		return time.June.String()
	case "Июня":
		return time.June.String()
	case "Июль":
		return time.July.String()
	case "Июля":
		return time.July.String()
	case "Август":
		return time.August.String()
	case "Августа":
		return time.August.String()
	case "Сентябрь":
		return time.September.String()
	case "Сентября":
		return time.September.String()
	case "Октябрь":
		return time.October.String()
	case "Октября":
		return time.October.String()
	case "Ноябрь":
		return time.November.String()
	case "Ноября":
		return time.November.String()
	case "Декабрь":
		return time.December.String()
	case "Декабря":
		return time.December.String()
	default:
		return ""
	}

}
