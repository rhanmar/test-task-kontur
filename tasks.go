package main

import (
	"fmt"
	"time"
)

// layEggs - задача для курицы
func layEggs(eggsToLay int) {
	EggsCount += eggsToLay
	fmt.Println(fmt.Sprintf("Курица снесла %d яиц. Стало %d", eggsToLay, EggsCount))
}

// getEggs - задача для фермера
func getEggs(eggsToGet int) {
	if EggsCount < eggsToGet {
		EggsCount = 0
	} else {
		EggsCount -= eggsToGet
	}
	fmt.Println(fmt.Sprintf("Фермер взял %d яиц. Осталось %d", eggsToGet, EggsCount))
}

// startChickenTask - начать выполнять задачи для курицы
func startChickenTask(config Config) {
	for range time.Tick(time.Second * time.Duration(config.layEggsSecInterval)) {
		eggsToLay := GetRandomIntForRange(config.layEggsMaxCount, config.layEggsMinCount)
		layEggs(eggsToLay)
	}
}

// startFarmerTask - начать выполнять задачи для фермера
func startFarmerTask(config Config) {
	for range time.Tick(time.Second * time.Duration(config.getEggsSecInterval)) {
		eggsToGet := GetRandomIntForRange(config.getEggsMaxCount, config.getEggsMinCount)
		getEggs(eggsToGet)
	}
}
