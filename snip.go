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
        if error != nil && os.IsNotExist(error) {
                errorDir := os.MkdirAll("directory", 0755)
                if errorDir != nil {
                        log.Fatal(error + " and " + errorDir)
                }
        }
}