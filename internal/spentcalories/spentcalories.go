package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
	trainingInfoFormat         = "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n"
)

func parseTraining(data string) (int, string, time.Duration, error) {
	fields := strings.Split(data, ",")
	if len(fields) != 3 {
		return 0, "", 0, fmt.Errorf("bad format: there must be exactly two commas separating three fields")
	}
	activity := fields[1]
	steps, err := strconv.Atoi(fields[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("bad format: could not parse int in first field")
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("bad data: step count must be positive")
	}
	duration, err := time.ParseDuration(fields[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("bad format: could not parse duration in third field")
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("bad data: duration must be positive")
	}
	return steps, activity, duration, nil

}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}
	calories := 0.0
	switch activity {
		case "Бег":
			calories, err = RunningSpentCalories(steps, weight, height, duration)
		case "Ходьба":
			calories, err = WalkingSpentCalories(steps, weight, height, duration)
		default:
			return "", fmt.Errorf("неизвестный тип тренировки")
	}
	result := fmt.Sprintf(trainingInfoFormat, activity, duration.Hours(), distance(steps, height), meanSpeed(steps, height, duration), calories)
	return result, nil
}

func caloriesCalculationHelper(steps int, weight, height float64, duration time.Duration) float64 {
	return weight * meanSpeed(steps, height, duration) * duration.Hours()
}

func caloriesParameterCheckHelper(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return fmt.Errorf("bad data: step count must be positive")
	}
	if weight <= 0 {
		return fmt.Errorf("bad data: weight must be positive")
	}
	if height <= 0 {
		return fmt.Errorf("bad data: height must be positive")
	}
	if duration <= 0 {
		return fmt.Errorf("bad data: duration must be positive")
	}
	return nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	err := caloriesParameterCheckHelper(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return caloriesCalculationHelper(steps, weight, height, duration), nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	err := caloriesParameterCheckHelper(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	calories := caloriesCalculationHelper(steps, weight, height, duration) * walkingCaloriesCoefficient
	return  calories, nil
}
