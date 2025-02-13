package server

import "time"

func getTimeOfDay() string {
	t := time.Now()
	hour := t.Hour()

	if hour >= 5 && hour < 12 {
		return "Morning"
	} else if hour >= 12 && hour < 17 {
		return "Afternoon"
	} else {
		return "Evening"
	}
}

func getFirstName() string{
  return "Andrew"
}
