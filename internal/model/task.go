package model

type Task struct {
	ProdOrderID        string
	SeqNr              string
	JobID              string
	OperationID        string
	ProdOrderStartDate string
	CTSec              string
	OriginalQty        string
	RemainingQty       string
	LineSequence       string
	ProcessTime        string
	WorkingGroup       string
	WorkingStation     string
}
