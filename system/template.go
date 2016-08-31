package system

import (
	//"fmt"
	"html/template"
	/*"regexp"
	  "runtime"
	  "russ/system/util"*/
	"io/ioutil"
	"strings"
)

type Template struct {
	*Context
	viewPath   string
	renderFile string
	Layout     string
	TplSuffix  string
	funcMap    template.FuncMap
}

func NewTemplate(vp string) *Template {
	t := &Template{
		viewPath:  vp,
		Layout:    "layout",
		TplSuffix: ".html",
		funcMap:   make(template.FuncMap),
	}
	t.funcMap["include"] = t.Include
	return t
}

func (t *Template) Render(data interface{}) {
	if t.renderFile == "" {
		t.renderFile = t.GetControllerName() + "/" + t.GetMethodName()
	}
	renderFilename := t.GetFilename(t.renderFile)
	layoutFilename := t.GetFilename("/layout/" + t.Layout)

	//go 的大坑 (/ □ \)，
	tpl := template.New(t.Layout + t.TplSuffix)

	tpl.Funcs(t.funcMap)
	tpl, err := tpl.ParseFiles(layoutFilename, renderFilename)
	if err != nil {
		t.RW.Write([]byte(err.Error()))
		return
	}

	tpl.Execute(t.RW, data)
}

// 如：controllerName/method
func (t *Template) SetView(filename string) {
	t.renderFile = filename
}

func (t *Template) SetFuncMap(key string, funcName interface{}) bool {
	_, ok := t.funcMap[key]
	if ok {
		return false
	}

	t.funcMap[key] = funcName
	return true
}

func (t *Template) GetFilename(filename string) string {
	return t.viewPath + "/" + strings.Trim(filename, "/") + t.TplSuffix
}

func (t *Template) Include(partial string) string {
	filename := t.GetFilename(partial)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return "Error: can not find tempalte: " + strings.TrimLeft(partial, "/") + t.TplSuffix
	}
	return string(b)
}
