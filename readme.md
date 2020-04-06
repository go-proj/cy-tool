# [Refer](https://blog.csdn.net/benben_2015/article/details/82227338)
    go mod init github.com/go-proj/cy-tool
    go mod tidy
    go mod verify
    go mod vendor

# [cobra](https://github.com/spf13/cobra/blob/master/cobra/README.md)
    cobra init --pkg-name github.com/spf13/newApp path/to/newApp
    cobra add serve
    cobra add config
    cobra add create -p 'configCmd'
