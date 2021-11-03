package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-clean-arch/adapters/controllers"
	"log"
	"net/http"

	"go-clean-arch/web/model"
)

type Server struct {
	userUseCases controllers.UserUseCasesAdapter
}

func NewServer(uc controllers.UserUseCasesAdapter) Server {
	return Server{userUseCases: uc}
}

func (s Server) Start() {
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/getuser", s.getUser).Methods("GET")
	http.Handle("/", router)

	//start and listen to requests
	err := http.ListenAndServe(
		":8080",
		router,
	)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s Server) getUser(w http.ResponseWriter, r *http.Request) {
	log.Println("entering get user")
	rs := model.FromCoreEntity(s.userUseCases.GetUser())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(rs)
	if err != nil {
		return
	}

	w.Write(jsonResponse) //nolint:errcheck
}
