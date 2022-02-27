package main

import (
        "fmt"
        "log"
        "os"
        "os/user"
        "path/filepath"
)

func main() {
        usr, err := user.Current()
        if err != nil {
                log.Fatal(err)
        }
        appdir := filepath.Join(usr.HomeDir, ".snipstore")
        fmt.Println(appdir)
        _, error := os.Stat(appdir)
        fmt.Println(error)
        if error != nil && os.IsNotExist(error) {
                errorDir := os.MkdirAll("directory", 0755)
                if errorDir != nil {
                        log.Fatal(fmt.Sprintf("%v and %s", error, errorDir))
                }
        }
        if len(os.Args) == 1 {
                log.Fatal("No arguments were provided. Type 'help' to get more info.")
                os.Exit(0)
        }
        var action = os.Args[1]
        switch action {

        }
        
}