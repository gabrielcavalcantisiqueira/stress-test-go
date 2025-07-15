/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newRunCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Performs a load test by sending multiple concurrent HTTP requests",
		Long: `Runs a stress test against a specified URL by sending a defined number of HTTP requests with configurable concurrency.

You can customize the url, the total number of requests, the number of concurrent workers.

Example usage:

stress-test-go run --url http://localhost:3000 --requests 1000 --concurrency 50`,
		Run:    newCmdRun(),
		PreRun: newCmdPreRun(),
	}
}

func newCmdRun() RUN {
	return func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")
		result := RunStressTest(url, requests, concurrency)
		printReport(*result)
	}
}

func newCmdPreRun() RUN {
	return func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrent, _ := cmd.Flags().GetInt("concurrency")
		fmt.Printf("Starting stress test on %s with %d requests and %d concurrent calls...\n", url, requests, concurrent)
	}
}

func init() {
	var runCmd = newRunCmd()
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("url", "u", "", "Chosse the url")
	runCmd.Flags().IntP("requests", "r", 10, "Chosse the total number of requests")
	runCmd.Flags().IntP("concurrency", "c", 1, "Chosse the number of simultaneos requests")
	runCmd.MarkFlagRequired("url")
}
