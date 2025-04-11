package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	list := strings.Split(data, ",")
	if len(list) != 2 {
		return 0, time.Duration(0), fmt.Errorf("Bad format: only one comma allowed")
	}
	steps, err := strconv.Atoi(list[0])
	if err != nil {
		return 0, time.Duration(0), fmt.Errorf("Bad format: could not parse int before the comma")
	}
	if steps <= 0 {
		return 0, time.Duration(0), fmt.Errorf("Bad format: step count must be positive")
	}
	duration, err := time.ParseDuration(list[1])
	if err != nil {
		return 0, time.Duration(0), fmt.Errorf("Bad format: could not parse duration")
	}
	if duration <= 0 {
		return 0, time.Duration(0), fmt.Errorf("Bad format: duration must be positive")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
}
