package main

import "fmt"

/*
	Visitor поведенческий паттерн, позволяет добавлять в программу новые операции,
	не изменяя классы объектов, над которыми эти операции могут выполняться.
*/

type Shape interface { // Элемент описывает метод принятия посетителя.
	getType() string
	accept(Visitor) // Этот метод должен иметь единственный параметр, объявленный с типом интерфейса посетителя.
}

type Square struct { // Конкретные элементы реализуют методы принятия посетителя.
	side int // Цель этого метода — вызвать тот метод посещения, который соответствует типу этого элемента.
	// Так посетитель узнает, с каким именно элементом он работает.
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

type Circle struct { // Конкретный элемент
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

type Rectangle struct { // Конкретный элемент
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	v.visitForrectangle(t)
}

func (t *Rectangle) getType() string {
	return "rectangle"
}

type Visitor interface { // Посетитель описывает общий интерфейс для всех типов посетителей.
	visitForSquare(*Square)       // Он объявляет набор методов, отличающихся типом входящего параметра,
	visitForCircle(*Circle)       // которые нужны для запуска операции для всех типов конкретных элементов.
	visitForrectangle(*Rectangle) // В языках, поддерживающих перегрузку методов,
	//эти методы могут иметь одинаковые имена, но типы их параметров должны отличаться.
}

type AreaCalculator struct { // Конкретный посетитель
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
}
func (a *AreaCalculator) visitForrectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}

type MiddleCoordinates struct { // Конкретные посетители реализуют особенное поведение для всех типов элементов,
	x int // которые можно подать через методы интерфейса посетителя.
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *MiddleCoordinates) visitForrectangle(t *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}

func main() { //  Клиентский код
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
