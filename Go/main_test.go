package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var rootURL = "https://localhost/LEAF_Request_Portal/"
var csrfToken string

var tr = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}

var cookieJar, _ = cookiejar.New(nil)
var client = &http.Client{
	Transport: tr,
	Timeout:   time.Second * 1,
	Jar:       cookieJar,
}

// TestMain performs initial setup and logs into the dev environment.
// In dev, the current username is set via REMOTE_USER docker environment
// Tests must not trigger Fatal. See teardownTestDB()
func TestMain(m *testing.M) {

	req, _ := http.NewRequest("GET", rootURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.46")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	body := string(bodyBytes)

	startIdx := strings.Index(body, "var CSRFToken = '") + 17
	endIdx := strings.Index(body[startIdx:], "';")
	csrfToken = body[startIdx : startIdx+endIdx]

	setupTestDB()

	code := m.Run()

	teardownTestDB()

	os.Exit(code)
}
