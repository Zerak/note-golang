package main

import (
	"flag"
	"github.com/golang/glog"
	"os"
)

func main() {
	//初始化命令行参数
	flag.Parse()
	//退出时调用，确保日志写入文件中
	defer glog.Flush()

	glog.Info(os.TempDir())

	glog.Info("hello, glog")
	glog.Warning("warning glog")
	glog.Error("error glog")

	glog.Infof("info %d", 1)
	glog.Warningf("warning %d", 2)
	glog.Errorf("error %d", 3)

	glog.V(1).Infoln("Processed1", "nItems1", "elements1")

	glog.V(2).Infoln("Processed2", "nItems2", "elements2")

	glog.V(3).Infoln("Processed3", "nItems3", "elements3")

	glog.V(4).Infoln("Processed4", "nItems4", "elements4")

	glog.V(5).Infoln("Processed5", "nItems5", "elements5")
}
