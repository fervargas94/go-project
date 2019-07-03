package utils

// CalculatePriority receives an x and y value and maps the coordinate it into a segment
func CalculatePriority(priority int, weight int) int {
	var prioritySize int
	if priority <= 5 && weight <= 5 {
		prioritySize = 1
	} else if priority > 5 && weight > 5 {
		prioritySize = 3
	} else {
		prioritySize = 2
	}
	return prioritySize
}
