package main

type HiredPerson interface {
	Salary(discount float64) float64
}
type Fulltime struct {
	name string
	baseSalary float64
}

type Seasonal struct {
	name string
	baseSalary float64
	hourlySalary float64
}

func (f Fulltime) Salary(discount float64) float64 {
	return f.baseSalary - (f.baseSalary * discount)/100;
}

func (s Seasonal) Salary(discount float64) float64 {
	return (s.baseSalary + s.hourlySalary) - ((s.baseSalary + s.hourlySalary)* discount)/100;
}

