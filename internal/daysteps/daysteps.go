package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// Переводит строку в слайс
	divLine := strings.Split(data, ",")
	if len(divLine) != 2 {
		return 0, 0, fmt.Errorf("Неверная формат данных")
	}

	// Проверяет шаги
	steps, err := strconv.Atoi(divLine[0])
	if err != nil || steps <= 0 {
		return 0, 0, fmt.Errorf("неверное количество шагов")
	}

	// Проверяет продолжительность
	t, err := time.ParseDuration(divLine[1])
	if err != nil || t <= 0 {
		return 0, 0, fmt.Errorf("неверная продолжительность")
	}
	return steps, t, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// Переводит строку в слайсы parsePackage и получает данные
	step, t, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	// Вычисляет дистанию в метрах
	distanceMetre := float64(step) * stepLength
	// Переводит дистанцию в километры
	distanceKiloMetre := distanceMetre / mInKm
	// Вычисляет калории
	calories, err := spentcalories.WalkingSpentCalories(step, weight, height, t)
	if err != nil {
		log.Println(err)
		return ""
	}
	//Формирует строку для вывода
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", step, distanceKiloMetre, calories)
	return result
}
