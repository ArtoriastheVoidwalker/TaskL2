package main

import "fmt"

type Department interface { // Интерфейс обработчика определяет общий для всех конкретных обработчиков интерфейс.
	execute(*Patient)   // Обычно достаточно описать единственный метод обработки запросов,
	setNext(Department) // но иногда здесь может быть объявлен и метод выставления следующего обработчика.
}

type Reception struct { // Конкретный обработчик содержат код обработки запросов.
	next Department // При получении запроса каждый обработчик решает, может ли он обработать запрос, а также стоит ли передать его следующему объекту.
	// В большинстве случаев обработчики могут работать сами по себе и быть неизменяемыми, получив все нужные детали через параметры конструктора.
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}

type Doctor struct { // Конкретный обработчик
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}

type Medical struct { // Конкретный обработчик
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}

type Cashier struct { // Конкретный обработчик
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}

type Patient struct { // Базовый обработчик — опциональный класс, который позволяет избавиться от дублирования одного и того же кода во всех конкретных обработчиках.
	name              string // Обычно этот класс имеет поле для хранения ссылки на следующий обработчик в цепочке.
	registrationDone  bool   // Клиент связывает обработчики в цепь, подавая ссылку на следующий обработчик через конструктор или сеттер поля.
	doctorCheckUpDone bool   // Также здесь можно реализовать базовый метод обработки, который бы просто перенаправлял запрос следующему обработчику, проверив его наличие.
	medicineDone      bool
	paymentDone       bool
}

func main() {
	/*
		Клиент может либо сформировать цепочку обработчиков единожды, либо перестраивать её динамически, в зависимости от логики программы.
		Клиент может отправлять запросы любому из объектов цепочки, не обязательно первому из них.
	*/
	cashier := &Cashier{}

	medical := &Medical{}
	medical.setNext(cashier)

	doctor := &Doctor{}
	doctor.setNext(medical)

	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	reception.execute(patient)
}
