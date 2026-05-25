package main

import (
	"fmt"
	"math"
)

	const inflationRate = 6.5 


func main() {
	var investmentAmount float64
	var years float64	
	var expectedReturnRate float64
 

	//fmt.Print("Investment Amount: ")
	outputText("Investiment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	//fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)

	//fmt.Printf("Future Value: %.2f\n", futureValue)
	//fmt.Printf("Future Real Value: %.2f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64)(fv float64, rfv float64){
	//fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//rfv := fv / math.Pow(1 + inflationRate/100, years)
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1 + inflationRate/100, years)
	//return fv, rfv 
	return
}