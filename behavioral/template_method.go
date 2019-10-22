package behavioral

import "fmt"

type ExportTemplate interface {
	SetHeader()
	SetBody()
	SetFooter()
}

type Export struct {
	ExportTemplate
}

func NewExporter(t ExportTemplate) Export {
	return Export{t}
}

func (w Export) ExportReport() {
	w.SetHeader()
	w.SetBody()
	w.SetFooter()
}

type TemplateImpl struct {}

func (w TemplateImpl) SetHeader() {
	fmt.Println("Set Header")
}

func (w TemplateImpl) SetBody() {
	fmt.Println("Set Body")
}

func (w TemplateImpl) SetFooter() {
	fmt.Println("Set Footer")
}