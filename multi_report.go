


package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
"math"
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
type PageDecorationConfig struct {
	ShowBackground bool
	ShowWatermark  bool
	ShowBorder     bool

	BackgroundImage string
	WatermarkImage  string

	BorderColor      string
	BorderWidth      float64 // in pixels
	WatermarkOpacity float64

	// Margins (in mm)
	MarginTop    float64
	MarginBottom float64
	MarginLeft   float64
	MarginRight  float64
}
type StudentDetailsConfig struct {
	IndexOrder       []int               // Order of fields from backend (like [1,7,4,3,...])
	VisibleFields    map[string]bool     // To control which fields are shown
	DisplayMode      string              // "single-column" or "two-column"
	ShowProfilePic   bool
	PicOnRight       bool
	ProfilePicBase64 string

	// Student data fields
	StudentName     string
	DateOfBirth     string
	FatherName      string
	MotherName      string
	Address         string
	Email           string
	Mobile          string
	AttendanceStats string
	StudentClass    string
	AcademicYear    string
	StudentRollNo   string
    ProfilePicWidth  int
    ProfilePicHeight int
}
func encodeImageToBase64(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error reading image %s: %v", path, err)
		return ""
	}
	return fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(data))
}
func pageSetup(pdfg *wkhtmltopdf.PDFGenerator, cfg PageDecorationConfig) {
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
pdfg.MarginTop.Set(uint(math.Round(cfg.MarginTop)))
pdfg.MarginBottom.Set(uint(math.Round(cfg.MarginBottom)))
pdfg.MarginLeft.Set(uint(math.Round(cfg.MarginLeft)))
pdfg.MarginRight.Set(uint(math.Round(cfg.MarginRight)))

}
func decoratePage(bgBase64, wmBase64 string) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<style>
  @page { margin: 0; }
  html, body {
    margin: 0;
    padding: 0;
    height: 100%%;
    width: 100%%;
  }

  body {
    background: url('%s') no-repeat center center;
    background-size: cover;
    position: relative;
  }

  /* Watermark centered */
  .watermark {
    position: absolute;
    top: 50%%;
    left: 50%%;
    transform: translate(-50%%, -50%%);
    width: 200px;
    opacity: 0.15;
    filter: blur(1px);
    z-index: 1;
  }

  /* Border over full page */
  .border {
    position: absolute;
    top: 10px;
    left: 10px;
    width: calc(100%% - 20px);
    height: calc(100%% - 20px);
    border: 1.5px solid #1a237e;
    box-sizing: border-box;
    z-index: 2;
  }
</style>
</head>
<body>
  <img class="watermark" src="%s" />
  <div class="border"></div>
</body>
</html>
`, bgBase64, wmBase64)
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
func generateStudentDetailsHTML(cfg StudentDetailsConfig) string {
	// Step 1: Map numeric indices to field keys
	indexFieldMap := map[int]string{
		1:  "Name",
		2:  "RollNo",
		3:  "FatherName",
		4:  "MotherName",
		5:  "ClassSection",
		6:  "AcademicYear",
		7:  "DateOfBirth",
		8:  "Attendance",
		9:  "Address",
		10: "Email",
		11: "Mobile",
	}

	// Step 2: Map field keys to actual values
	fieldValueMap := map[string]string{
		"Name":         cfg.StudentName,
		"RollNo":       cfg.StudentRollNo,
		"FatherName":   cfg.FatherName,
		"MotherName":   cfg.MotherName,
		"ClassSection": cfg.StudentClass,
		"AcademicYear": cfg.AcademicYear,
		"DateOfBirth":  cfg.DateOfBirth,
		"Attendance":   cfg.AttendanceStats,
		"Address":      cfg.Address,
		"Email":        cfg.Email,
		"Mobile":       cfg.Mobile,
	}

	// Step 3: Render helper
	renderField := func(label, value string) string {
		if strings.TrimSpace(value) == "" {
			return ""
		}
		return fmt.Sprintf(`<div style="margin: 0;"><b>%s:</b> %s</div>`, label, value)
	}

	// Step 4: Build list of enabled fields based on IndexOrder and VisibleFields
	var enabledFields []string
	for _, idx := range cfg.IndexOrder {
		if fieldKey, ok := indexFieldMap[idx]; ok && cfg.VisibleFields[fieldKey] {
			if val := fieldValueMap[fieldKey]; val != "" {
				enabledFields = append(enabledFields, renderField(fieldKey, val))
			}
		}
	}

	isSingleColumn := strings.ToLower(cfg.DisplayMode) == "single-column"

	// Step 5: Profile Picture
picHTML := ""
    if cfg.ShowProfilePic && cfg.ProfilePicBase64 != "" {
        // Assume cfg has fields like ProfilePicWidth and ProfilePicHeight (e.g., as int/float)
        picHTML = fmt.Sprintf(`
            <div style="text-align:center; margin-bottom:10px;">
                <img src="%s" alt="Student Photo"
                    style="width:%dpx; height:%dpx; border-radius:50%%; object-fit:cover; border:2px solid #ccc;"/>
            </div>
        `, cfg.ProfilePicBase64, cfg.ProfilePicWidth, cfg.ProfilePicHeight)
    }

	// Step 6: Build details layout
	var detailsContainer string
	if isSingleColumn {
		// --- Single-column layout ---
		detailsHTML := strings.Join(enabledFields, "")
		detailsContainer = fmt.Sprintf(`
			<div style="display:block; text-align:center; margin:0 auto;">
				%s
			</div>
		`, detailsHTML)
	} else {
		// --- Two-column layout ---
		if len(enabledFields) == 0 {
			detailsContainer = ""
		} else {
			mid := int(math.Ceil(float64(len(enabledFields)) / 2))
			leftCol := strings.Join(enabledFields[:mid], "")
			rightCol := ""
			if mid < len(enabledFields) {
				rightCol = strings.Join(enabledFields[mid:], "")
			}

			detailsContainer = fmt.Sprintf(`
				<div style="display:inline-block; width:50%%; vertical-align:top; text-align:left;">
					%s
				</div>
				<div style="display:inline-block; width:50%%; vertical-align:top; text-align:left; margin-left:0%%;">
					%s
				</div>
			`, leftCol, rightCol)
		}
	}

	// Step 7: Combine picture + details
	var contentHTML string
if isSingleColumn {
        contentHTML = fmt.Sprintf("%s%s", picHTML, detailsContainer)
    } else {
        // Retrieve the width from config, assuming it's an integer
        photoWidth := cfg.ProfilePicWidth 
        photoWidthPx := fmt.Sprintf("%dpx", photoWidth) // e.g., "250px"
        
        // Calculate the dynamic width string for the details container
        detailsWidthCalc := fmt.Sprintf("calc(85%% - %dpx)", photoWidth)
        
        if cfg.ShowProfilePic && cfg.ProfilePicBase64 != "" {
            if cfg.PicOnRight {
                contentHTML = fmt.Sprintf(`
                    <div style="display:inline-block; width:%s; vertical-align:middle;">%s</div>
                    <div style="display:inline-block; width:%s; vertical-align:middle; text-align:right; margin-left:10px;">%s</div>
                `, detailsWidthCalc, detailsContainer, photoWidthPx, picHTML)
            } else {
                contentHTML = fmt.Sprintf(`
                    <div style="display:inline-block; width:%s; vertical-align:middle; text-align:left; margin-right:10px;">%s</div>
                    <div style="display:inline-block; width:%s; vertical-align:middle;">%s</div>
                `, photoWidthPx, picHTML, detailsWidthCalc, detailsContainer)
            }
        } else {
            contentHTML = detailsContainer
        }
    }
	// Step 8: Outer wrapper
	wrapperStyle := "margin:0;"
	if isSingleColumn {
		wrapperStyle += "text-align:center; display:block;"
	} else {
		wrapperStyle += "white-space:nowrap; text-align:left;"
	}

	return fmt.Sprintf(`
		<div style="%s">
			%s
		</div>
	`, wrapperStyle, contentHTML)
}




func generateReportContentHTML(cfg PageDecorationConfig, headerHTML string, studentDetailsHTML string) string {
	bgStyle := ""
	if cfg.ShowBackground && cfg.BackgroundImage != "" {
		bgStyle = fmt.Sprintf(
			"background: url('%s') no-repeat center center; background-size: cover; padding:%.1fmm %.1fmm %.1fmm %.1fmm;",
			cfg.BackgroundImage, cfg.MarginTop, cfg.MarginRight, cfg.MarginBottom, cfg.MarginLeft)
	}

	watermarkHTML := ""
	if cfg.ShowWatermark && cfg.WatermarkImage != "" {
		watermarkHTML = fmt.Sprintf(
			`<img class="watermark" src="%s" style="opacity:%f;" />`,
			cfg.WatermarkImage, cfg.WatermarkOpacity)
	}

	borderHTML := ""
	if cfg.ShowBorder {
	
		borderHTML = fmt.Sprintf(`
			<div class="border" 
				style="
					position: absolute;
					top: 5mm;
					left: 5mm;
					right: 5mm;
					bottom: 5mm;
					border: %.2fmm solid %s;
					box-sizing: border-box;
					z-index: 2;">
			</div>`,
			cfg.BorderWidth, cfg.BorderColor)
	}

	return fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
	<meta charset="utf-8">
	<title>Eduate Report</title>
	<style>
		@page { margin: 0; }
		html, body {
			margin: 0;
			padding: 0;
			height: 100%%;
			width: 100%%;
			font-family: Arial, sans-serif;
			box-sizing: border-box;
		}

		body {
			%s
			position: relative;
		}

		.watermark {
			position: absolute;
			top: 50%%;
			left: 50%%;
			transform: translate(-50%%, -50%%);
			width: 250px;
			filter: blur(1px);
			z-index: 1;
		}

		.main-content-wrapper {
			/* Add 10mm padding to keep all content away from the border */
			padding: 10mm;
			position: relative; /* Ensure it respects Z-index */
			z-index: 3;
		}

		.header-zone {
			text-align: center;
			margin-bottom: 10px;
		}

		.content-zone {
			padding: 0;
			box-sizing: border-box;
		}
	</style>
	</head>
	<body>
		%s
		%s
		<div class="main-content-wrapper">
			<div class="header-zone">%s</div>

			<div class="content-zone">
				<hr style="margin:20px 0;"/>
				%s<div  class="std-details" style="margin:0 20px;text-align:left;"></div>

				<div style="margin:30px 20px;text-align:left;">
					<h3>Marks Summary</h3>
					<table style="width:100%%;border-collapse:collapse;">
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
			</div>
		</div>
	</body>
	</html>
	`, bgStyle, watermarkHTML, borderHTML, headerHTML,studentDetailsHTML)
}

func main() {
	photo1 := encodeImageToBase64("./assets/Arcadis_Logo.png")
	photo2 := encodeImageToBase64("./assets/Arcadis_Logo.png")
	instLogo := encodeImageToBase64("./assets/Arcadis_Logo.png")


	PrintPhoto1Config := false
	PrintPhoto2Config := true
	PrintInstLogo := true

	headerHTML := generateHeaderHTML(PrintPhoto1Config, PrintPhoto2Config, PrintInstLogo, photo1, photo2, instLogo)

    pageCfg := PageDecorationConfig{
	ShowBackground:   true,
	ShowWatermark:    false,
	ShowBorder:       false,
	BackgroundImage:  encodeImageToBase64("./assets/background.png"),
	WatermarkImage:   encodeImageToBase64("./assets/watermark.jpeg"),
	BorderColor:      "#1a237e",
	BorderWidth:      1.5,
	WatermarkOpacity: 0.15,
	MarginTop:       0,
	MarginBottom:    0,
	MarginLeft:      0,
	MarginRight:    0,
}
studentCfg := StudentDetailsConfig{
	ShowProfilePic:   true,
	PicOnRight:       false,
	ProfilePicBase64: encodeImageToBase64("./assets/colorwatermark.png"),

	StudentName:   "John Doe",
	StudentClass:  "10th Standard",
	StudentRollNo: "25",
	AcademicYear:  "2024 - 2025",
	DateOfBirth:   "15-08-2009",
	FatherName:    "Mr. Richard Doe",
	MotherName:    "Mrs. Emily Doe",
	Address:       "123 Main Street, Springfield",
	Email:         "john.doe@example.com",
	Mobile:        "9876543210",
	AttendanceStats: "95%",
	ProfilePicWidth: 80,
    ProfilePicHeight:80,
	DisplayMode: "two-column", // "single-column" or "two-column"

	// Backend-provided field order (using numeric indices)
	IndexOrder: []int{1, 7, 4, 3, 9, 10, 11, 8, 5, 6, 2},

	// Visibility map (which fields to show)
	VisibleFields: map[string]bool{
		"Name":         true,
		"DateOfBirth":  true,
		"MotherName":   true,
		"FatherName":   true,
		"Address":      true,
		"Email":        false,
		"Mobile":       false,
		"Attendance":   true,
		"ClassSection": false,
		"AcademicYear": true,
		"RollNo":       true,
	},
}

        studentDetailsHTML := generateStudentDetailsHTML(studentCfg)

	fullHTML := generateReportContentHTML(pageCfg, headerHTML,studentDetailsHTML)

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

pageSetup(pdfg, pageCfg) 
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
