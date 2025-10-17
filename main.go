package main

import (
	"fmt"
	"log"
"strings"
"github.com/disintegration/imaging"
	"github.com/jung-kurt/gofpdf"
)
type TestColumn struct {
	Name     string   // e.g. "UT 1"
	SubCols  []string // e.g. ["Max", "Min", "Obt", "Grade"]
	Flag  []bool   // e.g. [true, false, true, true] to hide Min
}

type SubjectData struct {
	SlNo    string
	Subject string
	Tests   map[string][]string // key = test name, value = only visible fields
}

type TotalFooterRow struct {
	Label   string
	Flag bool
	Values  map[string]string // test -> value
	VisibleCols map[string]bool
}

type FooterRow struct {
	Label   string
	Flag bool
	Values  map[string]string // test -> value

}


type ScholasticConfig struct {
	Title       string
	FontSize    float64
	Margin      float64
	TestColumns []TestColumn
	Subjects    []SubjectData
	TotalFooter      []TotalFooterRow
	Header []FooterRow
	Footer []FooterRow
	ShowMaxPerTest bool
	ShowMinPerTest bool

	ShowMaxPerSubject    bool
	ShowMinPersubject   bool
	showGradePerSubject bool
	ShowRemarksPerTest bool
	ShowConductPerTest bool
	ShowPercentage bool
	ShowGradePerTest bool
    ShowTotalsOfMaxMin bool

	ShowOverAllRemarks bool
	ShowOverAllConduct bool

	
}


func main() {
	config := ScholasticConfig{
		FontSize:           8,
		Margin:             10,
		Title:              "PART I - SCHOLASTIC AREA",
		ShowMaxPerTest : false,
		ShowMinPerTest: true,
		ShowMaxPerSubject:  true,
		ShowMinPersubject:  false,
		showGradePerSubject: true,
		ShowRemarksPerTest: true,
		ShowConductPerTest: true,
		ShowPercentage:     true,
		ShowGradePerTest:   true,
		ShowTotalsOfMaxMin: true,
		ShowOverAllRemarks: true,
		ShowOverAllConduct: false,
	
	}

		config.TestColumns = []TestColumn{
		{
			Name:    "UT 1",
			SubCols: []string{"Max", "Min", "Obt", "Grade"},
			Flag:    []bool{config.ShowMaxPerSubject, config.ShowMinPersubject, true, config.showGradePerSubject},
		},
		{
			Name:    "UT 2",
			SubCols: []string{"Max", "Min", "Obt", "Grade"},
			Flag:    []bool{config.ShowMaxPerSubject, config.ShowMinPersubject, true, config.showGradePerSubject},
		},
		{
			Name:    "UT 3",
			SubCols: []string{"Max", "Min", "Obt", "Grade"},
			Flag:    []bool{config.ShowMaxPerSubject, config.ShowMinPersubject, true, config.showGradePerSubject},
		},
	}

		config.Subjects = []SubjectData{
		{
			SlNo:    "1",
			Subject: "Mathematics",
			Tests: map[string][]string{
				"UT 1": {"30", "10", "25", "A"},
				"UT 2": {"30", "10", "28", "A+"},
				"UT 3": {"30", "10", "22", "B"},
			},
		},
		{
			SlNo:    "2",
			Subject: "Science",
			Tests: map[string][]string{
				"UT 1": {"30", "10", "18", "B"},
				"UT 2": {"30", "10", "26", "A"},
				"UT 3": {"30", "10", "20", "B"},
			},
		},
	}
	var totalFooter TotalFooterRow
	if config.ShowTotalsOfMaxMin {
		totalFooter = TotalFooterRow{
			Label: "Total",
			Values: map[string]string{
				"UT 1_Max":   "60",
				"UT 1_Min":   "20",
				"UT 1_Obt":   "43",
				"UT 1_Grade": "A",
				"UT 2_Max":   "60",
				"UT 2_Min":   "20",
				"UT 2_Obt":   "54",
				"UT 2_Grade": "A+",
				"UT 3_Max":   "60",
				"UT 3_Min":   "20",
				"UT 3_Obt":   "42",
				"UT 3_Grade": "A",
			},
			VisibleCols: map[string]bool{
				"UT 1_Max":   config.ShowMaxPerSubject,
				"UT 1_Min":   config.ShowMinPersubject,
				"UT 1_Obt":   true,
				"UT 1_Grade": config.ShowGradePerTest,
				"UT 2_Max":    config.ShowMaxPerSubject,
				"UT 2_Min":    config.ShowMinPersubject,
				"UT 2_Obt":   true,
				"UT 2_Grade": config.ShowGradePerTest,
				"UT 3_Max":    config.ShowMaxPerSubject,
				"UT 3_Min":   config.ShowMinPersubject,
				"UT 3_Obt":   true,
				"UT 3_Grade": config.ShowGradePerTest,
			},
		}
	} else {
		totalFooter = TotalFooterRow{
			Label: "Total",
			Values: map[string]string{
				"UT 1": "50",
				"UT 2": "54",
				"UT 3": "45",
			},
			Flag: true,
		}
	}

	config.TotalFooter = []TotalFooterRow{totalFooter}

	config.Header = []FooterRow{

	{
		Label: "Max",
		Values: map[string]string{
			"UT 1": "60",
			"UT 2": "60",
			"UT 3": "60",
		},
		Flag: config.ShowMaxPerTest,
	},
	{
		Label: "Min",
		Values: map[string]string{
			"UT 1": "10",
			"UT 2": "12",
			"UT 3": "15",
		},
		Flag: config.ShowMinPerTest,
	},

}
config.Footer = []FooterRow{


	{
		Label: "Percentage",
		Values: map[string]string{
			"UT 1": "83%", 
			"UT 2": "90%",
			"UT 3": "75%",
		},
		Flag: true,
	},
	{
		Label: "Grade",
		Values: map[string]string{
			"UT 1": "A",
			"UT 2": "A+",
			"UT 3": "A",
		},
		Flag: true,
	},
	{
		Label: "Remarks",
		Values: map[string]string{
			"UT 1": "Good",
			"UT 2": "Very Good",
			"UT 3": "Good",
		},
		Flag: true,
	},
		{
		Label: "Conduct",
		Values: map[string]string{
			"UT 1": "Good",
			"UT 2": "Very Good",
			"UT 3": "Good",
		},
		Flag: true,
	},
}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)
	pdf.SetTopMargin(10)
	pdf.SetAutoPageBreak(true, 10)

	decoratePage(pdf)
	addHeader(pdf)
	addTitle(pdf)
	addStudentInfo(pdf)
	addScholasticArea(pdf, config)
	addCoScholasticArea(pdf)
	addScholasticGraph(pdf)
	addGradeDetailsHorizontal(pdf)
	addFooter(pdf)

	err := pdf.OutputFileAndClose("report_card.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("✅ PDF generated successfully: report_card.pdf")
}


func decoratePage(pdf *gofpdf.Fpdf) {
	pageWidth, pageHeight := pdf.GetPageSize()
	bgPath := "assets/background.png"
	watermarkPath := "assets/colorwatermark.png"

	// ---- Load Background ----
	bgImg, err := imaging.Open(bgPath)
	if err != nil {
		log.Println("Error loading background:", err)
		return
	}

	// ---- Load Watermark ----
	wmImg, err := imaging.Open(watermarkPath)
	if err != nil {
		log.Println("Error loading watermark:", err)
		return
	}

	// ---- Resize watermark ----
	wmImg = imaging.Resize(wmImg, int(150), 0, imaging.Lanczos)

	wmImg = imaging.Blur(wmImg, 1.0) 

	
	blended := imaging.OverlayCenter(bgImg, wmImg, 0.6)

	// ---- Save temporary blended background ----
	tmpPath := "assets/tmp_blended_page.png"
	err = imaging.Save(blended, tmpPath)
	if err != nil {
		log.Println("Error saving blended image:", err)
		return
	}

	// ---- Draw final blended image on page ----
	pdf.ImageOptions(
		tmpPath,
		0, 0,
		pageWidth, pageHeight,
		false,
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
		0,
		"",
	)

	// ---- Optional Border ----
	margin := 8.0
	pdf.SetLineWidth(0.8)
	pdf.SetDrawColor(50, 50, 150)
	pdf.Rect(margin, margin, pageWidth-2*margin, pageHeight-2*margin, "D")
}
func addHeader(pdf *gofpdf.Fpdf) {
	// Define image paths
	logoPath := "./assets/Arcadis_Logo.png"
	photo1Path := "./assets/Arcadis_Logo.png"
	photo2Path := "./assets/Arcadis_Logo.png"

	// Page width (A4 = 210mm)
	pageWidth, _ := pdf.GetPageSize()
	margin := 10.0
	imgWidth := 10.0

	// Y-position for header images
	y := pdf.GetY()

	// Left photo
	pdf.ImageOptions(
		photo1Path,
		margin,
		y,
		imgWidth,
		0,
		false,
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
		0,
		"",
	)

	// Right photo
	pdf.ImageOptions(
		photo2Path,
		pageWidth-margin-imgWidth,
		y,
		imgWidth,
		0,
		false,
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
		0,
		"",
	)

	// Center logo
	centerX := (pageWidth / 2) - (imgWidth / 2)
	pdf.ImageOptions(
		logoPath,
		centerX,
		y,
		imgWidth,
		0,
		false,
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
		0,
		"",
	)

	// Move below the images
	pdf.SetY(y + imgWidth + 5)

	// Text content
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 8, "My-Eduate", "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 6, "User Group of Institutions", "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 6)
	pdf.CellFormat(0, 6, "Anjanadri School", "", 1, "C", false, 0, "")
	pdf.CellFormat(0, 6, "anjanadri@myeduate.com", "", 1, "C", false, 0, "")
	pdf.Ln(1)
	// Draw bottom border (horizontal line)
	x1 := margin
	x2 := pageWidth - margin
	yLine := pdf.GetY() // current Y after header
	pdf.SetDrawColor(0, 0, 0) // black line
	pdf.SetLineWidth(0.2)
	pdf.Line(x1, yLine, x2, yLine)
	pdf.Ln(2) // small gap after line for next section
}

func addTitle(pdf *gofpdf.Fpdf) {
// title
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 6, "Title", "", 1, "C", false, 0, "")
	pdf.Ln(2)
	// subtitle
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(0, 4, "Subtitle", "", 1, "C", false, 0, "")
	pdf.Ln(1)

}
func addStudentInfo(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "", 8)

	// Left column data (Label, Value)
	leftColumn := [][]string{
		{"Student Name", "AMRUTHAVARSHINI M"},
		{"Adm No.", "AS-51"},
		{"Father's / Guardian's Name", "MANJUNATHA"},
		{"Mother's Name", "LAKSHMIDEVI"},
	}

	// Right column data (Label, Value)
	rightColumn := [][]string{
		{"Class", "1 & A"},
		{"Academic Year", "2025-2026"},
		{"Attendance", ""},
		{"Date of Birth", "02-08-2005"},
	}

	cellHeight := 6.0
	margin := 10.0
	pageWidth, _ := pdf.GetPageSize()
	columnWidth := (pageWidth - 2*margin) / 2
	gap := 20.0

	labelWidth := 45.0  // width before colon
	colonWidth := 3.0   // width for the colon itself

	for i := 0; i < len(leftColumn); i++ {
		// ---- LEFT COLUMN ----
		pdf.SetX(margin)
		pdf.CellFormat(labelWidth, cellHeight, leftColumn[i][0], "", 0, "L", false, 0, "")
		pdf.CellFormat(colonWidth, cellHeight, ":", "", 0, "C", false, 0, "")
		pdf.CellFormat(columnWidth-labelWidth-colonWidth-gap, cellHeight, leftColumn[i][1], "", 0, "L", false, 0, "")

		// ---- RIGHT COLUMN ----
		pdf.SetX(margin + columnWidth)
		pdf.CellFormat(labelWidth, cellHeight, rightColumn[i][0], "", 0, "L", false, 0, "")
		pdf.CellFormat(colonWidth, cellHeight, ":", "", 0, "C", false, 0, "")
		pdf.CellFormat(columnWidth-labelWidth-colonWidth, cellHeight, rightColumn[i][1], "", 1, "L", false, 0, "")
	}

	pdf.Ln(1)
}


func addScholasticArea(pdf *gofpdf.Fpdf, cfg ScholasticConfig) {
	pdf.SetFont("Arial", "B", cfg.FontSize)
	pdf.CellFormat(0, 6, cfg.Title, "", 1, "L", false, 0, "")
	pdf.Ln(1)

	slWidth := 8.0
	subjectWidth := 50.0

	pageWidth, _ := pdf.GetPageSize()
	effectiveWidth := pageWidth - (2 * cfg.Margin)
	fixedWidth := slWidth + subjectWidth

	// count total visible subcolumns
	totalSubCols := 0
	for _, t := range cfg.TestColumns {
		for _, visible := range t.Flag {
			if visible {
				totalSubCols++
			}
		}
	}

	testWidth := effectiveWidth - fixedWidth
	testColWidth := testWidth / float64(totalSubCols)

	// --- Header Row 1 ---
	pdf.CellFormat(slWidth, 10, "Sl", "1", 0, "C", false, 0, "")
	pdf.CellFormat(subjectWidth, 10, "Subject", "1", 0, "C", false, 0, "")

	for _, t := range cfg.TestColumns {
		visibleCount := 0
		for _, v := range t.Flag {
			if v {
				visibleCount++
			}
		}
		pdf.CellFormat(testColWidth*float64(visibleCount), 5, t.Name, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// --- Header Row 2 ---
	pdf.CellFormat(slWidth, 0, "", "0", 0, "", false, 0, "")
	pdf.CellFormat(subjectWidth, 0, "", "0", 0, "", false, 0, "")

	for _, t := range cfg.TestColumns {
		for i, sub := range t.SubCols {
			if t.Flag[i] {
				pdf.CellFormat(testColWidth, 5, sub, "1", 0, "C", false, 0, "")
			}
		}
	}
	pdf.Ln(-1)

	// max _min rows per test
	pdf.SetFont("Arial", "B", cfg.FontSize)
	for _, f := range cfg.Header {
		if !f.Flag {
			continue // skip hidden footers
		}
		pdf.CellFormat(fixedWidth, 5, f.Label, "1", 0, "L", false, 0, "")
		for _, t := range cfg.TestColumns {
			visibleCount := 0
			for _, v := range t.Flag {
				if v {
					visibleCount++
				}
			}
			pdf.CellFormat(testColWidth*float64(visibleCount), 5, f.Values[t.Name], "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)
	}

	// --- Table Body ---
	pdf.SetFont("Arial", "", cfg.FontSize)
	for _, s := range cfg.Subjects {
		pdf.CellFormat(slWidth, 5, s.SlNo, "1", 0, "C", false, 0, "")
		pdf.CellFormat(subjectWidth, 5, s.Subject, "1", 0, "L", false, 0, "")

		for _, t := range cfg.TestColumns {
			values := s.Tests[t.Name]
			subIndex := 0
			for _, v := range t.Flag {
				if v {
					pdf.CellFormat(testColWidth, 5, values[subIndex], "1", 0, "C", false, 0, "")
				}
				subIndex++
			}
		}
		pdf.Ln(-1)
	}

	// ---totalFooter Rows ---
	pdf.SetFont("Arial", "B", cfg.FontSize)
if cfg.ShowTotalsOfMaxMin {
	// ✅ Show detailed totals (Max, Min, Obt, Grade per test)
	if len(cfg.TotalFooter) > 0 {
		totalFooter := cfg.TotalFooter[0]

		pdf.CellFormat(slWidth+subjectWidth, 5, totalFooter.Label, "1", 0, "L", false, 0, "")

		for _, test := range cfg.TestColumns {
			for i, subCol := range test.SubCols {
				if len(test.Flag) > i && test.Flag[i] {
					key := fmt.Sprintf("%s_%s", test.Name, subCol)
					value := totalFooter.Values[key]
					pdf.CellFormat(testColWidth, 5, value, "1", 0, "C", false, 0, "")
				}
			}
		}
		pdf.Ln(-1)
	}

} else {
	// ✅ Show simple footer rows (Obtained total only)
	for _, f := range cfg.TotalFooter {
		if !f.Flag {
			continue // skip hidden footer rows
		}

		pdf.CellFormat(slWidth+subjectWidth, 5, f.Label, "1", 0, "L", false, 0, "")

		for _, t := range cfg.TestColumns {
			visibleCount := 0
			for _, v := range t.Flag {
				if v {
					visibleCount++
				}
			}

			pdf.CellFormat(testColWidth*float64(visibleCount), 5, f.Values[t.Name], "1", 0, "C", false, 0, "")
		}

		pdf.Ln(-1)
	}
}

// otherfooter rows
	pdf.SetFont("Arial", "B", cfg.FontSize)


for _, f := range cfg.Footer {
	if !f.Flag {
		continue // skip hidden footers
	}

	pdf.CellFormat(fixedWidth, 5, f.Label, "1", 0, "L", false, 0, "")

	// --- check if this footer should merge across all tests ---
	shouldMerge := (cfg.ShowOverAllRemarks && strings.EqualFold(f.Label, "Remarks")) ||
		(cfg.ShowOverAllConduct && strings.EqualFold(f.Label, "Conduct"))

	if shouldMerge {
		// Calculate total width for all visible test columns
		totalVisibleWidth := 0.0
		for _, t := range cfg.TestColumns {
			for _, v := range t.Flag {
				if v {
					totalVisibleWidth += testColWidth
				}
			}
		}

		// Pick value (use UT 1 or first available)
		value := ""
		if val, ok := f.Values["UT 1"]; ok {
			value = val
		} else if len(f.Values) > 0 {
			for _, v := range f.Values {
				value = v
				break
			}
		}

		// Draw one merged cell
		pdf.CellFormat(totalVisibleWidth, 5, value, "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
		continue
	}

	// --- normal footer layout: per test ---
	for _, t := range cfg.TestColumns {
		visibleCount := 0
		for _, v := range t.Flag {
			if v {
				visibleCount++
			}
		}
		pdf.CellFormat(testColWidth*float64(visibleCount), 5, f.Values[t.Name], "1", 0, "C", false, 0, "")
	}

	pdf.Ln(-1)
}

}


func addCoScholasticArea(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(0, 6, "PART II - CO-SCHOLASTIC AREA", "", 1, "", false, 0, "")
	pdf.Ln(2)

	headers := []string{"Sl", "Subject", "Grade", "Remarks"}

	// --- Dynamic widths to fit page ---
	margin := 10.0
	pageWidth, _ := pdf.GetPageSize()
	totalWidth := pageWidth - 2*margin

	// Width ratios: Sl=8%, Subject=42%, Grade=15%, Remarks=35%
	widthRatios := []float64{0.08, 0.42, 0.15, 0.35}
	widths := make([]float64, len(widthRatios))
	for i, ratio := range widthRatios {
		widths[i] = totalWidth * ratio
	}

	// --- Header ---
	pdf.SetFont("Arial", "B", 8)
	headerAlign := []string{"C", "C", "C", "C"} // all centered
	for i, h := range headers {
		pdf.CellFormat(widths[i], 5, h, "1", 0, headerAlign[i], false, 0, "")
	}
	pdf.Ln(-1)

	// --- Data ---
	pdf.SetFont("Arial", "", 8)
	data := [][]string{
		{"1", "My Pronunciation is", "B", "BAD"},
		{"2", "I am Independent", "A+", ""},
		{"3", "I listen to instructions", "B", ""},
		{"4", "I can sing & dance", "C", ""},
	}

	for _, row := range data {
		for i, text := range row {
			align := "C" // default center
			if i == 1 {  // Subject column left-aligned
				align = "L"
			}
			pdf.CellFormat(widths[i], 5, text, "1", 0, align, false, 0, "")
		}
		pdf.Ln(-1)
	}

	pdf.Ln(2)
}
func addScholasticGraph(pdf *gofpdf.Fpdf) {
	// --- JSON Data ---
	type SubjectData struct {
		SlNo    string
		Subject string
		UT1     []string
		UT2     []string
		UT3     []string
	}
	data := []SubjectData{
		{"1", "KANNADA", []string{"30", "10", "20", "B"}, []string{"30", "10", "25", "A"}, []string{"30", "10", "22", "B"}},
		{"2", "ENGLISH", []string{"30", "10", "14", "C"}, []string{"30", "10", "18", "B"}, []string{"30", "10", "16", "C"}},
		{"3", "MATHEMATICS", []string{"30", "10", "15", "C"}, []string{"30", "10", "20", "B"}, []string{"30", "10", "18", "B"}},
		{"4", "GENERAL SCIENCE", []string{"30", "10", "20", "B"}, []string{"30", "10", "22", "B"}, []string{"30", "10", "24", "A"}},
		{"5", "SOCIAL STUDIES", []string{"30", "10", "14", "C"}, []string{"30", "10", "25", "A"}, []string{"30", "10", "26", "A"}},
	}

	// --- Chart Settings ---
	startX := 20.0
	startY := pdf.GetY() + 10
	chartHeight := 25.0
	barWidth := 8.0
	barGap := 0.1
	groupGap := 5.0
	maxMarks := 100.0 // Maximum marks per test
interval := 25.0

	bottomY := startY + chartHeight
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(0, 10, "Scholastic Area Graph", "", 1, "C", false, 0, "")
	pdf.Ln(2)

	// --- Draw Axis ---
	pdf.SetDrawColor(0, 0, 0)
	pdf.Line(startX, startY, startX, bottomY) // Y-axis
	pdf.Line(startX, bottomY, startX+180, bottomY) // X-axis

	// --- Y-axis Labels ---
	pdf.SetFont("Arial", "", 8)
for i := 0.0; i <= maxMarks; i += interval {
	y := bottomY - (i/maxMarks)*chartHeight
	pdf.Line(startX-1, y, startX+1, y)       // small tick line
	pdf.Text(startX-10, y+2, fmt.Sprintf("%.0f", i)) // label
}

	// --- Colors ---
	colors := [][]int{
		{79, 129, 189}, // UT1
		{192, 80, 77},  // UT2
		{155, 187, 89}, // UT3
	}

	// --- Draw Bars Subject-wise (Single Row) ---
	currentX := startX + 10
	pdf.SetFont("Arial", "", 6)

	for _, subj := range data {
		groupStartX := currentX
		tests := [][]string{subj.UT1, subj.UT2, subj.UT3}

		for testIdx, t := range tests {
			obtainedMarks := 0.0
			fmt.Sscanf(t[2], "%f", &obtainedMarks) // t[2] is "Obt" marks

			barHeight := (obtainedMarks / maxMarks) * chartHeight
			rgb := colors[testIdx%len(colors)]
			pdf.SetFillColor(rgb[0], rgb[1], rgb[2])
			x := currentX
			y := bottomY - barHeight
			pdf.Rect(x, y, barWidth, barHeight, "F")
			currentX += barWidth + barGap
		}

		// --- Subject Label below bars ---
		labelWidth := float64(len(tests))*(barWidth+barGap) - barGap
		pdf.SetXY(groupStartX, bottomY+2)
		pdf.MultiCell(labelWidth, 4, subj.Subject, "", "C", false)

		currentX += groupGap
	}
	// --- Legend ---

legendY := bottomY + 10
pdf.SetFont("Arial", "", 6)
testNames := []string{"UT1", "UT2", "UT3"}
legendGap := 20.0
rectSize := 4.0

// Calculate total legend width
totalLegendWidth := float64(len(testNames)-1)*legendGap + float64(len(testNames))*rectSize + float64(len(testNames))*6 // 6 = approx width of text

// Start X so legend is centered
legendX := startX + (180-totalLegendWidth)/2 // 180 = chart width

for i, test := range testNames {
	rgb := colors[i%len(colors)]
	pdf.SetFillColor(rgb[0], rgb[1], rgb[2])
	pdf.Rect(legendX, legendY, rectSize, rectSize, "F")
	pdf.Text(legendX+6, legendY+3, test)
	legendX += rectSize + 6 + legendGap // move to next legend item
	pdf.Ln(2)
}
pdf.Ln(2)
}


func addGradeDetailsVertical(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(0, 8, "Grade Details", "", 1, "", false, 0, "")
	pdf.Ln(2)

	headers := []string{"Grade", "Marks-Range"}
	widths := []float64{40, 60}
	pdf.SetFont("Arial", "B", 11)
	for i, h := range headers {
		pdf.CellFormat(widths[i], 8, h, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 6)
	data := [][]string{
		{"A+", "90.01 - 100%"},
		{"A1", "80.01 - 90%"},
		{"B+", "70.01 - 80%"},
		{"B", "60.01 - 70%"},
		{"C+", "50.01 - 60%"},
		{"C", "40.01 - 50%"},
		{"D", "32.01 - 40%"},
		{"E", "0 - 32%"},
	}
	for _, row := range data {
		for i, text := range row {
			pdf.CellFormat(widths[i], 8, text, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)
	}
	pdf.Ln(1)
}
func addGradeDetailsHorizontal(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(0, 5, "Grade Details", "", 1, "", false, 0, "")
	pdf.Ln(2)

	// --- Data ---
	data := [][]string{
		{"A+", "A1", "B+", "B", "C+", "C", "D", "E"},
		{"90.01 - 100%", "80.01 - 90%", "70.01 - 80%", "60.01 - 70%",
			"50.01 - 60%", "40.01 - 50%", "32.01 - 40%", "0 - 32%"},
	}

	numCols := len(data[0])

	// --- Calculate column widths to fit page ---
	margin := 10.0
	pageWidth, _ := pdf.GetPageSize()
	totalWidth := pageWidth - 2*margin
	colWidth := totalWidth / float64(numCols)

	// --- Header Row (Grades) ---
	pdf.SetFont("Arial", "B", 8)
	for _, grade := range data[0] {
		pdf.CellFormat(colWidth, 6, grade, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// --- Data Row (Marks Range) ---
	pdf.SetFont("Arial", "", 8)
	for _, markRange := range data[1] {
		pdf.CellFormat(colWidth, 5, markRange, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(5)
}

func addFooter(pdf *gofpdf.Fpdf) {
	pageWidth, pageHeight := pdf.GetPageSize()
	bottomMargin := 20.0
	cellHeight := 8.0
	startX := 10.0
	signCount := 3
	signWidth := (pageWidth - 2*startX) / float64(signCount)
	y := pageHeight - bottomMargin

	signatures := []string{"Class Teacher", "Principal", "Parent"}

	pdf.SetFont("Arial", "", 10)

	for i, sign := range signatures {
		x := startX + float64(i)*signWidth

		// --- Draw only top dashed line for this signature block ---
		// dashLen := 4.0
		// spaceLen := 2.0
		// lineX := x
		// for lineX < x+signWidth {
		// 	endX := lineX + dashLen
		// 	if endX > x+signWidth {
		// 		endX = x + signWidth
		// 	}
		// 	pdf.Line(lineX, y, endX, y) // top border only
		// 	lineX += dashLen + spaceLen
		// }

		// --- Add signature label below the line ---
		pdf.SetXY(x, y+2)
		pdf.CellFormat(signWidth, cellHeight, sign+" Signature", "", 0, "C", false, 0, "")
	}
}





