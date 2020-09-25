# Hello World program

This is the classical Hello World! program. To run the program do the following:

```
export GOPATH=<your workspace>/1-hello-world
cd $GOPATH
go run src/hello-world.go
```

Sometimes we'll want to build our programs into binaries. We can do this using ```go build```.

```
go build src/hello-world.go
./hello-world
```

Instead of generate the executable in the <workspace>, generate it in the bin directory.

```
go build -o bin/hello-world src/hello-world.go
bin/hello-world
```

An alternative way to build and install executable in the bin folder is to use the ```go install``` command.
It's important to define the GOBIN environment variable.


```
export GOBIN=$GOPATH/bin
go install src/hello-world.go
bin/hello-world
```
