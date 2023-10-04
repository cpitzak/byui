package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	var res []byte
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithBrowserOption())
	defer cancel()
	chromedpError := chromedp.Run(ctx,
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
	if chromedpError != nil {
		fmt.Println(chromedpError)
		os.Exit(1)
	}
	homeDir, homeDirError := os.UserHomeDir()
	if homeDirError != nil {
		fmt.Println(homeDirError)
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
	_, newFileLocationError := os.Stat(newFileLocation)
	newFileLocationExists := newFileLocationError == nil
	if newFileLocationExists {
		if renameError := os.Rename(newFileLocation, oldFileLocation); renameError != nil {
			fmt.Println("Error when renaming: " + newFileLocation + " to this location: " + oldFileLocation)
			os.Exit(1)
		}
	}
	os.WriteFile(newFileLocation, res, 0644)
	if !newFileLocationExists {
		fmt.Println("wrote screenshot of application status for the first time. Openining the screenshot...")
		cmd := exec.Command("xdg-open", newFileLocation)
		cmd.Run()
	} else {
		cmd := exec.Command("idiff", newFileLocation, oldFileLocation)
		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb
	
		if iDiffError := cmd.Run(); iDiffError != nil {
			fmt.Println(iDiffError)
			os.Exit(1)
		}
		if strings.Contains(outb.String(), "PASS") {
			fmt.Println("No status change")
		} else {
			fmt.Println("Status change! Openining screenshot of application status...")
			cmd2 := exec.Command("xdg-open", newFileLocation)
			cmd2.Run()
		}
	}
}
