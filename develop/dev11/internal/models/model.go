package models

import "time"

type Event struct {
	ID   string `json:"user_id"`
	Name string `json:"name"`
	Date Date   `json:"date"`
}

type Date struct {
	time.Time
}

func Parse(t string) (Date, error) {
	timeReq, err := time.Parse("2006-01-02", t)
	return Date{timeReq}, err
}

func (d *Date) UnmarshalJSON(date []byte) error {
	if string(date) == "" || string(date) == "null" {
		*d = Date{time.Now()}
		return nil
	}

	tm, err := time.Parse(`"`+"2006-01-02"+`"`, string(date))
	*d = Date{tm}
	return err
}
