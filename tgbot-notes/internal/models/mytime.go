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
	return fmt.Sprintf("%d %s %d %d:%d", year, getRUMonth(month), day, hour, minutes)
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

	monthNew := getIDMonth(month)
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

func getIDMonth(month string) time.Month {
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

func getRUMonth(month time.Month) string {
	switch month {
	case time.December:
		return "Декабря"
	case time.January:
		return "Января"
	case time.February:
		return "Февраля"
	case time.March:
		return "Марта"
	case time.April:
		return "Апреля"
	case time.May:
		return "Мая"
	case time.June:
		return "Июня"
	case time.July:
		return "Июля"
	case time.August:
		return "Августа"
	case time.September:
		return "Сентября"
	case time.October:
		return "Октября"
	case time.November:
		return "Ноября"
	default:
		return month.String()
	}
}
