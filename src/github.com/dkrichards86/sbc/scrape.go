package sbc

import (
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

// US Treasury Savings Bond Calculator URL
const sbc = "https://www.treasurydirect.gov/BC/SBCPrice"

var BondNotes = map[string]string{
	"NI": "Not Issued",
	"NE": "Not eligible for payment",
	"P5": "Includes 3 month interest penalty",
	"MA": "Matured and not earning interest",
}

// Build a MM/YYYY string from the current date
func currDateString() string {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	strMonth := intToStr(int64(currentMonth))
	strYear := intToStr(int64(currentYear))
	dateString := strMonth + "/" + strYear

	return dateString
}

// Sanitize float-like strings by stripping away all invalid characters.
// In this case, strip away anything that isn't 0-9 or '.'
func sanitizeFloat(input string) string {
	reg, _ := regexp.Compile("[^0-9.]+")

	return reg.ReplaceAllString(input, "")
}

// Convert a string to a 64-bit float
func toFloat(num string) float64 {
	sanitized := sanitizeFloat(num)
	f, _ := strconv.ParseFloat(sanitized, 64)
	return f
}

// Convert an integer to a string in base 10
func intToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}

func Scrape(bp BondParam) BondData {
	bow := surf.NewBrowser()

	redemptionDate := currDateString()

	params := url.Values{
		"RedemptionDate": {redemptionDate},
		"Series":         {bp.Series},
		"SerialNumber":   {bp.SerialNumber},
		"IssueDate":      {bp.IssueDate},
		"Denomination":   {bp.Denomination},
		"btnAdd.x":       {"CALCULATE"},
	}

	err := bow.PostForm(sbc, params)
	if err != nil {
		panic(err)
	}

	data := make([]string, 0)
	bow.Dom().Find("table.bnddata tbody > tr.altrow1 td").Each(func(_ int, s *goquery.Selection) {
		data = append(data, s.Text())
	})

	bond := BondData{
		SerialNumber:  data[0],
		Series:        data[1],
		Denomination:  toFloat(data[2]),
		IssueDate:     data[3],
		NextAccrual:   data[4],
		FinalMaturity: data[5],
		IssuePrice:    toFloat(data[6]),
		Interest:      toFloat(data[7]),
		InterestRate:  toFloat(data[8]),
		Value:         toFloat(data[9]),
		Note:          BondNotes[data[10]],
	}

	return bond
}
