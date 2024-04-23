package rutas

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/home.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	}
}

// func Home(response http.ResponseWriter, request *http.Request) {

// 	fmt.Fprintln(response, "hola mundo desde Golang")
// }

func Nosotros(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/nosotros.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	}
}

func Parametros(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/parametros.html")
	vars := mux.Vars(request)
	data := map[string]string{
		"id":   vars["id"],
		"slug": vars["slug"],
	}
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, data)
	}
	fmt.Fprintln(response, "ID ="+vars["id"]+" | SLUG="+vars["slug"])
}

func ParametrosQueryString(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/parametros-query-string.html")
	data := map[string]string{
		"id":   request.URL.Query().Get("Id"),
		"slug": request.URL.Query().Get("slug"),
	}
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, data)
	}
	fmt.Fprintln(response, request.URL)
	fmt.Fprintln(response, request.URL.RawQuery)
	fmt.Fprintln(response, request.URL.Query().Get("Id"))
	fmt.Fprintln(response, request.URL.Query().Get("slug"))
	id := request.URL.Query().Get("Id")
	slug := request.URL.Query().Get("slug")
	fmt.Fprintln(response, "----------------------------------------------")
	fmt.Fprintf(response, "id=%s | slug = %s", id, slug)
}

type Habilidades struct {
	Nombre string
}
type Datos struct {
	Nombre      string
	Edad        int
	Perfil      int
	Habilidades []Habilidades
}

func Estructuras(response http.ResponseWriter, request *http.Request) {
	template, _ := template.ParseFiles("templates/ejemplo/estructuras.html")
	habilidade1 := Habilidades{"Inteligencia"}
	habilidade2 := Habilidades{"Videojuegos"}
	habilidade3 := Habilidades{"Programacion"}
	habilidade4 := Habilidades{"Canto"}
	habilidades := []Habilidades{habilidade1, habilidade2, habilidade3, habilidade4}
	template.Execute(response, Datos{"Cesar", 42, 1, habilidades})
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	template.Execute(response, Datos{"cesar", 19, 2})
	// }
}
