package cmd

import (
	"fmt"
	"github.com/go-proj/lib/common"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// gitUpCmd represents the gitUp command
var gitUpCmd = &cobra.Command{
	Use:   "gitUp",
	Short: "cd $proj && git pull --recurse-submodules",
	Long:  `cd to the project directory, and then execute 'git pull --recurse-submodules'.`,

	Run: gitUpRun,
}

const GitUpWorkers = 20 // total goroutine

func gitUpRun(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
		for _, dir := range args[:] {
			gitUpDir(dir)
		}
	} else {
		dir_php := "/Users/clarkyu/code/php"
		fmt.Printf("start git_up %v\n", dir_php)
		gitUpDir(dir_php)
	}
}

// process each directory:
// 1. find a .git dir? execute cmd && exit
// 2. go through each dir inside...
func gitUpDir(dir string) bool {
	if strings.Contains(dir, "vendor") || strings.Contains(dir, "3rd_party") {
		return false
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("ReadDir_fail %v\n", err)
		return false
	}

	for _, f := range files {
		item := dir + "/" + f.Name()
		fi, err := os.Stat(item)
		if err != nil {
			fmt.Printf("Stat_fail %v\n", err)
			continue
		}

		switch mode := fi.Mode(); {
		case mode.IsDir():
			if strings.HasSuffix(item, ".git") && (!strings.Contains(item, "vendor")) {
				tmp := strings.Replace(item, "/.git", "", -1)
				fmt.Println(">>> start git up " + filepath.Base(tmp))
				ret, err := common.Run("git -C '"+tmp+"' pull --recurse-submodules ", true)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(string(ret))
				}
			} else {
				gitUpDir(item)
			}
		}
	}

	return true
}

func init() {
	rootCmd.AddCommand(gitUpCmd)

	//gitUpCmd.PersistentFlags().String("help", "", "A help for gitUp")
	//gitUpCmd.Flags().BoolP("info", "i", false, help())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// gitUpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// gitUpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
