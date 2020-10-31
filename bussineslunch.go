package main

import "fmt"
type Bl struct {
	first string
	second string
	desert string
	bibation string
}

func AddBl(array *[]Bl, NewBl Bl){
	*array = append(*array, NewBl)
	fmt.Println(array)
}

func main() {
	Bls :=[]Bl{
		{first: "Borsh",
			second:   "Tabaka",
			desert:   "Shakarchoy",
			bibation: "Lemon Choy",
		},
		{first: "Shurbo",
			second:   "Shashlik",
			desert:   "Kurut",
			bibation: "RC Cola",
		},
		{first: "Ugro",
			second:   "Gusht biryon",
			desert:   "Chakai turush",
			bibation: "Coca Cola",
		},

	}
	NewBl := Bl{
		first: 	  "Ugro",
		second:   "Gusht biryon",
		desert:   "Chakai turush",
		bibation: "Coca Cola",
	}
	fmt.Println(Bls)
	fmt.Println(&NewBl, Bls)
}