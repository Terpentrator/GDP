package main // Пакет main

import (
	"fmt" // Добавляем стандартный модуль fmt

	"github.com/lukesiler/stringutil" // Добавляем кастомный модуль stringutil
)

func main() {
	fmt.Println(stringutil.Reverse("Hello, OTUS!")) // Выводим результат
}
