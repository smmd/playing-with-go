package main

import (
	"fmt"
	"time"
)

const (
	anglePerHour   = 360 / 12
	anglePerMinute = 360 / 60
)

func main() {
	var hour int
	var minute int

	/*fmt.Println("Enter the hour:")
	hour, err := fmt.Scanln(hour)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter the minute:")
	minute, err = fmt.Scanln(minute)
	if err != nil {
		panic(err)
	}*/
	currentTime := time.Now()

	hour = currentTime.Hour() % 12
	minute = currentTime.Minute()

	result, err := clockAngle(hour, minute)
	if err != nil {
		panic(err)
	}

	fmt.Printf("For Hour: %d, Min: %d the Angle is: %d\n", hour, minute, result)
}

func clockAngle(hour, minute int) (int, error) {
	if hour > 12 {
		return 0, fmt.Errorf("hour %d invalid", hour)
	}

	if minute > 60 {
		return 0, fmt.Errorf("minute %d invalid", minute)
	}

	if hour == 12 {
		hour = 0
	}

	hourAngle := hour * anglePerHour
	minuteAngle := minute * anglePerMinute

	if minuteAngle == 0 {
		return 360 - hourAngle, nil
	}

	if hourAngle > minuteAngle {
		return hourAngle - minuteAngle, nil
	}

	return minuteAngle - hourAngle, nil
}
