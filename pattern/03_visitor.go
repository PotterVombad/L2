package pattern

import "fmt"

//паттерн Посетитель это поведенческий паттерн,
//который позволяет добавить новую операцию для целой иерархии классов,
//не изменяя код этих классов.

// Перенос операций в классы посетителей выгоден, когда
// -требуется множество несвязанных операций со структурой объекта,
// -классы, составляющие структуру объекта, известны, и ожидается,
//что они не изменятся,
// -необходимо часто добавлять новые операции,
// -алгоритм включает в себя несколько классов структуры объекта,
//но желательно управлять им в одном месте,
// -алгоритм должен работать в нескольких независимых иерархиях классов.

// Плюсы:
// -Открытость/закрытость (Open/Closed Principle): Паттерн соответствует принципу
// открытости/закрытости, так как позволяет добавлять новые функции,
//не модифицируя существующий код.
// -Увеличение гибкости: Позволяет легко добавлять новые функции к существующей
//иерархии классов, не нарушая ее структуру.
// -Централизация операций: Операции, выполняемые над объектами, централизованы
//в одном месте (в посетителе), что упрощает поддержку и изменение кода.

// Минусы:
// -Усложнение кода: Добавляет дополнительные интерфейсы и классы,
//что может усложнить код и структуру проекта.
// -Нарушение инкапсуляции: Посетитель должен знать о структуре объектов,
//что может нарушить инкапсуляцию и увеличить зависимость между посетителем и
//элементами.
// -Трудность добавления новых классов: Добавление новых классов требует изменения интерфейса
//посетителя и всех его конкретных реализаций.

// Интерфейс элемента, который будет посещен
type Element interface {
	Accept(visitor Visitor)
}

// Конкретный элемент A
type ElementA struct{}

func (e *ElementA) Accept(visitor Visitor) {
	visitor.VisitElementA(e)
}

func (e *ElementA) getType() string {
    return "ElementA"
}

// Конкретный элемент B
type ElementB struct{}

func (e *ElementB) Accept(visitor Visitor) {
	visitor.VisitElementB(e)
}

func (e *ElementB) getType() string {
    return "ElementB"
}

// Интерфейс посетителя
type Visitor interface {
	VisitElementA(element *ElementA)
	VisitElementB(element *ElementB)
}

// Конкретный посетитель
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitElementA(element *ElementA) {
	fmt.Println("Visit ElementA")
}

func (v *ConcreteVisitor) VisitElementB(element *ElementB) {
	fmt.Println("Visit ElementB")
}