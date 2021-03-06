package dataUtil

type Exercises map[string][]int

type ParticipantExercises map[string]Exercises

type TotalWorkouts []PracticeWorkout

type PracticeWorkout struct {
	Date         string               `json:"date"`
	Participants ParticipantExercises `json:"participants"`
}
