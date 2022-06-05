package main

import "fmt"

func main() {
	amount, err := paintNeed(1.6, -2.6)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%.2f litres needed \n", amount)

}
func paintNeed(width float64, height float64) (float64, error) {
	if width < 0 {
		// we are returning 0 here becoz to exit the function if width <0
		// it will not further check the next condition
		return 0, fmt.Errorf("%.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("%.2f is invalid", height)
	}
	area := width * height
	return area, nil // here "nil" indicates no error in function
}
