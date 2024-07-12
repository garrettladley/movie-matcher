package model

type Ranking struct {
	Movies []MovieID `json:"movies"`
}

type Prompt struct {
	People []Person  `json:"people"`
	Movies []MovieID `json:"movies"`
}

type Score struct {
	KendallTau int `json:"kendall_tau"`
}
