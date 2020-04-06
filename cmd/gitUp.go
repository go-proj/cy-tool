package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gitUpCmd represents the gitUp command
var gitUpCmd = &cobra.Command{
	Use:   "gitUp",
	Short: "cd $proj && git pull --recurse-submodules",
	Long: `cd to the project directory, and then execute 'git pull --recurse-submodules'. For example:
cd /Users/clarkyu/company/bigbase/code_php/pt-proj && git pull --recurse-submodules`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitUp called")
	},
}

func init() {
	rootCmd.AddCommand(gitUpCmd)
	gitUpCmd.PersistentFlags().String("help", "", "A help for gitUp")
	gitUpCmd.Flags().BoolP("help", "h", false, "Help message for gitUp")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gitUpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gitUpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
