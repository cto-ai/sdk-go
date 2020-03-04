package daemon

type GetSecretBody struct {
	Key string `json:"key"`
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
