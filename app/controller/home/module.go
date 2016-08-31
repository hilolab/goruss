package home

import (
	//"fmt"
	"russ/system"
)

func Register(app *system.App) {
	app.RegisterController(&IndexController{})
	//app.RegisterController(&PostController{})
}

type ControllerBase struct {
	system.Controller
}
