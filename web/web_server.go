package web

import (
	"check-ssl-service/internal/apiServer"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("web/templates/*"))

func CheckTLS(w http.ResponseWriter, r *http.Request) {
	check := r.URL.Query().Get("url_check")
	//fmt.Println(apiServer.CheckTLS(check))
	fmt.Println(check)

	if check != "" {
		errExecuteTmpl := tmpl.ExecuteTemplate(w, "Main", apiServer.CheckTLS(check))
		if errExecuteTmpl != nil {
			log.Println(errExecuteTmpl)
		}
	} else {
		res := apiServer.DayStruct{
			Day: "input endpoint",
		}
		errExecuteTmpl := tmpl.ExecuteTemplate(w, "Main", res)
		if errExecuteTmpl != nil {
			log.Println(errExecuteTmpl)
		}
	}

}

func Include(mux *http.ServeMux) {
	mux.HandleFunc("/web", CheckTLS)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets/"))))
}
