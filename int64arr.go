package main

import (
	"log"
	"time"
)

type int64arr []int64

func (a int64arr) Len() int {
	return len(a)
}

func (a int64arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a int64arr) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a int64arr) Pings(from string) int64arr {
	fromDate, err := time.Parse(DateLayoutISO, from)
	if err != nil {
		log.Print("Could not parse ISO date.")
		panic(err)
	}
	i := getLowerBound(a, fromDate)
	j := getUpperBoundFrom(a, fromDate)
	if i == -1 {
		a = make(int64arr, 0)
	} else {
		a = a[i:j]
	}
	return a
}

func getLowerBound(a int64arr, from time.Time) int {
	unixtime := from.Unix()
	for i := 0; i < a.Len(); i++ {
		if a[i] >= unixtime {
			return i
		}
	}
	return -1
}

func getUpperBoundFrom(a int64arr, from time.Time) int {
	from = from.AddDate(0, 0, 1).Add(-1 * time.Second)
	unixtime := from.Unix()
	for i := 0; i < a.Len(); i++ {
		if a[i] > unixtime {
			return i
		}
	}
	return a.Len()
}

func getUpperBound(a int64arr, to time.Time) int {
	unixtime := to.Unix()
	for i := 0; i < a.Len(); i++ {
		if a[i] >= unixtime {
			return i
		}
	}
	return a.Len()
}

func (a int64arr) PingsTo(from, to string) int64arr {
	fromFormat := GetDateFormat(from)
	toFormat := GetDateFormat(to)
	fromDate := ParseDate(fromFormat, from)
	toDate := ParseDate(toFormat, to)
	i := getLowerBound(a, fromDate)
	var j int
	if toFormat == FmtISO {
		j = getUpperBoundFrom(a, toDate)
	} else {
		j = getUpperBound(a, toDate)
	}
	if i == -1 {
		a = make(int64arr, 0)
	} else {
		a = a[i:j]
	}
	return a
}
