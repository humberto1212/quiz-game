package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func qa(s *int) int {

	f, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		var i int

		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", line[0])

		fmt.Scan(&i)
		resultToInt, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}

		if i == resultToInt {
			*s++
		}

	}

	return *s
} //end qa

func main() {
	sum := 0
	var total *int = &sum

	go qa(total)

	time.Sleep(5 * time.Second)

	fmt.Println("time is over: of 13 questions you answered", *total, "correct")
}
