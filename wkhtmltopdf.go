package main

import (
	"fmt"
	"log"
	"strings"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type Student struct {
	Name   string
	Class  string
	AdmNo  string
	Marks  map[string]string
}

func main() {
	// Example: 3 students, but can scale to 100+
	students := []Student{
		{"Sahana M", "1 A", "AS-51", map[string]string{"Math": "95", "Science": "88"}},
		{"Rahul K", "2 B", "AS-52", map[string]string{"Math": "89", "Science": "90"}},
		{"Lakshmi P", "3 A", "AS-53", map[string]string{"Math": "75", "Science": "85"}},
	}

	// Initialize PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Loop over all students
for _, s := range students {
    pdfg := wkhtmltopdf.NewPDFGenerator()
    html := generateStudentHTML(s)
    page := wkhtmltopdf.NewPageReader(strings.NewReader(html))
    pdfg.AddPage(page)
    pdfg.Create()
    filename := fmt.Sprintf("%s.pdf", s.Name)
    pdfg.WriteFile(filename)
}
	// PDF settings
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.MarginBottom.Set(10)
	pdfg.MarginTop.Set(10)
	pdfg.MarginLeft.Set(10)
	pdfg.MarginRight.Set(10)

	// Create and save the PDF
	if err := pdfg.Create(); err != nil {
		log.Fatal("Failed to create PDF:", err)
	}

	if err := pdfg.WriteFile("All_Students_Report.pdf"); err != nil {
		log.Fatal("Failed to write PDF:", err)
	}

	fmt.Println("✅ PDF for all students generated successfully!")
}

// --- Helper ---
func generateStudentHTML(s Student) string {
	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
	<meta charset="utf-8">
	<style>
		body { font-family: Arial; margin: 20px; }
		h2 { text-align: center; color: #2b2b2b; }
		table { width: 100%%; border-collapse: collapse; }
		th, td { border: 1px solid #000; padding: 5px; text-align: center; }
	</style>
	</head>
	<body>
	<h2>Report Card - %s (%s)</h2>
	<p>Adm No: %s</p>
	<table>
	<tr><th>Subject</th><th>Marks</th></tr>`, s.Name, s.Class, s.AdmNo)

	for subject, marks := range s.Marks {
		html += fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>", subject, marks)
	}

	html += `
	</table>
	<div style="page-break-after: always;"></div>
	</body>
	</html>`

	return html
}
