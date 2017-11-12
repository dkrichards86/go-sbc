package sbc

type BondData struct {
	SerialNumber  string
	Series        string
	Denomination  float64
	IssueDate     string
	NextAccrual   string
	FinalMaturity string
	IssuePrice    float64
	Interest      float64
	InterestRate  float64
	Value         float64
	Note          string
}

type BondParam struct {
	Series       string
	SerialNumber string
	IssueDate    string
	Denomination string
}
