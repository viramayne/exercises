package main

import "fmt"

func collatzesHypothesis() {
	/* Гипотеза Коллатца
	Найдите число шагов, за которые получится единица, используя следующий процесс:
	берём любое натуральное число n больше единицы. Если оно чётное, то делим его на 2,
	а если нечётное, то умножаем на 3 и прибавляем 1*/
	n := 30
	count := 0
	for {
		if n == 1 {
			break
		}
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		count++
	}
	fmt.Println(count)
}