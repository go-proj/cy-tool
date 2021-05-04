package cmd

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// ymlCmd represents the yml command
var ymlCmd = &cobra.Command{
	Use:   "yml",
	Short: "Read yaml config file, analyze redis",
	Long:  ``,
	Run:   ymlRun,
}

type TT_Cfg struct {
	RedisUrl string `yaml:"redis"` // redis://localhost:6379/dbname
	MySQLUrl string `yaml:"mysql"` // mysql://username:password@protocol(address)/dbname?param=value
}

var cmdConfPath string //
var cmdCtx = context.Background()

func ymlRun(cmd *cobra.Command, args []string) {
	//if (len(args) <= 0) {
	// process : sub command ...
	//	fmt.Println("please input config file path")
	//	return
	//}
	if len(cmdConfPath) == 0 {
		err := cmd.Help()
		if err != nil {
			log.Fatalf("cmdConfPath.Get err #%v", err)
		}
		return
	}

	//try to read yaml file content
	yamlContent, err := ioutil.ReadFile(cmdConfPath)
	if err != nil {
		log.Fatalf("yamlContent.Get err #%v", err)
	}

	var cfg TT_Cfg
	err = yaml.Unmarshal(yamlContent, &cfg)
	if err != nil {
		log.Fatalf("Yaml Unmarshal: %v", err)
	}

	fmt.Println(cfg.RedisUrl)
	fmt.Println(cfg.MySQLUrl)
}

func RedisClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rdb.Ping(cmdCtx).Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func init() {
	rootCmd.AddCommand(ymlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ymlCmd.PersistentFlags().String("foo", "", "A help for foo")
	// ymlCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.demo.yaml)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ymlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	ymlCmd.Flags().StringVarP(&cmdConfPath, "yaml", "f", "", "yaml config file path")
}
