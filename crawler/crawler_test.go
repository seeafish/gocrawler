package main

import "testing"

var sourceURL = "https://gist.githubusercontent.com/seeafish/9e05075f913d3b2bebdaa91a9d0b303b/raw/46817bcb723b79629ec97ab154a0f76e8e41aae2/simple.html"

func TestGetSource(t *testing.T) {
	actualResult := getSource(sourceURL)

	expectedResult := `<html>
  <head><title>Test page</title></head>
  <body>
    <p>Some text</p>
    <a href="/someurl">A link</a>
  </body>
</html>`

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestFindLinks(t *testing.T) {
	actualResult := findLinks(getSource(sourceURL))

	linksArr := make([]string, 1)
	linksArr[0] = "/someurl"
	expectedResult := linksArr

	if len(actualResult) != len(expectedResult) {
		t.Fatalf("Expected length of %d but got length of %d", len(expectedResult), len(actualResult))
	}

	if actualResult[0] != expectedResult[0] {
		t.Fatalf("Expected value %s but got value %s", expectedResult, actualResult)
	}
}
