package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"time"
)

var ResultsDirty atomic.Bool
var ResultsString string

func SetupResultCreator() {
	ResultsString = buildResults()
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

func HandleResults(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, ResultsString)
}

const HTML_TEMPLATE_START = `
<!doctype html>
<html lang="en">
<head>
<meta charset="UTF-8" />
<link rel="icon" type="image/svg+xml" href="/vite.svg" />
<meta name="viewport" content="width=device-width, initial-scale=1.0" />
<title>Geländespiel Ergebnisse</title>
<style>
table, th, td {
	border: 1px solid black;
}
html {
	text-align: center;
}
body {
	display: flex;
	flex-direction: column;
	align-items: center;
}
</style>
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
	sb.WriteString("<table>")
	sb.WriteString("<thead><tr><th>Gruppe</th>")
	for i := range NUM_ITERATIONS {
		fmt.Fprintf(&sb, "<th>Iteration %d</th>", i+1)
	}
	sb.WriteString("<th>Gesamt</th></tr></thead><tbody>")

	for index := range DBGroups {
		station_indicies := DBGroups[index].Stations
		fmt.Fprintf(&sb, "<tr><td>Gruppe %d</td>", index+1)
		total := 0
		for it_idx := range NUM_ITERATIONS {
			entry_idx := it_idx + int(station_indicies[it_idx])*NUM_ITERATIONS
			entry := DBScores[entry_idx]
			if entry == SCORE_STUDENT {
				total++
			}
			fmt.Fprintf(&sb, "<td>%s</td>", getEntryString(entry))
		}
		fmt.Fprintf(&sb, "<td>%d</td>", total)
		sb.WriteString("</tr>")
	}

	fmt.Fprintf(&sb, "</tbody></table>")

	sb.WriteString(HTML_TEMPLATE_END)
	fmt.Println("Finished writing results!")
	return sb.String()
}

func getEntryString(entry uint8) string {
	switch entry {
	case SCORE_STUDENT:
		return "1"
	case SCORE_TUTOR:
		return "0"
	default:
		return "N/A"
	}
}
