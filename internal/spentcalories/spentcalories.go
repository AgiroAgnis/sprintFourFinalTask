package spentcalories

import (
	"errors"
	"fmt"
	"log"
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
)

var ErrInvalidLength = errors.New("неверный ввод данных")
var ErrNonNumeric = errors.New("требуется число")
var ErrInvalidSteps = errors.New("количество шагов не может быть отрицательным или нулевым")
var ErrInvalidTime = errors.New("неккоректный формат времени")
var ErrInvalidWeight = errors.New("вес должен быть в диапазоне от 40 до 150 кг")
var ErrInvalidHeight = errors.New("рост должен быть в диапазоне от 140 до 250")
var ErrInvalidActivity = errors.New("неизвестный тип тренировки`")

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию

	parts := strings.Split(data, ",")

	if len(parts) != 3 {
		return 0, "", 0, ErrInvalidLength
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, ErrNonNumeric
	}

	if steps <= 0 {
		return 0, "", 0, ErrInvalidSteps
	}

	activity := parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, ErrInvalidTime
	}

	if duration <= 0 {
		return 0, "", 0, ErrInvalidTime
	}

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distance := (stepLength * float64(steps)) / float64(mInKm)
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	dist := distance(steps, height)
	meanSpeed := dist / duration.Hours()
	return meanSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию

	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	durationHours := duration.Hours()

	switch activity {
	case "Бег":
		calories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}

		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
			activity, durationHours, dist, speed, calories)
		return result, nil

	case "Ходьба":
		calories, err := RunningSpentCalories(steps, weight, height, duration)

		if err != nil {
			log.Println(err)
			return "", err
		}
		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
			activity, durationHours, dist, speed, calories)
		return result, nil

	default:
		err := ErrInvalidActivity
		log.Println(err)
		return "", err
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	if steps <= 0 {
		return 0, ErrInvalidSteps
	}

	if duration <= 0 {
		return 0, ErrInvalidTime
	}

	if weight < 40.0 || weight > 150.0 {
		return 0, ErrInvalidWeight
	}

	if height <= 1.40 || height > 2.50 {
		return 0, ErrInvalidHeight
	}

	speed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := (weight * speed * float64(durationInMinutes)) / float64(minInH)
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	if steps <= 0 {
		return 0, ErrInvalidSteps
	}

	if duration <= 0 {
		return 0, ErrInvalidTime
	}

	if weight < 40.0 || weight > 150.0 {
		return 0, ErrInvalidWeight
	}

	if height <= 1.40 || height > 2.50 {
		return 0, ErrInvalidHeight
	}

	speed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	caloriesProcessing := (weight * speed * float64(durationInMinutes)) / float64(minInH)
	calories := caloriesProcessing * walkingCaloriesCoefficient

	return calories, nil
}
