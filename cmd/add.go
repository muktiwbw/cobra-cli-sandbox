/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

This command is used to add two numbers together
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func add(a, b int) int {
	return a + b
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A command to add two numbers together",
	Long: `This command is used to add two numbers together.

The add command takes two numbers as arguments and returns their sum.

Example:
  mynewapp add 5 3
  Output: 8

Usage:
  mynewapp add <number1> <number2>

The numbers can be integers or decimal values.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if the number of arguments is 2
		if len(args) != 2 {
			return fmt.Errorf("this command requires exactly 2 arguments, but got %d", len(args))
		}

		// Convert the first argument
		num1, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("first argument '%s' is not a valid number", args[0])
		}

		// Convert the second argument
		num2, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("second argument '%s' is not a valid number", args[1])
		}

		// Add the integers and print the result
		result := add(num1, num2)
		fmt.Printf("The sum of %d and %d is %d\n", num1, num2, result)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
