package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// Configuration for page decoration
type PageDecorationConfig struct {
	ShowBackground   bool
	BackgroundImage  string // path or base64
	ShowWatermark    bool
	WatermarkImage   string // path or base64
	ShowBorder       bool
	BorderColor      string
	BorderWidth      float64
	BorderType       string // solid, dashed, etc.
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

// --- Helper: convert local image file to base64 string ---
func toBase64(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("⚠️ Warning: could not read %s: %v", filePath, err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(data)
}

// --- Generates full HTML content ---
func generateHTML(cfg PageDecorationConfig) string {
	var bgBase64, wmBase64 string

	if cfg.ShowBackground && cfg.BackgroundImage != "" {
		bgBase64 = cfg.BackgroundImage
	}
	if cfg.ShowWatermark && cfg.WatermarkImage != "" {
		wmBase64 = cfg.WatermarkImage
	}

	borderStyle := "none"
	if cfg.ShowBorder {
		borderStyle = fmt.Sprintf("%.1fpx %s %s", cfg.BorderWidth, cfg.BorderType, cfg.BorderColor)
	}

	return fmt.Sprintf(`<!DOCTYPE html>
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
}

.page {
	height: 100%%;
	width: 100%%;
	box-sizing: border-box;
}

/* Outer wrapper (border + background) */
.page-wrapper {
	width: 100%%;
	height: 100%%;
	box-sizing: border-box;
	border: %s;
	padding: %s %s %s %s;
	background-color: white;
	%s
}

/* Watermark layer */
.content-with-watermark {
	width: 100%%;
	height: 100%%;
	box-sizing: border-box;
	%s
	padding: %s %s %s %s;
}

/* Main content area */
.page-content {
	width: 100%%;
	height: 100%%;
	box-sizing: border-box;
}

.page-break {
	page-break-after: always;
}
</style>
</head>
<body>
	<div class="page">
		<div class="page-wrapper">
			<div class="content-with-watermark">
				<div class="page-content">
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
		</div>
	</div>
	<div class="page-break"></div>
</body>
</html>`,
		// CSS interpolation
		borderStyle,
		cfg.MarginTop, cfg.MarginRight, cfg.MarginBottom, cfg.MarginLeft,

		// background CSS
		func() string {
			if cfg.ShowBackground && bgBase64 != "" {
				return fmt.Sprintf("background-image: url('data:image/png;base64,%s'); background-repeat: no-repeat; background-size: cover; background-position: center;", bgBase64)
			}
			return ""
		}(),

		// watermark CSS
		func() string {
			if cfg.ShowWatermark && wmBase64 != "" {
				return fmt.Sprintf("background-image: url('data:image/png;base64,%s'); background-repeat: no-repeat; background-size: 30%%; background-position: center center; background-blend-mode: normal;", wmBase64)
			}
			return ""
		}(),

		cfg.TopPadding, cfg.RightPadding, cfg.BottomPadding, cfg.LeftPadding,

		cfg.TitleColor, cfg.TitleFontSize,
		cfg.SubtitleColor, cfg.SubtitleFontSize,
	)
}

// --- Generate PDF using Chromedp ---
func generatePDF(htmlContent string, output string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Give Chrome a few seconds to render
	ctx, cancel = context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate("data:text/html,"+url.PathEscape(htmlContent)),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithPaperWidth(8.27).  // A4
				WithPaperHeight(11.7). // A4
				Do(ctx)
			return err
		}),
	)
	if err != nil {
		return err
	}

	return os.WriteFile(output, buf, 0644)
}

// --- Main ---
func main() {
	cfg := PageDecorationConfig{
		ShowBackground:   true,
		BackgroundImage:  toBase64("./assets/background.png"), // local image file
		ShowWatermark:    true,
		WatermarkImage:   toBase64("./assets/watermark.jpeg"),
		ShowBorder:       false,
		BorderColor:      "#333333",
		BorderWidth:      3,
		BorderType:       "solid",
		MarginTop:        "20px",
		MarginRight:      "30px",
		MarginBottom:     "20px",
		MarginLeft:       "30px",
		TopPadding:       "15px",
		RightPadding:     "20px",
		BottomPadding:    "15px",
		LeftPadding:      "20px",
		TitleColor:       "#1a237e",
		TitleFontSize:    22,
		SubtitleColor:    "#424242",
		SubtitleFontSize: 16,
	}

	html := generateHTML(cfg)

	if err := generatePDF(html, "report.pdf"); err != nil {
		log.Fatal("❌ PDF generation failed:", err)
	}

	fmt.Println("✅ PDF generated successfully: report.pdf")
}
