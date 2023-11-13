package main

import (
	"os"
	"strconv"
)

// Config - Структура со значениями для диапазонов
type Config struct {
	layEggsMinCount    int
	layEggsMaxCount    int
	layEggsSecInterval int
	getEggsMinCount    int
	getEggsMaxCount    int
	getEggsSecInterval int
}

// InitConfig - функция инициализации конфига
func InitConfig() Config {
	layEggsMinCount, err := strconv.Atoi(os.Getenv("layEggsMinCount"))
	if err != nil {
		layEggsMinCount = 2
	}
	layEggsMaxCount, err := strconv.Atoi(os.Getenv("layEggsMaxCount"))
	if err != nil {
		layEggsMaxCount = 5
	}
	layEggsSecInterval, err := strconv.Atoi(os.Getenv("layEggsSecInterval"))
	if err != nil {
		layEggsSecInterval = 2
	}
	getEggsMinCount, err := strconv.Atoi(os.Getenv("getEggsMinCount"))
	if err != nil {
		getEggsMinCount = 10
	}
	getEggsMaxCount, err := strconv.Atoi(os.Getenv("getEggsMaxCount"))
	if err != nil {
		getEggsMaxCount = 20
	}
	getEggsSecInterval, err := strconv.Atoi(os.Getenv("getEggsSecInterval"))
	if err != nil {
		getEggsSecInterval = 10
	}
	return Config{
		layEggsMinCount:    layEggsMinCount,
		layEggsMaxCount:    layEggsMaxCount,
		layEggsSecInterval: layEggsSecInterval,
		getEggsMinCount:    getEggsMinCount,
		getEggsMaxCount:    getEggsMaxCount,
		getEggsSecInterval: getEggsSecInterval,
	}
}
