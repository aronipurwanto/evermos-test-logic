package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func initLayout() ([][]string, []string, string) {
	array := make([][]string,6)
	for i := range array{
		array[i] = make([]string,8)
	}

	var prop []string
	var position string

	for i := range array{
		for j:= range array[i] {
			if i== 0 || i==5 || j==0 || j==7 {
				array[i][j] = "#"
			}else if i==2 && j>=2 && j <= 4 {
				array[i][j] = "#"
			}else if (i==3 && j== 4) || (i ==3 && j ==6) {
				array[i][j] = "#"
			}else if i==4 && j==2{
				array[i][j] = "#"
			}else if i==4 && j==1{
				array[i][j] = "X"
				position = strconv.Itoa(i)+"|"+strconv.Itoa(j)
			} else {
				array[i][j] = "."
				point := strconv.Itoa(i)+"|"+strconv.Itoa(j)
				prop = append(prop, point)
			}
		}
	}

	return array,prop, position
}

func printArray(array [][]string)  {
	for i:= range array {
		for j := range array[i] {
			fmt.Print(array[i][j] +"\t")
		}
		fmt.Println("\n")
	}
}

func main() {
	array,prop, point := initLayout()

	//print array
	printArray(array)

	fmt.Println("empty koordinat: ", prop)
	fmt.Println("current positin: ", point)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("A -> UP")
	fmt.Println("B -> RIGHT")
	fmt.Println("C -> DOWN")
	
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}
	}
}
