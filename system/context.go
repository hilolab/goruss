package system

import (
	"net/http"
	"russ/system/session"
)

type Context struct {
	RW             http.ResponseWriter
	R              *http.Request
	App            *App
	moduleName     string
	controllerName string
	methodName     string
	Session        *session.Session
}

func (c *Context) GetModuleName() string {
	return c.moduleName
}

func (c *Context) GetControllerName() string {
	return c.controllerName
}

func (c *Context) GetMethodName() string {
	return c.methodName
}
