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
    // Example student list (can be 100+)
    students := []Student{
        {"Sahana M", "1 A", "AS-51", map[string]string{"Math": "95", "Science": "88"}},
        {"Rahul K", "2 B", "AS-52", map[string]string{"Math": "89", "Science": "90"}},
        {"Lakshmi P", "3 A", "AS-53", map[string]string{"Math": "75", "Science": "85"}},
    }

    // 1. **CONCATENATE ALL HTML into a single string**
    var combinedHTML strings.Builder
    
    // Add the starting HTML structure
    combinedHTML.WriteString(`<!DOCTYPE html><html><head><meta charset="utf-8">
    <style>
        body { font-family: Arial; margin: 20px; }
        h2 { text-align: center; color: #2b2b2b; }
        table { width: 100%%; border-collapse: collapse; }
        th, td { border: 1px solid #000; padding: 5px; text-align: center; }
    </style>
    </head><body>`)

    for _, s := range students {
        // Append the body content for each student
        combinedHTML.WriteString(generateStudentContent(s)) 
    }

    // Add the closing HTML structure
    combinedHTML.WriteString(`</body></html>`)
    
    // 2. **Initialize PDF generator**
    pdfg, err := wkhtmltopdf.NewPDFGenerator()
    if err != nil {
        log.Fatal(err)
    }

    // 3. **Add the ONE combined page**
    page := wkhtmltopdf.NewPageReader(strings.NewReader(combinedHTML.String()))
    pdfg.AddPage(page)


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

// NOTE: This new function only returns the body content, not the full HTML/HEAD/BODY structure
func generateStudentContent(s Student) string {
    content := fmt.Sprintf(`
    <h2>Report Card - %s (%s)</h2>
    <p>Adm No: %s</p>
    <table>
    <tr><th>Subject</th><th>Marks</th></tr>`, s.Name, s.Class, s.AdmNo)

    for subject, marks := range s.Marks {
        content += fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>", subject, marks)
    }

    content += `
    </table>
    <div style="page-break-after: always;"></div>`

    return content
}

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
