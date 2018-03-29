package scheduler

import (
	"math"
	"time"

	. "../config"
	. "../models"
)

func Init() error {
	go func() {
		for {
			UpdateMemorizations()
			time.Sleep(10 * time.Second)
		}
	}()

	return nil
}

func UpdateMemorizations() {
	var items []Memorization
	last_update_seletor := uint64(time.Now().Unix()) - uint64(Settings.MemorizationsUpdateTimeDelta)

	err := DB.Model(&items).
		Where("last_update_timestamp < ?", last_update_seletor).
		Limit(1024).
		Select()

	if err != nil {
		panic(err)
	}

	for _, memorization := range items {
		currentTimestamp := time.Now().Unix()

		dt := currentTimestamp - memorization.LastUpdateTimestamp
		forgettingSpeed := 1.0 / Settings.MemorizationFullForgettingInDays / 24 / 3600
		forgettingSpeed *= 1 - math.Sqrt(memorization.MemorizationCoefficient)

		if forgettingSpeed < Settings.MemorizationMinimumForgettingSpeed {
			forgettingSpeed = Settings.MemorizationMinimumForgettingSpeed
		}

		memorizationDelta := float64(dt) * forgettingSpeed

		memorization.MemorizationCoefficient -= memorizationDelta
		if memorization.MemorizationCoefficient < 0 {
			memorization.MemorizationCoefficient = 0
		}

		memorization.LastUpdateTimestamp = currentTimestamp
		_, err = DB.Model(&memorization).
			Column("memorization_coefficient", "last_update_timestamp").
			Update()

		if err != nil {
			panic(err)
		}
	}
}
