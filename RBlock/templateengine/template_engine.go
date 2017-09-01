package templateengine

import (
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strings"
	"bytes"
	"errors"
)

type TemplateEngine struct{
	templateDir string
	//key - file name without ext.
	TemplateMap map[string]*template.Template
}

func NewTemplateEngine (templateDir string) (*TemplateEngine, error){
	//var err error = nil


	return &TemplateEngine{templateDir, templateMap}, nil
}

//execute template
func (self TemplateEngine) Execute (tmpName string, object interface{}) (*bytes.Buffer, error){




	return buf, tmp.Execute(buf, object)
}