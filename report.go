package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
"sort"
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
	PrintHeader           string
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


// -----------------------------
// MAIN
// -----------------------------
func main() {
	cfg := PageDecorationConfig{
		ShowBackground:   true,
		BackgroundImage:  toBase64("./assets/background.png"),
		ShowWatermark:    true,
		WatermarkImage:   toBase64("./assets/watermark.jpeg"),
		ShowBorder:       true,
		BorderColor:      "#333",
		BorderWidth:      2,
		BorderType:       "solid",
		MarginTop:        "20px",
		MarginRight:      "20px",
		MarginBottom:     "20px",
		MarginLeft:       "20px",
		
		
	}

	instCfg := InstitutionDetailsConfig{
		PrintPhoto1Config:     true,
		PrintPhoto2Config:     true,
		PrintInstLogo:         true,
		PrintHeader:           "EDUATE PRIVATE LIMITED",
		EnableHeader:          true,
		Photo1Base64:          toBase64("./assets/Arcadis_Logo.png"),
		Photo2Base64:          toBase64("./assets/Arcadis_Logo.png"),
		InstLogoBase64:        toBase64("./assets/Arcadis_Logo.png"),
		HeaderColor:           "#1A237E",
		HeaderFontSize:        20,
		PrintCustName:         true,
		CustName:              "Eduate ERP System",
		CustomerNameColor:     "#000",
		CustomerNameFontSize:  16,
		PrintInstName:         true,
		InstName:              "Eduate International School",
		InstNameColor:         "#111",
		InstNameFontSize:      18,
		EnableAffiliated:      true,
		PrintAffiliatedTo:     "Affiliated to CBSE",
		AffiliatedColor:       "#444",
		AffiliatedFontSize:    14,
		PrintInstAddress:      true,
		InstAddress:           "123 Eduate Street",
		InstPlace:             "Bangalore",
		InstPin:               "560001",
		AddressColor:          "#555",
		AddressFontSize:       12,
		PrintInstWebsite:      true,
		InstWebsite:           "www.eduate.com",
		WebsiteColor:          "#444",
		WebsiteFontSize:       12,
		PrintInstEmail:        true,
		InstEmail:             "info@eduate.com",
		EmailColor:            "#444",
		EmailFontSize:         12,
	}
	titleCfg := TitleConfig{
		TitleColor:       "#1a237e",
		TitleFontSize:    22,
		SubtitleColor:    "#424242",
		SubtitleFontSize: 16,
		TitleText : "Title",
		SubTitleText:"subtitle",
			EnableTitle:true,
	EnableSubTitle:true,
	}
studentCfg := StudentDetailsConfig{
	StudentName:      "John Doe",
	StudentRollNo:    "45",
	FatherName:       "Mr. Doe",
	MotherName:       "Mrs. Doe",
	StudentClass:     "10-A",
	AcademicYear:     "2024-25",
	DateOfBirth:      "01-01-2010",
	AttendanceStats:  "95%",
	Address:          "123 Main Street",
	Email:            "john@example.com",
	Mobile:           "9999999999",
	ShowPhoto:        true,
	PhotoBase64:      toBase64("./assets/colorwatermark.png"),
	PhotoOnRight:     false,
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
	FontSize:         14,
	FontColor:        "#000",
	DisplayTwoColumn: true,
	StudentPhotoX :80,
	StudentPhotoY:80,
	

}
	html := generateHTML(cfg, instCfg,titleCfg,studentCfg)

	if err := generatePDF(html, "report.pdf"); err != nil {
		log.Fatal("❌ PDF generation failed:", err)
	}
	fmt.Println("✅ PDF generated successfully: report.pdf")
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
func generateHTML(cfg PageDecorationConfig, instCfg InstitutionDetailsConfig , titleCfg TitleConfig,studentCfg StudentDetailsConfig) string {
	instHTML := generateInstitutionDetailsHTML(instCfg)
	titleHTML := generateTitleHTML(titleCfg)
studentDetailsHTML := generateStudentDetailsHTML(studentCfg)
	var bgCSS, wmCSS string
	if cfg.ShowBackground && cfg.BackgroundImage != "" {
		bgCSS = fmt.Sprintf(`background-image: url('data:image/png;base64,%s'); background-repeat: no-repeat; background-size: cover; background-position: center;`, cfg.BackgroundImage)
	}
	if cfg.ShowWatermark && cfg.WatermarkImage != "" {
		wmCSS = fmt.Sprintf(`background-image: url('data:image/jpeg;base64,%s'); background-repeat: no-repeat; background-size: 30%%; background-position: center center; `, cfg.WatermarkImage)
	}
	borderStyle := "none"
	if cfg.ShowBorder {
		borderStyle = fmt.Sprintf("%.1fpx %s %s", cfg.BorderWidth, cfg.BorderType, cfg.BorderColor)
	}

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
.page-wrapper {
	width: 100%%;
	height: 100%%;
	border: %s;
	padding: %s %s %s %s;
	box-sizing: border-box;
	background-color: white;
	%s
}
.content-with-watermark {
	width: 100%%;
	height: 100%%;
	box-sizing: border-box;
	position: relative;
	padding:10px;
	%s
}
.header-section {
	text-align: center;
	margin-bottom: 10px;
}
.header-section img {
	height: 60px;
	margin: 0 10px;
}
.info {
	font-size: 14px;
	margin-top: 8px;
}
</style>
</head>
<body>
	<div class="page-wrapper">
		<div class="content-with-watermark">
			%s <!-- Institution Header -->
			<hr style="border:0.5px solid black">
			 <b>%s</b>
			
			<div>%s</div>
			

			<h3>Academic Performance</h3>
			<table border="1" width="100%%" cellspacing="0" cellpadding="6">
				<tr><th>Subject</th><th>Marks</th></tr>
				<tr><td>Math</td><td>90</td></tr>
				<tr><td>Science</td><td>85</td></tr>
				<tr><td>English</td><td>88</td></tr>
			</table>

			<h3>Remarks</h3>
			<p>Excellent performance overall.</p>

			<h3>Result</h3>
			<p>Promoted to next class.</p>
		</div>
	</div>
</body>
</html>
`, borderStyle, cfg.MarginTop, cfg.MarginRight, cfg.MarginBottom, cfg.MarginLeft,
		bgCSS, wmCSS,
		instHTML,
		titleHTML,studentDetailsHTML)
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
		%s
		<div style="color:%s; font-size:%dpx;">%s</div>
		<div style="color:%s; font-size:%dpx;">%s</div>
		<div style="color:%s; font-size:%dpx;">%s, %s - %s</div>
		<div style="color:%s; font-size:%dpx;">%s</div>
				<div style="color:%s; font-size:%dpx;">%s</div>

	</div>
	`,
		cfg.EnableHeaderText(),
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
func (i InstitutionDetailsConfig) EnableHeaderText() string {
	if i.EnableHeader {
		return fmt.Sprintf(`<h2 style="margin:0; color:%s; font-size:%dpx;">%s</h2>`,
			i.HeaderColor, i.HeaderFontSize, i.PrintHeader)
	}
	return ""
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