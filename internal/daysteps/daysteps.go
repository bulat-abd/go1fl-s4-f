package daysteps

import (
	"fmt"
	"log"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
	actionInfoFormat = "Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n"
)

func parsePackage(data string) (int, time.Duration, error) {
	fields := strings.Split(data, ",")
	if len(fields) != 2 {
		return 0, 0, fmt.Errorf("bad format: exactly one comma must be present, separating exactly two fields")
	}
	steps, err := strconv.Atoi(fields[0])
	if err != nil {
		return 0, 0, fmt.Errorf("bad format: could not parse int in the first field - %w", err)
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("bad data: step count must be positive")
	}
	duration, err := time.ParseDuration(fields[1])
	if err != nil {
		return 0, 0, fmt.Errorf("bad format: could not parse duration - %w", err)
	}
	if duration <= 0 {
		return 0, 0, fmt.Errorf("bad data: duration must be positive")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	distance := float64(steps) * stepLength / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	result := fmt.Sprintf(actionInfoFormat, steps, distance, calories)
	return result
}
