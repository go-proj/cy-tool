package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gitSearchCmd represents the gitUp command
var gitSearchCmd = &cobra.Command{
	Use:   "gitSearch",
	Short: "gitSearch dir_name",
	Long: ``,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitSearchCmd called")
	},
}

func init() {
	rootCmd.AddCommand(gitSearchCmd)

	//gitUpCmd.PersistentFlags().String("help", "", "A help for gitUp")
	//gitUpCmd.Flags().BoolP("info", "i", false, help())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// gitUpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// gitUpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
