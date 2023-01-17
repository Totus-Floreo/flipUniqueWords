package main

import (
	"fmt"
	"math/rand"
)

// начало решения

// генерит случайные слова из 5 букв
// с помощью randomWord(5)
func generate(cancel <-chan) <-chan {
	out := make(chan)
	go func() {
		// ...
	}()
	return out
}

// выбирает слова, в которых не повторяются буквы,
// abcde - подходит
// abcda - не подходит
func takeUnique(cancel <-chan, in <-chan) <-chan  {
	out := make(chan)
	go func() {
		// ...
	}()
	return out
}

// переворачивает слова
// abcde -> edcba
func reverse(cancel <-chan, in <-chan) <-chan {
	out := make(chan)
	go func() {
		// ...
	}()
	return out
}

// объединяет c1 и c2 в общий канал
func merge(cancel <-chan, c1, c2 <-chan) <-chan {
	out := make(chan)
	go func() {
		// ...
	}()
	return out
}

// печатает первые n результатов
func print(cancel <-chan, in <-chan, n int) {
	// ...
}

// конец решения

// генерит случайное слово из n букв
func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	c1 := generate(cancel)
	c2 := takeUnique(cancel, c1)
	c3_1 := reverse(cancel, c2)
	c3_2 := reverse(cancel, c2)
	c4 := merge(cancel, c3_1, c3_2)
	print(cancel, c4, 10)
}
