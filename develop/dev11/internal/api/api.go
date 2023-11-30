package api

import (
	"dev11/internal/db"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type API struct {
	db db.Storage
}

func (a API) createEvent(w http.ResponseWriter, r *http.Request) {
	events, err := a.parsePostMethod(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "error method", http.StatusBadRequest)
		return
	}
	for _, event := range events {
		err = a.db.CreateEvent(event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (a API) updateEvent(w http.ResponseWriter, r *http.Request) {
	events, err := a.parsePostMethod(r)
	if err != nil {
		http.Error(w, "error method", http.StatusBadRequest)
		return
	}
	err = a.db.UpdateEvent(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a API) deleteEvent(w http.ResponseWriter, r *http.Request) {
	events, err := a.parsePostMethod(r)
	if err != nil {
		http.Error(w, "error method", http.StatusBadRequest)
		return
	}
	for _, event := range events {
		err = a.db.DeleteEvent(event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}
	}
}

func (a API) getDayEvents(w http.ResponseWriter, r *http.Request) {
	event, err := a.parseGetMethod(r)
	if err != nil {
		http.Error(w, "error method", http.StatusBadRequest)
		return
	}
	events, err := a.db.GetEvent(*event, "day")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	answer, err := makeJson(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	w.Write(answer)
}

func (a API) getWeekEvents(w http.ResponseWriter, r *http.Request) {
	event, err := a.parseGetMethod(r)
	if err != nil {
		http.Error(w, "error method", http.StatusBadRequest)
		return
	}
	events, err := a.db.GetEvent(*event, "week")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	answer, err := makeJson(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	w.Write(answer)
}

func (a API) getMonthEvents(w http.ResponseWriter, r *http.Request) {
	event, err := a.parseGetMethod(r)
	if err != nil {
		http.Error(w, "error method", http.StatusBadRequest)
		return
	}
	events, err := a.db.GetEvent(*event, "month")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	answer, err := makeJson(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	w.Write(answer)
}

func loggingMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request: Method " + r.Method + ", Url " + r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

func (a *API) NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/create_event", loggingMiddleware(http.HandlerFunc(a.createEvent)))
	router.HandleFunc("/update_event", loggingMiddleware(http.HandlerFunc(a.updateEvent)))
	router.HandleFunc("/delete_event", loggingMiddleware(http.HandlerFunc(a.deleteEvent)))
	router.HandleFunc("/events_for_day", loggingMiddleware(http.HandlerFunc(a.getDayEvents)))
	router.HandleFunc("/events_for_week", loggingMiddleware(http.HandlerFunc(a.getWeekEvents)))
	router.HandleFunc("/events_for_month", loggingMiddleware(http.HandlerFunc(a.getMonthEvents)))

	return router
}

func (a API) Run(host string, port string) error {

	server := &http.Server{
		Addr:         host + ":" + port,
		Handler:      a.NewRouter(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Println("Run server, host:", host, "port:", port)
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func New(db db.Storage) API {
	return API{
		db: db,
	}
}
