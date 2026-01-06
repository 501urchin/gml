```go
package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/501urchin/gml"
	"github.com/501urchin/gml/gmlx"
)

const stylesheet = `/* Global layout */
body {
	font-family: Arial, sans-serif;
	margin: 20px;
	background-color: #f4f4f4;
	color: #222;
}

/* Job card */
.job-card {
	background: white;
	padding: 20px;
	margin-bottom: 20px;
	border-radius: 5px;
	box-shadow: 0 2px 4px rgba(0,0,0,0.08);
}

.job-card h2 {
	margin-top: 0;
	font-size: 1.25rem;
}

.company {
	font-weight: bold;
	color: #333;
}

.meta {
	color: #666;
	font-size: 0.95rem;
}

.remote {
	color: #28a745;
	font-weight: bold;
	margin-left: 8px;
}

/* Homepage list */
.home-container {
	display: grid;
	grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
	gap: 16px;
	align-items: start;
}

.home-item {
	background: #fff;
	padding: 12px 14px;
	border-radius: 6px;
	box-shadow: 0 1px 3px rgba(0,0,0,0.06);
}

.job-link {
	color: #0366d6;
	text-decoration: none;
	font-weight: 600;
}

.job-link:hover {
	text-decoration: underline;
}

.logo {
	font-size: 1.5rem;
	font-weight: 700;
	margin-bottom: 12px;
}

`

type job struct {
	Id       int
	Title    string
	Company  string
	Location string
	Remote   bool
}

func jobCard(info job) gml.GmlElement {
	return gml.Article().
		Attributes(gml.Class("job-card")).
		Children(
			gml.H2().Children(
				gml.Content(info.Title),
			),

			gml.P().Attributes(gml.Class("company")).Children(
				gml.Content(info.Company),
			),

			gml.P().Attributes(gml.Class("meta")).Children(
				gml.Content(info.Location),

				gmlx.If(info.Remote, func() gml.GmlElement {
					return gml.Span().Attributes(gml.Class("remote")).Children(
						gml.Content(" • Remote"),
					)
				}),
			),

			gml.A().Attributes(
				gml.Href("/jobs/"+strconv.Itoa(info.Id)),
			).Children(
				gml.Content("View job →"),
			),
		)
}

func InvalidJob(id int) gml.GmlElement {
	return gml.Div().Children(
		gml.H1().Children(
			gml.Content("Invalid Job ID"),
		),
		gml.P().Children(
			gml.Content("Job ID "+strconv.Itoa(id)+" does not exist."),
		),
	)
}

func Layout(child gml.GmlElement) gml.GmlElement {
	return gmlx.Group(
		gml.DocTypeHtml(),
		gml.Html().Children(
			gml.Head().Children(
				gml.Script().Attributes(
					gml.Src("/main.js"),
					gml.Attribute("defer"),
				),

				gml.Style().Children(
					gml.Content(stylesheet),
				),
			),

			gml.Body().Children(
				gmlx.If(child != nil, func() gml.GmlElement {
					return child
				}).Else(func() gml.GmlElement {
					return gml.Content("hello world")
				}),
			),
		),
	)
}

func HomePage(jobIds []int) gml.GmlElement {
	return gml.Div().Attributes(gml.Class("home-container")).Children(
		gml.Div().Attributes(gml.Class("logo")).Children(
			gml.Content("Open Roles"),
		),
		gmlx.Map(jobIds, func(_, id int) gml.GmlElement {
			return gml.Div().Attributes(gml.Class("home-item")).Children(
				gml.A().Attributes(
					gml.Href("/"+strconv.Itoa(id)),
					gml.Class("job-link"),
				).Children(
					gml.Content("Job ID " + strconv.Itoa(id)),
				),
			)
		}),
	)
}

func main() {
	jobMap := map[int]job{
		1: {Id: 1, Title: "Software Engineer", Company: "Tech Corp", Location: "San Francisco, CA", Remote: true},
		2: {Id: 2, Title: "Data Analyst", Company: "DataWorks", Location: "New York, NY", Remote: false},
		3: {Id: 3, Title: "Product Manager", Company: "Innovate LLC", Location: "Austin, TX", Remote: true},
		4: {Id: 4, Title: "DevOps Engineer", Company: "CloudNet", Location: "Seattle, WA", Remote: false},
		5: {Id: 5, Title: "UX Designer", Company: "DesignHub", Location: "Boston, MA", Remote: true},
	}
	http.HandleFunc("/{jobid}", func(w http.ResponseWriter, r *http.Request) {
		var jobID int
		_, err := fmt.Sscanf(r.URL.Path, "/%d", &jobID)
		if err != nil {
			Layout(InvalidJob(0)).Render(r.Context(), w)
			return
		}

		job, exists := jobMap[jobID]
		if !exists {
			Layout(InvalidJob(jobID)).Render(r.Context(), w)
			return
		}

		Layout(jobCard(job)).Render(r.Context(), w)
	})

	jobIDs := make([]int, 0, len(jobMap))
	for id := range jobMap {
		jobIDs = append(jobIDs, id)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Layout(HomePage(jobIDs)).Render(r.Context(), w)
	})

	fmt.Println("Server running at http://localhost:8049")
	err := http.ListenAndServe(":8049", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

```