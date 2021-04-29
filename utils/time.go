package utils

var railwayTime = map[string]int{
	// AM
	"12AM": 0,
	"1AM":  1,
	"2AM":  2,
	"3AM":  3,
	"4AM":  4,
	"5AM":  5,
	"6AM":  6,
	"7AM":  7,
	"8AM":  8,
	"9AM":  9,
	"10AM": 10,
	"11AM": 11,
	// PM
	"12PM": 12,
	"1PM":  13,
	"2PM":  14,
	"3PM":  15,
	"4PM":  16,
	"5PM":  17,
	"6PM":  18,
	"7PM":  19,
	"8PM":  20,
	"9PM":  21,
	"10PM": 22,
	"11PM": 23,
}

func TimeStrToRailwayTime(ts string) int {
	return railwayTime[ts]
}
