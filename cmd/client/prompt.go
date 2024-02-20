package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

// transaction currency prompt
func PromptCurrency() string {
	prompt := promptui.Select{
		Label: "Select transaction currency",
		Items: []string{
			"INR", "USD",
		},
	}
	_, currency, _ := prompt.Run()
	return currency
}

// transaction amount prompt
func PromptAmount(currency string) int64 {
	validate := func(input string) error {
		_, err := strconv.ParseInt(input, 10, 64)
		return err
	}
	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("Amount (in %s)", currency),
		Validate: validate,
	}
	for {
		result, err := prompt.Run()
		if err != nil {
			fmt.Println("invalid amount value!")
			continue
		}
		amount, _ := strconv.ParseInt(result, 10, 64)
		return amount
	}
}

// service availed prompt
func PromptService() string {
	validate := func(input string) error {
		if strings.TrimSpace(input) == "" {
			return fmt.Errorf("service name is required")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Which service was availed?",
		Validate: validate,
	}
	for {
		result, err := prompt.Run()
		if err != nil {
			fmt.Println(err)
			continue
		}
		return result
	}
}

// person name prompt
func PromptPerson() string {
	validate := func(input string) error {
		if strings.TrimSpace(input) == "" {
			return fmt.Errorf("person name is required")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Who availed the service?",
		Validate: validate,
	}
	for {
		result, err := prompt.Run()
		if err != nil {
			fmt.Println(err)
			continue
		}
		return result
	}
}

// invoice id prompt
func PromptInvoiceId() string {
	validate := func(input string) error {
		if strings.TrimSpace(input) == "" {
			return fmt.Errorf("invoice id is required")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Invoice ID",
		Validate: validate,
	}
	for {
		result, err := prompt.Run()
		if err != nil {
			fmt.Println(err)
			continue
		}
		return result
	}
}
