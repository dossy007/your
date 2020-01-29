package handle

import (
	"net/http"

	"text/template"
)

// type Movie struct {
// 	ID         int
// 	Url        string
// 	CategoryID int
// }

// type Category struct {
// 	ID     int
// 	Name   string
// 	Movies []Movie
// }

func Showindex(w http.ResponseWriter, r *http.Request) {
	tem, _ := template.ParseFiles("index.html")

	v := 2
	tem.Execute(w, v)
	//execute is template to act and http.RequestWriter に書き出す
}
