package structure

type Artists struct {
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    Locations
	ConcertDates Dates
	Relations    Relation
	Language     Language
	Suggestion   []Artists
}

type Language struct {
	En          []string
	Fr          []string
	Es          []string
	Ge          []string
	CurrentLang []string
}

type Base struct {
	Nbr       int
	Data      []Artists
	Show      []Artists
	Language  Language
	DataGenre map[string][]Artists
	ShowGenre map[string][]Artists
}

type Relation struct {
	ID       int64               `json:"id"`
	Relation map[string][]string `json:"datesLocations"`
}

type Locations struct {
	ID        int64    `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	ID    int64    `json:"id"`
	Dates []string `json:"dates"`
}
