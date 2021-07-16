package main

var gMessageStructList []*sMessage
var gMessageIDList []*sMessageIDNode
var gCommandItems []*sCommandItem
var gTypeKeyWord = []string{"BOOL", "N8", "N16", "N32", "N64", "U8", "U16", "U32", "U64", "F32", "F64", "STR"}

func main() {
	gMessageStructList = make([]*sMessage, 0, 1024)
	gMessageIDList = make([]*sMessageIDNode, 0, 1024)

	gCommandItems = make([]*sCommandItem, 0, 16)
	gCommandItems = append(gCommandItems, &sCommandItem{mCmd: "-md", mBuilder: new(sBuildMessage), mParm: make([]string, 0, 2), mCanExecute: false})
	gCommandItems = append(gCommandItems, &sCommandItem{mCmd: "-go", mBuilder: new(sBuildGoFile), mParm: make([]string, 0, 2), mCanExecute: false})
	gCommandItems = append(gCommandItems, &sCommandItem{mCmd: "-cs", mBuilder: new(sBuildCSFile), mParm: make([]string, 0, 2), mCanExecute: false})
	gCommandItems = append(gCommandItems, &sCommandItem{mCmd: "-laya", mBuilder: new(sBuildLayaFile), mParm: make([]string, 0, 2), mCanExecute: false})
	gCommandItems = append(gCommandItems, &sCommandItem{mCmd: "-cocos", mBuilder: new(sBuildCocosFile), mParm: make([]string, 0, 2), mCanExecute: false})

	initConsoleColor()

	//解析命令
	if analysisArgs() == false {
		return
	}

	//初始化builder参数
	for _, v := range gCommandItems {
		if v.mCmd == "-md" && v.mCanExecute == false {
			logErr("lack necessary option -md.")
			return
		}
		if v.mCanExecute == true && v.mBuilder.initCommandParm(v.mParm) == false {
			return
		}
	}

	//解析并生成
	for _, v := range gCommandItems {
		if v.mCanExecute == false {
			continue
		}
		if v.mBuilder.init() == false {
			return
		}
		if v.mBuilder.buildMessageStruct() == false {
			return
		}
		if v.mBuilder.buildMessageID() == false {
			return
		}
		v.mBuilder.clear()
	}

	log("[SUCCESS] Generate message completely!")
}
