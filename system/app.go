package system

import (
	//"fmt"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"russ/system/session"
	"russ/system/util"
	"strings"
)

const (
	Release = iota
	Debug
)

type App struct {
	workDir           string
	DefaultModule     string
	DefaultController string
	DefaultAction     string
	Router            *Router
	FileServer        *FileServer
	Debug             int
	SessionManager    *session.SessionManager
}

type ControllerInfo struct {
	cType  reflect.Type
	method string
}

func NewApp() *App {
	return &App{
		workDir:           os.Getenv("GOPATH") + "/src",
		DefaultModule:     "home",
		DefaultController: "index",
		DefaultAction:     "index",
		Router:            NewRouter(),
		FileServer:        NewFileServer(),
		Debug:             Release,
	}
}

//默认注册所有HTTP method
func (a *App) RegisterController(ctrls ...ControllerInterface) {
	for _, c := range ctrls {
		cVal := reflect.ValueOf(c)
		cType := cVal.Type()
		cRealType := reflect.Indirect(cVal).Type()

		pkgPath := cRealType.PkgPath()

		modulePath := strings.Split(pkgPath, "controller")
		if len(modulePath) != 2 {
			panic("register " + cRealType.String() + " error")
		}

		pkgInfo := strings.Split(cRealType.String(), ".")
		file := strings.ToLower(pkgInfo[1])
		pos := strings.Index(file, "controller")

		countMethod := cType.NumMethod()

		for i := 0; i < countMethod; i++ {
			method := cType.Method(i).Name
			if _, ok := internalMethod[method]; ok || method[0] > 96 {
				continue
			}

			pattern := ""
			if strings.Index(modulePath[1], a.DefaultModule) != 1 {
				pattern += modulePath[1] + "/"
			}

			pattern += file[0:pos] + "/" + strings.ToLower(method)
			ctrlInfo := &ControllerInfo{
				cType:  cRealType,
				method: method,
			}

			a.Router.Add(pattern, ctrlInfo)
		}

	}
}

func (a *App) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	url := req.URL.Path

	if filePath, ok := a.FileServer.IsStaticDir(url); ok {
		http.ServeFile(rw, req, filePath)
		return
	}

	ctrlInfo := a.Router.Get(req.Method, url)

	if ctrlInfo == nil {
		http.NotFound(rw, req)
		return
	}

	cVal := reflect.New(ctrlInfo.cType)

	/*in := reflect.ValueOf(context)
	cVal.Elem().FieldByName("Context").Set(in)*/

	ctrl, ok := cVal.Interface().(ControllerInterface)
	if !ok {
		http.NotFound(rw, req)
		return
	}

	pkgPath := ctrlInfo.cType.PkgPath()
	ctrlName := ctrlInfo.cType.Name()

	modulePos := strings.LastIndex(pkgPath, "/")
	ctrlPos := strings.LastIndex(ctrlName, "Controller")

	context := &Context{
		RW:             rw,
		R:              req,
		App:            a,
		moduleName:     pkgPath[modulePos+1:],
		controllerName: util.Ucfirst(ctrlName[:ctrlPos]),
		methodName:     util.Ucfirst(ctrlInfo.method),
	}

	//启动session
	/*if a.SessionManager != nil {
		context.Session = a.SessionManager.Start(rw, req)
	}*/

	viewPath := a.GetWorkDir() + "/" + strings.Replace(pkgPath, "controller", "view", 1)
	tmpl := NewTemplate(viewPath)

	ctrl.beforeInit(context, tmpl)                      //系统初始化控制器调用的函数
	cVal.MethodByName(internalMethod["init"]).Call(nil) //用户初始化控制器调用的函数
	cVal.MethodByName(ctrlInfo.method).Call(nil)        //用户访问URL调用控制器的函数
	cVal.MethodByName(internalMethod["end"]).Call(nil)  //用户结束控制器调用的函数
	ctrl.afterFinish()                                  //系统结束控制器调用的函数
}

//注册单个HTTP method
func (a *App) AddRouter(hm, pattern string, ctrlInfo *ControllerInfo) {
	segment := util.PathSplit(pattern)
	a.Router.AddBase(hm, segment, ctrlInfo)
}

func (a *App) AddStaticDir(url, dir string) bool {
	return a.FileServer.SetStaticDir(url, dir)
}

//获取工作目录
func (a *App) GetWorkDir() string {
	return a.workDir
}

func (a *App) Run(addr string) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.ListenAndServe(addr, a)
}

func (a *App) SetDebug(mode int) {
	a.Debug = mode
}
