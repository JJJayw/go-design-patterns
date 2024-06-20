package lib

type Gun struct {
	name  string
	power int
}

var _ IGun = (*Gun)(nil)

func (g *Gun) SetName(name string) {
	g.name = name
}

func (g *Gun) GetName() string {
	return g.name
}

func (g *Gun) SetPower(power int) {
	g.power = power
}

func (g *Gun) GetPower() int {
	return g.power
}
