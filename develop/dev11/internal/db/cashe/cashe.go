package cashe

import (
	"dev11/internal/models"
	"fmt"
	"errors"
)

type Cashe struct {
	db map[string]map[models.Date][]string
}

func (c *Cashe) GetEvent(event models.Event, days string) (map[models.Date][]string, error) {
	err := c.checkUser(event.ID)
	if err != nil {
		return nil, err
	}
	answer := make(map[models.Date][]string)
	if days == "day"{
		answer[event.Date] = append(answer[event.Date], c.db[event.ID][event.Date]...)
	} else if days == "month"{
		for key, value := range c.db[event.ID]{
			if event.Date.Month() == key.Month(){
				answer[event.Date] = append(answer[event.Date], value...)
			}
		}
	} else{
		eventYear, eventWeek := event.Date.ISOWeek()
		for key, value := range c.db[event.ID]{
			keyYear, keyWeek := key.ISOWeek()
			if eventYear == keyYear && eventWeek == keyWeek{
				answer[event.Date] = append(answer[event.Date], value...)
			}
		}
	}
	return answer, nil
}

func (c *Cashe) CreateEvent(event models.Event) error {
	err := c.checkUser(event.ID)
	if err != nil {
		c.db[event.ID] = make(map[models.Date][]string)
	}
	c.db[event.ID][event.Date] = append(c.db[event.ID][event.Date], event.Name)
	return nil
}

func (c *Cashe) UpdateEvent(events []models.Event) error {
	if len(events) != 2 {
		return errors.New("there are too munch events to update, need 2")
	}
	for _, event := range events {
		err := c.checkUser(event.ID)
		if err != nil {
			return err
		}
	}
	err := c.DeleteEvent(events[0])
	if err != nil {
		return err
	}
	err = c.CreateEvent(events[1])
	if err != nil {
		return err
	}
	return fmt.Errorf("there is no such event in this day")
}

func (c *Cashe) DeleteEvent(event models.Event) error {
	err := c.checkUser(event.ID)
	if err != nil {
		return err
	}
	index := 0
	flag := false
	for num, value := range c.db[event.ID][event.Date] {
		if event.Name == value {
			index = num
			flag = true
			break
		}
	}
	if flag {
		c.db[event.ID][event.Date] = append(c.db[event.ID][event.Date][:index], c.db[event.ID][event.Date][index+1:]...)
		return nil
	}
	return fmt.Errorf("there is no such event in this day")
}

func (c *Cashe) checkUser(uid string) error {
	_, ok := c.db[uid]
	if !ok {
		return fmt.Errorf("there is no such user")
	}
	return nil
}

func NewCashe() *Cashe {
	return &Cashe{
		db: make(map[string]map[models.Date][]string),
	}
}
