//Задача: Вы разрабатываете текстовый редактор для
//программистов и хотите реализовать проверку
//корректности расстановки скобок. В коде могут
//встречаться скобки []{}(). Из них скобки [,{
//и ( считаются открывающими, а соответствующими им закрывающими скобками являются
//],} и ).
// случае, если скобки расставлены неправильно, редактор должен
//также сообщить пользователю первое место, где обнаружена ошибка.
//В первую очередь необходимо найти закрывающую скобку, для которой либо нет соответствующей открывающей (например, скобка ] в
//строке “]()”), либо же она закрывает не соответствующую ей открывающую скобку (пример: “()[}”). Если таких ошибок нет,
//необходимо найти первую открывающую скобку, для которой нет соответствующей закрывающей (пример: скобка ( в строке “{}([]”).
//Помимо скобок, исходный код может содержать символы латинского алфавита, цифры и знаки препинания
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func vivodIndex(otv []int, stek []string) {
	var k int
	var ok string
	if ok != "ok" {
		for i := 0; i != len(stek); i++ {
			if stek[i] == "}" || stek[i] == ")" || stek[i] == "]" {
				fmt.Println(otv[k])
				ok = "ok"

				break
			} else {
				continue
			}
			k = k + 1
			break
		}
	} else {
		fmt.Println(otv[0])
	}
}

func checkStek(vhod string) {
	sl := strings.Split(vhod, "")
	var stek []string
	var nom []int
	var n int
	var k int

	for i := 0; i != len(sl); i++ {
		n = len(stek)
		if sl[i] == "{" || sl[i] == "}" || sl[i] == "(" || sl[i] == ")" || sl[i] == "[" || sl[i] == "]" {
			stek = append(stek, sl[i])
			nom = append(nom, i)
			k = len(nom)
			if n > 0 {
				switch {
				case stek[n] == ")":
					if stek[n-1] == "(" {
						stek = append(stek[0 : n-1])
						nom = append(nom[0 : k-2])
					} else {
						continue
					}
				case stek[n] == "]":
					if stek[n-1] == "[" {
						stek = append(stek[0 : n-1])
						nom = append(nom[0 : k-2])
					} else {
						continue
					}
				case stek[n] == "}":
					if stek[n-1] == "{" {
						stek = append(stek[0 : n-1])
						nom = append(nom[0 : k-2])
					} else {
						continue
					}
				default:
					continue
				}
			} else {
				continue
			}
		} else {
			continue
		}
	}
	if len(stek) == 0 {
		fmt.Println("Success")
	} else {

		fmt.Println(nom, stek)
	}
}
func main() {
	var vhod string
	vhod, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	if strings.Contains(vhod, ")") || strings.Contains(vhod, "(") || strings.Contains(vhod, "]") || strings.Contains(vhod, "[") || strings.Contains(vhod, "{") || strings.Contains(vhod, "}") {
		checkStek(vhod)
	} else {
		fmt.Println("0")

	}
}
