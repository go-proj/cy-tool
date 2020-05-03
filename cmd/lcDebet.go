package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lcDebetCmd represents the lcDebet command
var lcDebetCmd = &cobra.Command{
	Use:   "lcDebet",
	Short: "fetch debet info from EastMoney",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lcDebet called")
	},
}

func init() {
	rootCmd.AddCommand(lcDebetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lcDebetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lcDebetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
