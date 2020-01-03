package handle

import (
	"net/http"
	"text/template"
)

func Showindex(w http.ResponseWriter, r *http.Request) {
	tem, _ := template.ParseFiles("index.html")

	v := 1
	tem.Execute(w, v)
	//execute is template to act and http.RequestWriter に書き出す
}
