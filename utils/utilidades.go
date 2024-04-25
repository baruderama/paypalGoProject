package utilidades

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Frontend string = "templates/layout/frontend.html"

var Store = sessions.NewCookieStore([]byte("session-name"))

func RetornarMensajesFlash(response http.ResponseWriter, request *http.Request) (string, string) {
	session, _ := Store.Get(request, "flash-session")
	fm := session.Flashes("css")
	session.Save(request, response)
	css_session := ""
	if len(fm) == 0 {
		css_session = ""
	} else {
		css_session = fm[0].(string)
	}
	fm2 := session.Flashes("mensaje")
	session.Save(request, response)
	css_mensaje := ""
	if len(fm2) == 0 {
		css_mensaje = ""
	} else {
		css_mensaje = fm2[0].(string)
	}
	return css_session, css_mensaje

}

func CrearMensajesFlash(response http.ResponseWriter, request *http.Request, css string, mensaje string) {

	session, err := Store.Get(request, "flash-session")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	session.AddFlash(css, "css")
	session.AddFlash(mensaje, "mensaje")
	session.Save(request, response)
}
