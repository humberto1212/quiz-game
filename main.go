package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	sum := 0

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
		//fmt.Printf("var1 = %T\n", line)

		// for l, _ := range line {
		// 	fmt.Println(l)
		// }

		fmt.Scan(&i)
		resultToInt, err := strconv.Atoi(line[1])

		if i == resultToInt {
			sum++
		}

	}

	fmt.Println(sum)

}
