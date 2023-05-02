package utils

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"log"
)

func GbkToUtf8(resp io.ReadCloser) string {
	gbkBodyText := transform.NewReader(resp, simplifiedchinese.GBK.NewDecoder())
	bodyText, err := io.ReadAll(gbkBodyText)
	if err != nil {
		log.Panicln(bodyText)
	}
	return string(bodyText)
}
