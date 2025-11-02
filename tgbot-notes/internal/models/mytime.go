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

func (t *MyTime) Set(ctx context.Context, currentTime string) error {
	var (
		err     error
		newTime time.Time
	)

	if ctx.Value("set_note") != nil {
		temp := parseTime(currentTime)
		newTime, err = time.Parse("_2 January 2006 15:04:05", temp)
		t.t = newTime
	} else {
		newTime, err = time.Parse(time.RFC3339, currentTime)
		t.t = newTime
	}

	return err
}

func (t *MyTime) String() string {
	return t.t.Format("_2 January 2006 15:04:05")
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
	if monthNew > time.Now().Month() {
		newTime := fmt.Sprintf("%s %s %d %s:%s:%s", day, monthNew.String(), time.Now().Year(), hour, minute, second)
		return newTime
	}
	if monthNew == time.Now().Month() {
		tmp, _ := strconv.Atoi(day)
		if tmp > time.Now().Day() {
			newTime := fmt.Sprintf("%s %s %d %s:%s:%s", day, monthNew.String(), time.Now().Year(), hour, minute, second)
			return newTime
		} else {
			newTime := fmt.Sprintf("%s %s %d %s:%s:%s", day, monthNew.String(), time.Now().Year()+1, hour, minute, second)
			return newTime
		}
	}
	if monthNew < time.Now().Month() {
		newTime := fmt.Sprintf("%s %s %d %s:%s:%s", day, monthNew.String(), time.Now().Year()+1, hour, minute, second)
		return newTime
	}

	return ""
}

func getMonth(month string) time.Month {
	switch month {
	case "Январь":
		return time.January
	case "Января":
		return time.January
	case "Февраль":
		return time.February
	case "Февраля":
		return time.February
	case "Март":
		return time.March
	case "Марта":
		return time.March
	case "Апрель":
		return time.April
	case "Апреля":
		return time.April
	case "Май":
		return time.May
	case "Мая":
		return time.May
	case "Июнь":
		return time.June
	case "Июня":
		return time.June
	case "Июль":
		return time.July
	case "Июля":
		return time.July
	case "Август":
		return time.August
	case "Августа":
		return time.August
	case "Сентябрь":
		return time.September
	case "Сентября":
		return time.September
	case "Октябрь":
		return time.October
	case "Октября":
		return time.October
	case "Ноябрь":
		return time.November
	case "Ноября":
		return time.November
	case "Декабрь":
		return time.December
	case "Декабря":
		return time.December
	default:
		return time.Now().Month()
	}

}
