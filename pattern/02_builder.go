package main

import "fmt"

/*
	Builder - порождающий паттерн,позволяющий создавать сложные объекты пошагово.
	Builder даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.
	Плюсы:
	1) Позволяет создавать продукты пошагово.
	2) Позволяет использовать один и тот же код для создания различных продуктов.
	3) Изолирует сложный код сборки продукта от его основной бизнес-логики.
	Минусы:
	1) Усложняет код программы из-за введения дополнительных классов.
	2) Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.
*/

type IBuilder interface { // Интерфейс IBuilder объявляет шаги конструирования, общие для всех видов билдеров.
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IBuilder { // Проверка типа билдера. Конкретные билдеры реализуют строительные шаги, каждый по-своему.
	// Конкретные билдеры могут производить разнородные объекты(в данном случае иглу и обычноый дом), не имеющие общего интерфейса.
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}

// Билдер для производства продуктов первого типа
type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

// Пошаговая сборка частей объекта первого типа
func (b *NormalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

// Сборка продукта первого типа
func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// Билдер для производства продуктов второго типа

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

// Пошаговая сборка частей объекта второго типа
func (b *IglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

// Сборка продукта второго типа
func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

type House struct {
	windowType string
	doorType   string
	floor      int
}

// Директор определяет порядок вызова отдельных шагов для сборки продуктов
type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func main() {
	normalBuilder := getBuilder("normal") // Создание билдера для первого продукта
	iglooBuilder := getBuilder("igloo")   // Создание билдера для второго продукта

	director := newDirector(normalBuilder) // Создание директора для второго продукта
	normalHouse := director.buildHouse()   // Определение порядка выполнения шагов

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}
