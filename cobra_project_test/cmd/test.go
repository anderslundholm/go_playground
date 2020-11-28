package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	testProjectFlag  string
	testInstanceFlag string
	testDatabaseFlag string
	// testCmd represents the test command
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test connection to a cloud database.",
		Long:  `Test connection to a cloud database.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("testing /project/%s/instance/%s/database/%s/\n", testProjectFlag, testInstanceFlag, testDatabaseFlag)
		},
	}
)

func init() {
	rootCmd.AddCommand(testCmd)

	testCmd.Flags().StringVarP(&testProjectFlag, "project", "p", "", "Project ID")
	testCmd.Flags().StringVarP(&testInstanceFlag, "instance", "i", "", "Instance Name")
	testCmd.Flags().StringVarP(&testDatabaseFlag, "database", "d", "", "Database Name")
	testCmd.MarkFlagRequired("project")
	testCmd.MarkFlagRequired("instance")
	testCmd.MarkFlagRequired("database")
}
