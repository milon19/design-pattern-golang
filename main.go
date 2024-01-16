package main

import (
	"design-pattern/creational/singleton/config"
	"fmt"
)

func main() {
	config1 := config.GetConfig()
	config2 := config.GetConfig()
	fmt.Println("configManager1 == configManager2:", config1 == config2)
	if appName, ok := config1.Get("app_name"); ok {
		fmt.Println("appName:", appName)
	}
}
