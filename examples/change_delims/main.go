package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"z3ntl3.com/autoenv/pkg/autoenv"
)

func main() {
	viper.AddConfigPath("../../tests")

	// viper.AutomaticEnv()
	viper.SetConfigName("test2")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	autoenv.New().SetDelim("<", ">").Execute() // interpolates all the placeholder variables with the delims
	fmt.Println(viper.GetString("WEBSITE")) // see the change
}