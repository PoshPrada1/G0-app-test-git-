package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	for {
		fmt.Println("КАЛЬКУЛЯТОР IMT")
		userHeight, userWeight := getUserInput()
		IMT, err := calculateIMT(userWeight, userHeight)
		if err != nil {
			fmt.Println("Вы ввели неккоректные вес или рост")
			continue
		}
		outputResult(IMT)
		userCheck := stillDo()
		if !userCheck {
			break
		}
	}

}

func outputResult(imt float64) {
	result := fmt.Sprintf("ВАШ IMT: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("Вы дрыщ")
	case imt < 18.5:
		fmt.Println("Вы почти дрыщ")
	case imt < 25:
		fmt.Println("Вы нормис")
	case imt < 30:
		fmt.Println("Вы почти жирный")
	default:
		fmt.Println("Вы жирнич")
	}
}

func calculateIMT(userWeight, userHeight float64) (float64, error) {
	const IMTPower = 2
	if userWeight <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	IMT := userWeight / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Print("Введите рост: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите вес: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}

func stillDo() bool {
	var userChoise string
	fmt.Println("Продолжить? (y/n):")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
