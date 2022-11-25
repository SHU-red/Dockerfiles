package main

// Imports
import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"time"
)

// Create Expanded FileInfo Struct
type ExtFileInfo struct {
    FileInfo os.FileInfo
    AbsPath string
}


// Check if Environment Variable is set
func isEnvExist(key string) bool {
    if _, ok := os.LookupEnv(key); ok {
        return true
    }
    return false
}

// Starting application
func main() {
    fmt.Println("===== Starting fmon2telegram =====")

    // CONFIGURATION ==============
    envvar := "DUMMY"

    //Shell definition
    shell := os.Getenv("SHELL")

    // Directory to check
    dir := "./files"
    envvar = "FMONTG_DIR"
    if isEnvExist(envvar) {
        fmt.Println("Environment Variable " + envvar + " used!")
        dir = os.Getenv(envvar)
    }

    // Message text
    txt := "New File"
    envvar = "FMONTG_TXT"
    if isEnvExist(envvar) {
        fmt.Println("Environment Variable " + envvar + " used!")
        txt = os.Getenv(envvar)
    }

    // Transform filepath into absolute path
    dir, err := filepath.Abs(dir)
    if err != nil {
        log.Fatal(err)
    }

    // Number of kept files
    k := 0
    envvar = "FMONTG_NUM"
    if isEnvExist(envvar) {
        fmt.Println("Environment Variable " + envvar + " used!")
        k, _ = strconv.Atoi(os.Getenv(envvar))
    }

    // Check frequency in Milliseconds
    f := 1000
    envvar = "FMONTG_FRQ"
    if isEnvExist(envvar) {
        fmt.Println("Environment Variable " + envvar + " used!")
        f, _ = strconv.Atoi(os.Getenv(envvar))
    }
    
    // DECLARATIONS =================

    // Loop
    l := 0

    // Number of files
    n := -1
    n_old := -1

    // Infinite Loop
    for {

        // Increment loop
        l++

        // Provide Timestamp
        current_time := time.Now()
        t := current_time.Format("060102-150405")

        // Loop Header
        fmt.Println("\n===== ", t)

        // Declare slices
        var p []ExtFileInfo
        var tmp ExtFileInfo

        // Declare RegEx for Filetypes
        re_vid := regexp.MustCompile(".*\\.mp4$")
        re_anim := regexp.MustCompile(".*\\.gif$")
        re_img := regexp.MustCompile(".*\\.jpg$|.*\\.jpeg$|.*\\.png$")

        // Scan folders animation 
        err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
            
            // Ignore Directories
            if !info.IsDir() {

                // Check Type via RegEx
                if re_vid.MatchString(path) || re_img.MatchString(path) {

                    // Bundle information in ExtFileInfo struct
                    tmp.FileInfo = info
                    tmp.AbsPath = path

                    // Append Information
                    p = append(p, tmp)

                }

            }
            return nil
        })
        if err != nil {
            panic(err)
        }

        // Number of files
        n = len(p)

        // Sort FileList
        sort.Slice(p, func(i,j int) bool{
            return p[i].FileInfo.ModTime().Unix() < p[j].FileInfo.ModTime().Unix()
        })

        // If not in first run
        if n_old > -1 {

            // Calculate difference
            d := math.Max(float64(n-n_old), 0)

            // If new files are found
            if d > 0 {

                // Notify about new files
                fmt.Println("Found new files:")

                // Pushfiles
                for x := n-int(d); x < n; x++ {

                    // Print Name
                    fmt.Println(p[x].FileInfo.Name())
                    fmt.Println(p[x].AbsPath)

                    // Default command as image
                    command := "-i"

                    // File extension related part of command
                    if re_vid.MatchString(p[x].AbsPath) {

                        // Command for videos and animations
                        command = "--video"

                    } else if re_anim.MatchString(p[x].AbsPath) {

                        // Command for animations such as gif
                        command = "--animation"
                    }

                    // Send via Telegram-send (installed via pip, thats why path to binary has to be used)
                    command = "telegram-send " + command + " " + p[x].AbsPath + " --caption '<b>" + txt + ":</b>\n\nTimestamp: <i> " + t + "</i>\n" + p[x].FileInfo.Name() + "' --format html"
                    fmt.Println(command)
                    out, err := exec.Command(shell,"-c",command).Output()
                    if err != nil {
                        log.Fatal(err)
                    }
                    fmt.Println(string(out))

                }
            }

            // If exceeded number of files to keep
            if n > k && k > 0 {

                // Loop over files which are too much
                for x := 0; x < n - k; x++ {

                    // Delete File
                    err := os.Remove(p[x].AbsPath)
                    if err != nil {
                        fmt.Println("Could not delete file: ", p[x].FileInfo.Name())
                    } else {
                        fmt.Println("Deleted file: ", p[x].FileInfo.Name())
                    }

                }

            }

        }

        // Store file number
        n_old = n

        // Sleep
        time.Sleep(time.Duration(f) * time.Millisecond)

    }

}
