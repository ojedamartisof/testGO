package main

import (
"fmt"
)

func main() {
	deporteFav := "béisbol"
	switch deporteFav {
	case "béisbol":
		fmt.Println("Ve al estadio")
	case "natación":
		fmt.Println("Ve a la pileta")
	case "crossfit":
		fmt.Println("Te quiero ver en los sports.")
	}
}