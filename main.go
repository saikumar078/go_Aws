// package main

// import (
// 	"context"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"time"

// 	"github.com/chromedp/chromedp"
// )

// func main() {
// 	// Create a new context
// 	ctx, cancel := chromedp.NewContext(context.Background())
// 	defer cancel()

// 	// Navigate to a URL
// 	url := "https://www.netflix.com/login?nextpage=https%3A%2F%2Fwww.netflix.com%2Fbrowse"
// 	if err := chromedp.Run(ctx,
// 		chromedp.Navigate(url),
// 	); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Wait for a few seconds (optional)
// 	time.Sleep(5 * time.Second)

// 	// Capture a screenshot
// 	var screenshot []byte
// 	if err := chromedp.Run(ctx, chromedp.CaptureScreenshot(&screenshot)); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Save the screenshot to a file
// 	if err := ioutil.WriteFile("screenshot.png", screenshot, 0644); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Browser page opened successfully. Screenshot saved as screenshot.png.")
// }

package main

import (
	"os/exec"
	"runtime"
	"time"
)

func main() {
	// URL to open
	url := "https://www.netflix.com/login?nextpage=https%3A%2F%2Fwww.netflix.com%2Fbrowse"

	// Open web page based on the operating system
	switch runtime.GOOS {
	case "windows":
		exec.Command("cmd", "/c", "start", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	default: // Linux
		exec.Command("xdg-open", url).Start()
	}

	// Wait for 10 seconds
	time.Sleep(10 * time.Second)
	
}
