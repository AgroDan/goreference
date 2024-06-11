# Go Reference

This repository will act as a means that I can use to refer back to techniques and methods I've used before with Go. Since I usually go months between projects that require coding in Go, despite years of experience it only takes a few weeks of not using it to forget everything. So this will be a way to document things. This isn't a "learn go" tutorial or anything, but rather references to methods that I can potentially use in the future so I at least know how to do things.

Probably only useful for me, but hey if you can use this then knock yourself out.

## Starting a Go Project

There are two aspects to go projects. Modules and Workspaces. A module is your most basic "atomic" project, where it consists of a "main" package (the default package where the expected entrypoint exists, aka `func main()`). If you are working with multiple modules that all require their own `main` namespace, a Workspace will help tie everything together.

In this repo, I will set the current workspace to `./`, and then each subdirectory will be its own separate module. That way, you can type `go run <module_name>` to run the `main()` function in that specific module from the current directory.

### Start a workspace

To start a workspace, type:

```terminal
$ go work init
```

This creates a `go.work` file which only states the version of Go that we are running. Now I'll create a "Hello World" application in the `./hello/` directory.

### Create a module

I'll create the `hello` module by setting up the directory, `cd`'ing into it, and initializing the module.

```terminal
$ mkdir -p ./hello
$ pushd ./hello
$ go mod init hello
go: creating new go.mod: module hello
```

This creates the `go.mod` file, which consists of the version of Go we are using as well as the name of the particular module.

Now we can go back to the main directory and tell the workspace that the `hello` module is part of this workspace.

```terminal
$ popd
$ go work use hello
```

This modifies the `go.work` file so that now in this current directory, I can run the `hello` program just by typing `go run hello`.

I made a quick "Hello World" application to show that if you type `go run hello` it will print "Hello World!"

# Table of Contents

I'll try to add this as much as possible

- [Working with Files and String Manipulation](./filesandstrings/)
- [Working with Structs](./structs/)
- [Finite Sets](./finitesets/)
- [Merge Sort](./mergesort/)
- [Concurrency: Channels, Waitgroups, and Mutexes](./channels/)
- [Working with Files and Directories](./files/)
- [Bytes Buffer](./buffers/)
- [Web Requests](./webreq/)
- [Hashing and Base64'ing in Go](./hashing)