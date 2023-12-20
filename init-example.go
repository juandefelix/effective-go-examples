package main

import (
    "log"
    "flag"
)

var (
    user, home, gopath string
)

func init() {
    log.Println("In init")
    if user == "" {
        log.Fatal("$USER not set")
    }
    if home == "" {
        home = "/home/" + user
    }
    if gopath == "" {
        gopath = home + "/go"
    }
    // gopath may be overridden by --gopath flag on command line.
    flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}


func main() {
    log.Println("In main")
}

func finally() {
    log.Println("In finally")
}
