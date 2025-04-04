package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"text/template"

	"github.com/alewtschuk/pfmt"
)

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Job struct {
	Title  string   `json:"title"`
	Start  string   `json:"start"`
	End    string   `json:"end"`
	Points []string `json:"points"`
}

type Reference struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Email    string `json:"email"`
}

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func parseProjects() (projects []Project) {
	file, err := os.Open("projects.json")
	CheckErr(err)
	data, err := io.ReadAll(file)
	CheckErr(err)
	CheckErr(json.Unmarshal(data, &projects))
	file.Close()
	return
}

func parseJobs() (jobs []Job) {
	file, err := os.Open("jobs.json")
	CheckErr(err)
	data, err := io.ReadAll(file)
	CheckErr(err)
	CheckErr(json.Unmarshal(data, &jobs))
	file.Close()
	return
}

func parseReferences() (references []Reference) {
	file, err := os.Open("references.json")
	CheckErr(err)
	data, err := io.ReadAll(file)
	CheckErr(err)
	CheckErr(json.Unmarshal(data, &references))
	file.Close()
	return
}

func fOut(fbuf bytes.Buffer) {
	file, err := os.Create("out.html")
	CheckErr(err)
	defer file.Close()

	_, err = fbuf.WriteTo(file)
	CheckErr(err)
	pfmt.Printcln("Resume written to out.html!", 2)
}

func main() {
	var buf bytes.Buffer
	t, _ := template.ParseFiles("template.tmpl")
	t.Execute(&buf, struct {
		Projects   []Project
		Jobs       []Job
		References []Reference
	}{
		Projects:   parseProjects(),
		Jobs:       parseJobs(),
		References: parseReferences(),
	})
	fOut(buf)
}
