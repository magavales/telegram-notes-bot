package models

import (
	"context"
	"fmt"
	"strconv"
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

// !!!
func (t *MyTime) Set(ctx context.Context, currentTime string) error {
	var (
		err     error
		newTime time.Time
	)
	newTime, err = time.Parse(time.RFC3339, currentTime)
	if err != nil {
		temp := parseTime(currentTime)
		newTime, err = time.Parse("_2 January 2006 15:04:05", temp)
		if err != nil {
			return err
		}
	}

	t.t = newTime

	return err
}

func (t *MyTime) String() string {
	year, month, day := t.t.Date()
	hour, minutes, _ := t.t.Clock()
	return fmt.Sprintf("%d %s %d %d:%d", year, getMonth(month.String()), day, hour, minutes)
}

func parseTime(currentTime string) string {
	strs := strings.Split(currentTime, " ")
	day := strs[0]
	month := strs[1]
	cloks := strings.Split(strs[2], ":")
	hour := cloks[0]
	minute := cloks[1]
	second := ""
	if len(cloks) > 2 {
		second = cloks[2]
	} else {
		second = "00"
	}

	monthNew := getMonth(month)
	if monthNew > time.Now().Month().String() {
		newTime := fmt.Sprintf("%s %s %d %s:%s:%s", day, monthNew, time.Now().Year(), hour, minute, second)
		return newTime
	}
	if monthNew == time.Now().Month().String() {
		tmp, _ := strconv.Atoi(day)
		if tmp > time.Now().Day() {
			newTime := fmt.Sprintf("%s %s %d %s:%s:%s", day, monthNew, time.Now().Year(), hour, minute, second)
			return newTime
		} else {
			newTime := fmt.Sprintf("%s %s %d %s:%s:%s", day, monthNew, time.Now().Year()+1, hour, minute, second)
			return newTime
		}
	}
	if monthNew < time.Now().Month().String() {
		newTime := fmt.Sprintf("%s %s %d %s:%s:%s", day, monthNew, time.Now().Year()+1, hour, minute, second)
		return newTime
	}

	return ""
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
	case time.December.String():
		return "Декабря"
	case time.January.String():
		return "Января"
	case time.February.String():
		return "Февраля"
	case time.March.String():
		return "Марта"
	case time.April.String():
		return "Апреля"
	case time.May.String():
		return "Мая"
	case time.June.String():
		return "Июня"
	case time.July.String():
		return "Июля"
	case time.August.String():
		return "Августа"
	case time.September.String():
		return "Сентября"
	case time.October.String():
		return "Октября"
	case time.November.String():
		return "Ноября"
	default:
		return time.Now().Month().String()
	}

}
