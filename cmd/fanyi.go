package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// fyYoudaoCmd represents the fyYoudao command
var fanyiCmd = &cobra.Command{
	Use:   "fanyi",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: fanyiCmdRun,
}

func init() {
	rootCmd.AddCommand(fanyiCmd)

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// fyYoudaoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// fyYoudaoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func fanyiCmdRun(cmd *cobra.Command, args []string) {
	fmt.Println("fanyi called", args)
}
