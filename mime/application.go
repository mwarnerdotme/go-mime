package mime

const (
	// .7z: 7-zip archive
	APPLICATION_7Z = "application/x-7z-compressed"
	// .abw: abiword document
	APPLICATION_ABW = "application/x-abiword"
	// .arc: archive document
	APPLICATION_ARC = "application/x-freearc"
	// .azw: amazon kindle ebook
	APPLICATION_AZW = "application/vnd.amazon.ebook"
	// .bin: binary data
	APPLICATION_BIN = "application/octet-stream"
	// .bz: bzip archive
	APPLICATION_BZ = "application/x-bzip"
	// .bz2: bzip2 archive
	APPLICATION_BZ2 = "application/x-bzip2"
	// .cda: cd audio
	APPLICATION_CDA = "application/x-cdf"
	// .csh: c-shell script
	APPLICATION_CSH = "application/x-csh"
	// .doc: microsoft word document
	APPLICATION_DOC = "application/msword"
	// .docx: microsoft word document (openxml)
	APPLICATION_DOCX = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	// .eot: embedded opentype fonts
	APPLICATION_EOT = "application/vnd.ms-fontobject"
	// .epub: electronic publication
	APPLICATION_EPUB = "application/epub+zip"
	// .gz: gzip compressed archive
	APPLICATION_GZ = "application/gzip"
	// .jar: java archive
	APPLICATION_JAR = "application/java-archive"
	// .json: javascript object notation
	APPLICATION_JSON = "application/json"
	// .jsonld: javascript object notation ld
	APPLICATION_JSONLD = "application/ld+json"
	// .mpkg: Apple installer package
	APPLICATION_MPKG = "application/vnd.apple.installer+xml"
	// .odp: opendocument presentation
	APPLICATION_ODP = "application/vnd.oasis.opendocument.presentation"
	// .ods: opendocument spreadsheet
	APPLICATION_ODS = "application/vnd.oasis.opendocument.spreadsheet"
	// .odt: opendocument text
	APPLICATION_ODT = "application/vnd.oasis.opendocument.text"
	// .ogx: ogg
	APPLICATION_OGX = "application/ogg"
	// .pdf: adobe portable document
	APPLICATION_PDF = "application/pdf"
	// .php: hypertext preprocessor - personal home page
	APPLICATION_PHP = "application/x-httpd-php"
	// .ppt: microsoft powerpoint
	APPLICATION_PPT = "application/vnd.ms-powerpoint"
	// .ppt: microsoft powerpoint openxml
	APPLICATION_PPTX = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	// .rar: rar archive
	APPLICATION_RAR = "application/vnd.rar"
	// .rtf: rich text
	APPLICATION_RTF = "application/rtf"
	// .sh: bourne shell script
	APPLICATION_SH = "application/x-sh"
	// .swf: adobe flash
	APPLICATION_SWF = "application/x-shockwave-flash"
	// .tar: tape archive
	APPLICATION_TAR = "application/x-tar"
	// .vsd: Microsoft Visio
	APPLICATION_VSD = "application/vnd.visio"
	// .xhtml: xhtml
	APPLICATION_XHTML = "application/xhtml+xml"
	// .xls: Microsoft Excel
	APPLICATION_XLS = "applicatioon/vnd.ms-excel"
	// .xlsx: Microsoft Excel (openxml)
	APPLICATION_XLSX = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	// .xml: extensible markup
	APPLICATION_XML = "application/xml"
	// .xul: XUL
	APPLICATION_XUL = "application/vnd.mozilla.xul+xml"
	// .zip: zip archive
	APPLICATION_ZIP = "application/zip"

	// aliases/alternatives

	APPLICATION_EXCEL        = APPLICATION_XLS
	APPLICATION_EXCEL_X      = APPLICATION_XLSX
	APPLICATION_OGG          = APPLICATION_OGX
	APPLICATION_POWERPOINT   = APPLICATION_PPT
	APPLICATION_POWERPOINT_X = APPLICATION_PPTX
	APPLICATION_VISIO        = APPLICATION_VSD
	APPLICATION_WORD         = APPLICATION_DOC
	APPLICATION_WORD_X       = APPLICATION_DOCX
)

// slice containing all available application MIME types
func GetApplicationMIMETypes() []string {
	return []string{
		APPLICATION_7Z,
		APPLICATION_ABW,
		APPLICATION_ARC,
		APPLICATION_AZW,
		APPLICATION_BIN,
		APPLICATION_BZ,
		APPLICATION_BZ2,
		APPLICATION_CDA,
		APPLICATION_CSH,
		APPLICATION_DOC,
		APPLICATION_DOCX,
		APPLICATION_EOT,
		APPLICATION_EPUB,
		APPLICATION_EXCEL,
		APPLICATION_EXCEL_X,
		APPLICATION_GZ,
		APPLICATION_JAR,
		APPLICATION_JSON,
		APPLICATION_JSONLD,
		APPLICATION_MPKG,
		APPLICATION_ODP,
		APPLICATION_ODS,
		APPLICATION_ODT,
		APPLICATION_OGX,
		APPLICATION_PDF,
		APPLICATION_PHP,
		APPLICATION_PPT,
		APPLICATION_PPTX,
		APPLICATION_RAR,
		APPLICATION_RTF,
		APPLICATION_SH,
		APPLICATION_SWF,
		APPLICATION_TAR,
		APPLICATION_VSD,
		APPLICATION_WORD,
		APPLICATION_WORD_X,
		APPLICATION_XHTML,
		APPLICATION_XLS,
		APPLICATION_XLSX,
		APPLICATION_XML,
		APPLICATION_XUL,
		APPLICATION_ZIP,
	}
}
