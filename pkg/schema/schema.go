package schema

type Request struct {
	Id      string `json:"id,omitemptye"`
	Message string `json:"message"`
}

// Response schema
type Response struct {
	Code    int    `json:"code,omitempty"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResultsSchema struct {
	Project   string
	StartTime string
	StopTime  string
	Items     []ItemInfo
}

type ItemInfo struct {
	Name        string
	Fail        bool
	Pass        bool
	Time        int64
	DisplayTime string
	Log         string
}
