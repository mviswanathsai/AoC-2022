package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	data, err := os.Open("testdata/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)

	for scanner.Scan(){
		fmt.Println(string(scanner.Text()[2]))
	}

}