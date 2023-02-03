package main

import "fmt"

/*
	Фабричный метод — это порождающий паттерн проектирования,
	который определяет общий интерфейс для создания объектов в суперклассе,
	позволяя подклассам изменять тип создаваемых объектов.
*/

type IGun interface { // Интерфейс продукта определяет общий интерфейс объектов, которые может произвести создатель и его подклассы.
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct { // Конкретный продукт содержат код различных продуктов. Продукты будут отличаться реализацией, но интерфейс у них будет общий.
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

type Ak47 struct { // Конкретный продукт по-своему реализуют фабричный метод, производя те или иные конкретные продукты.
	Gun // Фабричный метод не обязан всё время создавать новые объекты.
	// Его можно переписать так, чтобы возвращать существующие объекты из какого-то хранилища или кэша.
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type musket struct { // Конкретный продукт
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

/*
	Объявляет фабричный метод, который должен возвращать новые объекты продуктов. Важно, чтобы тип результата совпадал с общим интерфейсом продуктов.
Зачастую фабричный метод объявляют абстрактным, чтобы заставить все подклассы реализовать его по-своему.
Но он может возвращать и некий стандартный продукт.
Несмотря на название, важно понимать, что создание продуктов не является единственной функцией создателя.
Обычно он содержит и другой полезный код работы с продуктом. Аналогия: большая софтверная компания может иметь центр подготовки программистов,
но основная задача компании — создавать программные продукты, а не готовить программистов.
*/
func getGun(gunType string) (IGun, error) {

	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
