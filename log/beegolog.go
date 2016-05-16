package main

import "github.com/astaxie/beego/logs"

func main() {
	log := logs.NewLogger(1024)
	defer log.Flush()

	//log.EnableFuncCallDepth(true)
	//log.SetLogFuncCallDepth(0)

	//log.SetLogger("console", "")
	log.SetLogger("file", `{"filename":"blog.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)

	log.Emergency("emergency")
	log.Emergency("emergency1")
	log.Emergency("emergency2")
	log.Emergency("emergency3")

	log.Alert("alert")
	log.Alert("alert1")
	log.Alert("alert2")
	log.Alert("alert3")

	log.Critical("critical")
	log.Critical("critical1")
	log.Critical("critical2")
	log.Critical("critical3")

	log.Error("error")
	log.Error("error1")
	log.Error("error2")

	log.DelLogger("console")

	log.Warn("warning")
	log.Warn("warning1")
	log.Warn("warning2")

	log.Notice("notice")
	log.Notice("notice1")
	log.Notice("notice2")
	log.Notice("notice3")

	log.Trace("trace")
	log.Trace("trace1")
	log.Trace("trace2")
	log.Trace("trace3")

	log.Info("info")
	log.Info("info1")
	log.Info("info2")
	log.Info("info3")

	log.Info("info1")
	log.Info("info2")
	log.Info("info3")
	log.Info("info4")

}
