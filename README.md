# Hello glide

```
$ echo 'export $GOPATH=~/go' >> ~/.zshrc
$ mkdir ~/go/{src,pkg,bin}
```

```
$ brew install glide
$ glide init
$ # edit glide.yaml to add dependencies
$ glide up
```

```
$ go run main.go hello go
hello go

$ go run main.go goodbye go
```
