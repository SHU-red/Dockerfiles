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
    fmt.Println("===== Starting fmon2pb =====")

    // CONFIGURATION ==============
    envvar := "DUMMY"

    //Shell definition
    shell := os.Getenv("SHELL")

    // Directory to check
    dir := "./files"
    envvar = "FMONPB_DIR"
    if isEnvExist(envvar) {
        fmt.Println("Environment Variable " + envvar + " used!")
        dir = os.Getenv(envvar)
    }

    // Transform filepath into absolute path
    dir, err := filepath.Abs(dir)
    if err != nil {
        log.Fatal(err)
    }

    // Number of kept files
    k := 10
    envvar = "FMONPB_NUM"
    if isEnvExist(envvar) {
        fmt.Println("Environment Variable " + envvar + " used!")
        k, _ = strconv.Atoi(os.Getenv(envvar))
    }

    // Check frequency in Milliseconds
    f := 1000
    envvar = "FMONPB_FRQ"
    if isEnvExist(envvar) {
        fmt.Println("Environment Variable " + envvar + " used!")
        f, _ = strconv.Atoi(os.Getenv(envvar))
    }

    // Configured channel
    c := ""
    envvar = "FMONPB_CHN"
    if isEnvExist(envvar) {
        fmt.Println("Environment Variable " + envvar + " used!")
        c = " -c " + os.Getenv(envvar) + " "
    }

    // Configured device
    s := ""
    envvar = "FMONPB_DEV"
    if isEnvExist(envvar) {
        fmt.Println("Environment Variable " + envvar + " used!")
        s = " -d " + os.Getenv(envvar) + " "
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

                    // Send via Pushbullet
                    out, err := exec.Command(shell,"-c", "pb push -f " + dir + "/" + p[x].Name() + c + s + " '" + t + " " + p[x].Name() + "'").Output()
                    if err != nil {
                        log.Fatal(err)
                    }
                    fmt.Println(string(out))

                }
            }

            // If exceeded number of files to keep
            if n > k {

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
