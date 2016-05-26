package main

import (
    "time"
    "strconv"
    "regexp"
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

