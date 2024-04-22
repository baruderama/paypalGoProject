package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"paypalGoProject/rutas"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	//Router
	//----------------------------------------------------------------------------------------
	mux := mux.NewRouter()
	mux.HandleFunc("/", rutas.Home)
	mux.HandleFunc("/nosotros", rutas.Nosotros)
	mux.HandleFunc("/parametros/{id:.*}/{slug:.*}", rutas.Parametros)      //formato path
	mux.HandleFunc("/paramteros-querystring", rutas.ParametrosQueryString) //formato query-String
	mux.HandleFunc("/estructuras", rutas.Estructuras)
	//----------------------------------------------------------------------------------------

	//Ejecucion del servidor
	//----------------------------------------------------------------------------------------
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
		return
	}

	server := &http.Server{
		Addr:         "192.168.1.7:" + os.Getenv("PORT"),
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("corriendo servidor desde http://192.168.1.7:" + os.Getenv("PORT"))
	log.Fatal(server.ListenAndServe())
	//----------------------------------------------------------------------------------------

}

/*
func main() {
	//mux := http.NewServeMux()
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "hola mundo")
	},
	)
	fmt.Println("corriendo el servidor en http://192.168.1.7:8000")
	log.Fatal(http.ListenAndServe("192.168.1.7:8000", nil))
}
*/
