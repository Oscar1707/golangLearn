package main

import "fmt"

func main()  {
	
	fmt.Println("-----------------[ FIBONACCI ]----------------")
	n1 :=0
	n2 :=1
	r  :=0
	if n2 == 0 {
		fmt.Println("Hello")
	
	 for i := 0; i < 10; i++ {
		 r = n1+n2
		 fmt.Println("n1:_ ",n1," n2:_ ",n2, "Res:_ ",r)
		 n1 = r+n2
		 r = n1 + n2
	 }
	}

	fmt.Println("Me duele la espalda de cargar el peso")

}

