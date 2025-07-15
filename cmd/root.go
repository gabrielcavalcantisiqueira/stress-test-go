/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type RUN func(cmd *cobra.Command, args []string)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stress-test-go",
	Short: "A CLI tool for performing HTTP stress tests",
	Long: `Stress Test Go is a command-line tool designed to help developers
perform HTTP stress testing on web services and APIs.

With this tool, you can:
  - Send a configurable number of HTTP requests to a target URL
  - Control the level of concurrency (parallel requests)
  - Get a detailed status code report after execution

Built with Go and Cobra for simplicity and performance.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
