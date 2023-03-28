package main

import (
	"fmt"
	"math"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
func PagIbig_Calc(salary float64) float64 {
	if salary <= 1500 {
		return toFixed(salary*0.01, 2)
	} else if (salary * 0.02) >= 100 {
		return 100
	} else {
		return toFixed(salary*0.02, 2)
	}
}
func PhilHealth_Calc(salary float64) float64 {
	if salary <= 10000 {
		return toFixed((450 / 2), 2)
	}
	if salary >= 10000.01 && salary <= 89999.99 {
		return toFixed(((salary * 0.045) / 2), 2)
	}
	return 4050
}
func SSS_Calc(salary float64) float64 {
	var baseSalary float64 = 202.5
	if salary < 4250 {
		return 180
	}
	for i := 0; i < 56; i++ {
		if salary >= 4250.0+float64(i)*500.0 && salary <= 4749.99+float64(i)*500.0 {
			return baseSalary + float64(i)*22.5
		}
	}
	return 1350
}

func incomeTax_Calc(salary float64) float64 {
	if salary <= 20833 {
		return 0
	}
	if salary >= 20833 && salary <= 33332 {
		return toFixed(((salary - 20733) * 0.15), 2)
	}
	if salary >= 33333 && salary <= 66666 {
		return toFixed((((salary - 33333) * 0.2) + 1875), 2)
	}
	if salary >= 66667 && salary <= 166666 {
		return toFixed((((salary - 66667) * 0.25) + 8541.8), 2)
	}
	if salary >= 166667 && salary <= 666666 {
		return toFixed((((salary - 166667) * 0.3) + 33541.8), 2)
	}

	return toFixed((((salary - 666667) * 0.35) + 183541.8), 2)

}
func main() {
	var grossSalary float64

	fmt.Print("Please enter a number: ")
	fmt.Scan(&grossSalary)

	var SSS_Value float64 = SSS_Calc(grossSalary)
	var philHealth_Value float64 = PhilHealth_Calc(grossSalary)
	var pagibig_Value float64 = PagIbig_Calc(grossSalary)
	var monthlyContribution_Value float64 = SSS_Value + philHealth_Value + pagibig_Value
	var incomeTax_Value = incomeTax_Calc(grossSalary - monthlyContribution_Value)

	fmt.Println("Pagibig: ", pagibig_Value)
	fmt.Println("PhilHealth: ", philHealth_Value)
	fmt.Println("SSS: ", SSS_Value)
	fmt.Println("Total: ", monthlyContribution_Value)
	fmt.Println("gross salary - monthly contribution: ", grossSalary-monthlyContribution_Value)
	fmt.Println("Income Tax: ", incomeTax_Value)
	fmt.Println("net salary: ", grossSalary-incomeTax_Value-monthlyContribution_Value)
}
