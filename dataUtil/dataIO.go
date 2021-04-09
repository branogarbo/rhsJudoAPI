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

func WriteTotalWorkout(mutatedTotalWorkouts TotalWorkouts) error {
	var (
		err                       error
		mutatedTotalWorkoutsBytes []byte
	)

	mutatedTotalWorkoutsBytes, err = json.Marshal(mutatedTotalWorkouts)
	if err != nil {
		return err
	}

	err = os.WriteFile("./data/judoWorkoutLog.json", mutatedTotalWorkoutsBytes, 0666)
	if err != nil {
		return err
	}

	return nil
}
