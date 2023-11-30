package api

import (
	"dev11/internal/models"
	"encoding/json"
	"errors"
	"net/http"
)

func (a API) parsePostMethod(r *http.Request) ([]models.Event, error) {
	if r.Method != http.MethodPost {
		return nil, errors.New("invalid method: " + r.Method)
	}
	return a.parseJson(r)
}

func (a API) parseGetMethod(r *http.Request) (*models.Event, error) {
	if r.Method != http.MethodGet {
		return nil, errors.New("invalid method: " + r.Method)
	}

	date, err := models.Parse(r.URL.Query().Get("date"))
	if err != nil {
		return nil, err
	}
	
	m := models.Event{
		ID:   r.URL.Query().Get("user_id"),
		Date: date,
	}
	return &m, nil
}

func (a API) parseJson(r *http.Request) ([]models.Event, error) {
	events := []models.Event{}
	dec := json.NewDecoder(r.Body)
	for dec.More() {
		var event models.Event
		err := dec.Decode(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func makeJson(res map[models.Date][]string) ([]byte, error){
	result, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return result, nil
}
