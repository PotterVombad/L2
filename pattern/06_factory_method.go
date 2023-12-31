package pattern

//паттерн Фабричный метод это порождающий паттерн проектирования, который решает проблему 
//создания различных продуктов, без указания конкретных классов продуктов.
//Фабричный метод задаёт метод, который следует использовать вместо вызова оператора new для 
//создания объектов-продуктов. Подклассы могут переопределить этот метод, 
//чтобы изменять тип создаваемых продуктов.

//Использование паттерна может решить эти проблемы:
//-Как создать объект, чтобы подклассы могли переопределять, какой класс инстанцировать?
//-Как класс может отложить инстанцирование для подклассов?

//Использование паттерна описывает следующее решение:
//-Определите отдельную операцию (метод фабрики) для создания объекта.
//-Создайте объект, вызвав метод фабрики.

// Плюсы:
// -Расширяемость: Легко добавлять новые подклассы, реализующие фабричный метод, 
//и тем самым расширять функциональность системы.
//-Устранение зависимости от конкретных классов: Клиентский код зависит от интерфейса 
//фабричного метода, а не от конкретных классов, что улучшает гибкость и поддерживаемость кода.
//-Отделение создания объектов от их использования: Клиентский код может использовать объекты, 
//не заботясь о том, как они создаются.

// Минусы:
// -Усложнение кода: Внедрение фабричного метода может усложнить структуру кода, 
//особенно если у вас нет необходимости в создании новых подклассов.
//-Неявность взаимосвязей: Иногда может быть трудно понять, 
//какой фабричный метод вызывается и какие именно объекты создаются, особенно в сложных иерархиях классов.


// Интерфейс продукта
type ProductInterface interface {
	Use() string
}

// Конкретный продукт A
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Product A used"
}

// Конкретный продукт B
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Product B used"
}

// Интерфейс создания продукта
type Creator interface {
	CreateProduct() ProductInterface
}

// Конкретный создатель A
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) CreateProduct() ProductInterface {
	return &ConcreteProductA{}
}

// Конкретный создатель B
type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) CreateProduct() ProductInterface {
	return &ConcreteProductB{}
}