package web

import (
	"check-ssl-service/internal/apiServer"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("web/templates/*"))

func parseDocs() DocStruct {
	jsonFile, errOpenJson := os.Open("./docs/docs.json")
	if errOpenJson != nil {
		log.Println("Error open json file", errOpenJson)
	}

	defer func(jsonFile *os.File) {
		errCloseJson := jsonFile.Close()
		if errCloseJson != nil {
			log.Println("Error close config file", errCloseJson)
		}
	}(jsonFile)

	byteConfig, errByteConfig := ioutil.ReadAll(jsonFile)
	if errByteConfig != nil {
		log.Println("Err read byte", errByteConfig)
	}
	docs := DocStruct{}
	err := json.Unmarshal(byteConfig, &docs)
	if err != nil {
		log.Println("Error parse config", err)
	}
	return docs
}

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
