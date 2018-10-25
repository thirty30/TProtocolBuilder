package main

import (
	"fmt"
	"os"
)

var gBuildGo sBuildGoFile
var gBuildLaya sBuildLayaFile
var gBuildCs sBuildCsFile
var gBuildCocos sBuildCocosFile

func main() {
	//解析命令
	if analysisArgs(os.Args[1:]) == false {
		return
	}

	//初始化构造器
	if gBuildGo.init() == false {
		return
	}
	if gBuildLaya.init() == false {
		return
	}
	if gBuildCs.init() == false {
		return
	}
	if gBuildCocos.init() == false {
		return
	}

	//解析消息结构
	if buildStruct() == false {
		return
	}
	//解析消息ID
	if buildID() == false {
		return
	}

	//清理构造器
	gBuildGo.clear()
	gBuildLaya.clear()
	gBuildCs.clear()
	gBuildCocos.clear()

	log("[SUCCESS] Generate message completely!")

}

func log(aContent string) {
	fmt.Println(aContent)
}

func logErr(aContent string) {
	fmt.Println("[ERROR] " + aContent)
}
