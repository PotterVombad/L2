package pattern

import (
	"fmt"
	"time"
)

/*
фасад - упростить инфтерфейс для того чтобы пользователь, разработчик которой
работает с системой, скрытие бизнес логики, пользователю миимум функционала
+
изолирует клиентов от системы
-
сам интерфейс может стать суперклассом по итогу все последующие функции будут проходитб через него
*/

type Product struct {
	Name string
	Cost float32
}

type Shop struct {
	Name string
	Products []Product
}

type Bank struct {
	Name string
	Cards []Card
}

type Card struct {
	Name string
	Bank *Bank
	Balance float32
}

type User struct {
	Name string
	Card *Card
}

func (u *User) GetBalance() float32 { 
	return u.Card.Balance
}

func(c *Card) CheckBalance() err {
	
}

func (s *Shop) Sell(user User, product Product) error {
	fmt.Println("запрос к карте")
	time.Sleep(time.Second * 5)

	return nil
}