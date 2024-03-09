package hw03frequencyanalysis

import (
	"strings"
	"unicode"
)

func Top10(str string) []string {
	// Place your code here.
	// ----- Проверка на пустую строку
	if str == "" {
		return nil
	}
	// ----- Делаем все буквы строчными
	refStringNew := strings.ToLower(str)
	// ----- Удаляем ' - ' из строки
	refStringNew = strings.ReplaceAll(refStringNew, " - ", " ")
	// ----- Создаём мапу по словам
	mapwords := make(map[string]int)
	// ----- Создаём слайс по самым популярным словам
	strout := make([]string, 0, 10)
	// ----- Выделяем из текста слова
	f := func(c rune) bool {
		return unicode.IsSpace(c) || c == '.' || c == ',' || c == '!' || c == '\'' || c == '"'
	}
	words := strings.FieldsFunc(refStringNew, f)
	// ----- Cчитаем количество слов
	for _, word := range words {
		mapwords[word]++
	}
	// ----- Запоминаем длину мапы
	lenmapwords := len(mapwords)
	// ----- Считаем слова и количество повторений
	for i := 0; (i < 10) && (i < lenmapwords); i++ {
		maxvol := 0
		for mp, vol := range mapwords {
			if maxvol < vol {
				maxvol = vol
				if len(strout) <= i {
					strout = append(strout, mp)
				} else {
					strout[i] = mp
				}
			}
			if (maxvol == vol) && (mp < strout[i]) {
				if len(strout) <= i {
					strout = append(strout, mp)
				} else {
					strout[i] = mp
				}
			}
		}
		delete(mapwords, strout[i])
	}
	return strout
}
