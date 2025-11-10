// package main

// import (
//     "fmt"
//     "log"
//     "strings"
//     _ "embed"

//     wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
// )

// // --- EMBED EXTERNAL HTML FILES ---

// // Assume base.html and header.html are created and embedded as in the previous solution
// //go:embed templates/base.html
// var baseHTML string

// //go:embed templates/header.html
// // var headerHTMLTemplate string 
// // This template needs a lot more placeholders now!


// var styleCSS string
// // --- DATA STRUCTURES ---

// type Student struct {
//     Name  string
//     Class string
//     AdmNo string
//     Marks map[string]string
// }

// // ReportCardConfig holds all boolean flags and styling parameters from the JSX config.
// type ReportCardConfig struct {
//     // Content Flags (Booleans)
//     PrintInstWebsite bool
//     PrintInstAddress bool
//     PrintInstEmail   bool
//     PrintInstLogo    bool
//     PrintInstName    bool
//     PrintCustName    bool
//     EnableHeader     bool
//     EnableAffiliated bool
//     PrintPhoto1Config bool
//     PrintPhoto2Config bool
//     InstNamePrintType string // Corresponds to InstNamePrintType

//     // Content Text
//     PrintHeader       string
//     PrintAffiliatedTo string

//     // Style Parameters (Sizes/Colors)
//     InstNameFontSize     int
//     CustomerNameFontSize int
//     AddressFontsize      int
//     SetLogoHeight        int
//     SetLogoWidth         int
//     CustomerNameColor    string
//     InstAddrColor        string
//     InstNameColor        string
//     EmailColor           string
//     EmailFontSize        int
//     WebsiteColor         string
//     WebsiteFontSize      int
//     HeaderColor          string
//     HeaderFontSize       int
//     AffiliatedColor      string
//     AffiliatedFontSize   int
// }

// // InstitutionData holds the dynamic data typically available at runtime (InstFormData in JSX).
// type InstitutionData struct {
//     CustName           string
//     InstName           string
//     InstAddress        string
//     InstPlace          string
//     InstPin            string
//     InstURL            string
//     InstEmail          string
//     BranchDesc         string // For instNamePrint logic
//     DefaultLogo1       string // Photo1.defaultLogo in JSX
//     DefaultLogo2       string // Photo2.defaultLogo in JSX
//     FallbackLogoSrc    string // imgSrc in JSX
// }

// // --- CONSTANTS ---
// const (
//     InstIDCustomPhone = "352187318389"
//     PrintInstNameTypeInstName = "INST_NAME"
//     PrintInstNameTypeBranchDesc = "BRANCH_DESC"
// )


// // generateHeaderHTML reconstructs the complex, conditional header based on config and data.
// func generateHeaderHTML(cfg ReportCardConfig, data InstitutionData, instID string) string {
// 	var b strings.Builder

// 	// --- Determine which logos are active ---
// 	hasLeft := cfg.PrintPhoto1Config && data.DefaultLogo1 != ""
// 	hasRight := cfg.PrintPhoto2Config && data.DefaultLogo2 != ""
// 	hasCenter := cfg.PrintInstLogo && data.FallbackLogoSrc != ""

// 	// --- Determine layout type ---
// 	var layout string
// 	switch {
// 	case hasLeft && hasRight && !hasCenter:
// 		layout = "two-side-layout" // [Left] [Text] [Right]
// 	case hasLeft && hasRight && hasCenter:
// 		layout = "three-logo-centered" // [Left] [CenterLogo] [Right], text centered below
// 	case (hasLeft && !hasRight && !hasCenter) ||
// 		(!hasLeft && hasRight && !hasCenter) ||
// 		(!hasLeft && !hasRight && hasCenter):
// 		layout = "single-logo-center-top" // Only one logo centered above text
// 	case (hasLeft && hasCenter) || (hasRight && hasCenter):
// 		layout = "multi-logos-center-top" // Left/Right + Center together on top, text below
// 	default:
// 		layout = "text-only-center" // No logos, just centered text
// 	}

// 	// --- Start wrapper ---
// 	b.WriteString(fmt.Sprintf(`<div class="%s progress-report__header-wrapper">`, layout))

// 	// --- Logo and text wrapper ---
// 	b.WriteString(`<div class="header__logo-row">`)

// 	logoStyleStr := fmt.Sprintf("height: %dpx; width: %dpx;", cfg.SetLogoHeight, cfg.SetLogoWidth)

// 	// Helper to render an <img> tag
// 	renderImg := func(src, alt string) string {
// 		return fmt.Sprintf(`<img src="%s" alt="%s" style="%s" />`, src, alt, logoStyleStr)
// 	}

// 	// -------------------------------
// 	//  Layout-specific rendering
// 	// -------------------------------

// 	switch layout {

// 	// --- CASE 1: Two Side Layout ---
// 	case "two-side-layout":
// 		// Left logo
// 		b.WriteString(fmt.Sprintf(`<div class="header__logo-slot header__logo-slot--left">%s</div>`, renderImg(data.DefaultLogo1, "Logo 1")))
// 		// Center details
// 		b.WriteString(`<div class="header__details-container">`)
// 		b.WriteString(generateInstitutionDetails(cfg, data, instID))
// 		b.WriteString(`</div>`)
// 		// Right logo
// 		b.WriteString(fmt.Sprintf(`<div class="header__logo-slot header__logo-slot--right">%s</div>`, renderImg(data.DefaultLogo2, "Logo 2")))

// 	// --- CASE 2: Three logos centered ---
// 	case "three-logo-centered":
// 		b.WriteString(`<div class="header__three-logos-center">`)
// 		b.WriteString(renderImg(data.DefaultLogo1, "Logo 1"))
// 		b.WriteString(renderImg(data.FallbackLogoSrc, "Fallback Logo"))
// 		b.WriteString(renderImg(data.DefaultLogo2, "Logo 2"))
// 		b.WriteString(`</div>`)
// 		b.WriteString(`<div class="header__details-container center-text">`)
// 		b.WriteString(generateInstitutionDetails(cfg, data, instID))
// 		b.WriteString(`</div>`)

// 	// --- CASE 3: Single logo centered top ---
// 	case "single-logo-center-top":
// 		b.WriteString(`<div class="header__single-logo-center">`)
// 		if hasLeft {
// 			b.WriteString(renderImg(data.DefaultLogo1, "Logo 1"))
// 		} else if hasRight {
// 			b.WriteString(renderImg(data.DefaultLogo2, "Logo 2"))
// 		} else if hasCenter {
// 			b.WriteString(renderImg(data.FallbackLogoSrc, "Fallback Logo"))
// 		}
// 		b.WriteString(`</div>`)
// 		b.WriteString(`<div class="header__details-container center-text">`)
// 		b.WriteString(generateInstitutionDetails(cfg, data, instID))
// 		b.WriteString(`</div>`)

// 	// --- CASE 4: Multi logos (left/right + center) ---
// 	case "multi-logos-center-top":
// 		b.WriteString(`<div class="header__multi-logos-center">`)
// 		if hasLeft {
// 			b.WriteString(renderImg(data.DefaultLogo1, "Logo 1"))
// 		}
// 		if hasCenter {
// 			b.WriteString(renderImg(data.FallbackLogoSrc, "Fallback Logo"))
// 		}
// 		if hasRight {
// 			b.WriteString(renderImg(data.DefaultLogo2, "Logo 2"))
// 		}
// 		b.WriteString(`</div>`)
// 		b.WriteString(`<div class="header__details-container center-text">`)
// 		b.WriteString(generateInstitutionDetails(cfg, data, instID))
// 		b.WriteString(`</div>`)

// 	// --- CASE 5: Text only ---
// 	default:
// 		b.WriteString(`<div class="header__details-container center-text">`)
// 		b.WriteString(generateInstitutionDetails(cfg, data, instID))
// 		b.WriteString(`</div>`)
// 	}

// 	// --- Close the logo row ---
// 	b.WriteString(`</div>`)
// 	b.WriteString(`<hr class="header-separator">`)
// 	b.WriteString(`</div>`)

// 	return b.String()
// }

// // generateStudentContent remains the same, but signature changes to include config and data
// func generateStudentContent(s Student, cfg ReportCardConfig, data InstitutionData, instID string) string {
// 	var content strings.Builder
    
//     // Pass config and data to the header generation
//     content.WriteString(generateHeaderHTML(cfg, data, instID, ))

// 	content.WriteString(fmt.Sprintf(`
//     <div style="margin-top: 20px;">
//         <h2>Report Card - %s (%s)</h2>
//         <p>Adm No: %s</p>
//         <table>
//         <tr><th>Subject</th><th>Marks</th></tr>`, s.Name, s.Class, s.AdmNo))

// 	for subject, marks := range s.Marks {
// 		content.WriteString(fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>", subject, marks))
// 	}

// 	content.WriteString(`
//         </table>
//         </div>`)
    
// 	content.WriteString(`<div style="page-break-after: always;"></div>`)

// 	return content.String()
// }


// // --- MAIN FUNCTION ---

// func main() {
//     // --- 1. SETUP DYNAMIC CONFIGURATION AND DATA ---
    
//     // Mock the dynamic InstId that triggers the custom phone number
//     const currentInstID = InstIDCustomPhone // Change this to test the phone number visibility

//     // Mock the runtime data (InstFormData in JSX)
//     instData := InstitutionData{
//         CustName:        "Managed by XYZ Group",
//         InstName:        "Example Public School",
//         InstAddress:     "123 Main Street",
//         InstPlace:       "Bengaluru",
//         InstPin:         "560001",
//         InstURL:         "www.exampleschool.edu",
//         InstEmail:       "info@exampleschool.edu",
//         BranchDesc:      "East Campus - Whitefield",
//         DefaultLogo1:    "https://picsum.photos/60/60?random=1", // Photo1.defaultLogo
//         DefaultLogo2:    "https://picsum.photos/60/60?random=2", // Photo2.defaultLogo
//         FallbackLogoSrc: "https://picsum.photos/60/60?random=3", // imgSrc
//     }
    
//     // Mock the ReportCardConfig (ReportCardConfig() in JSX)
//     config := ReportCardConfig{
//         // Content Flags
//         PrintInstWebsite: true,
//         PrintInstAddress: true,
//         PrintInstEmail:   true,
//        // Controls FallbackLogoSrc
//         PrintInstName:    true,
//         PrintCustName:    true,
//         EnableHeader:     true,
//         EnableAffiliated: true,
// 		  PrintInstLogo:    true,
//         PrintPhoto1Config: true, // Show Logo 1
//         PrintPhoto2Config: true, // Show Logo 2
//         InstNamePrintType: PrintInstNameTypeInstName, // Change to PrintInstNameTypeBranchDesc to see branch name

//         // Content Text
//         PrintHeader:       "A Unit of the Excellence Education Trust",
//         PrintAffiliatedTo: "Affiliated to CBSE, New Delhi",

//         // Style Parameters
//         InstNameFontSize:     24,
//         CustomerNameFontSize: 16,
//         AddressFontsize:      12,
//         SetLogoHeight:        60,
//         SetLogoWidth:         60,
//         CustomerNameColor:    "#333333",
//         InstAddrColor:        "#666666",
//         InstNameColor:        "#000080",
//         EmailColor:           "#666666",
//         EmailFontSize:        12,
//         WebsiteColor:         "#666666",
//         WebsiteFontSize:      12,
//         HeaderColor:          "#000000",
//         HeaderFontSize:       18,
//         AffiliatedColor:      "#555555",
//         AffiliatedFontSize:   14,
//     }

// 	// Example student list
// 	students := []Student{
// 		{"Sahana M", "1 A", "AS-51", map[string]string{"Math": "95", "Science": "88"}},
// 		{"Rahul K", "2 B", "AS-52", map[string]string{"Math": "89", "Science": "90"}},
// 		{"Lakshmi P", "3 A", "AS-53", map[string]string{"Math": "75", "Science": "85"}},
// 	}

// 	// --- 2. COMBINE ALL HTML ---

// 	var studentContent strings.Builder

// 	// Loop through all students and append their report content
// 	for _, s := range students {
// 		// Wrap content in a div for individual page margins/spacing
// 		studentContent.WriteString(`<div class="page-content">`)
// 		// Pass configuration and data to the content function
// 		studentContent.WriteString(generateStudentContent(s, config, instData, currentInstID))
// 		studentContent.WriteString(`</div>`)
// 	}

//     // Insert the generated student content into the embedded base HTML
//     // (Assuming base.html has the "" placeholder)
// 	finalHTML := strings.Replace(baseHTML, "", studentContent.String(), 1)


// 	// --- 3. GENERATE PDF ---

// 	pdfg, err := wkhtmltopdf.NewPDFGenerator()
// 	if err != nil {
// 		log.Fatal("Error initializing PDF generator:", err)
// 	}

// 	// Add the ONE combined HTML string as a single PageReader
// 	page := wkhtmltopdf.NewPageReader(strings.NewReader(finalHTML))
// 	pdfg.AddPage(page)

// 	// PDF settings
// 	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
// 	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
// 	pdfg.MarginBottom.Set(10)
// 	pdfg.MarginTop.Set(10)
// 	pdfg.MarginLeft.Set(10)
// 	pdfg.MarginRight.Set(10)

// 	// Create and save the PDF
// 	if err := pdfg.Create(); err != nil {
// 		log.Fatal("Failed to create PDF:", err)
// 	}

// 	if err := pdfg.WriteFile("Students_Report.pdf"); err != nil {
// 		log.Fatal("Failed to write PDF:", err)
// 	}

// 	fmt.Println("✅ PDF for all students generated successfully!")
// }

// // generateInstitutionDetails builds the institution text block (name, address, contact info)
// func generateInstitutionDetails(cfg ReportCardConfig, data InstitutionData, instID string) string {
// 	var d strings.Builder

// 	// Header line
// 	if cfg.EnableHeader {
// 		d.WriteString(fmt.Sprintf(`<span style="font-size:%dpx;color:%s;">%s</span><br/>`,
// 			cfg.HeaderFontSize, cfg.HeaderColor, cfg.PrintHeader))
// 	}

// 	// Customer name
// 	if cfg.PrintCustName {
// 		d.WriteString(fmt.Sprintf(`<span style="font-size:%dpx;color:%s;">%s</span><br/>`,
// 			cfg.CustomerNameFontSize, cfg.CustomerNameColor, data.CustName))
// 	}

// 	// Institution name
// 	if cfg.PrintInstName {
// 		instNameValue := data.InstName
// 		if cfg.InstNamePrintType == PrintInstNameTypeBranchDesc {
// 			instNameValue = data.BranchDesc
// 		}
// 		d.WriteString(fmt.Sprintf(`<b style="font-size:%dpx;color:%s;">%s</b><br/>`,
// 			cfg.InstNameFontSize, cfg.InstNameColor, instNameValue))
// 	}

// 	// Affiliated To
// 	if cfg.EnableAffiliated {
// 		d.WriteString(fmt.Sprintf(`<span style="font-size:%dpx;color:%s;">%s</span><br/>`,
// 			cfg.AffiliatedFontSize, cfg.AffiliatedColor, cfg.PrintAffiliatedTo))
// 	}

// 	// Address
// 	if cfg.PrintInstAddress {
// 		address := strings.Join([]string{data.InstAddress, data.InstPlace, data.InstPin}, ", ")
// 		address = strings.Replace(address, ", "+data.InstPin, " - "+data.InstPin, 1)
// 		d.WriteString(fmt.Sprintf(`<span style="font-size:%dpx;color:%s;">%s</span><br/>`,
// 			cfg.AddressFontsize, cfg.InstAddrColor, address))
// 	}

// 	// Website and email
// 	if cfg.PrintInstWebsite || cfg.PrintInstEmail {
// 		d.WriteString(`<div>`)
// 		if cfg.PrintInstWebsite {
// 			d.WriteString(fmt.Sprintf(`<span style="font-size:%dpx;color:%s;margin-right:10px;">Website: %s</span>`,
// 				cfg.WebsiteFontSize, cfg.WebsiteColor, data.InstURL))
// 		}
// 		if cfg.PrintInstEmail {
// 			d.WriteString(fmt.Sprintf(`<span style="font-size:%dpx;color:%s;">Email: %s</span>`,
// 				cfg.EmailFontSize, cfg.EmailColor, data.InstEmail))
// 		}
// 		d.WriteString(`</div>`)
// 	}

// 	// Hardcoded phone (specific instID)
// 	if instID == InstIDCustomPhone {
// 		d.WriteString(fmt.Sprintf(`<span style="font-size:%dpx;">Phone: 080-28486734, 9611317506</span>`,
// 			cfg.AddressFontsize))
// 	}

// 	return d.String()
// }
package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func encodeImageToBase64(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error reading image %s: %v", path, err)
		return ""
	}
	return fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(data))
}

// readCSS reads CSS file contents and returns a string safe to place in <style>
func readCSS(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Warning: couldn't read CSS %s: %v", path, err)
		return ""
	}
	return string(b)
}

func generateInstitutionTextHTML(title, name, address, contact string) string {
	return fmt.Sprintf(`
		<div class="institution-text" style="text-align:center; line-height:1.4;">
			<h2 style="margin:0; font-size:22px;">%s</h2>
			<p style="margin:4px 0;">%s</p>
			<p style="margin:4px 0;">%s</p>
			<p style="margin:4px 0;">%s</p>
		</div>
	`, title, name, address, contact)
}
// generateHeaderHTML dynamically builds the header layout
// based on how many images (1, 2, or 3) are enabled.
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
	institutionText := generateInstitutionTextHTML(
		"Eduate Report Card",
		"Eduate Private Limited",
		"NHCE Campus, Bengaluru - 560045",
		"Email: info@eduate.in | Website: www.eduate.in",
	)

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


func generateFullHTML(headerHTML string, cssContent string) string {
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
	cssBlock := ""
	if cssContent != "" {
		cssBlock = fmt.Sprintf("<style>\n%s\n</style>", cssContent)
	}

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
	</html>`, cssBlock, headerHTML, body)
}

func main() {
	photo1 := encodeImageToBase64("./assets/Arcadis_Logo.png")
	photo2 := encodeImageToBase64("./assets/Arcadis_Logo.png")
	instLogo := encodeImageToBase64("./assets/Arcadis_Logo.png")

	// choose config
	PrintPhoto1Config := false
	PrintPhoto2Config := true
	PrintInstLogo := true

	// read CSS file and inline it
	cssPath := "./templates/style.css" // adjust if your CSS is elsewhere
	cssContent := readCSS(cssPath)
	if cssContent == "" {
		fmt.Println("Warning: CSS content empty or not found - PDF will use only inline styles.")
	}

	// generate header and full HTML (use your existing header generator)
	headerHTML := generateHeaderHTML(PrintPhoto1Config, PrintPhoto2Config, PrintInstLogo, photo1, photo2, instLogo)
	fullHTML := generateFullHTML(headerHTML, cssContent)

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	page := wkhtmltopdf.NewPageReader(strings.NewReader(fullHTML))
	// Do NOT rely on page.CustomArgs or AllowLocalFileAccess (they may not exist in your wrapper)
	pdfg.AddPage(page)

	if err := pdfg.Create(); err != nil {
		log.Fatal(err)
	}
	if err := pdfg.WriteFile("./final_layout.pdf"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ PDF generated successfully: final_layout.pdf")
}
