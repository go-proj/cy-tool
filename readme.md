# ref
	https://ordina-jworks.github.io/development/2018/10/20/make-your-own-cli-with-golang-and-cobra.html
	https://zhangguanzhang.github.io/2019/06/02/cobra

# [go mod](https://blog.csdn.net/benben_2015/article/details/82227338)
    go mod init github.com/go-proj/cy-tool
    go mod tidy
    go mod verify
    go mod vendor
    go list -m -u all <https://www.reddit.com/r/golang/comments/b1dd3r/update_go_mod_dependencies>

# [cobra usage](https://github.com/spf13/cobra/blob/master/cobra/README.md)
    cobra init --pkg-name github.com/spf13/newApp path/to/newApp
    cobra add serve
    cobra add config
    cobra add create -p 'configCmd'
