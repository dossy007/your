package handle

import (
	"fmt"
	"net/http"
)

func Showindex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Helloopopop, World")

}
