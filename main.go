package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	URL         = "https://www.sec.gov/cgi-bin/cik_lookup"
	UserAgent   = "Mozilla/5.0 (X11; Linux x86_64; rv:120.0) Gecko/20100101 Firefox/120.0"
	Accept      = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"
	AcceptLang  = "en-US,en;q=0.5"
	AcceptEnc   = "gzip, deflate, br"
	ContentType = "application/x-www-form-urlencoded"
	Referer     = "https://www.sec.gov/edgar/searchedgar/cik"
	Origin      = "https://www.sec.gov"

	Timeout     = 10 * time.Second
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage:go run main.go <company_name>")
	}

	company := getCompanyName(os.Args[1:])
	respBody := makeRequest(company)
	parseHTML(respBody)
}

func getCompanyName(args []string) string {
	return url.QueryEscape(strings.Join(args, " "))
}

func makeRequest(company string) string {
	reqBody := bytes.NewBufferString(fmt.Sprintf("company=%s", company))
	req, err := http.NewRequest("POST", URL, reqBody)
	if err != nil {
		log.Fatalf("problem making POST: %v", err)
	}

	setRequestHeaders(req)

	client := &http.Client{Timeout: Timeout}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("problem with the response: %v", err)
	}
	defer resp.Body.Close()

	return readResponseBody(resp)
}

func setRequestHeaders(req *http.Request) {
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", Accept)
	req.Header.Set("Accept-Language", AcceptLang)
	req.Header.Set("Accept-Encoding", AcceptEnc)
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("Referer", Referer)
	req.Header.Set("Origin", Origin)
}

func readResponseBody(resp *http.Response) string {
	reader, err := decompressResponse(resp)
	if err != nil {
		log.Fatalf("problem decompressing response: %v", err)
	}
	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalf("problem reading response body: %v", err)
	}
	return string(body)
}

func decompressResponse(response *http.Response) (io.ReadCloser, error) {
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		return gzip.NewReader(response.Body)
	default:
		return response.Body, nil
	}
}

func parseHTML(html string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("pre").Each(func(index int, preHtml *goquery.Selection) {
		if index > 0 {
			fmt.Println(strings.TrimSpace(preHtml.Text()))
		}
	})
}
