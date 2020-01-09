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
	// db, err := gorm.Open("mysql", "root:@/your?charset=utf8&parseTime=True&loc=Local")

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	//db接続,category,youtube table作成model

	// movie := Movie{Url: "https://www.youtube.com/watch?v=dQaLSNWAKtQ", CategoryID: 2}
	// db.Create(&movie)

	//all gets

	// movie := []Movie{}
	// db.Where("category_id = ?", 2).Find(&movie)
	// for _, a := range movie {
	// 	fmt.Println(a.Url) //idはIDで取得 unit
	// }

	// category := []Category{}
	// var category Category
	// db.Where("category_id = ?", 2).Find(&movie)
	// fmt.Println(movie)
	// for _, a := range category {
	// 	fmt.Println(a)
	// }

	tem, _ := template.ParseFiles("index.html")

	// var sli []Category
	// var m []Movie
	// var v1 Category
	// var v2 Movie

	// sli = append(sli, v1)

	// user := Product{Text: "aaa", Image: 3333}
	// v := 1
	tem.Execute(w, v)
}

func Connected() []Category { //2重slice 全件取得
	db := ConnectDB()
	defer db.Close()

	//sql
	// query := `SELECT c.name,m.url FROM categories c left join movies m on c.id = m.category_id`
	query := `SELECT c.id, c.name, m.id, m.url FROM categories c left join movies m on c.id = m.category_id ORDER BY c.id ASC`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var sli []Category
	var c Category //structをv1で使用する宣言

	var m Movie

	for rows.Next() { //nextはscanを使う為
		// if err := rows.Scan(&v1.Name, &v1.Movies); err != nil {
		if err := rows.Scan(&c.ID, &c.Name, &m.ID, &m.Url); err != nil {

			log.Fatal(err)
		}

		// sli = append(sli, v1)
		// fmt.Println(m.CategoryID)
		// fmt.Println(c.ID)
		m.CategoryID = c.ID
		l := len(sli)
		if l > 0 && sli[l-1].ID == c.ID {
			sli[l-1].Movies = append(sli[l-1].Movies, m)
			// fmt.Println(sli[l-1].Movies, true)

			//[]Moviesに入れる処理
		} else {
			// fmt.Println(len(c.Movies))
			if len(c.Movies) != 0 {
				// c.Movies
				c.Movies = remove(c.Movies, c.Movies[0])
			}
			fmt.Printf("%T\n", c.Movies) //初期値をからに設定し直す
			// fmt.Println(c.Movies, "cmovies")
			// c.Movies = []handle.Movie
			c.Movies = append(c.Movies, m) //cの[]moviesにappend
			// fmt.Println(m, "qwqwqw")
			sli = append(sli, c) //sliにcの塊が入る
			//lが+される
			// fmt.Println(sli)
			// fmt.Println(sli, false)

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
