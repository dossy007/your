package handle

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Movie struct {
	ID         int
	Url        string
	CategoryID int
}

type Category struct {
	ID     int
	Name   string
	Movies []Movie
}

func Showindex(w http.ResponseWriter, r *http.Request) {
	tem, _ := template.ParseFiles("index.html")

	v := 1
	tem.Execute(w, v)
	//execute is template to act and http.RequestWriter に書き出す
}

func Showdb(w http.ResponseWriter, r *http.Request) {

	v := Connected()

	tem, _ := template.ParseFiles("index.html")

	tem.Execute(w, v)
}

func Connected() []Category { //2重slice 全件取得
	db := ConnectDB()
	defer db.Close()

	//sql
	query := `SELECT c.id, c.name, m.id, m.url FROM categories c left join movies m on c.id = m.category_id ORDER BY c.id ASC`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var sli []Category
	var c Category

	var m Movie

	for rows.Next() { //nextはscanを使う為
		if err := rows.Scan(&c.ID, &c.Name, &m.ID, &m.Url); err != nil {

			log.Fatal(err)
		}
		m.CategoryID = c.ID
		l := len(sli)
		if l > 0 && sli[l-1].ID == c.ID {
			sli[l-1].Movies = append(sli[l-1].Movies, m)
			//[]Moviesに入れる処理
		} else {
			if len(c.Movies) != 0 { //Initialize c.Movies
				c.Movies = remove(c.Movies, c.Movies[0])
			}
			c.Movies = append(c.Movies, m) //[]c.Moviesにappend
			sli = append(sli, c)
		}
	}
	fmt.Println(sli)
	return sli
}

var db *sql.DB

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@/your?parseTime=true")

	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func remove(ints []Movie, search Movie) []Movie {
	result := []Movie{}
	for _, v := range ints {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}
