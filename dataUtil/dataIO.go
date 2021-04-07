package dataUtil

import (
	"encoding/json"
	"os"
)

func ReadTotalWorkoutStruct() (TotalWorkouts, error) {
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

func WriteTotalWorkoutStruct(mutatedTotalWorkouts interface{}) error {
	var (
		err       error
		jsonBytes []byte
	)
	jsonBytes, err = json.Marshal(mutatedTotalWorkouts)
	if err != nil {
		return err
	}

	err = os.WriteFile("./data/judoWorkoutLog", jsonBytes, 0666)
	if err != nil {
		return err
	}

	return nil
}
