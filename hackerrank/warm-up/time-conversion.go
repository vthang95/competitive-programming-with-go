package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var time string
	var list []string
	fmt.Scan(&time)
	if time == "12:00:00AM" {
		fmt.Println("00:00:00")
	} else if time == "12:00:00PM" {
		fmt.Println("12:00:00")
	} else {
		timePostfix := string(time[(len(time) - 2):(len(time))])
		list = strings.Split(time, ":")
		if timePostfix == "PM" && list[0] != "12" {
			num, _ := strconv.Atoi(list[0])
			list[0] = strconv.Itoa(num + 12)
		}
		if timePostfix == "AM" && list[0] == "12" {
			list[0] = "00"
		}
		time = strings.Join(list, ":")
		fmt.Println(string(time[0:(len(time) - 2)]))
	}
}
