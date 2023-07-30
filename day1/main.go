package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	
	data, err := os.Open("testdata/day1.txt")

	defer data.Close();


	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(data)

	largestSum := 0

	findSum(scanner, &largestSum);

	
}

func findSum(scanner *bufio.Scanner, largestSum *int ){



	sum := 0


	for scanner.Scan() {
		
		if(scanner.Text() == ""){
				setLargestSum(sum, largestSum)
				sum = 0
				scanner.Scan();
		}


		num, err := strconv.Atoi(scanner.Text());

		if err != nil {
			log.Fatal(err)
		}

	    sum += num		
	}

    setLargestSum(sum , largestSum);

	fmt.Println(*largestSum);



}

func setLargestSum(sum int, largestSum *int){

	if (sum > *largestSum) {
		*largestSum = sum
	}

}