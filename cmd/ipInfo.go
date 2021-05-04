package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ipInfoCmd = &cobra.Command{
	Short: "输出ip相关信息",
	Long: `输出ip相关信息:
1.输出本机IP
2.输出当前代理信息
3.输出当前「外网ip」/「输入ip」相关信息`,
	Run: ipInfoCmdRun,
}

func init() {
	rootCmd.AddCommand(ipInfoCmd)
	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// ipInfoCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// ipInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ipInfoCmdRun(cmd *cobra.Command, args []string) {
	fmt.Println("ipInfo called", args)
}
