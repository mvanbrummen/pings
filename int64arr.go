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
    t := fromDate.Unix()
    for i := 0; i < a.Len(); i++ {
	if a[i] >= t {
            fmt.Printf("\na[i] is greater than t %v", i)
	    fmt.Printf("\na was %v", a)
	    a = a[i:]
	    fmt.Printf("\na is now %v", a)
	    break
   	}
    }
    return a
}

func (a int64arr) PingsTo(from, to string) int64arr {
    return a
}
