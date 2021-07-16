package main

import "os"

//eg. MsgBuilder.exe -s MsgStruct -i MsgID -go ./MsgDefine.go -ts./MsgDefine.ts -cs ./MsgDefine.cs
func analysisArgs() bool {
	//没有选项的情况显示帮助信息
	if len(os.Args) <= 1 {
		for _, v := range gCommandItems {
			log(v.mBuilder.getCommandDesc())
		}
		return false
	}

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
			if pItem == nil {
				logErr("illegal command:" + parm)
				return false
			}
			pItem.mParm = append(pItem.mParm, parm)
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
