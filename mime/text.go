package mime

/*
Text MIME types
*/

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
	// .yaml: YAML ain't markup language
	TEXT_YAML = "text/yaml"
)

// aliases/alternatives

const (
	TEXT_HTM = TEXT_HTML
	TEXT_MJS = TEXT_JS
)

/*
Text content types (HTTP only MIME types) - no file extension
*/

const ()

// slice containing all available text MIME types
var TextMIMETypes = []MimeType{
	TEXT_CSS,
	TEXT_CSV,
	TEXT_HTML,
	TEXT_ICS,
	TEXT_JS,
	TEXT_PLAIN,
	TEXT_YAML,
}

/*
Text file extensions
*/

const (
	EXTENSION_CSS  = ".css"
	EXTENSION_CSV  = ".csv"
	EXTENSION_HTML = ".html"
	EXTENSION_HTM  = ".htm"
	EXTENSION_ICS  = ".ics"
	EXTENSION_JS   = ".js"
	EXTENSION_TXT  = ".txt"
	EXTENSION_YAML = ".yaml"
	EXTENSION_YML  = ".yml"
)

// slice containing all available text file extensions
var TextFileExtensions = []FileExtension{
	EXTENSION_CSS,
	EXTENSION_CSV,
	EXTENSION_HTML,
	EXTENSION_HTM,
	EXTENSION_ICS,
	EXTENSION_JS,
	EXTENSION_TXT,
	EXTENSION_YAML,
	EXTENSION_YML,
}
