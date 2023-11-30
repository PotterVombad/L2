package pattern


// паттерн Строитель это порождающий паттерн проектирования, 
//который позволяет создавать объекты пошагово.
//В отличие от других порождающих паттернов, 
//Строитель позволяет производить различные продукты, 
//используя один и тот же процесс строительства.

// Достоинства:
// Позволяет изменять внутреннее представление продукта.
// Инкапсулирует код для построения и представления.
// Обеспечивает контроль над этапами процесса строительства.

// Недостатки:
// Для каждого типа изделия должен быть создан отдельный конструктор.
// Классы Builder должны быть изменяемыми.
// Может затруднять внедрение зависимостей.


// Builder - интерфейс строителя
type Builder interface {
	BuildPart1()
	BuildPart2()
	BuildPart3()
	GetResult() House
}

// Product - конечный продукт, который мы строим
type House struct {
	Part1 string
	Part2 string
	Part3 string
}

// ConcreteBuilder - конкретная реализация строителя
type ConcreteBuilder struct {
	house House
}

func (b *ConcreteBuilder) BuildPart1() {
	b.house.Part1 = "Part1 built"
}

func (b *ConcreteBuilder) BuildPart2() {
	b.house.Part2 = "Part2 built"
}

func (b *ConcreteBuilder) BuildPart3() {
	b.house.Part3 = "Part3 built"
}

func (b *ConcreteBuilder) GetResult() House {
	return b.house
}

// директор управляет процессом построения
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() House {
	d.builder.BuildPart1()
	d.builder.BuildPart2()
	d.builder.BuildPart3()
	return d.builder.GetResult()
}


