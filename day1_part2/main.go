package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	
	data, err := os.Open("testdata/day1.txt")
	if err != nil {
		log.Fatal(err)
	}


	defer data.Close();


	scanner := bufio.NewScanner(data)

	largestSum := []int{0,0,0};

	findSum(scanner, &largestSum);

	
}

func findSum(scanner *bufio.Scanner, largestSum *[]int ){

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

	fmt.Println(*largestSum)


}

func setLargestSum(sum int, largestSum *[]int){

	if(!sort.SliceIsSorted(*largestSum, func(i, j int) bool {
		return (*largestSum)[i] > (*largestSum)[j]
	})){
		sort.Sort(sort.Reverse(sort.IntSlice(*largestSum)))
	}

	i := sort.Search(len(*largestSum), func(i int) bool { return (*largestSum)[i] < sum })

	if (i <= 2){
		(*largestSum)[2] = sum
	} 

}


