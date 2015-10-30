package alertbaseutil

type ROW struct {
	Subtab       string
	Level        string
	Subject      string
	Escalate     string
	Escalatemin1 string
	Escalatemin2 string
	Subjectnum   string
	Doneat       string
	Openat       string
	Owner        string
	Assigner     string
	Status       string
	Comment      string
}

type OPSROW struct {
	Opsname      string
	Takenoverat  string
}

type ROWS []*ROW
