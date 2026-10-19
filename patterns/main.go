package main

import "fmt"


// ПАТТЕРН ФАСАД


type CPU struct{}

func (c *CPU) Start()    { fmt.Println("CPU: запуск") }
func (c *CPU) Shutdown() { fmt.Println("CPU: выключение") }

type Memory struct{}

func (m *Memory) Load() { fmt.Println("Memory: загрузка данных") }

type Disk struct{}

func (d *Disk) Read() { fmt.Println("Disk: чтение с диска") }

type ComputerFacade struct {
	cpu    *CPU
	memory *Memory
	disk   *Disk
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{cpu: &CPU{}, memory: &Memory{}, disk: &Disk{}}
}

func (c *ComputerFacade) Start() {
	fmt.Println("Запуск компьютера")
	c.cpu.Start()
	c.memory.Load()
	c.disk.Read()
}

func (c *ComputerFacade) Shutdown() {
	fmt.Println("Выключение компьютера")
	c.cpu.Shutdown()
}

func demoFacade() {
	pc := NewComputerFacade()
	pc.Start()
	pc.Shutdown()
}

// ПАТТЕРН СТРОИТЕЛЬ

type Pizza struct {
	dough     string
	sauce     string
	cheese    bool
	pepperoni bool
}

type PizzaBuilder struct {
	pizza Pizza
}

func NewPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{}
}

func (b *PizzaBuilder) SetDough(d string) *PizzaBuilder {
	b.pizza.dough = d
	return b
}

func (b *PizzaBuilder) SetSauce(s string) *PizzaBuilder {
	b.pizza.sauce = s
	return b
}

func (b *PizzaBuilder) AddCheese() *PizzaBuilder {
	b.pizza.cheese = true
	return b
}

func (b *PizzaBuilder) AddPepperoni() *PizzaBuilder {
	b.pizza.pepperoni = true
	return b
}

func (b *PizzaBuilder) Build() Pizza {
	return b.pizza
}

func demoBuilder() {
	pizza := NewPizzaBuilder().
		SetDough("тонкое").
		SetSauce("томатный").
		AddCheese().
		AddPepperoni().
		Build()

	fmt.Printf("Пицца: %+v\n", pizza)
}

// ПАТТЕРН ПОСЕТИТЕЛЬ

type Shape interface {
	Accept(v Visitor)
}

type Circle struct{ Radius float64 }
type Square struct{ Side float64 }

func (c *Circle) Accept(v Visitor) { v.VisitCircle(c) }
func (s *Square) Accept(v Visitor) { v.VisitSquare(s) }

type Visitor interface {
	VisitCircle(c *Circle)
	VisitSquare(s *Square)
}

type AreaVisitor struct{ total float64 }

func (a *AreaVisitor) VisitCircle(c *Circle) {
	area := 3.14 * c.Radius * c.Radius
	fmt.Printf("Площадь круга: %.2f\n", area)
	a.total += area
}

func (a *AreaVisitor) VisitSquare(s *Square) {
	area := s.Side * s.Side
	fmt.Printf("Площадь квадрата: %.2f\n", area)
	a.total += area
}

func demoVisitor() {
	shapes := []Shape{&Circle{Radius: 5}, &Square{Side: 4}}
	v := &AreaVisitor{}
	for _, s := range shapes {
		s.Accept(v)
	}
	fmt.Printf("Общая площадь: %.2f\n", v.total)
}

// ПАТТЕРН КОМАНДА

type Command interface {
	Execute()
	Undo()
}

type Light struct{ on bool }

func (l *Light) TurnOn()  { l.on = true; fmt.Println("Свет включён") }
func (l *Light) TurnOff() { l.on = false; fmt.Println("Свет выключен") }

type TurnOnCommand struct{ light *Light }
type TurnOffCommand struct{ light *Light }

func (c *TurnOnCommand) Execute() { c.light.TurnOn() }
func (c *TurnOnCommand) Undo()    { c.light.TurnOff() }

func (c *TurnOffCommand) Execute() { c.light.TurnOff() }
func (c *TurnOffCommand) Undo()    { c.light.TurnOn() }

type RemoteControl struct {
	history []Command
}

func (r *RemoteControl) Press(c Command) {
	c.Execute()
	r.history = append(r.history, c)
}

func (r *RemoteControl) UndoLast() {
	if len(r.history) == 0 {
		fmt.Println("Нечего отменять")
		return
	}
	last := r.history[len(r.history)-1]
	r.history = r.history[:len(r.history)-1]
	last.Undo()
}

func demoCommand() {
	light := &Light{}
	remote := &RemoteControl{}

	remote.Press(&TurnOnCommand{light: light})
	remote.Press(&TurnOffCommand{light: light})
	remote.UndoLast()
}

// ПАТТЕРН ЦЕПОЧКА ВЫЗОВОВ

type Handler interface {
	SetNext(h Handler) Handler
	Handle(request string)
}

type BaseHandler struct{ next Handler }

func (b *BaseHandler) SetNext(h Handler) Handler {
	b.next = h
	return h
}

func (b *BaseHandler) PassToNext(request string) {
	if b.next != nil {
		b.next.Handle(request)
	}
}

type AuthHandler struct{ BaseHandler }

func (h *AuthHandler) Handle(req string) {
	if req == "login" {
		fmt.Println("Auth: пользователь авторизован")
	}
	h.PassToNext(req)
}

type LogHandler struct{ BaseHandler }

func (h *LogHandler) Handle(req string) {
	fmt.Printf("Log: получен запрос '%s'\n", req)
	h.PassToNext(req)
}

type DataHandler struct{ BaseHandler }

func (h *DataHandler) Handle(req string) {
	fmt.Println("Data: обработка запроса")
}

func demoChain() {
	auth := &AuthHandler{}
	log := &LogHandler{}
	data := &DataHandler{}

	auth.SetNext(log).SetNext(data)
	auth.Handle("login")
}

// ПАТТЕРН ФАБРИЧНЫЙ МЕТОД

type Transport interface {
	Deliver() string
}

type Truck struct{}

func (t *Truck) Deliver() string { return "Доставка грузовиком по земле" }

type Ship struct{}

func (s *Ship) Deliver() string { return "Доставка кораблём по морю" }

type Plane struct{}

func (p *Plane) Deliver() string { return "Доставка самолётом по воздуху" }

func NewTransport(kind string) (Transport, error) {
	switch kind {
	case "truck":
		return &Truck{}, nil
	case "ship":
		return &Ship{}, nil
	case "plane":
		return &Plane{}, nil
	default:
		return nil, fmt.Errorf("неизвестный тип транспорта: %s", kind)
	}
}

func demoFactory() {
	for _, kind := range []string{"truck", "ship", "plane"} {
		t, err := NewTransport(kind)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		fmt.Println(t.Deliver())
	}
}

// ПАТТЕРН СТРАТЕГИЯ

type PaymentStrategy interface {
	Pay(amount float64)
}

type CardPayment struct{}

func (c *CardPayment) Pay(amount float64) {
	fmt.Printf("Оплата картой: %.2f руб.\n", amount)
}

type CashPayment struct{}

func (c *CashPayment) Pay(amount float64) {
	fmt.Printf("Оплата наличными: %.2f руб.\n", amount)
}

type CryptoPayment struct{}

func (c *CryptoPayment) Pay(amount float64) {
	fmt.Printf("Оплата криптой: %.2f руб.\n", amount)
}

type Shop struct {
	strategy PaymentStrategy
}

func (s *Shop) SetStrategy(p PaymentStrategy) {
	s.strategy = p
}

func (s *Shop) Checkout(amount float64) {
	s.strategy.Pay(amount)
}

func demoStrategy() {
	shop := &Shop{}
	shop.SetStrategy(&CardPayment{})
	shop.Checkout(1500)

	shop.SetStrategy(&CryptoPayment{})
	shop.Checkout(3000)
}

// ПАТТЕРН СОСТОЯНИЕ


type State interface {
	Next(*Order)
	PrintStatus()
}

type Order struct {
	state State
}

func (o *Order) SetState(s State) { o.state = s }
func (o *Order) Next()            { o.state.Next(o) }
func (o *Order) Print()           { o.state.PrintStatus() }

type CreatedState struct{}

func (s *CreatedState) Next(o *Order) { o.SetState(&PaidState{}) }
func (s *CreatedState) PrintStatus()  { fmt.Println("Заказ создан") }

type PaidState struct{}

func (s *PaidState) Next(o *Order) { o.SetState(&ShippedState{}) }
func (s *PaidState) PrintStatus()  { fmt.Println("Заказ оплачен") }

type ShippedState struct{}

func (s *ShippedState) Next(o *Order) { o.SetState(&DeliveredState{}) }
func (s *ShippedState) PrintStatus()  { fmt.Println("Заказ отправлен") }

type DeliveredState struct{}

func (s *DeliveredState) Next(o *Order) { fmt.Println("Заказ уже доставлен") }
func (s *DeliveredState) PrintStatus()  { fmt.Println("Заказ доставлен") }

func demoState() {
	order := &Order{state: &CreatedState{}}
	order.Print()
	order.Next()
	order.Print()
	order.Next()
	order.Print()
	order.Next()
	order.Print()
}

// ТОЧКА ВХОДА

func main() {
	fmt.Println("\n Фасад ")
	demoFacade()

	fmt.Println("\n Строитель ")
	demoBuilder()

	fmt.Println("\n Посетитель ")
	demoVisitor()

	fmt.Println("\n Команда ")
	demoCommand()

	fmt.Println("\n Цепочка вызовов ")
	demoChain()

	fmt.Println("\n Фабричный метод ")
	demoFactory()

	fmt.Println("\n Стратегия ")
	demoStrategy()

	fmt.Println("\n Состояние ")
	demoState()
}