package main

import (
	"errors"
	"fmt"
	"fs-helper/utils"
	"math"
	"os"
	"strconv"
)

const balanceFileName = "balance.txt"

func readOperationFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 0.0, errors.New("Couldn't read the file")
	}

	balance, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 0.0, errors.New("Could't parse data correctly")
	}

	return balance, err
}

func writeDataToFS(testingData string) {
	wantedDataType := []byte(testingData);
	fmt.Print(wantedDataType)
	os.WriteFile("data.txt", wantedDataType, 0644)
}

func main() {
	const inflationRate = 2.5
	var investmentRate, years float64
	var expectedRate float64

	utils.MainLogger("Data printed")
	
	investmentRate = requestInput("Enter investment rate : ")
	years = requestInput("Enter years : ")
	expectedRate = requestInput("Expected rate : ")

	futureValue, realValue := calculateInvestments(investmentRate, years, expectedRate, inflationRate)

	fmt.Println("Future value : ", futureValue)
	fmt.Println("Real value : ", realValue)
}

func calculateInvestments(investmentRate, years, expectedRate, inflationRate float64) (float64, float64) {
	ir := investmentRate * math.Pow(1 + (expectedRate / 100), years)
	fv := ir / math.Pow(1 + (inflationRate / 100), years)

	return ir, fv
}

func requestInput(message string) float64 {
	var enteredValue float64
	fmt.Print(message);
	fmt.Scan(&enteredValue)
	return enteredValue
}

func outputText(text string) string {
	formattedString := fmt.Sprintf("<|%s|>", text);
	return formattedString
}