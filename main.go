package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	const dataIso = "2006-01-02"

	date := "2022-11-21"

	data := strings.TrimPrefix(date, " 12:00:00")
	fmt.Println(data)

	t, _ := time.Parse(dataIso, data)
	fmt.Println(t)

	upt := t.Format("2006-01-02")
	fmt.Println(upt)
}

