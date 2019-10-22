package behavioral

import "fmt"

type IStrategy interface {
	Export()
}

type Excel struct {}
func (Excel) Export() {
	fmt.Println("EXPORT REPORT to EXCEL")
}

type PDF struct {}
func (PDF) Export() {
	fmt.Println("EXPORT REPORT to PDF")
}

type Implementation struct {
	ReportStrat IStrategy
}

func (i Implementation) Export() {
	i.ReportStrat.Export()
}