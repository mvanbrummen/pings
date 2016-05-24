package main

import (
    "testing"
    "reflect"
) 

type testpair struct {
    date string  
    expectation int64arr    
}

var tests = []testpair {
  {"2016-05-23", int64arr{1464015600}},
  {"2016-05-24", int64arr{1464083339, 1464083341}},
  {"2016-05-25", int64arr{1464188400}},
}

var a = int64arr {
    1464015600,
    1464083339,
    1464083341,
    1464188400,
}

func TestInt64arr(t *testing.T) {
    for _, pair := range tests {
        v := a.Pings(pair.date)
        if !reflect.DeepEqual(pair.expectation, v) {
    	    t.Error(
		"For", pair.date,
		"expected", pair.expectation,
		"got", v,
 	    )
        }
    } 
}
