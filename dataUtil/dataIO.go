package dataUtil

import (
	"encoding/json"
	"os"
)

func ReadTotalWorkoutFile() (TotalWorkouts, error) {
	var (
		totalWorkouts TotalWorkouts
		jsonFile      []byte
		err           error
	)

	jsonFile, err = os.ReadFile("./data/judoWorkoutLog.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonFile, &totalWorkouts)
	if err != nil {
		return nil, err
	}

	return totalWorkouts, nil
}

func WriteTotalWorkout(mutatedTotalWorkouts []byte) error {
	var (
		err error
	)

	err = os.WriteFile("./data/judoWorkoutLog.json", mutatedTotalWorkouts, 0666)
	if err != nil {
		return err
	}

	return nil
}
