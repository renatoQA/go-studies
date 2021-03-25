package main

import (
	"fmt"

	"github.com/tebeka/selenium"
)

func main() {
	port := 4723
	caps := selenium.Capabilities{
		"platformName": "Android",
		"deviceName":   "device",
		"app":          "C:/Users/renato.neto/Desktop/go-appium/myworten.apk",
	}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))

	if err != nil {
		fmt.Errorf("Not possible create new remote connection with Selenium: %v", err)
	}

	fmt.Println(wd.SessionID())

}
