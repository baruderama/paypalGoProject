package rutas

import (
	"fmt"
	"html/template"
	"net/http"
	utilidades "paypalGoProject/utils"
	"paypalGoProject/validations"
)

func Formularios_get(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/formulario/formulario.html", utilidades.Frontend))
	template.Execute(response, nil)
}

func Formularios_post(response http.ResponseWriter, request *http.Request) {
	// template := template.Must(template.ParseFiles("templates/formulario/formulario.html", utilidades.Frontend))
	// template.Execute(response, nil)
	mensaje := ""
	correo := request.FormValue("correo")
	if !validations.Regex_correo.MatchString(correo) {
		mensaje = mensaje + " . El E-mail ingresado no es valido"
		fmt.Fprintln(response, request.FormValue("correo"))
	}

	if !validations.ValidarPassword(request.FormValue("password")) {
		mensaje = mensaje + " . El password debe tener al menos 1 numero, una mayuscula, y un largo entre 6 y 20 caracteres"
	}
	if mensaje != "" {
		fmt.Fprintln(response, mensaje)
		return
	}
	fmt.Fprintln(response, request.FormValue("nombre"))
}
