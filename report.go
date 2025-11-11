package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"

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
	TitleColor       string
	TitleFontSize    int
	SubtitleColor    string
	SubtitleFontSize int
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
		
		TitleColor:       "#1a237e",
		TitleFontSize:    22,
		SubtitleColor:    "#424242",
		SubtitleFontSize: 16,
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

	html := generateHTML(cfg, instCfg)

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
func generateHTML(cfg PageDecorationConfig, instCfg InstitutionDetailsConfig) string {
	instHTML := generateInstitutionDetailsHTML(instCfg)

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
			<hr>
			<h1 style="text-align:center; color:%s; font-size:%dpx;">Progress Report</h1>
			<h3 style="text-align:center; color:%s; font-size:%dpx;">Academic Year 2024-2025</h3>
			<hr>
			<p><strong>Student Name:</strong> John Doe</p>
			<p><strong>Roll No:</strong> 45</p>

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
		cfg.TitleColor, cfg.TitleFontSize, cfg.SubtitleColor, cfg.SubtitleFontSize)
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


func (i InstitutionDetailsConfig) EnableHeaderText() string {
	if i.EnableHeader {
		return fmt.Sprintf(`<h2 style="margin:0; color:%s; font-size:%dpx;">%s</h2>`,
			i.HeaderColor, i.HeaderFontSize, i.PrintHeader)
	}
	return ""
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