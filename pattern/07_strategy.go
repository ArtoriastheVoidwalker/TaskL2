package main

import "fmt"

/*
	Стратегия — это поведенческий паттерн, выносит набор алгоритмов в собственные классы и делает их взаимозаменимыми.
	Применимость:
	1) Когда необходимо использовать разные вариации какого-то алгоритма внутри одного объекта.
	2) Когда у есть множество похожих классов, отличающихся только некоторым поведением.
	3) Когда вы не хотите обнажать детали реализации алгоритмов для других классов.
	Плюсы:
	1) Горячая замена алгоритмов на лету.
	2) Изолирует код и данные алгоритмов от остальных классов.
	3) Уход от наследования к делегированию.
	Минусы:
	1) Усложняет программу за счёт дополнительных классов.
	2) Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

  Паттерн реализован на примере разработки «In-Memory-Cache». Поскольку он находится внутри памяти, его размер ограничен.
	Как только он полностью заполнится, какие-то записи придется убрать для освобождения места.
	Эту функцию можно реализовать с помощью нескольких алгоритмов, самые популярные среди них:
   Наиболее давно использовавшиеся (Least Recently Used – LRU): убирает запись, которая использовалась наиболее давно.
   «Первым пришел, первым ушел» (First In, First Out — FIFO): убирает запись, которая была создана раньше остальных
   Наименее часто использовавшиеся (Least Frequently Used — LFU): убирает запись, которая использовалась наименее часто.
*/

type EvictionAlgo interface { // Интерфейс стратегии определяет интерфейс, общий для всех вариаций алгоритма.
	evict(c *Cache) // Контекст использует этот интерфейс для вызова алгоритма.
	//Для контекста неважно, какая именно вариация алгоритма будет выбрана, так как все они имеют одинаковый интерфейс.
}

type Fifo struct { // Конкретная стратегия реализует одну из вариаций алгоритма.
}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}

type Lru struct { // Конкретная стратегия реализует одну из вариаций алгоритма.
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}

type Lfu struct { // Конкретная стратегия реализует одну из вариаций алгоритма.
}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strtegy")
}

type Cache struct { // Контекст хранит ссылку на объект конкретной стратегии, работая с ним через общий интерфейс стратегий.
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

// Во время выполнения программы контекст получает вызовы от клиента и делегирует их объекту конкретной стратегии.

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")

}
