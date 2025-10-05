package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync/atomic"
	"time"
)

var ResultsDirty atomic.Bool
var ResultsString string

func SetupResultCreator() {
	ticker := time.NewTicker(time.Minute)
	go func() {
		for range ticker.C {
			if ResultsDirty.Load() {
				ResultsString = buildResults()
				ResultsDirty.Store(false)
			}
		}
	}()
}

const HTML_TEMPLATE_START = `
<!doctype html>
<html lang="en">
<head>
<meta charset="UTF-8" />
<link rel="icon" type="image/svg+xml" href="/vite.svg" />
<meta name="viewport" content="width=device-width, initial-scale=1.0" />
<title>Geländespiel Ergebnisse</title>
</head>
<body>
<p>Ergebnisse &mdash; Stand 
`
const HTML_TEMPLATE_END = `
</body>
</html>
`

func buildResults() string {
	var sb strings.Builder
	sb.WriteString(HTML_TEMPLATE_START)
	sb.WriteString(time.Now().Local().Format(time.TimeOnly))
	sb.WriteString("</p>")
	sb.WriteString("<table><tr><th>Gruppe</th><th>Iteration 1</th><th>Iteration 2</th><th>Iteration 3</th></tr>")
	fmt.Fprintf(&sb, "")

	sb.WriteString(HTML_TEMPLATE_END)
	return sb.String()
}

func HandleResults(w http.ResponseWriter, r *http.Request) {

}
