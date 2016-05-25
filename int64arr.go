package main

import (
    "fmt"
    "time"
    "regexp"
    "strconv"
)

const (
    DateLayoutISO string = "2006-01-02"
    FmtISO = "^\\d{4}-\\d{2}-\\d{2}$"
    FmtUnix = "^[0-9]\\d*$"
)

func IsDateFormat(format, date string) bool {
    b, err := regexp.MatchString(format, date)
    if err != nil {
        return false
    }   
    return b
} 

func GetDateFormat(date string) string {
    var format string
    if IsDateFormat(FmtISO, date) {
 	format = FmtISO
    } else if IsDateFormat(FmtUnix, date) {
  	format = FmtUnix
    } else {
        format = ""
    }
    return format
}

func ParseDate(format, dateStr string) time.Time {
    if format == FmtISO {
	date, _ := time.Parse(DateLayoutISO, dateStr)
	return date
    } else if format == FmtUnix {
        unixts, _ := strconv.ParseInt(dateStr, 0, 64) 
  	date := time.Unix(unixts, 0)  
    	return date
    } else {
	return time.Time{}
    }
}

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
 	fmt.Printf("Could not parse ISO date.")
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
