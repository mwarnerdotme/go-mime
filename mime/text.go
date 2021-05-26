package mime

const (
	// .css: cascading style sheet
	TEXT_CSS = "text/css"
	// .csv: comma-separated values
	TEXT_CSV = "text/csv"
	// .html: hypertext markup language
	TEXT_HTML = "text/html"
	// .ics: iCalendar
	TEXT_ICS = "text/calendar"
	// .js: javascript
	TEXT_JS = "text/javascript"
	// .txt: plain text (ASCII, ISO 8859-n)
	TEXT_PLAIN = "text/plain"

	// http only
	TEXT_YAML = "text/yaml"

	// aliases/alternatives

	TEXT_HTM = TEXT_HTML
	TEXT_MJS = TEXT_JS
)

// slice containing all available text MIME types
func GetTextMIMETypes() []string {
	return []string{
		TEXT_CSS,
		TEXT_CSV,
		TEXT_HTML,
		TEXT_ICS,
		TEXT_JS,
		TEXT_PLAIN,
		TEXT_YAML,
	}
}
