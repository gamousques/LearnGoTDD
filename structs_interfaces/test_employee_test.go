package main

import "testing"

func TestEmployeeInterface(t *testing.T) {

	employeeTests := []struct {
		employee HiredPerson
		discount float64
		wantSalary float64
	} {
		{employee: Fulltime{name: "gastón", baseSalary: 100.0}, discount: 30, wantSalary: 70.0},
		{employee: Seasonal {name: "Maria", baseSalary: 100.0, hourlySalary: 20.0}, discount: 30, wantSalary: 84.0},
	}

	for _, tt := range employeeTests  {
		got := tt.employee.Salary(tt.discount)
		if got != tt.wantSalary {
			t.Errorf("got: %g want:%g employee:%v", got, tt.wantSalary, tt.employee)
		}
	}


}