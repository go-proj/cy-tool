package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fmtCurlCmd represents the fmtCurl command
var fmtCurlCmd = &cobra.Command{
	Use:   "fmtCurl",
	Short: "fmt浏览器导出的curl命令语句",
	Long: `将浏览器导出的curl语句，进行简单的格式化. 例如:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: fmtCurlRun,
}

func fmtCurlRun(cmd *cobra.Command, args []string) {
	fmt.Println("fmtCurl called", args)
}

func init() {
	rootCmd.AddCommand(fmtCurlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// fmtCurlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// fmtCurlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
