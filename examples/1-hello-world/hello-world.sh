# To run the program, put the code in `hello-world.go` and
# use `go run`.
$ export GOPATH=<your workspace>/1-hello-world
$ go run src/hello-world.go
Hello, World!

# Sometimes we'll want to build our programs into
# binaries. We can do this using `go build`.
$ go build src/hello-world.go
$ ls
hello-world

# We can then execute the built binary directly.
$ ./hello-world
Hello, World!

# Instead of generate the executable in the <workspace>,
# generate it in the bin directory
$ go build -o bin/hello-world src/hello-world.go
$ ls bin
hello-world 

# An alternative way to build and install  executable
# in the bin folder is to use the go install command.
# It's important to define the GOBIN environment variable.
$ export GOBIN=$GOPATH/bin
$ go install src/hello-world.go
$ ls bin
hello-world
