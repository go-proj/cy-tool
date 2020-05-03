package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// proxyXiciCmd represents the proxyXici command
var proxyXiciCmd = &cobra.Command{
	Use:   "proxyXici",
	Short: "fetch available proxy from xici.com",
	Long: ``,
	Run: proxyXiciRun,
}

func init() {
	rootCmd.AddCommand(proxyXiciCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// proxyXiciCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// proxyXiciCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

const (
	XiciHome = "https://www.xicidaili.com"
	XiciGaoNi = "https://www.xicidaili.com/nn"
	XiciPuTong = "https://www.xicidaili.com/nt"
	XiciHttps = "https://www.xicidaili.com/wn"
	XiciHttp = "https://www.xicidaili.com/wt"
)

func proxyXiciRun(cmd *cobra.Command, args []string) {
	fmt.Println("proxyXici called")
}
