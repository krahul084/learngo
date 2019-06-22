package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		speed int
		now   time.Time
	)
	speed, now = 100, time.Now()
	fmt.Println(speed, now)
	swap()
}

func swap() {
	var (
		speed     int
		prevSpeed int
	)
	speed = 100
	prevSpeed = speed
	speed = 50
	speed, prevSpeed = prevSpeed, speed
	fmt.Println(speed, prevSpeed)
}
