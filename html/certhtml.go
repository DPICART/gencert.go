package html

import (
	"fmt"
	"gencert/cert"
	"html/template"
	"os"
	"path"
)

type HtmlSaver struct {
	OutputDir string
}

func New(outputDir string) (*HtmlSaver, error) {
	var p *HtmlSaver
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return p, err
	}
	p = &HtmlSaver{
		OutputDir: outputDir,
	}
	return p, nil
}

func (h *HtmlSaver) Save(c cert.Cert) error {
	t, err := template.New("certificate").Parse(tpl)
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("%v.html", c.LabelTitle)
	path := path.Join(h.OutputDir, filename)
	f, err := os.Create(path)
	if nil != err {
		return err
	}
	defer f.Close()
	err = t.Execute(f, c)
	if nil != err {
		return err
	}
	fmt.Printf("Saved certificate to '%v'\n", path)
	return nil
}

var tpl = `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.LabelTitle}}</title>
	</head>
	<body>
		<h1>{{.LabelCompletion}}</h1>
		<h2><em>{{.LabelPresented}}</em></h2>
		<h1>{{.Name}}</h1>
		<h2>{{.LabelParticipation}}</h2>
		<p>
			<em>{{.LabelDate}}</em>
		</p>
	</body>
	</html>
`
