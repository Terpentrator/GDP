package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")
var ErrInvalidAfterSlashStick = errors.New("wrong after the slash stick")
var ErrInvalidNotSymbolOrSlashStick = errors.New("not a symbol or a slash stick")

func Unpack(str string) (string, error) {
	// Place your code here.

	strout := []string{} // Выходная строка
	var count byte		 // Счётчик количества повторений

	// ----- Цикл от забора и до заката по str
	for i := 0; i < len(str); i++ {
		// ----- Если косая
		if string(str[i]) == `\` {
			// ----- Если стоим на предпоследнем символе
			if i == len(str)-2 {
				// ----- Сохраняем символ и сваливаем
				strout = append(strout, string(str[i]))
				// ----- Сохраняем символ после неё, и сваливаем
				strout = append(strout, string(str[i+1]))
				break
				// ----- Если после символа идёт символ
			}
			// ----- Если после косой цифра
			if (string(str[i+1]) >= "1") && (string(str[i+1]) <= "9") {
				// ----- Если после косой c цифрой - цифра
				if (string(str[i+2]) >= "1") && (string(str[i+2]) <= "9") {
					// ----- Сохраняем цифру
					strout = append(strout, string(str[i+1]))
					// ----- Определяем кол-во повторений
					count = str[i+2] - 48
					// ----- Добавляем в выводим количество цифр
					for j := 0; j < int(count)-1; j++ {
						strout = append(strout, string(str[i+1]))
					}
					// ----- Увеличиваем счетчик на 2
					i += 2
					// ----- Переходим к новой итерации цикла
					continue
					// ----- Если после косой c цифрой - не цифра
				} else {
					// ----- Сохраняем цифру
					strout = append(strout, string(str[i+1]))
					// ----- Увеличиваем счетчик на 1
					i++
					// ----- Переходим к новой итерации цикла
					continue
				}
				// ----- Если после косой - косая
			} else if string(str[i+1]) == `\` {
				// ----- Если после косой c косой (!) - цифра
				if (string(str[i+2]) >= "1") && (string(str[i+2]) <= "9") {
					// ----- Сохраняем косую
					strout = append(strout, `\`)
					// ----- Определяем кол-во повторений
					count = str[i+2] - 48
					// ----- Добавляем в выводим количество косых
					for j := 0; j < int(count)-1; j++ {
						strout = append(strout, `\`)
					}
					// ----- Увеличиваем счетчик на 2
					i += 2
					// ----- Переходим к новой итерации цикла
					continue
					// ----- Если после косой c косой (!) - не цифра
				} else {
					// ----- Сохраняем косую
					strout = append(strout, `\`)
					// ----- Увеличиваем счетчик на 1
					i++
					// ----- Переходим к новой итерации цикла
					continue
				}
				// ----- Если после косой - n
			} else if string(str[i+1]) == `n` {
				// ----- Если после косой c n - цифра
				if (string(str[i+2]) >= "1") && (string(str[i+2]) <= "9") {
					// ----- Сохраняем косую
					strout = append(strout, `\`)
					// ----- Сохраняем n
					strout = append(strout, `n`)
					// ----- Определяем кол-во повторений
					count = str[i+2] - 48
					// ----- Добавляем в выводим количество косых c n
					for j := 0; j < int(count)-1; j++ {
						strout = append(strout, `\`)
						strout = append(strout, `n`)
					}
					// ----- Увеличиваем счетчик на 2
					i += 2
					// ----- Переходим к новой итерации цикла
					continue
					// ------ Если после косой c n - не цифра
				} else {
					// ----- Сохраняем косую
					strout = append(strout, `\`)
					// ----- Сохраняем n
					strout = append(strout, `n`)
					// ----- Увеличиваем счетчик на 1
					i++
					// ----- Переходим к новой итерации цикла
					continue
				}
				// ----- Если после косой всё остальное
			} else {
				return "",ErrInvalidAfterSlashStick
			}
		}

		// ----- Если символ
		if (string(str[i]) >= "a") && (string(str[i]) <= "z") {
			// ----- Если стоим на последнем символе
			if i == len(str)-1 {
				// ----- Сохраняем его и сваливаем
				strout = append(strout, string(str[i]))
				break
				// ----- Если после символа идёт символ
			} else if (string(str[i+1]) >= "a") && (string(str[i+1]) <= "z") {
				// ----- Созраняем символ и переходим к следующей итерации цикла
				strout = append(strout, string(str[i]))
				continue
				// ----- Если после символа идёт цифра
			} else if (string(str[i+1]) >= "1") && (string(str[i+1]) <= "9") {
				// ----- Созраняем символ
				strout = append(strout, string(str[i]))
				// ----- Определяем кол-во повторений
				count = str[i+1] - 48
				// ----- Добавляем в выводим количество символов
				for j := 0; j < int(count)-1; j++ {
					strout = append(strout, string(str[i]))
				}
				// ----- Увеличиваем счетчик на 1
				i++
				// ----- Переходим к новой итерации цикла
				continue
				// ----- Если после символа идёт ноль
			} else if string(str[i+1]) == "0" {
				// ----- Увеличиваем счетчик на 1
				i++
				// ----- Переходим к новой итерации цикла
				continue
			} else {
				// ----- Созраняем символ
				strout = append(strout, string(str[i]))
				// ----- Переходим к новой итерации цикла
				continue
			}
		} else {
			return "",ErrInvalidNotSymbolOrSlashStick
		}
	}
	// ----- Делаем из выходного слайса strout выходную строку 
	return strings.Join(strout, ""),nil
}
