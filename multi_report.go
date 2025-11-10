


package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)
type InstitutionTextConfig struct {
	ShowTrustName bool
	ShowAddress   bool
	ShowInstEmail bool
	ShowWebsite   bool
	ShowInstName  bool

	ShowHeader bool
	ShowCaption bool

	TrustName string
	Header    string
	InstName  string
	Address   string
	Email     string
	Website   string
	Caption   string

	TitleFontSize   int
	NameFontSize    int
	AddressFontSize int
	ContactFontSize int

	TitleColor   string
	NameColor    string
	AddressColor string
	ContactColor string
}

func encodeImageToBase64(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error reading image %s: %v", path, err)
		return ""
	}
	return fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(data))
}


func decoratePDFWithWkhtmltopdf(outputPath string) error {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	// ---- HTML content with background, watermark, border ----
	html := `
<!DOCTYPE html>
<html>
<head>
<style>
  @page {
    margin: 0;
  }

  body {
    margin: 0;
    width: 100%;
    height: 100%;
    position: relative;
    font-family: sans-serif;
  }

  /* Background image */
  .bg {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: url('assets/background.png') no-repeat center center;
    background-size: cover;
    z-index: 0;
  }

  /* Watermark */
  .watermark {
    position: absolute;
    top: 50%;
    left: 50%;
    width: 150px; /* resize as needed */
    transform: translate(-50%, -50%);
    opacity: 0.6;
    filter: blur(1px);
    z-index: 1;
  }

  /* Optional border */
  .border {
    position: absolute;
    top: 8px;
    left: 8px;
    width: calc(100% - 16px);
    height: calc(100% - 16px);
    border: 0.8px solid rgb(50,50,150);
    box-sizing: border-box;
    z-index: 2;
  }

  /* Content goes above background and watermark */
  .content {
    position: relative;
    z-index: 3;
    padding: 20px;
  }
</style>
</head>
<body>
  <div class="bg"></div>
  <img class="watermark" src="assets/colorwatermark.png"/>
  <div class="border"></div>

  <div class="content">
    <h1>Student Report</h1>
    <p>This is sample content on top of the background and watermark.</p>
  </div>
</body>
</html>
`

	// Create a new page
	page := wkhtmltopdf.NewPageReader(strings.NewReader(html))
	pdfg.AddPage(page)

	// Set PDF options
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.MarginTop.Set(0)
	pdfg.MarginBottom.Set(0)

	// Generate PDF
	err = pdfg.Create()
	if err != nil {
		return err
	}

	// Write PDF to file
	return pdfg.WriteFile(outputPath)
}
func generateInstitutionTextHTML(cfg InstitutionTextConfig) string {
	html := `<div class="institution-text" style="text-align:center; line-height:1.4;">`
	if cfg.ShowHeader {
		html += fmt.Sprintf(
			`<h2 style="margin:0; font-size:%dpx; color:%s;">%s</h2>`,
			cfg.TitleFontSize, cfg.TitleColor, cfg.Header)
	}
    if cfg.ShowTrustName {
		html += fmt.Sprintf(
			`<h2 style="margin:0; font-size:%dpx; color:%s;">%s</h2>`,
			cfg.TitleFontSize, cfg.TitleColor, cfg.TrustName)
	}



	if cfg.ShowInstName {
		html += fmt.Sprintf(
			`<p style="margin:4px 0; font-size:%dpx; color:%s;">%s</p>`,
			cfg.NameFontSize, cfg.NameColor, cfg.InstName)
	}

	if cfg.ShowAddress {
		html += fmt.Sprintf(
			`<p style="margin:4px 0; font-size:%dpx; color:%s;">%s</p>`,
			cfg.AddressFontSize, cfg.AddressColor, cfg.Address)
	}

	if cfg.ShowInstEmail {
		html += fmt.Sprintf(
			`<p style="margin:4px 0; font-size:%dpx; color:%s;">%s</p>`,
			cfg.Email, cfg.ContactColor, cfg.Email)
	}
if cfg.ShowWebsite {
		html += fmt.Sprintf(
			`<p style="margin:4px 0; font-size:%dpx; color:%s;">%s</p>`,
			cfg.Website, cfg.ContactColor, cfg.Website)
	}
	html += `</div>`

	return html
}
// generateHeaderHTML dynamically builds the header layout
func generateHeaderHTML(printPhoto1Config, printPhoto2Config, printInstLogo bool, photo1, photo2, instLogo string) string {
	count := 0
	if printPhoto1Config {
		count++
	}
	if printPhoto2Config {
		count++
	}
	if printInstLogo {
		count++
	}

	// ✅ Generate the common institution text block

instTextCfg := InstitutionTextConfig{
	ShowTrustName:  false,
	ShowAddress:    true,
	ShowInstEmail:  true,
	ShowWebsite:    true,
	ShowInstName:   true,
	ShowHeader:     false,
	ShowCaption:    true,

	TrustName: "Om Sri",
	Header:    "Eduate Report Card",
	InstName:  "Eduate Private Limited",
	Address:   "NHCE Campus, Bengaluru - 560045",
	Email:     "info@eduate.in",
	Website:   "www.eduate.in",
	Caption:   "caption",

	TitleFontSize:   22,
	NameFontSize:    14,
	AddressFontSize: 14,
	ContactFontSize: 13,

	TitleColor:   "#000000",
	NameColor:    "#333333",
	AddressColor: "#555555",
	ContactColor: "#777777",
}


institutionText := generateInstitutionTextHTML(instTextCfg)

	var layout string

	switch count {

	case 1:
		src := ""
		if printPhoto1Config {
			src = photo1
		} else if printInstLogo {
			src = instLogo
		} else if printPhoto2Config {
			src = photo2
		}

		layout = fmt.Sprintf(`
			<div class="single-image">
				<img src="%s" alt="single-logo" class="logo-image" style="height:80px; margin-bottom:8px;"/>
				%s
			</div>
		`, src, institutionText)
// two-images
	case 2:
		left := ""
		right := ""
		if printPhoto1Config {
			left = photo1
		}
		if printInstLogo {
			if left == "" {
				left = instLogo
			} else {
				right = instLogo
			}
		}
		if printPhoto2Config {
			if left == "" {
				left = photo2
			} else if right == "" {
				right = photo2
			}
		}

layout = fmt.Sprintf(`
	<div style="width:100%%; text-align:center; white-space:nowrap;">
		<div style="display:inline-block; width:20%%; text-align:left; vertical-align:middle;">
			<img src="%s" alt="left-logo" style="height:80px; width:auto; object-fit:contain;"/>
		</div>

		<div style="display:inline-block; width:58%%; text-align:center; vertical-align:middle; line-height:1.4;">
			%s
		</div>

		<div style="display:inline-block; width:20%%; text-align:right; vertical-align:middle;">
			<img src="%s" alt="right-logo" style="height:80px; width:auto; object-fit:contain;"/>
		</div>
	</div>
`, left, institutionText, right)

	case 3:
layout = fmt.Sprintf(`
	<div style="width:100%%; ">
		<div style="display:inline-block; width:30%%; text-align:left;">
			<img src="%s" alt="logo1" style="height:80px; width:80px; object-fit:contain;"/>
		</div>
		<div style="display:inline-block; width:30%%; text-align:center;">
			<img src="%s" alt="logo2" style="height:80px; width:80px; object-fit:contain;"/>
		</div>
		<div style="display:inline-block; width:30%%; text-align:right;">
			<img src="%s" alt="logo3" style="height:80px; width:80px; object-fit:contain;"/>
		</div>
	</div>
	<div style="text-align:center; margin-top:10px; line-height:1.4;">
		%s
	</div>
`, photo1, instLogo, photo2, institutionText)

	default:
		layout = fmt.Sprintf(`
			<div class="institution-only" style="text-align:center;">%s</div>
		`, institutionText)
	}

	return fmt.Sprintf(`<div class="header-container" style="margin-bottom:10px;">%s</div>`, layout)
}


func generateFullHTML(headerHTML string) string {
	body := `
		<hr style="margin:30px 0;"/>

		<div style="margin:0 50px;text-align:left;">
			<h3>Student Details</h3>
			<p><b>Name:</b> John Doe</p>
			<p><b>Class:</b> 10th Standard</p>
			<p><b>Roll No:</b> 25</p>
			<p><b>Academic Year:</b> 2024 - 2025</p>
		</div>

		<div style="margin:30px 50px;text-align:left;">
			<h3>Marks Summary</h3>
			<table style="width:100%;border-collapse:collapse;">
				<tr style="background:#f2f2f2;">
					<th style="border:1px solid #ccc;padding:8px;">Subject</th>
					<th style="border:1px solid #ccc;padding:8px;">Marks</th>
					<th style="border:1px solid #ccc;padding:8px;">Grade</th>
				</tr>
				<tr><td style="border:1px solid #ccc;padding:8px;">Maths</td><td style="border:1px solid #ccc;padding:8px;">95</td><td style="border:1px solid #ccc;padding:8px;">A+</td></tr>
				<tr><td style="border:1px solid #ccc;padding:8px;">Science</td><td style="border:1px solid #ccc;padding:8px;">88</td><td style="border:1px solid #ccc;padding:8px;">A</td></tr>
				<tr><td style="border:1px solid #ccc;padding:8px;">English</td><td style="border:1px solid #ccc;padding:8px;">92</td><td style="border:1px solid #ccc;padding:8px;">A+</td></tr>
			</table>
		</div>
	`

	// Insert CSS content directly inside a <style> tag


	return fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
	  <meta charset="utf-8">
	  <title>Eduate Report</title>
	  %s
	</head>
	<body style="font-family:Arial, sans-serif; margin:30px;">
	%s
	%s
	</body>
	</html>`, headerHTML, body)
}

func main() {
	photo1 := encodeImageToBase64("./assets/Arcadis_Logo.png")
	photo2 := encodeImageToBase64("./assets/Arcadis_Logo.png")
	instLogo := encodeImageToBase64("./assets/Arcadis_Logo.png")

	// choose config
	PrintPhoto1Config := false
	PrintPhoto2Config := true
	PrintInstLogo := true



	// generate header and full HTML (use your existing header generator)
	headerHTML := generateHeaderHTML(PrintPhoto1Config, PrintPhoto2Config, PrintInstLogo, photo1, photo2, instLogo)
	fullHTML := generateFullHTML(headerHTML)

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	page := wkhtmltopdf.NewPageReader(strings.NewReader(fullHTML))
	pdfg.AddPage(page)

	if err := pdfg.Create(); err != nil {
		log.Fatal(err)
	}
	if err := pdfg.WriteFile("./final_layout.pdf"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ PDF generated successfully: final_layout.pdf")
}
