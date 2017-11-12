# Go-SBC
Go-SBC is a savings bond analysis tool. It ingests a CSV of bond data, POSTs them
to the [TreasuryDirect website]("https://www.treasurydirect.gov/BC/SBCPrice)
(a product of Bureau of the Fiscal Service and US Treasury), scrapes the resulting
page for bond data, and performs some simple calculations on them.

## Background
I have a handful of US Savings Bonds lying around, approaching maturity. 
[TreasuryDirect]("https://www.treasurydirect.gov/BC/SBCPrice) has a tool to
calculate savings bond values, but this tool does not support bulk operations,
and no API was available. I wanted to create a tool for bulk calculations.

I had been looking to learn [Go](https://golang.org/) for some time now, and this
project seemed like a perfect candidate. Pardon any discrepancies, as my Go may
not be the most idiomatic.

## Usage
First, run `go get github.com/dkrichards86/sbc` to fetch dependencies.

You will need to include a CSV file in the root directory. The CSV needs 4 columns,
with no header row, with the columns representing series, serial numnber, issue
date and denomination. The order is important here, as a `BondParam` struct is 
created based on the values at specific indexes. The main script uses a default 
name of `bonds.csv`, but you are free to change the name. 

Once you've added that file, run `go run main.go (--filename=YOUR_FILENAME.csv)`

`main.go` will generate output similar to the following: 
```
Values
=====
Interest Accrued: $xxx.xx
Total Value: $xxxxx.xx

Notes
=====
L123456789EE: Matured and not earning interest
L246801357EE: Not eligible for payment
C987654321EE: Includes 3 month interest penalty
```

## Credits
Under the hood, Go-SBC uses [Surf](https://github.com/headzoo/surf) for stateful
browsing and [goquery](https://github.com/PuerkitoBio/goquery) for easy HTML
selection. It was written on [C9.io](https://c9.io/).

I relied on [Go by Example](https://gobyexample.com/) and [Stack Overflow](https://stackoverflow.com/)
to learn the syntax and idioms associated with Go.