package main

import (
	//"fmt"
	//"os"
	//"path/filepath"
	"russ/app/controller/admin"
	"russ/app/controller/home"
	"russ/system"
	"russ/system/session"
	//"time"
)

func main() {
	/*dir, _ := filepath.Glob("g:/golang/bin/*")
	fmt.Println(dir)*/
	app := system.NewApp()
	app.SetDebug(system.Debug)
	//fmt.Println(app.GetWorkDir())
	home.Register(app)
	admin.Register(app)
	app.AddStaticDir("static", app.GetWorkDir()+"/russ")
	fileStore := session.NewFileStore()
	app.SessionManager = session.NewSessionManager(fileStore)


	// fmt.Println(sessid)
	app.Run(":9999")
	//time.Sleep(time.Second * 5)
}
