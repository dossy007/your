package handle

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"text/template"
)

func Showindex(w http.ResponseWriter, r *http.Request) {
	tem, _ := template.ParseFiles("index.html")

	v := 1
	tem.Execute(w, v)
	//execute is template to act and http.RequestWriter に書き出す
}

func Showdb(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:@/your?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
