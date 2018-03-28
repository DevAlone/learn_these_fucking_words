package scheduler

import (
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
	last_update_seletor := time.Now().Unix() - int64(Settings["memorizations_update_time_delta"].(float64))

	err := DB.Model(&items).
		Where("last_update_timestamp < ?", last_update_seletor).
		Limit(1024).
		Select()

	if err != nil {
		panic(err)
	}

	for _, memorization := range items {
		memorization.MemorizationCoefficient -= 0.00001
		if memorization.MemorizationCoefficient < 0 {
			memorization.MemorizationCoefficient = 0
		}

		memorization.LastUpdateTimestamp = uint64(time.Now().Unix())
		_, err = DB.Model(&memorization).
			Column("memorization_coefficient", "last_update_timestamp").
			Update()

		if err != nil {
			panic(err)
		}
	}

	/*_, err = DB.Model(&items).
	Column("memorization_coefficient").
	Update()

	if err != nil {
		panic(err)
	}*/
}
