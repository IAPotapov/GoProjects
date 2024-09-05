package main

import (
	"Skyscrapers/models"
	"Skyscrapers/views"
	"fmt"
	"net/http"
	"time"
)

var field *models.Field
var body string

func root(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%v\n", body)
}
func next(w http.ResponseWriter, req *http.Request) {
	r := field.FindAnyElimination()
	for i := 0; i < len(r); i++ {
		field.CommitEllimination(&r[i])
	}
	body = views.GetTable(field)
	fmt.Fprintf(w, "%v\n", body)
}

func main() {
	t1 := time.Now()
	fmt.Println("Start time:", t1.Format(time.DateTime))

	//clues := []int{1, 3, 2, 2, 2, 2, 1, 3, 2, 2, 1, 3, 2, 3, 2, 1}
	//clues := []int{1, 4, 2, 2, 2, 1, 2, 3, 3, 2, 1, 3, 2, 3, 2, 1}
	//clues := []int{0, 0, 4, 0, 0, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	// multiple solutions:
	//clues := []int{0, 2, 0, 0, 0, 2, 0, 0, 0, 2, 0, 2, 2, 0, 0, 0}

	clues := []int{0, 0, 0, 0, 0, 0, 2, 0, 1, 0, 0, 3, 3, 0, 3, 0}
	field = models.NewField(clues)
	//list := field.FindAllElliminations()
	//fmt.Println("Number of found eliminations:", len(list))
	g := field.Guess()
	if g == models.Fail {
		fmt.Println("Fail")
	}
	if g == models.Solved {
		fmt.Println("Solved")
	}
	if g == models.Unknown {
		fmt.Println("Unknown")
	}

	fmt.Println(field.Serialize())
	// debug web server
	/*body = views.GetTable(field)
	http.HandleFunc("/", root)
	http.HandleFunc("/next", next)
	http.ListenAndServe(":8090", nil)*/

	t2 := time.Now()
	fmt.Println("Finish time:", t2.Format(time.DateTime))
	d := t2.Sub(t1)

	s := d.Seconds()
	fmt.Printf("Duration: %v sec\n", s)
}
