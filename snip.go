package main

import (
        "fmt"
        "log"
        "os"
        "os/user"
        "path/filepath"
)

func main() {
        fmt.Println("Snip v0.0.1 made with love by @Octalbyte (https://github.com/Octalbyte)")
        usr, err := user.Current()
        if err != nil {
                log.Fatal(err)
        }
        appdir := filepath.Join(usr.HomeDir, ".snipstore")
        //fmt.Println(appdir)
        _, error := os.Stat(appdir)
        if error != nil && os.IsNotExist(error) {
                errorDir := os.MkdirAll(appdir, 0755)
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
                case "help":
                help()
                default: 
                fmt.Println("Unknown action. run 'snip help' to get help.")                
        }
        
}

func help(){
        fmt.Println(
                `
                Hi! This is the help.
                Here are some arguments for the code:

                Any issues might be reported at https://github.com/Octalbyte/snip
                `)
}