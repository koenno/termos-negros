package portal

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/koenno/termos-negros/domain"
)

type MenuParser interface {
	Parse(r io.Reader) ([]domain.DayMenu, error)
}

type BodyParser struct {
	menuParser MenuParser
}

func NewBodyParser(menuParser MenuParser) BodyParser {
	return BodyParser{
		menuParser: menuParser,
	}
}

func (p BodyParser) Parse(bb []byte, headers http.Header) ([]domain.DayMenu, error) {
	if !p.validContentType(headers.Get("content-type")) {
		return nil, fmt.Errorf("unsupported content type: %s", headers.Get("content-type"))
	}

	bytesReader := bytes.NewReader(bb)
	return p.menuParser.Parse(bytesReader)
}

func (p BodyParser) validContentType(contentType string) bool {
	if contentType == "" {
		return false
	}
	elems := strings.Split(contentType, ";")
	return elems[0] == "text/html"
}
