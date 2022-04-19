package main

type LbcResponse struct {
	Count   int64    `json:"count"`
	Results []Result `json:"results"`
}

type Result struct {
	Company            Company            `json:"company"`
	RequestForProposal RequestForProposal `json:"requestForProposal"`
}

type Company struct {
	EncodedID                   string `json:"encodedId"`
	Name                        string `json:"name"`
	IsCurrentSupplierReferenced bool   `json:"isCurrentSupplierReferenced"`
}

type RequestForProposal struct {
	ID                  string              `json:"id"`
	EncodedID           string              `json:"encodedId"`
	Title               string              `json:"title"`
	Location            Location            `json:"location"`
	WorkUnit            WorkUnit            `json:"workUnit"`
	AverageDailyRate    float64             `json:"averageDailyRate"`
	Currency            string              `json:"currency"`
	StartingDate        string              `json:"startingDate"`
	PublicationDate     string              `json:"publicationDate"`
	MissionDuration     int64               `json:"missionDuration"`
	RemoteWork          string              `json:"remoteWork"`
	Applications        string              `json:"applications"`
	Competences         string              `json:"competences"`
	Confidential        bool                `json:"confidential"`
	MaxBudgetDisplay    bool                `json:"maxBudgetDisplay"`
	Favorite            bool                `json:"favorite"`
	SupplierApplication SupplierApplication `json:"supplierApplication"`
}

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type SupplierApplication struct {
	Applied         bool   `json:"applied"`
	ApplicationDate string `json:"applicationDate"`
}

type WorkUnit struct {
	WorkUnitType string `json:"workUnitType"`
	WorkUnits    int64  `json:"workUnits"`
}
