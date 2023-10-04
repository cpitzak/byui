package main

import (
	"strings"
    "context"
    "fmt"
    "os"
    "time"
    "github.com/chromedp/chromedp"
    "path/filepath"
    "os/exec"
    "bytes"
    "errors"
)

func main() {
    
    var res []byte
    ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithBrowserOption())
    defer cancel()
    err := chromedp.Run(ctx,
        chromedp.Navigate("https://employment.byui.net/login"),
        chromedp.WaitReady("body"),
        chromedp.SetValue(`user_username`, "", chromedp.ByID),
        chromedp.SetValue(`user_password`, "", chromedp.ByID),
        chromedp.Click(`//*[@value="Log In"]`, chromedp.BySearch),
        chromedp.Sleep(time.Second*2),
        chromedp.Navigate("https://employment.byui.net/job_applications"),
        chromedp.Sleep(time.Second*2),
        chromedp.FullScreenshot(&res, 100),
    )
    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
    homeDir, err9 := os.UserHomeDir()
    if err9 != nil {
        fmt.Println(err9)
    }
    storageFolder := filepath.Join(homeDir, ".byui")
    if _, err := os.Stat(storageFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(storageFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
    oldFileLocation := filepath.Join(storageFolder, "old_appstatus.png")
    newFileLocation := filepath.Join(storageFolder, "appstatus.png")
    if _, err3 := os.Stat(newFileLocation); err3 == nil {
        if e := os.Rename(newFileLocation, oldFileLocation); e != nil {
            os.Exit(1)
        }
    }

    os.WriteFile(newFileLocation, res, 0644)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    cmd := exec.Command("idiff", newFileLocation, oldFileLocation)
    var outb, errb bytes.Buffer
    cmd.Stdout = &outb
    cmd.Stderr = &errb
    
    if err5 := cmd.Run(); err5 != nil {
        fmt.Println(err5)
        os.Exit(1)
    }
    if strings.Contains(outb.String(), "PASS") {
        fmt.Println("No status change")
    } else {
        fmt.Println("Status change! Openining app status...")
        cmd2 := exec.Command("xdg-open", newFileLocation)
        cmd2.Run()
    }
}