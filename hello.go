package main

import "fmt"

func main()  {
	
	//Add fibonacci serie
	//fibonacci()

	//Hacker Rank
	weirdN(3)

}

func weirdN(num float32)  {
	fmt.Println("-----------------[ WEIRDN ]----------------")
	fmt.Println("num ",num)
	
	var isPair float32 = num / 2
	fmt.Println("Vamos para sigapr:_ ", isPair)
	

}

func fibonacci()  {
	fmt.Println("-----------------[ FIBONACCI ]----------------")

	n1 :=0
	n2 :=1
	r  :=0
	
	 for i := 0; i < 10; i++ {
		 r = n1+n2
		 fmt.Println(n1," + ",n2, "= ",r)
		n1 = n2
		n2 = r
	}	
}

