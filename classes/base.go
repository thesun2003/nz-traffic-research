package BaseClass

type BaseInterface interface {
	Init(string)
}

type Coordinates struct {
	Lat, Long float64
}

type Base struct {
	Id       float64
	Name     string
	Position Coordinates
	Weight   float64
}

func (item *Base) generate_id() {
	item.Id = 35
}

func (item *Base) Init() {
	item.generate_id()
	item.Position.Lat = 0
	item.Position.Long = 0
	item.Weight = 0
	item.Name = "A car"
}

func (item *LightCar) Init(name string) {
	item.Base.Init()
	if len(name) > 0 {
		item.Name = name
	}
	item.Weight = 1000
}

func (item *HeavyCar) Init(name string) {
	item.Base.Init()
	if len(name) > 0 {
		item.Name = name
	}
	item.Weight = 2000
}

type LightCar struct {
	Base
}

type HeavyCar struct {
	Base
}
