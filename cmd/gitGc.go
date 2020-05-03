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

// gitGcCmd represents the gitGc command
var gitGcCmd = &cobra.Command{
	Use:   "gitGc",
	Short: "git gc --prune=now --aggressive ",
	Long: `execute git gc command, to prune history.`,
	Run: gitGcRun,
}

func init() {
	rootCmd.AddCommand(gitGcCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// gitGcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// gitGcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

const GitGcWorkers = 20 // total goroutine

func gitGcRun(cmd *cobra.Command, args []string) {
	if (len(args) != 0) {
		for _, dir := range args[:] {
			gitGcDir(dir)
		}
	} else {
		bigbase := "/Users/clarkyu/company/bigbase"
		fmt.Printf("start git_gc %v\n", bigbase)
		gitGcDir(bigbase)
	}
}

// process each directory:
// 1. find a .git dir? execute cmd && exit
// 2. go through each dir inside...
func gitGcDir(dir string) bool {
	if (strings.Contains(dir, "vendor") || strings.Contains(dir, "3rd_party")) {
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
			if (strings.HasSuffix(item, ".git") && (!strings.Contains(item, "vendor"))) {
				tmp := strings.Replace(item, "/.git", "", -1)
				fmt.Println(">>> start git gc " + filepath.Base(tmp))
				ret, err := common.Run("git -C '" + tmp + "' gc --prune=now --aggressive ", true)
				if (err != nil) {
					fmt.Println( err )
				} else {
					fmt.Println( string(ret) )
				}
			} else {
				gitGcDir(item) //gothrough folder
			}
		case mode.IsRegular():
			//TODO file?
		}
	}

	return true
}
