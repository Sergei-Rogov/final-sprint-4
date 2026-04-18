package spentcalories

import (
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

func parseTraining(data string) (int, string, time.Duration, error) {
	// Разделяет строку на слайсы
	divLine := strings.Split(data, ",")
	// Проверяет длину слайса
	if len(divLine) != 3 {
		return 0, "", 0, fmt.Errorf("Неверная длина")
	}
	// Преобразует первый элемент в int, проверяет на ошибки
	steps, err := strconv.Atoi(divLine[0])
	if err != nil {
		return 0, "", 0, err
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("неверный шаг")
	}
	// Преобразует третий элемент в time.Duration, проверяет на ошибки
	t, err := time.ParseDuration(divLine[2])
	if err != nil {
		return 0, "", 0, err
	}
	if t <= 0 {
		return 0, "", 0, fmt.Errorf("неверная продолжительность")
	}
	active := divLine[1]
	return steps, active, t, nil
}

func distance(steps int, height float64) float64 {
	// Расчитывает дистанцию
	stepLength := height * stepLengthCoefficient
	v := float64(steps) * stepLength
	dist := v / mInKm
	return dist
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Проверяет на неверные параметры
	if duration <= 0 || steps <= 0 {
		return 0
	}
	// Расчитывает ср.скорость
	dist := distance(steps, height)
	result := dist / duration.Hours()
	return result
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// Получает значения из строки
	steps, active, t, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// Проверяет вид тренировки и возвращает сроку
	var calories float64
	switch active {
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, t)
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, t)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, t)

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		active,
		t.Hours(),
		dist,
		speed,
		calories,
	)

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Неверные параметры")
	}
	mSpeed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	if minutes <= 0 {
		return 0, fmt.Errorf("неверная продолжительность")
	}
	// Уможает вес пользователя на Ср.скорость и Продолжительность в минутах
	v := mSpeed * minutes
	// Делит v на число в минутах для получения калорий
	calories := (weight * v) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Неверные параметры")
	}
	mSpeed := meanSpeed(steps, height, duration)
	calories := (weight * mSpeed * duration.Minutes()) / minInH
	result := calories * walkingCaloriesCoefficient
	return result, nil
}
