package initialization

import (
	. "../models"

	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InitDb() error {
	file, err := os.Open("data/words_frequency/en.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	words := make(map[string]int)
	maximumFrequency := 0

	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, " ")
		word, _frequency := s[0], s[1]

		frequency, err := strconv.Atoi(_frequency)
		if err != nil {
			panic(err)
		}
		words[word] = frequency
		if frequency < 1 {
			panic(frequency)
		}

		if frequency > maximumFrequency {
			maximumFrequency = frequency
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println("Number of new words: ", maximumFrequency)

	for word, frequency := range words {
		fmt.Println(frequency, word)

	}

	return nil
}
