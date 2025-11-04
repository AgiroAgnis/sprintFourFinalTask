package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	spentaclories "github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

var ErrInvalidLength = errors.New("неверный ввод данных")
var ErrNonNumeric = errors.New("требуется число")
var ErrInvalidSteps = errors.New("количество шагов не может быть отрицательным или нулевым")
var ErrInvalidTime = errors.New("неккоректный формат времени")

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию

	parts := strings.SplitN(data, ",", 2)
	if len(parts) != 2 {
		return 0, 0, ErrInvalidLength
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, ErrNonNumeric
	}

	if steps <= 0 {
		return 0, 0, ErrInvalidSteps
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, ErrInvalidTime
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию

	steps, duration, err := parsePackage(data)

	if err != nil {
		switch err {
		case ErrInvalidLength:
			return "Неверный ввод данных"
		case ErrNonNumeric:
			return "Требуется число"
		case ErrInvalidSteps:
			return "Количество шагов не может быть отрицательным или нулевым"
		case ErrInvalidTime:
			return "Неккоректный формат времени"
		default:
			return "Неизвестная ошибка"
		}
	}

	distance := (float64(steps) * stepLength) / float64(mInKm)

	calories, err := spentaclories.WalkingSpentCalories(steps, weight, height, duration)

	if err != nil {
		switch err {
		case spentaclories.ErrInvalidWeight:
			return "Вес должен быть в диапазоне от 40 до 150 кг"
		case spentaclories.ErrInvalidHeight:
			return "Рост должен быть в диапазоне от 140 до 250"
		default:
			return "Ошибка расчета калорий"
		}
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distance, calories)
	return result
}
