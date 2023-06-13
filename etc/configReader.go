package util

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func ConfigReader() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./etc/")

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		// Handle error reading the configuration file
		fmt.Println(err)
	}

}

// Code to read the server urls

func ReadMultiServerConfig() *[]string {
	// Check if a value exists in the configuration
	var serverList []string
	exists := viper.IsSet("servers")
	fmt.Println("Multiple server are present")

	if exists {

		// []interface{} any type
		// string
		// int
		// []string
		// [24]string

		servers := viper.Get("servers").([]interface{})

		for _, server := range servers {
			serverMap := server.(map[string]interface{})
			name := serverMap["name"].(string)
			url := serverMap["url"].(string)
			fmt.Printf("Server Name: %s, URL: %s\n", name, url)
			serverList = append(serverList, url)
		}

	}
	return &serverList
}
