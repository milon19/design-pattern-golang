package main

import (
	"design-pattern/creational/builder/query"
	"design-pattern/creational/singleton/config"
	"design-pattern/creational/singleton/logger"
	"fmt"
)

func main() {
	config1 := config.GetConfig()
	config2 := config.GetConfig()
	fmt.Println("configManager1 == configManager2:", config1 == config2)
	if appName, ok := config1.Get("app_name"); ok {
		fmt.Println("appName:", appName)
	}

	logging := logger.GetLogger()
	logging.Info("This is an info message")
	logging.Debug("This is a debug message")
	logging.Error("This is an error message")
	logging.Critical("This is a critical message")

	queryBuilder := query.NewSQLQueryBuilder()

	q := queryBuilder.
		Select("name, age").
		From("users").
		Where("age > 18").
		Build()

	fmt.Println("Constructed SQL Query:")
	fmt.Println(q)
}
