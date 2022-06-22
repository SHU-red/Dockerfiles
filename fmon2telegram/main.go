package main

// Imports
import (
    "os"
    "fmt"
    "time"
    "io/ioutil"
    "log"
    "sort"
    "math"
    "os/exec"
    "path/filepath"
    "strconv"
)

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

    // List for ignoring directories
    var p []os.FileInfo


    // Infinite Loop
    for {

        // Increment loop
        l++

        // Provide Timestamp
        current_time := time.Now()
        t := current_time.Format("060102-150405")

        // Loop Header
        fmt.Println("\n===== ", t)

        // Get files
        files, err := ioutil.ReadDir(dir)
        if err != nil {
            log.Fatal(err)
        }

        // Empty FileList
        p = nil

        // Pick only files and ignore directories
        for _, file := range files {

            // If File
            if !file.IsDir() {

                // Append File to list
                p = append(p, file)
            }
        }

        // Number of files
        n = len(p)

        // Sort FileList
        sort.Slice(p, func(i,j int) bool{
            return p[i].ModTime().Unix() < p[j].ModTime().Unix()
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
                    fmt.Println(p[x].Name())

                    // Get extension
                    ext := filepath.Ext(dir + "/" + p[x].Name())

                    // Create filepath string
                    p_path := dir + "/" + p[x].Name()

                    // Default command as image
                    command := "-i"

                    // File extension related part of command
                    if ext == ".gif" || ext == ".mp4" || ext == ".mkv" || ext == ".avi" || ext == ".mpeg" {

                        // Command for videos and animations
                        command = "--animation"

                    }

                    // Send via Telegram-send (installed via pip, thats why path to binary has to be used)
                    command = "telegram-send " + command + " " + p_path + " --caption '<b>" + txt + ":</b>\n\nTimestamp: <i> " + t + "</i>\n" + p[x].Name() + "' --format html"
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
                    err := os.Remove(dir + "/" + p[x].Name())
                    if err != nil {
                        log.Fatal(err)
                    } else {
                        fmt.Println("Deleted file: ", p[x].Name())
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
