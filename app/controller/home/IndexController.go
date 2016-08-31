package home

import (
//"fmt"
)

type IndexController struct {
	ControllerBase
}

func (this *IndexController) Index() {
	this.T.Render(nil)
	this.Session.Set("name", "hilo")
	this.Session.Set("Name", "dylan")
	this.Session.Save()

	//this.Context.RW.Write([]byte("Hello, IndexController/Index"))
}

func (this *IndexController) Banner() {
	this.Context.RW.Write([]byte("Hi,I am Banner action"))
}

func (this *IndexController) internal() {

}

func Include(partial string) string {
	return "include"
}
