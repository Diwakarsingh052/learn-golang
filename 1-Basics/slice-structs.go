package main

import "fmt"

type movie struct {
	name   string
	rating int
}

type movies []movie

func main() {

	//a := movie{
	//	name:   "marvel",
	//	rating: 13,
	//}
	allMovies := movies{ // []movie
		{
			name:   "marvel",
			rating: 13,
		},
		{
			name:   "batman",
			rating: 16,
		},
	}
	fmt.Println(allMovies[0].name)
	for _, v := range allMovies {
		fmt.Println(v)
	}

}
