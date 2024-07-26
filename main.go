/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/jiang-kevin/airline-reservation-assignment/cmd"
	"github.com/jiang-kevin/airline-reservation-assignment/data"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var data data.Data = data.AraData{}

	rootCmd := cmd.RootCmd()
	rootCmd.AddCommand(cmd.BookCmd(data))
	rootCmd.AddCommand(cmd.CancelCmd(data))

	cmd.Execute(rootCmd)
}
