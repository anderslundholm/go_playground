package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var (
	localRootFlag   bool
	persistRootFlag bool
	times           int
	rootCmd         = &cobra.Command{
		Use:   "cobra_test",
		Short: "Example cobra CLI.",
		Long:  "This is an example of the cobra package. \nIt will have multiple subcommands and flags.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from the root command!")
		},
	}
	echoCmd = &cobra.Command{
		Use:   "echo [strings to echo]",
		Short: "prints given strings to stdout",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}
	timesCmd = &cobra.Command{
		Use:   "times [strings to echo]",
		Short: "prints given strings to stdout multiple times",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if times == 0 {
				return errors.New("times can not be 0")
			}
			for i := 0; i < times; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
			return nil
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persist-flag", "p", false, "a persistant use flag")
	rootCmd.Flags().BoolVarP(&localRootFlag, "local-flag", "l", false, "a local root flag")
	timesCmd.Flags().IntVarP(&times, "times", "t", 1, "number of times to echo to stdout")
	timesCmd.MarkFlagRequired("times")
	rootCmd.AddCommand(echoCmd)
	echoCmd.AddCommand(timesCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
