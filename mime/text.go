package mime

/*
Text MIME types
*/

const (
	// .css: cascading style sheet
	TEXT_CSS MIMEType = "text/css"
	// .csv: comma-separated values
	TEXT_CSV MIMEType = "text/csv"
	// .html: hypertext markup language
	TEXT_HTML MIMEType = "text/html"
	// .ics: iCalendar
	TEXT_ICS MIMEType = "text/calendar"
	// .js: javascript
	TEXT_JS MIMEType = "text/javascript"
	// .txt: plain text (ASCII, ISO 8859-n)
	TEXT_PLAIN MIMEType = "text/plain"
	// .yaml: YAML ain't markup language
	TEXT_YAML MIMEType = "text/yaml"
)

// Aliases/alternatives

const (
	TEXT_HTM = TEXT_HTML
	TEXT_MJS = TEXT_JS
)

/*
Text content types (HTTP only MIME types) - no file extension
*/

const ()

// Slice containing all available text MIME types.
var TextMIMETypes = []MIMEType{
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
	EXTENSION_MJS  = ".mjs"
	EXTENSION_TXT  = ".txt"
	EXTENSION_YAML = ".yaml"
	EXTENSION_YML  = ".yml"
)

// Slice containing all available text file extensions.
var TextFileExtensions = []FileExtension{
	EXTENSION_CSS,
	EXTENSION_CSV,
	EXTENSION_HTML,
	EXTENSION_HTM,
	EXTENSION_ICS,
	EXTENSION_JS,
	EXTENSION_MJS,
	EXTENSION_TXT,
	EXTENSION_YAML,
	EXTENSION_YML,
}

/*
Map of file extensions to MIME types
*/

var TextMIMETypeMap = MIMETypeMap{
	EXTENSION_CSS:  TEXT_CSS,
	EXTENSION_CSV:  TEXT_CSV,
	EXTENSION_HTML: TEXT_HTML,
	EXTENSION_HTM:  TEXT_HTM,
	EXTENSION_ICS:  TEXT_ICS,
	EXTENSION_JS:   TEXT_JS,
	EXTENSION_MJS:  TEXT_JS,
	EXTENSION_TXT:  TEXT_PLAIN,
	EXTENSION_YAML: TEXT_YAML,
	EXTENSION_YML:  TEXT_YAML,
}
