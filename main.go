package main

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
)

func main() {
	// initialize a Chrome browser instance on port 4444
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer service.Stop()

	// configure the browser options
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless-new", // comment out this line for testing
	}})

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")

	if err != nil {
		log.Fatal("Error:", err)
	}

	err = driver.Get("https://scrapingclub.com/exercise/list_infinite_scroll/")
	if err != nil {
		log.Fatal("Error:", err)
	}
}
