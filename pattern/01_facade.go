package pattern

//паттерн Фасад это структурный паттерн, который предоставляет простой (но урезанный) интерфейс 
//к сложной системе объектов, библиотеке или фреймворку.
//Кроме того, что Фасад позволяет снизить общую сложность программы, он также помогает вынести код, 
//зависимый от внешней системы в единственное место.

//Использование паттерна может решить эти проблемы:
//-Чтобы упростить использование сложной подсистемы, для набора 
//интерфейсов в подсистеме должен быть предоставлен простой интерфейс.
//-Зависимости от подсистемы должны быть сведены к минимуму.

//Использование паттерна описывает следующее решение:
//-реализует простой интерфейс с точки зрения (путем делегирования) интерфейсов в подсистеме и
//-может выполнять дополнительные функциональные возможности до / после пересылки запроса.

// Плюсы:
// -Простота использования: Позволяет клиентскому коду использовать сложную подсистему, 
//не беспокоясь о деталях ее внутренней реализации.
//-Снижение зависимостей: Уменьшает зависимость клиентского кода от сложной структуры подсистемы, 
//что делает код более гибким и легко поддерживаемым.
//-Улучшение структуры кода: Создает четкую иерархию, что улучшает структуру кода и делает его более понятным.

// Минусы:
// -Ограничение гибкости: Фасад может не предоставлять всех возможностей подсистемы, 
//что может ограничить гибкость клиентского кода в случае необходимости более сложных операций.
//-Добавление слоя абстракции: Некоторые разработчики считают, что фасад добавляет дополнительный слой абстракции, 
//что может усложнить систему для опытных разработчиков.

import "fmt"

// Подсистема A
type SubsystemA struct {
}

func (s *SubsystemA) OperationA1() {
	fmt.Println("Subsystem A, Operation A1")
}

func (s *SubsystemA) OperationA2() {
	fmt.Println("Subsystem A, Operation A2")
}

// Подсистема B
type SubsystemB struct {
}

func (s *SubsystemB) OperationB1() {
	fmt.Println("Subsystem B, Operation B1")
}

func (s *SubsystemB) OperationB2() {
	fmt.Println("Subsystem B, Operation B2")
}

// Фасад
type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
	}
}

// Методы фасада, предоставляющие унифицированный интерфейс для клиента
func (f *Facade) Operation() {
	fmt.Println("Facade operation:")
	f.subsystemA.OperationA1()
	f.subsystemA.OperationA2()
	f.subsystemB.OperationB1()
	f.subsystemB.OperationB2()
}