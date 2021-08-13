package main

import "fmt"

//1. A good credit score is a score of 450 or above.
//2. For a good credit score, your interest rate is 15%.
//3. For a less than good score, your interest rate is 20%.
//4. For a good credit score, your monthly payment must be no more than 20% of your
//monthly income.
//5. For a less than good credit score, your monthly payment must be no more than 10%
//of your monthly income.
//6. If a credit score, monthly income, loan amount, or loan term is less than 0, return
//an error.
//7. If the term of the loan if not divisible by 12 months, return an error.
//8. The interest payment will be a simple calculation of loan amount * interest rate *
//loan term.

var (
	CreditScore    = "Credit Score: "
	Income         = "Income: "
	LoanAmount     = "Loan Amount: "
	MonthlyPayment = " Monthly Payment: "
	Rate           = "Rate: "
	TotalCost      = " total: "
	Approved       = false
)

func main() {
	loanCal(500, 1000, 1000, 24, 5)

}

func loanCal(creditScore, monthlyIncome, loanAmount, loanTerm, interestRate int) {
	var monthlyPayment int
	var totalCost int

	//p := monthlyIncome
	//t := loanTerm
	//i := interestRate
	// for a good credit Score
	goodCreditScoreMonthlyIncomePercent := 15 / 100 * monthlyIncome
	fitResult := goodCreditScoreMonthlyIncomePercent / 100
	finaResult1 := fitResult * interestRate
	goodCreditInterestRate := (12) * finaResult1 / 12

	//good CreditScore for monthlyPayment
	percentGoodCreditScore := 20 / 100 * monthlyIncome

	//less than a good credit score

	lessThanGoodCreditScoreMonthlyIncomePercent := 10 / 100 * monthlyIncome
	fit1 := interestRate / 100
	fitLessthanResult := fit1 * interestRate
	lessThanInterestRate := (12) * fitLessthanResult / 12

	//less than  a good credit Score
	percentLessthanGoodCredit := 10 / 100 * monthlyIncome

	//MonthlyPayment cal
	monthlyPayment = monthlyIncome * lessThanInterestRate

	switch {

	case creditScore >= 450 && interestRate == goodCreditInterestRate && loanTerm%12 == 0 && percentGoodCreditScore <= monthlyPayment:
		application1()
		totalCost = loanAmount * interestRate * loanTerm
		Approved = true
		fmt.Println(CreditScore, creditScore, Income, monthlyIncome, LoanAmount, loanAmount, MonthlyPayment, monthlyPayment, Rate, interestRate, totalCost, Approved)
	case 0 <= creditScore || 0 <= interestRate || 0 <= monthlyIncome ||  0 <= loanAmount:
		fmt.Println("Error input does not Match!!")
	case loanTerm%12 != 0:
		fmt.Println("Error!!")

	case 450 < creditScore && interestRate == lessThanInterestRate && lessThanGoodCreditScoreMonthlyIncomePercent <= monthlyIncome && loanTerm%12 == 0 && percentLessthanGoodCredit <= monthlyPayment:

		totalCost = loanAmount * interestRate * loanTerm
		application1()
		Approved = false
		fmt.Println(CreditScore, creditScore, Income, monthlyIncome, LoanAmount, loanAmount, MonthlyPayment, monthlyPayment, Rate, interestRate, totalCost, Approved)
	default:
		fmt.Println("Error! Match not found Exit...")
	}

}

func application1() {

	fmt.Println("Application 1")
	fmt.Println("----------------")

}

