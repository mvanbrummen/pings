package main

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
