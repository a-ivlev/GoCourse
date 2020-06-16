package getspeed

type Status interface {
	Speed() int
}

func GetSpeed(s Status) int {
	return s.Speed()
}

type Auto struct {
	Marka string
	Way   int
	Time  int
}

type Human struct {
	Name  string
	Age   int
	Phone int
	Way   int
	Time  int
}

func (a Auto) Speed() int {
	return a.Way / a.Time
}

func (h Human) Speed() int {
	return h.Way / h.Time
}
