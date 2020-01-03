package main

import (
	"fate-cloud-agent/pkg"
	"fmt"
)

func main() {
	//
	//pkg.Install([]string{"fate-10000", "E:\\machenlong\\AI\\github\\owlet42\\KubeFATE\\k8s-deploy\\fate-10000"})
	//pkg.Install([]string{"fate-9999", "E:\\machenlong\\AI\\github\\owlet42\\KubeFATE\\k8s-deploy\\fate-9999"})
	//pkg.List("Table")
	//pkg.Delete([]string{"fate-10000"})
	//pkg.List("Table")
	//db.Db()
	fmt.Println(pkg.Namespace("allnamespaces"))
}
