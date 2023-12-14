package main

import(
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/SylvanasGr/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting Go Api service...")
	err := http.ListenAndServe("localhost:9000",r)
	if err != nil {
		log.Error(err)
	}
}