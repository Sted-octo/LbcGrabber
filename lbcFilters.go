package main

type LbcFilters struct {
	SortedBy         string      `json:"sortedBy"`
	Keywords         []string    `json:"keywords"`
	JobTitles        []string    `json:"jobTitles"`
	Companies        []string    `json:"companies"`
	Location         LbcLocation `json:"location"`
	AverageDailyRate interface{} `json:"averageDailyRate"`
	StartingDate     interface{} `json:"startingDate"`
	MissionDurations interface{} `json:"missionDurations"`
	RemoteWork       interface{} `json:"remoteWork"`
	PublicationDate  string      `json:"publicationDate"`
}

type LbcLocation struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Radius    string  `json:"radius"`
}
