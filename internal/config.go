package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

type Info struct {
	Bees     string
	HiveSlot string
	Field    string
}

type Config struct {
	Info Info `toml:"info"`
}

func config_init() {
	// Path to your TOML file
	tomlFile := "config.toml"

	// Open and read the TOML file
	file, err := os.Open(tomlFile)
	if err != nil {
		fmt.Printf("Error opening TOML file: %v\n", err)
		return
	}
	defer file.Close()

	// Parse the TOML data into a Config struct
	var config Config
	decoder := toml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		fmt.Printf("Error decoding TOML: %v\n", err)
		return
	}

	// Access the extracted information
	fmt.Println("Info:")
	fmt.Printf("Bees: %s\n", config.Info.Bees)
	fmt.Printf("Hive Slot: %s\n", config.Info.HiveSlot)
	fmt.Printf("Field: %s\n", config.Info.Field)
}
