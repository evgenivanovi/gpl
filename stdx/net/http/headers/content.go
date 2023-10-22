package headers

/* __________________________________________________ */

const (
	ContentTypeKey     HeaderKey = "Content-Type"
	ContentLengthKey   HeaderKey = "Content-Length"
	ContentEncodingKey HeaderKey = "Content-Encoding"

	TypeAny             ContentType = "*/*"
	TypeApplicationJSON ContentType = "application/json"
	TypeApplicationXML  ContentType = "application/xml"

	EncodingAny     ContentEncoding = "*"
	EncodingZSTD    ContentEncoding = "zstd"
	EncodingLZ4     ContentEncoding = "lz4"
	EncodingGZIP    ContentEncoding = "gzip"
	EncodingDeflate ContentEncoding = "deflate"

	TypeTextPlain    ContentType = "text/plain"
	TypeTextHTML     ContentType = "text/html"
	TypeTextCSV      ContentType = "text/csv"
	TypeTextCmd      ContentType = "text/cmd"
	TypeTextCSS      ContentType = "text/css"
	TypeTextXML      ContentType = "text/xml"
	TypeTextMarkdown ContentType = "text/markdown"
)

/* __________________________________________________ */

type ContentType HeaderValue

func (ct ContentType) String() string {
	return string(ct)
}

type ContentEncoding string

func (ce ContentEncoding) String() string {
	return string(ce)
}

/* __________________________________________________ */
