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
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if steps <= 0 {
		// redundant due to same check in parsePackage function
		return ""
	}
	distance := float64(steps) * stepLength / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	result := fmt.Sprintf(actionInfoFormat, steps, distance, calories)
	return result
}
