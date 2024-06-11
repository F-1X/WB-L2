package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

// Базовая задача
// Создать программу печатающую точное время с использованием
// NTP -библиотеки. Инициализировать как go module. Использовать
// библиотеку github.com/beevik/ntp. Написать программу
// печатающую текущее время / точное время с использованием этой
// библиотеки.
// Требования:
// 1. Программа должна быть оформлена как go module
// 2. Программа должна корректно обрабатывать ошибки
// библиотеки: выводить их в STDERR и возвращать ненулевой
// код выхода в OS

// Решение - модуль ntp

func main() {
	GetCurrentTime()
}

func GetCurrentTime() {
	time, err := ntp.Time("0bee1vik-ntp.pool.nt1p.org")
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}
	fmt.Println(time)
}
