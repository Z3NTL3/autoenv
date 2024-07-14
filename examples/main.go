package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"z3ntl3.com/autoenv/pkg/autoenv"
)

func main() {
	viper.AddConfigPath("../tests")

	// viper.AutomaticEnv()
	viper.SetConfigName("test")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	autoenv.New().Execute()
	fmt.Printf("Evaluation of WEBSITE: %s", viper.GetString("website"))

}