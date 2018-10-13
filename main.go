package main

import (
	"bufio"
	"fmt"
	"os"
)

type ding struct {
	encKey int
	msg    string
	score  int
}

func main() {
	highScores := make([]ding, 0)
	file, err := os.Open("challenge4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decFile, err := os.Create("dec.txt")
	if err != nil {
		panic(err)
	}
	defer decFile.Close()

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		m := DecryptWithSingleByteXorCipher(line)
		for key, val := range m {
			valAsString := string(val)
			score := GetLetterFrequencyScore(valAsString)
			_, err := decFile.WriteString(valAsString)
			if err != nil {
				panic(err)
			}
			if score > 80 {
				d := ding{key, valAsString, score}
				highScores = append(highScores, d)
			}
		}
	}
	/* dinge := make([]ding, 0)


	for key, val := range m {
		valAsString := string(val)
		d := ding{encKey: key, msg: valAsString, score: GetLetterFrequencyScore(valAsString)}
		dinge = append(dinge, d)
	}
	sort.Slice(dinge, func(i, j int) bool {
		return dinge[i].score > dinge[j].score
	})

	for i := 0; i < 10; i++ {
		fmt.Println(dinge[i])
	} */
	//fmt.Printf("Key: %d as Char: %c \nValue: %s\n \n", 88, 88, m[88])

}
