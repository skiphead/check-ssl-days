package main

import (
	"check-ssl-service/internal/apiServer"
	"check-ssl-service/web"
	"log"
	"net/http"
	"time"
)

func main() {
	apiServer.CheckTLS("yandex.ru:443")

	mux := http.NewServeMux()
	web.Include(mux)

	//mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./docs/assets/"))))

	//Configuration server
	server := &http.Server{
		Addr:           ":" + apiServer.Conf.Port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// Check TLS on/off

	//		log.Println("TLS HTTP Server listen Port", conf.Port, "version", version())
	//		log.Fatal(server.ListenAndServeTLS(conf.ServerCrt, conf.ServerKey))

	//		log.Println("HTTP Server listen Port", apiServer.Conf.Port)
	log.Fatal(server.ListenAndServe())

}
