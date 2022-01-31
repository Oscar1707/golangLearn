package main

import "fmt"

func main()  {
	
	//Add fibonacci serie
	fibonacci()

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

