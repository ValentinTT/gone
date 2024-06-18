package cmd

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type Date struct {
	day   string
	month string
	year  string
	date  time.Time
}

func (date *Date) parseToday() {
	date.date = time.Now()
	date.day = fmt.Sprintf("%0*d", 2, date.date.Day())
	date.month = fmt.Sprintf("%0*d", 2, int(date.date.Month()))
	date.year = strconv.Itoa(date.date.Year())
}

func (date *Date) parseDate(arg, layout_flag string) error {
	params, err := getParams(arg)
	if err != nil {
		return err
	}

	di, mi, yi, err := getDateElementsIndex(layout_flag)
	if err != nil {
		return err
	}

	date.day = fmt.Sprintf("%0*v", 2, params[di])
	date.month = fmt.Sprintf("%0*v", 2, params[mi])

	// input in this format: yy-(dd o mm) month or day missing
	if len(params) == 2 && yi != 2 {
		return errors.New("date doesn't follow the layout")
	}
	currentYear := strconv.Itoa(time.Now().Year())
	date.year = currentYear
	if len(params) == 3 {
		date.year = currentYear[:4-len(params[yi])] + params[yi]
	}

	dateS := date.day + "-" + date.month + "-" + date.year
	layout := layouts["dmy"]
	time, err := time.Parse(layout, dateS)
	if err != nil {
		return err
	}
	date.date = time
	return nil
}

func getParams(arg string) ([]string, error) {
	re := regexp.MustCompile(`[ .\-/]`)
	params := re.Split(arg, -1)

	if l := len(params); l < 2 || l > 3 { // it has more or less numbers than d, m, y
		return nil, errors.New("date with wrong format")
	}

	return params, nil
}

func getDateElementsIndex(layout_flag string) (day, month, year int, err error) {
	for i, v := range layout_flag {
		switch v {
		case 'd':
			day = i
		case 'm':
			month = i
		case 'y':
			year = i
		}
	}
	return
}
