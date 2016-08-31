package system

import (
//"fmt"
//"reflect"
//"runtime"
)

var internalMethod = map[string]string{
	"beforeInit":  "beforeInit",
	"afterFinish": "afterFinish",
	"init":        "Init",
	"end":         "End",
}

type ControllerInterface interface {
	beforeInit(*Context, *Template)
	afterFinish()
	Init()
	End()
}

type Controller struct {
	*Context
	T *Template
}

func (c *Controller) beforeInit(ctx *Context, tmpl *Template) {
	c.Context = ctx
	c.T = tmpl
	c.T.Context = ctx

	//开启session
	if c.Context.App.SessionManager != nil {
		c.Context.Session = c.Context.App.SessionManager.Start(c.Context.RW, c.Context.R)
	}
}

func (c *Controller) Init() {

}

func (c *Controller) End() {

}

func (c *Controller) afterFinish() {
	if c.Context.Session != nil {
		c.Context.Session.Save()
	}
}
