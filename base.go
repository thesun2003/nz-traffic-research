package main
import "fmt"

type Coordinates struct {
	x, y float64
}

type Base struct {
	id       float64
	name     string
	position Coordinates
	weight   float64
}

func (item *Base) generate_id() {
	item.id = 35
}

func (item *Base) __init() {
	item.generate_id()
	item.position.x = 0
	item.position.y = 0
	item.weight = 0
	item.name = "A car"
}

func (item *LightCar) __init(name string) {
	item.Base.__init()
	if len(name) > 0 {
		item.name = name
	}
	item.weight = 1000
}

func (item *HeavyCar) __init(name string) {
	item.Base.__init()
	if len(name) > 0 {
		item.name = name
	}
	item.weight = 2000
}

type LightCar struct {
	Base
}

type HeavyCar struct {
	Base
}

func main() {
	car1 := LightCar{}
	car2 := HeavyCar{}

	car1.__init("Toyota Starlet")
	car2.__init("Mazda Axela")

    fmt.Println(car1, car2)
}