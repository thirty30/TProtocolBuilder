package main

import (
	"fmt"
)

var gMessageStructList []*sMessage
var gMessageIDList []*sMessageIDNode
var gCommandItems []*sCommandItem

func main() {
	gMessageStructList = make([]*sMessage, 0, 1024)
	gMessageIDList = make([]*sMessageIDNode, 0, 1024)

	gCommandItems = make([]*sCommandItem, 0, 16)
	gCommandItems = append(gCommandItems, &sCommandItem{mCmd: "-md", mBuilder: new(sBuildMessage), mParm: make([]string, 0, 2), mCanExecute: false})
	gCommandItems = append(gCommandItems, &sCommandItem{mCmd: "-go", mBuilder: new(sBuildGoFile), mParm: make([]string, 0, 2), mCanExecute: false})
	gCommandItems = append(gCommandItems, &sCommandItem{mCmd: "-cs", mBuilder: new(sBuildCSFile), mParm: make([]string, 0, 2), mCanExecute: false})
	gCommandItems = append(gCommandItems, &sCommandItem{mCmd: "-laya", mBuilder: new(sBuildLayaFile), mParm: make([]string, 0, 2), mCanExecute: false})
	gCommandItems = append(gCommandItems, &sCommandItem{mCmd: "-cocos", mBuilder: new(sBuildCocosFile), mParm: make([]string, 0, 2), mCanExecute: false})

	//解析命令
	if analysisArgs() == false {
		return
	}

	for _, v := range gCommandItems {
		if v.mCanExecute == false {
			continue
		}
		if v.mBuilder.init() == false {
			return
		}
		v.mBuilder.buildMessageStruct()
		v.mBuilder.buildMessageID()
		v.mBuilder.clear()
	}

	log("[SUCCESS] Generate message completely!")
}

func log(aContent string) {
	fmt.Println(aContent)
}

func logErr(aContent string) {
	fmt.Println("[ERROR] " + aContent)
}
