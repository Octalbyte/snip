package main
import (
    "os/user"
    "fmt"
    "log"
	"path/filepath"
)
func main() {
    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }
	appdir := filepath.Join(usr.HomeDir, ".snipstore")
    fmt.Println( usr.HomeDir )
	fmt.Println(appdir)
	
}