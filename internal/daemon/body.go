package daemon

// GetSecretBody is the JSON body expected by the /secret/get endpoint
type GetSecretBody struct {
	Key    string `json:"key"`
	Hidden bool   `json:"hidden"`
}

type SetSecretBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PrintBody struct {
	Text string `json:"text"`
}

type SpinnerStartBody = PrintBody

type SpinnerStopBody struct {
	Text string `json:"text,omitempty"`
}

type ProgressBarStartBody struct {
	Length  int    `json:"length"`
	Initial int    `json:"initial"`
	Text    string `json:"text"`
}

type ProgressBarAdvanceBody struct {
	Increment int `json:"increment,omitempty"`
}

type ProgressBarStopBody = SpinnerStopBody

type EventsBody struct {
	Start string `json:"start"`
	End   string `json:"end"`
}
