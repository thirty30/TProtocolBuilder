package main

import "os"

//eg. MsgBuilder.exe -s MsgStruct -i MsgID -go ./MsgDefine.go -ts./MsgDefine.ts -cs ./MsgDefine.cs
func analysisArgs() bool {
	//解析指令
	args := os.Args[1:]
	var pItem *sCommandItem = nil
	for i := 0; i < len(args); i++ {
		parm := args[i]
		if parm[0] == '-' {
			pItem = findCommandItem(parm)
			if pItem == nil {
				logErr("illegal command:" + parm)
				return false
			}
			pItem.mCanExecute = true
		} else {
			pItem.mParm = append(pItem.mParm, parm)
		}
	}

	//没有选项的情况显示帮助信息
	{
		t := 0
		for _, v := range gCommandItems {
			if v.mCanExecute == false {
				t++
			}
		}
		if t == len(gCommandItems) {
			for _, v := range gCommandItems {
				log(v.mBuilder.getCommandDesc())
			}
			return false
		}
	}

	//检查指令和参数是否匹配
	for _, v := range gCommandItems {
		if v.mCmd == "-md" && v.mCanExecute == false {
			logErr("lack necessary option -md.")
			return false
		}
		if v.mCanExecute == true && v.mBuilder.verifyCommandParm(v.mParm) == false {
			return false
		}
	}

	return true
}

func findCommandItem(aCmd string) *sCommandItem {
	for _, v := range gCommandItems {
		if aCmd == v.mCmd {
			return v
		}
	}
	return nil
}
