package main

import (
	"fmt"
	"math/rand"
)

// начало решения

// генерит случайные слова из 5 букв
// с помощью randomWord(5)
func generate(cancel <-chan struct{}) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case out <- randomWord(5):
			case <-cancel:
				return
			}
		}
	}()
	return out
}

// выбирает слова, в которых не повторяются буквы,
// abcde - подходит
// abcda - не подходит
func takeUnique(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case str, ok := <-in:
				if !ok {
					return
				}
				if isUnique(str) {
					select {
					case out <- str:
					case <-cancel:
						return
					}
				} else {
					continue
				}
			case <-cancel:
				return
			}
		}
	}()
	return out
}

func isUnique(word string) bool {
	uniqueTable := make(map[rune]int)
	for _, val := range word {
		uniqueTable[val]++
	}
	for count := range uniqueTable {
		if uniqueTable[count] > 1 {
			return false
		}
	}
	return true
}

// переворачивает слова
// abcde -> edcba
func reverse(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case str, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- reverseWord(str):
				case <-cancel:
					return
				}
			case <-cancel:
				return
			}
		}
	}()
	return out
}

func reverseWord(str string) string {
	rns := []rune(str)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

// объединяет c1 и c2 в общий канал
func merge(cancel <-chan struct{}, c1, c2 <-chan string) <-chan string {
	out := make(chan string)
	closer := make(chan struct{})
	go func() {
		for {
			select {
			case out <- <-c1:
			case <-cancel:
				closer <- struct{}{}
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case out <- <-c2:
			case <-cancel:
				closer <- struct{}{}
				return
			}
		}
	}()
	go func() {
		for i := 0; i < 2; i++ {
			<-closer
		}
		close(closer)
		close(out)
	}()
	return out
}

// печатает первые n результатов
func print(cancel <-chan struct{}, in <-chan string, n int) {
	for i := 0; i < n; i++ {
		reversedWord := <-in
		originalWord := reverseWord(reversedWord)
		fmt.Print(originalWord, " -> ", reversedWord, "\n")
	}
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
