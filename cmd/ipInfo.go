package cmd

import (
	"fmt"
	"github.com/go-proj/cy-tool/util"
	"github.com/spf13/cobra"
	"log"
	"net"
	"os"
	"sync"
)

var ipInfoCmd = &cobra.Command{
	Use:   "ipInfo",
	Short: "输出ip相关信息",
	Long: `输出ip相关信息:
1.输出本机IP
2.输出当前代理信息
3.输出当前「外网ip」/「输入ip」相关信息`,
	Run: ipInfoRun,
}

func ipInfoRun(cmd *cobra.Command, args []string) {
	var hostname string
	if tmp, err := os.Hostname(); err == nil {
		hostname = tmp
	}
	fmt.Printf("hostname: %#v\n\n", hostname)

	ips_local := localIp()
	fmt.Printf("内网ip: %+v\n\n", ips_local)

	fmt.Println("公网ip信息:")
	if len(args) != 0 {
		ip := args[0]
		resp := QueryIp(ip)
		for _, raw := range resp {
			fmt.Println(raw)
		}
	} else {
		resp := GlobalIp()
		for _, raw := range resp {
			fmt.Println(raw)
		}
	}
}

func init() {
	rootCmd.AddCommand(ipInfoCmd)
	// Cobra supports Persistent Flags whimch will work for this command and all subcommands, e.g.:
	// ipInfoCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// ipInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func localIp() []string {
	var ips []string

	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal("net.Interfaces failed.")
	}

	for _, i := range ifaces {
		if "en0" != i.Name {
			continue
		}

		addrs, err := i.Addrs()
		if err != nil {
			log.Fatal("net.Addrs failed.")
		}

		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ips = append(ips, v.IP.String())
			case *net.IPAddr:
				ips = append(ips, v.IP.String())
			}
		}
	}
	return ips
}

//外网ip
func GlobalIp() []string {
	var info []string

	drivers := []util.IpSvc{
		util.IpFmSvc,
		util.IpInfoSvc,
		util.IpipSvc,
	}

	c := make(chan string, len(drivers))

	wg := sync.WaitGroup{}
	wg.Add(3)

	for _, driver := range drivers {
		go func(d util.IpSvc) {
			c <- d.LocalInfo()
			wg.Done()
		}(driver)
	}
	wg.Wait()

	for len(c) > 0 {
		info = append(info, <- c)
	}
	return info
}

//查询ip信息
func QueryIp(ip string) []string {
	var info []string

	drivers := []util.IpSvc{
		util.IpFmSvc,
		util.IpInfoSvc,
		util.IpipSvc,
	}

	c := make(chan string, len(drivers))

	wg := sync.WaitGroup{}
	wg.Add(3)

	for _, driver := range drivers {
		go func(d util.IpSvc) {
			c <- d.GlobalInfo(ip)
			wg.Done()
		}(driver)
	}
	wg.Wait()

	for len(c) > 0 {
		info = append(info, <- c)
	}
	return info
}
