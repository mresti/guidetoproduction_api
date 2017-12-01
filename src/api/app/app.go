package app

import (
	m "api/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var (
	counter int
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(user, password, dbname string) {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	//http.HandleFunc("/favicon.ico", func(_ http.ResponseWriter, _ *http.Request) {})
	a.Router.HandleFunc("/", a.count)
	a.Router.HandleFunc("/stats", a.stats)
	a.Router.HandleFunc("/numbers/{number}", a.getNumbers).Methods(http.MethodGet, http.MethodPost)
}

func (a *App) count(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Invalid HTTP Method")
		return
	}
	counter++

	myMessage := m.MessageAPI{
		Message: "Hello DevFest Cordoba 2017",
	}
	respondWithJSON(w, http.StatusOK, myMessage)
}

func (a *App) stats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Invalid HTTP Method")
		return
	}
	myVisits := m.VisitAPI{
		Visits: counter,
	}

	respondWithJSON(w, http.StatusOK, myVisits)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) getNumbers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number, _ := strconv.Atoi(vars["number"])

	if number > 10 || number < 1 {
		number = 10
	}

	myNumber := m.NumberAPI{
		Number: number,
	}

	respondWithJSON(w, http.StatusOK, myNumber)
}
