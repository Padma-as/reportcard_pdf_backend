package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
"sort"
"strings"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// -----------------------------
// STRUCTS
// -----------------------------
type PageDecorationConfig struct {
	ShowBackground   bool
	BackgroundImage  string
	ShowWatermark    bool
	WatermarkImage   string
	ShowBorder       bool
	BorderColor      string
	BorderWidth      float64
	BorderType       string
	MarginTop        string
	MarginRight      string
	MarginBottom     string
	MarginLeft       string
	TopPadding       string
	RightPadding     string
	BottomPadding    string
	LeftPadding      string

}
type StudentReportData struct {
	StudentCfg     StudentDetailsConfig
	Tests          []Test

}
type TitleConfig struct {
    TitleColor       string
	TitleFontSize    int
	SubtitleColor    string
	SubtitleFontSize int
	TitleText string
	SubTitleText string
	EnableTitle bool
	EnableSubTitle bool
}

type InstitutionDetailsConfig struct {
	PrintPhoto1Config     bool
	PrintPhoto2Config     bool
	PrintInstLogo         bool
	HeaderName           string
	EnableHeader          bool
	Photo1Base64          string
	Photo2Base64          string
	InstLogoBase64        string
	HeaderColor           string
	HeaderFontSize        int
	PrintCustName         bool
	CustName              string
	CustomerNameColor     string
	CustomerNameFontSize  int
	PrintInstName         bool
	InstName              string
	InstNameColor         string
	InstNameFontSize      int
	EnableAffiliated      bool
	PrintAffiliatedTo     string
	AffiliatedColor       string
	AffiliatedFontSize    int
	PrintInstAddress      bool
	InstAddress           string
	InstPlace             string
	InstPin               string
	AddressColor          string
	AddressFontSize       int
	PrintInstWebsite      bool
	InstWebsite           string
	WebsiteColor          string
	WebsiteFontSize       int
	PrintInstEmail        bool
	InstEmail             string
	EmailColor            string
	EmailFontSize         int
}
type StudentDetailsConfig struct {
	StudentName        string
	StudentRollNo      string
	FatherName         string
	MotherName         string
	StudentClass       string
	AcademicYear       string
	DateOfBirth        string
	AttendanceStats    string
	Address            string
	Email              string
	Mobile             string
	PhotoBase64        string
	PhotoOnRight       bool  // if false, photo will be on left
	ShowPhoto          bool
	ShowName           bool
	ShowRollNo         bool
	ShowFatherName     bool
	ShowMotherName     bool
	ShowClassSection   bool
	ShowAcademicYear   bool
	ShowDateOfBirth    bool
	ShowAttendance     bool
	ShowAddress        bool
	ShowEmail          bool
	ShowMobile         bool
	FontSize           int
	FontColor          string

	DisplayTwoColumn   bool
	StudentPhotoX int
	StudentPhotoY int
}

type ReportConfig struct {
	ShowMaxPerSubject        bool
	ShowMinPerSubject        bool
	ShowTotal      bool
	ShowPercentage bool
	ShowGradePerSubject       bool

	EnableSlNo bool
	ShowMaxPerTest bool
	ShowMinPerTest bool
	ShowGradePerTest bool
	ShowRemarksPerTest bool
	ShowConductPerTest bool
	
	PrintRemarks string
	PrintConduct string
	EnableGradeForLastTestOnly bool
	PrintOnlyGrade bool
	Table1Title string
	Table2Tittle string
	TableTitleFontSize int
	TableDataFontSize int
	RemarksText string
	ConductText string
	PercentageText string
	Fontsize int

	ShowTestName bool
	Subjects            []string

}

type Test struct {
	Name    string
	Marks   map[string]int // subject -> marks
	Max     map[string]int
	Min     map[string]int
	Remarks map[string]string
	Grade   map[string]string
}
// -----------------------------
// MAIN
// -----------------------------
func main() {
	cfg := PageDecorationConfig{
		ShowBackground:  true,
		BackgroundImage: toBase64("./assets/background.png"),
		ShowWatermark:   true,
		WatermarkImage:  toBase64("./assets/watermark.jpeg"),
		ShowBorder:      true,
		BorderColor:     "#333",
		BorderWidth:     2,
		BorderType:      "solid",
		MarginTop:       "20px",
		MarginRight:     "20px",
		MarginBottom:    "20px",
		MarginLeft:      "20px",
	}

	instCfg := InstitutionDetailsConfig{
		PrintPhoto1Config:    true,
		PrintPhoto2Config:    true,
		PrintInstLogo:        true,
		HeaderName:           "EDUATE PRIVATE LIMITED",
		EnableHeader:         true,
		Photo1Base64:         toBase64("./assets/Arcadis_Logo.png"),
		Photo2Base64:         toBase64("./assets/Arcadis_Logo.png"),
		InstLogoBase64:       toBase64("./assets/Arcadis_Logo.png"),
		HeaderColor:          "#1A237E",
		HeaderFontSize:       20,
		PrintCustName:        true,
		CustName:             "Eduate ERP System",
		CustomerNameColor:    "#000",
		CustomerNameFontSize: 16,
		PrintInstName:        true,
		InstName:             "Eduate International School",
		InstNameColor:        "#111",
		InstNameFontSize:     18,
		EnableAffiliated:     true,
		PrintAffiliatedTo:    "Affiliated to CBSE",
		AffiliatedColor:      "#444",
		AffiliatedFontSize:   14,
		PrintInstAddress:     true,
		InstAddress:          "123 Eduate Street",
		InstPlace:            "Bangalore",
		InstPin:              "560001",
		AddressColor:         "#555",
		AddressFontSize:      12,
		PrintInstWebsite:     true,
		InstWebsite:          "www.eduate.com",
		WebsiteColor:         "#444",
		WebsiteFontSize:      12,
		PrintInstEmail:       true,
		InstEmail:            "info@eduate.com",
		EmailColor:           "#444",
		EmailFontSize:        12,
	}

	titleCfg := TitleConfig{
		TitleColor:       "#1a237e",
		TitleFontSize:    22,
		SubtitleColor:    "#424242",
		SubtitleFontSize: 16,
		TitleText:        "Title",
		SubTitleText:     "subtitle",
		EnableTitle:      true,
		EnableSubTitle:   true,
	}

	// ✅ Base Student Config
	baseStudentCfg := StudentDetailsConfig{
		ShowName:         true,
		ShowRollNo:       true,
		ShowFatherName:   true,
		ShowMotherName:   true,
		ShowClassSection: true,
		ShowAcademicYear: true,
		ShowDateOfBirth:  true,
		ShowAttendance:   true,
		ShowAddress:      true,
		ShowEmail:        true,
		ShowMobile:       true,
		ShowPhoto:        true,
		DisplayTwoColumn: true,
		FontSize:         14,
		FontColor:        "#000",
		PhotoBase64:      toBase64("./assets/colorwatermark.png"),
		StudentPhotoX:    80,
		StudentPhotoY:    80,
	}

	// ✅ Students
studentsData := []StudentReportData{
	{
		StudentCfg: newStudent(
			baseStudentCfg,
			"John Doe", "45", "Mr. David Doe", "Mrs. Sarah Doe",
			"10-A", "2024-25", "01-01-2010", "95%",
			"123 Main Street, Bangalore", "john@example.com", "9999999999", true,
		),
		Tests: []Test{
			{
				Name: "Test 1",
				Marks: map[string]int{
					"Math": 85, "Science": 90, "English": 88,
				},
				Max: map[string]int{
					"Math": 100, "Science": 100, "English": 100,
				},
				Grade: map[string]string{
					"Math": "A", "Science": "A+", "English": "A",
				},
				Remarks: map[string]string{
					"Math": "Good", "Science": "Excellent", "English": "Good",
				},
			},
			{
				Name: "Test 2",
				Marks: map[string]int{
					"Math": 78, "Science": 84, "English": 90,
				},
				Max: map[string]int{
					"Math": 100, "Science": 100, "English": 100,
				},
				Grade: map[string]string{
					"Math": "B+", "Science": "A", "English": "A+",
				},
				Remarks: map[string]string{
					"Math": "Fair", "Science": "Good", "English": "Excellent",
				},
			},	{
				Name: "Test 3",
				Marks: map[string]int{
					"Math": 78, "Science": 84, "English": 90,
				},
				Max: map[string]int{
					"Math": 100, "Science": 100, "English": 100,
				},
				Grade: map[string]string{
					"Math": "B+", "Science": "A", "English": "A+",
				},
				Remarks: map[string]string{
					"Math": "Fair", "Science": "Good", "English": "Excellent",
				},
			},
		},
	},
	{
		StudentCfg: newStudent(
			baseStudentCfg,
			"Jane Smith", "72", "Mr. Michael Smith", "Mrs. Olivia Smith",
			"9-B", "2024-25", "", "92%", "", "jane@example.com", "", true,
		),
		Tests: []Test{
			{
				Name: "Test 1",
				Marks: map[string]int{
					"Math": 78, "Science": 84, "English": 90,
				},
				Max: map[string]int{
					"Math": 100, "Science": 100, "English": 100,
				},
				Grade: map[string]string{
					"Math": "B+", "Science": "A", "English": "A+",
				},
				Remarks: map[string]string{
					"Math": "Fair", "Science": "Good", "English": "Excellent",
				},
			},
			{
				Name: "Test 2",
				Marks: map[string]int{
					"Math": 78, "Science": 84, "English": 90,
				},
				Max: map[string]int{
					"Math": 100, "Science": 100, "English": 100,
				},
				Grade: map[string]string{
					"Math": "B+", "Science": "A", "English": "A+",
				},
				Remarks: map[string]string{
					"Math": "Fair", "Science": "Good", "English": "Excellent",
				},
			},
				{
				Name: "Test 2",
				Marks: map[string]int{
					"Math": 78, "Science": 84, "English": 90,
				},
				Max: map[string]int{
					"Math": 100, "Science": 100, "English": 100,
				},
				Grade: map[string]string{
					"Math": "B+", "Science": "A", "English": "A+",
				},
				Remarks: map[string]string{
					"Math": "Fair", "Science": "Good", "English": "Excellent",
				},
				
			},
			{
				Name: "Test 2",
				Marks: map[string]int{
					"Math": 78, "Science": 84, "English": 90,
				},
				Max: map[string]int{
					"Math": 100, "Science": 100, "English": 100,
				},
				Grade: map[string]string{
					"Math": "B+", "Science": "A", "English": "A+",
				},
				Remarks: map[string]string{
					"Math": "Fair", "Science": "Good", "English": "Excellent",
				},
				
			},
				{
				Name: "Test 2",
				Marks: map[string]int{
					"Math": 78, "Science": 84, "English": 90,
				},
				Max: map[string]int{
					"Math": 100, "Science": 100, "English": 100,
				},
				Grade: map[string]string{
					"Math": "B+", "Science": "A", "English": "A+",
				},
				Remarks: map[string]string{
					"Math": "Fair", "Science": "Good", "English": "Excellent",
				},
				
			},
				{
				Name: "Test 2",
				Marks: map[string]int{
					"Math": 78, "Science": 84, "English": 90,
				},
				Max: map[string]int{
					"Math": 100, "Science": 100, "English": 100,
				},
				Grade: map[string]string{
					"Math": "B+", "Science": "A", "English": "A+",
				},
				Remarks: map[string]string{
					"Math": "Fair", "Science": "Good", "English": "Excellent",
				},
				
			},
		},
	},
}

	acdCfg := ReportConfig{
		ShowMaxPerSubject:        false,
		ShowMinPerSubject:        false,
		ShowTotal:                true,
		ShowPercentage:           true,
		ShowGradePerSubject:      true,
		ShowRemarksPerTest:       true,
		ShowConductPerTest:       true,
		EnableSlNo:               true,
		PrintRemarks:     "over-all",
		PrintConduct:     "over-all",
		EnableGradeForLastTestOnly: true,
		PrintOnlyGrade:             false,
		ShowMaxPerTest:true,
		ShowMinPerTest:true,
		Table1Title:                "Part A",
		Table2Tittle:               "Part B",
		TableTitleFontSize:         16,
		TableDataFontSize:          14,
		RemarksText:                "Remarks",
		ConductText:                "Conduct",
		PercentageText:             "Percentage",
		ShowTestName:               false,
		Subjects:                   []string{"English", "Math", "Science"},
	}

html := generateAllStudentsHTML(cfg, instCfg, titleCfg, acdCfg, studentsData)

	// ✅ Generate a single PDF file containing all pages
	if err := generatePDF(html, "All_Students_Report.pdf"); err != nil {
		log.Fatal("❌ PDF generation failed:", err)
	}
	fmt.Println("✅ Generated: All_Students_Report.pdf")
}

// -----------------------------
// HELPERS
// -----------------------------
func toBase64(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("⚠️ Could not read file %s: %v", path, err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(data)
}

// -----------------------------
// HTML GENERATION
// -----------------------------
func generateAllStudentsHTML(
	cfg PageDecorationConfig,
	instCfg InstitutionDetailsConfig,
	titleCfg TitleConfig,
	acdConfig ReportConfig,
	students []StudentReportData,
) string {
	var allReportsHTML string

	for _, studentData := range students {
		instHTML := generateInstitutionDetailsHTML(instCfg)
		titleHTML := generateTitleHTML(titleCfg)
		studentDetailsHTML := generateStudentDetailsHTML(studentData.StudentCfg)
		academicDetailsHTML := generateAcademicDetails(acdConfig, studentData.Tests)
        chartHTML := generateStudentChartHTML(studentData.Tests)

		var bgCSS, wmCSS string
		if cfg.ShowBackground && cfg.BackgroundImage != "" {
			bgCSS = fmt.Sprintf(`background-image: url('data:image/png;base64,%s'); background-repeat: no-repeat; background-size: cover; background-position: center;`, cfg.BackgroundImage)
		}
		if cfg.ShowWatermark && cfg.WatermarkImage != "" {
			wmCSS = fmt.Sprintf(`background-image: url('data:image/jpeg;base64,%s'); background-repeat: no-repeat; background-size: 30%%; background-position: center center;`, cfg.WatermarkImage)
		}

		borderStyle := "none"
		if cfg.ShowBorder {
			borderStyle = fmt.Sprintf("%.1fpx %s %s", cfg.BorderWidth, cfg.BorderType, cfg.BorderColor)
		}

		// Combine per-student page
		studentPageHTML := fmt.Sprintf(`
		<div class="report-page">
			<div class="page-wrapper" style="border:%s; padding:%s %s %s %s; %s">
				<div class="content-with-watermark" style="%s">
					%s <!-- Institution Header -->
					<hr style="border:0.5px solid black">
					<b>%s</b>
					
					<div>%s</div>
					<div>%s</div>
<div>%s</div>
					
				</div>
			</div>
		</div>
		`, borderStyle, cfg.MarginTop, cfg.MarginRight, cfg.MarginBottom, cfg.MarginLeft,
			bgCSS, wmCSS, instHTML, titleHTML, studentDetailsHTML, academicDetailsHTML,chartHTML)

		allReportsHTML += studentPageHTML
	}

	// Wrap all pages in a single HTML document
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<style>
html, body {
	margin: 0;
	padding: 0;
	height: 100%%;
	width: 100%%;
	-webkit-print-color-adjust: exact;
	print-color-adjust: exact;
	font-family: 'Arial', sans-serif;
}
.report-page {
	width: 210mm; /* A4 width */
	height: 297mm; /* A4 height */
	box-sizing: border-box;
	page-break-after: always;
	margin: 0;
	padding: 0;
}
.report-page:last-child {
	page-break-after: avoid;
}
.page-wrapper {
	width: 100%%;
	height: 100%%;
	box-sizing: border-box;
	background-color: white;
}
.content-with-watermark {
	width: 100%%;
	height: 100%%;
	box-sizing: border-box;
	position: relative;
	padding:10px;
}
.header-section {
	text-align: center;
	margin-bottom: 10px;
}
.header-section img {
	height: 60px;
	margin: 0 10px;
	object-fit: contain;
}
.info {
	font-size: 14px;
	margin-top: 8px;
}
</style>
</head>
<body>
	%s
</body>
</html>
`, allReportsHTML)
}

// -----------------------------
// INSTITUTION SECTION
// -----------------------------
func generateInstitutionDetailsHTML(cfg InstitutionDetailsConfig) string {
	images := []string{}
	if cfg.PrintPhoto1Config && cfg.Photo1Base64 != "" {
		images = append(images, fmt.Sprintf(`<img src="data:image/png;base64,%s" alt="photo1" style="height:60px;">`, cfg.Photo1Base64))
	}
	if cfg.PrintInstLogo && cfg.InstLogoBase64 != "" {
		images = append(images, fmt.Sprintf(`<img src="data:image/png;base64,%s" alt="logo" style="height:60px;">`, cfg.InstLogoBase64))
	}
	if cfg.PrintPhoto2Config && cfg.Photo2Base64 != "" {
		images = append(images, fmt.Sprintf(`<img src="data:image/png;base64,%s" alt="photo2" style="height:60px;">`, cfg.Photo2Base64))
	}

	// Build institution text HTML
	instTextHTML := fmt.Sprintf(`
	<div style="text-align:center; margin-top:10px;">
		<div style="color:%s; font-size:%dpx;">%s</div>
		<div style="color:%s; font-size:%dpx;">%s</div>
		<div style="color:%s; font-size:%dpx;">%s</div>
		<div style="color:%s; font-size:%dpx;">%s, %s - %s</div>
		<div style="color:%s; font-size:%dpx;">%s</div>
		<div style="color:%s; font-size:%dpx;">%s</div>

	</div>
	`,
		cfg.HeaderColor, cfg.HeaderFontSize, cfg.HeaderName,
		cfg.InstNameColor, cfg.InstNameFontSize, cfg.InstName,
		cfg.AffiliatedColor, cfg.AffiliatedFontSize, cfg.PrintAffiliatedTo,
		cfg.AddressColor, cfg.AddressFontSize, cfg.InstAddress, cfg.InstPlace, cfg.InstPin,
		cfg.WebsiteColor, cfg.WebsiteFontSize, cfg.InstWebsite,
		cfg.EmailColor, cfg.EmailFontSize, cfg.InstEmail,
	)

	var headerRowHTML string

	switch len(images) {
	case 0:
		// No image → show text only
		headerRowHTML = instTextHTML
	case 1:
		// Single image → image centered, text below
		headerRowHTML = fmt.Sprintf(`
		<div style="text-align:center; margin-bottom:10px;">%s</div>
		%s
		`, images[0], instTextHTML)
	case 2:
		// Two images → left image, right image, text in same row
		headerRowHTML = fmt.Sprintf(`
		<div style="display:flex; align-items:center; justify-content:space-between; margin-bottom:10px;">
			<div>%s</div>
			<div style="flex:1; text-align:center;">
				%s
			</div>
			<div>%s</div>
		</div>
		`, images[0], instTextHTML, images[1])
	case 3:
		// Three images → images justify space-between, text below
		headerRowHTML = fmt.Sprintf(`
		<div style="display:flex; justify-content:space-between; margin-bottom:10px;">
			%s%s%s
		</div>
		%s
		`,
			images[0], images[1], images[2],
			instTextHTML)
	}

	return fmt.Sprintf(`<div class="header-section">%s</div>`, headerRowHTML)
}

func generateTitleHTML(titleCfg TitleConfig) string {
	titleHTML := ""
	subtitleHTML := ""

	if titleCfg.EnableTitle {
		titleHTML = fmt.Sprintf(
			`<h1 style="text-align:center; color:%s; font-size:%dpx; margin:5px 0;">%s</h1>`,
			titleCfg.TitleColor, titleCfg.TitleFontSize, titleCfg.TitleText,
		)
	}

	if titleCfg.EnableSubTitle {
		subtitleHTML = fmt.Sprintf(
			`<h3 style="text-align:center; color:%s; font-size:%dpx; margin:2px 0;">%s</h3>`,
			titleCfg.SubtitleColor, titleCfg.SubtitleFontSize, titleCfg.SubTitleText,
		)
	}

	return titleHTML + subtitleHTML
}

func generateStudentDetailsHTML(cfg StudentDetailsConfig) string {
	// Convert all values to string safely
	fields := []struct {
		Index int
		Label string
		Value string
		Show  bool
	}{
		{4, "Name", cfg.StudentName, cfg.ShowName},
		{2, "Roll No", fmt.Sprintf("%v", cfg.StudentRollNo), cfg.ShowRollNo},
		{3, "Father Name", cfg.FatherName, cfg.ShowFatherName},
		{1, "Mother Name", cfg.MotherName, cfg.ShowMotherName},
		{5, "Class & Section", cfg.StudentClass, cfg.ShowClassSection},
		{6, "Academic Year", cfg.AcademicYear, cfg.ShowAcademicYear},
		{7, "Date of Birth", cfg.DateOfBirth, cfg.ShowDateOfBirth},
		{8, "Attendance", fmt.Sprintf("%v", cfg.AttendanceStats), cfg.ShowAttendance},
		{9, "Address", cfg.Address, cfg.ShowAddress},
		{10, "Email", cfg.Email, cfg.ShowEmail},
		{11, "Mobile", fmt.Sprintf("%v", cfg.Mobile), cfg.ShowMobile},
	}

	// Sort fields by Index
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Index < fields[j].Index
	})

	// Generate HTML columns
	var leftColumn, rightColumn string
	var allFieldsHTML string
	displayTwo := cfg.DisplayTwoColumn
	mid := (len(fields) + 1) / 2

	for i, f := range fields {
		if !f.Show {
			continue
		}

		value := f.Value
		if value == "" {
			value = "-"
		}

    fieldHTML := fmt.Sprintf(
        `<div style="display:grid; grid-template-columns:1fr 0.01fr 1fr; gap:5px; margin:2px 0; font-size:%dpx; color:%s;">
            <div style="font-weight:bold;">%s</div>
			<span>:</span>
            <div>%s</div>
        </div>`,
        cfg.FontSize, cfg.FontColor, f.Label, value,
    )
		if displayTwo {
			if i < mid {
				leftColumn += fieldHTML
			} else {
				rightColumn += fieldHTML
			}
		} else {
			allFieldsHTML += fieldHTML
		}
	}

	// Photo HTML
	photoHTML := ""
	if cfg.ShowPhoto && cfg.PhotoBase64 != "" {
		photoHTML = fmt.Sprintf(
			`<img src="data:image/png;base64,%s" style="height:%dpx; width:%dpx; object-fit:cover; margin-bottom:10px;">`,
			cfg.PhotoBase64, cfg.StudentPhotoY, cfg.StudentPhotoX,
		)
	}
	detailsWidth := fmt.Sprintf("calc(100%% - %dpx)", cfg.StudentPhotoX)

	// Two-column layout
	if displayTwo {
		// Left content includes photo if photo is on left
		leftSideHTML := ""
		if cfg.ShowPhoto && !cfg.PhotoOnRight {
			leftSideHTML = photoHTML
		}
		leftSideHTML += fmt.Sprintf(`<div style="flex:1;">%s</div>`, leftColumn)
		rightSideHTML := fmt.Sprintf(`<div style="flex:1;">%s</div>`, rightColumn)

		contentHTML := fmt.Sprintf(`
			<div style="display:flex; align-items:flex-start; width:%s;">
				%s
				%s
			</div>
		`,detailsWidth, leftSideHTML, rightSideHTML)

		// Photo on right
		if cfg.ShowPhoto && cfg.PhotoOnRight {
			contentHTML = fmt.Sprintf(`
				<div style="display:flex; align-items:flex-start;width:%s;">
					%s
					<div>%s</div>
				</div>
			`,detailsWidth, contentHTML, photoHTML)
		}

		return contentHTML
	}

	// Single-column layout
	if cfg.ShowPhoto {
		return fmt.Sprintf(`
			<div style="display:flex; flex-direction:column; align-items:center;">
				<div>%s</div>
				<div style="align-self:flex-start; margin-top:10px;">%s</div>
			</div>
		`, photoHTML, allFieldsHTML)
	}

	return allFieldsHTML
}

func generateAcademicDetails(cfg ReportConfig, tests []Test) string {
	fontSize := cfg.TableDataFontSize
	if fontSize == 0 {
		fontSize = 12
	}

	SlcolSpan := 1
if cfg.EnableSlNo {
    SlcolSpan = 2
}

	html := fmt.Sprintf(`<html><head><style>
		table { border-collapse: collapse; width: 100%%; margin-top: 20px; font-size: %dpx; }
		th, td { border: 1px solid #000; padding: 2px; text-align: center; }
		th.subject { text-align: left !important; padding-left: 8px; }
		th { background-color: #eee; }
	</style></head><body><table>`, fontSize)

	// --- Helper: count number of columns per test ---
	getColCount := func() int {
		colCount := 0
		if !cfg.PrintOnlyGrade {
			if cfg.ShowMaxPerSubject {
				colCount++
			}
			if cfg.ShowMinPerSubject {
				colCount++
			}
			colCount++ // Obt column
		}
		if cfg.ShowGradePerSubject {
			colCount++
		}
		return colCount
	}

	// --- Helper: total columns across tests ---
	totalCols := func() int {
		base := 1
		if cfg.EnableSlNo {
			base = 2
		}
		return base + len(tests)*getColCount()
	}

	// --- Header ---
rowSpan := 1
if cfg.ShowTestName {
	rowSpan = 2
}

if cfg.EnableSlNo {
	html += fmt.Sprintf(`<th rowspan="%d">Sl</th>`, rowSpan)
}
html += fmt.Sprintf(`<th rowspan="%d" class="subject">Subject</th>`, rowSpan)

	if cfg.ShowTestName {
		for _, test := range tests {
			html += fmt.Sprintf(`<th colspan="%d">%s</th>`, getColCount(), test.Name)
		}
		html += "</tr><tr>"
		for range tests {
			if !cfg.PrintOnlyGrade {
				if cfg.ShowMaxPerSubject {
					html += "<th>Max</th>"
				}
				if cfg.ShowMinPerSubject {
					html += "<th>Min</th>"
				}
				html += "<th>Obt</th>"
			}
			if cfg.ShowGradePerSubject {
				html += "<th>Grade</th>"
			}
		}
	} else {
		for range tests {
			if !cfg.PrintOnlyGrade {
				if cfg.ShowMaxPerSubject {
					html += "<th>Max</th>"
				}
				if cfg.ShowMinPerSubject {
					html += "<th>Min</th>"
				}
				html += "<th>Obt</th>"
			}
			if cfg.ShowGradePerSubject {
				html += "<th>Grade</th>"
			}
		}
	}
	html += "</tr>"

	// --- Optional Max/Min per Test ---
	addSummaryRow := func(label string, perTestData func(test Test) int) {
		html += "<tr>"
		
html += fmt.Sprintf(`<td colspan="%d"><b>%s</b></td>`, SlcolSpan, label)
		for _, test := range tests {
			total := perTestData(test)
			html += fmt.Sprintf(`<td colspan="%d">%d</td>`, getColCount(), total)
		}
		html += "</tr>"
	}

	if cfg.ShowMaxPerTest {
		addSummaryRow("Max Marks", func(t Test) int {
			sum := 0
			for _, subj := range cfg.Subjects {
				sum += t.Max[subj]
			}
			return sum
		})
	}

	if cfg.ShowMinPerTest {
		addSummaryRow("Min Marks", func(t Test) int {
			sum := 0
			for _, subj := range cfg.Subjects {
				sum += t.Min[subj]
			}
			return sum
		})
	}

	// --- Body ---
	for i, subj := range cfg.Subjects {
		html += "<tr>"
		if cfg.EnableSlNo {
			html += fmt.Sprintf("<td>%d</td>", i+1)
		}
		html += fmt.Sprintf("<td>%s</td>", subj)

		for _, test := range tests {
			if !cfg.PrintOnlyGrade {
				if cfg.ShowMaxPerSubject {
					html += fmt.Sprintf("<td>%d</td>", test.Max[subj])
				}
				if cfg.ShowMinPerSubject {
					html += fmt.Sprintf("<td>%d</td>", test.Min[subj])
				}
				html += fmt.Sprintf("<td>%d</td>", test.Marks[subj])
			}
			if cfg.ShowGradePerSubject {
				html += "<td>A+</td>"
			}
		}
		html += "</tr>"
	}

	// --- Footer Rows ---
	addFooterRow := func(label, value string) {
		// Handle "over-all"
		if (label == cfg.RemarksText && cfg.PrintRemarks == "over-all") ||
			(label == cfg.ConductText && cfg.PrintConduct == "over-all") {
			html += fmt.Sprintf(`<tr><td colspan="%d"><b>%s</b> </td><td colspan="%d">%s</td></tr>`, SlcolSpan, label,totalCols, value)
			return
		}

		html += "<tr>"
		if cfg.EnableSlNo {
			html += fmt.Sprintf(`<td colspan="%d"><b>%s</b></td>`,SlcolSpan, label)
		} else {
			html += fmt.Sprintf(`<td><b>%s</b></td>`, label)
		}

		for range tests {
			html += fmt.Sprintf(`<td colspan="%d">%s</td>`, getColCount(), value)
		}
		html += "</tr>"
	}

	if cfg.ShowTotal {
		addFooterRow("Total", "450")
	}
	if cfg.ShowPercentage {
		addFooterRow(cfg.PercentageText, "90%")
	}
	if cfg.ShowRemarksPerTest {
		addFooterRow(cfg.RemarksText, "Excellent")
	}
	if cfg.ShowConductPerTest {
		addFooterRow(cfg.ConductText, "Good")
	}

	html += "</table></body></html>"
	return html
}


func generateStudentChartHTML(tests []Test) string {
	if len(tests) == 0 {
		return `
<div class="chart-section" style="page-break-inside: avoid; text-align: center; margin-top: 20px;">
	<p style="color: #666;">Not enough test data to generate a performance chart.</p>
</div>`
	}

	// --- Chart Config ---
	const (
		maxHeight   = 90.0 // Max bar height = 100 marks
		baseY       = 100.0 // Y position of X-axis (0 marks)
		chartHeight = 250   // Increased to accommodate legends
		barWidth    = 25.0
		barGap      = 5.0
		groupGap    = 20.0
	)

	// --- Extract Subjects Dynamically ---
	subjects := []string{}
	if len(tests[0].Marks) > 0 {
		for subject := range tests[0].Marks {
			subjects = append(subjects, subject)
		}
	}
	sort.Strings(subjects)

	// --- Color Palette (auto loops if more tests) ---
	colors := []string{
		"#1A237E", "#4CAF50", "#F44336", "#FF9800", "#9C27B0",
		"#03A9F4", "#795548", "#009688", "#E91E63", "#607D8B",
	}

	// --- Helper: Axis Drawing (Y intervals: 0, 25, 50, 75, 100) ---
// --- Helper: Axis Drawing (Y intervals: 0, 25, 50, 75, 100) ---
buildAxes := func() string {
    var ticks strings.Builder
    for i := 0; i <= 4; i++ {
        value := float64(i) * 25
        y := baseY - (value * (maxHeight / 100.0))
        
        // Add the Y-axis label and a small tick mark
        ticks.WriteString(fmt.Sprintf(
            `<text x="30" y="%.0f" font-size="12" fill="#666">%d</text>
              <line x1="50" y1="%.0f" x2="55" y2="%.0f" stroke="#999" stroke-width="1"/>`,// y line  small ticks
            y+4, int(value), y, y,
        ))

        // ADDITION: Check if value is 50 (i=2) to draw the full horizontal grid line
        if value == 50.0 {
            ticks.WriteString(fmt.Sprintf(
                `<line x1="50" y1="%.0f" x2="680" y2="%.0f" stroke="#ccc" stroke-width="1" stroke-dasharray="4,4"/>`,
                y, y,
            ))
        }
    }

    return fmt.Sprintf(`
        <line x1="50" y1="%.0f" x2="50" y2="10" stroke="#999" stroke-width="1" /> // <-- FIXED Y2 HERE
        %s

        <line x1="50" y1="%.0f" x2="680" y2="%.0f" stroke="#999" stroke-width="1" />// bottom line
    `, baseY, ticks.String(), baseY, baseY)
}
	var barGroup, xLabels, legends strings.Builder

	// --- Draw Bars and Labels ---
	for i, subj := range subjects {
		groupStart := 60.0 + float64(i)*(float64(len(tests))*(barWidth+barGap) + groupGap)

		for j, test := range tests {
			color := colors[j%len(colors)]
			mark := test.Marks[subj]
			h := float64(mark) * maxHeight / 100.0
			x := groupStart + float64(j)*(barWidth+barGap)

			barGroup.WriteString(fmt.Sprintf(
				`<rect x="%.0f" y="%.0f" width="%.0f" height="%.0f" fill="%s" />
				 <text x="%.0f" y="%.0f" font-size="10" fill="#333" text-anchor="middle">%d</text>`,//sub names
				x, baseY-h, barWidth, h, color, x+(barWidth/2), baseY-h-5, mark,
			))
		}

		groupCenter := groupStart + float64(len(tests))*(barWidth+barGap)/2
		xLabels.WriteString(fmt.Sprintf(
			`<text x="%.0f" y="%.0f" font-size="10" fill="#000" font-weight="bold" text-anchor="middle">%s</text>`,
			groupCenter, baseY+15, subj,
		))
	}

	// --- Legend (Below the Chart, Centered) ---
	legendStartX := 50.0
	legendY := baseY + 25
	for i, test := range tests {
		color := colors[i%len(colors)]
		x := legendStartX + float64(i)*100.0
		legends.WriteString(fmt.Sprintf(
			`<rect x="%.0f" y="%.0f" width="12" height="12" fill="%s" />
			 <text x="%.0f" y="%.0f" font-size="12" fill="#333">%s</text>`,
			x, legendY, color, x+20, legendY+10, test.Name,
		))
	}

	// --- Return Final HTML ---
	return fmt.Sprintf(`
<div class="chart-section" style="page-break-inside: avoid; text-align: center; margin-top: 20px;">
	<div style="width:100%%; height: %dpx; max-width: 95%%; margin: 10px auto;">
		<svg width="100%%" height="100%%" viewBox="0 0 800 %d" xmlns="http://www.w3.org/2000/svg">
			<g transform="translate(10, 10)">
				%s
				%s
				%s
				%s
			</g>
		</svg>
	</div>
</div>`,
		chartHeight, chartHeight,
		buildAxes(),
		barGroup.String(),
		xLabels.String(),
		legends.String(),
	)
}


// -----------------------------
// PDF GENERATION
// -----------------------------
func generatePDF(htmlContent string, filename string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	tmpFile := "temp.html"
	if err := os.WriteFile(tmpFile, []byte(htmlContent), 0644); err != nil {
		return err
	}
	defer os.Remove(tmpFile)

	absPath, err := filepath.Abs(tmpFile)
	if err != nil {
		return err
	}
	url := "file://" + absPath

	var pdfBuf []byte
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				WithPaperWidth(8.27).  // A4 width
				WithPaperHeight(11.7). // A4 height
				Do(ctx)
			if err != nil {
				return err
			}
			pdfBuf = buf
			return nil
		}),
	); err != nil {
		return err
	}

	return os.WriteFile(filename, pdfBuf, 0644)
}





// Helper function 
func columnCount(cfg ReportConfig) int {
	count := 2 // Subject + Obt.
	if cfg.EnableSlNo {
		count++
	}

	return count
}
func countEnabledColumns(cfg ReportConfig) int {
	count := 1 // Obtained is always shown
	if cfg.ShowMaxPerSubject {
		count++
	}
	if cfg.ShowMinPerSubject {
		count++
	}
	if cfg.ShowGradePerSubject {
		count++
	}

	return count
}



func addFooterRow(html *string, cfg ReportConfig, label string, value any) {
	colSpan := columnCount(cfg)



	dataColSpan := colSpan - 2
	*html += "<tr>"

	if cfg.EnableSlNo {
		// Merge SlNo + Subject columns
		*html += fmt.Sprintf("<td colspan='2'><b>%s</b></td>", label)
	} else {
		// Only Subject column
		*html += fmt.Sprintf("<td><b>%s</b></td>", label)
		dataColSpan = colSpan - 1
	}

	// Value cell (spanning the rest)
	*html += fmt.Sprintf("<td colspan='%d'>%v</td>", dataColSpan, value)
	*html += "</tr>"
}


func newStudent(base StudentDetailsConfig, name, rollNo, father, mother, class, year, dob, attendance, address, email, mobile string, showPhoto bool) StudentDetailsConfig {
	cfg := base
	cfg.StudentName = name
	cfg.StudentRollNo = rollNo
	cfg.FatherName = father
	cfg.MotherName = mother
	cfg.StudentClass = class
	cfg.AcademicYear = year
	cfg.DateOfBirth = dob
	cfg.AttendanceStats = attendance
	cfg.Address = address
	cfg.Email = email
	cfg.Mobile = mobile
	cfg.ShowPhoto = showPhoto
	return cfg
}