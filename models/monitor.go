package models

type StatusReq struct {
	TestString string
}

type ProposeReq struct {
	Candidate string
	Value     string
}

type RestartNetReq struct {
	NodeType string
	NodeName string
}
