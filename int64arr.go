package main

import (
    "fmt"
    "time"
)

const DateLayoutISO string = "2006-01-02"

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
    j := getUpperBound(a, fromDate)
    if i == -1 {
  	a = nil
    } else {
        fmt.Printf("\na was %v", a)
        a = a[i:j]
        fmt.Printf("\na is now %v", a)
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

func getUpperBound(a int64arr, from time.Time) int {
    from = from.AddDate(0, 0, 1)
    unixtime := from.Unix()
    for i := 0; i < a.Len(); i++ {
	if a[i] > unixtime {
	    return i
 	} 
    }
    return a.Len()
}
func (a int64arr) PingsTo(from, to string) int64arr {
    return a
}
