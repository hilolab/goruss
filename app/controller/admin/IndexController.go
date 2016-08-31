package admin

type IndexController struct {
	ControllerBase
}

func (this *IndexController) beforeInit() {

	this.Context.RW.Write([]byte("beforeInit admin <br />"))
}
func (this *IndexController) Init() {
	//this.beforeInit()
	this.Context.RW.Write([]byte("init"))
}

func (this *IndexController) Index() {
	this.Context.RW.Write([]byte("Hello, admin/IndexController/Index\r\n"))
	name := this.Session.Get("name")
	this.RW.Write([]byte(name.(string)))
	this.Session.Set("name", "pong")
	this.Session.Save()
}

func (this *IndexController) Banner() {
	this.Context.RW.Write([]byte("Hi,I am Banner action in admin package"))
}

func (this *IndexController) End() {
	this.Context.RW.Write([]byte("end"))
}

func (this *IndexController) internal() {

}
