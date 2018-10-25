package main

var gCommand sCommand //指令
type sCommand struct {
	StructPath    string
	IDPath        string
	GoFileName    string
	CsFileName    string
	LayaFileName  string
	CocosFileName string
}

func init() {
	gCommand.StructPath = ""
	gCommand.IDPath = ""
	gCommand.GoFileName = ""
	gCommand.CsFileName = ""
	gCommand.LayaFileName = ""
	gCommand.CocosFileName = ""
}

//eg. MsgBuilder.exe -s MsgStruct -i MsgID -go ./MsgDefine.go -ts./MsgDefine.ts -cs ./MsgDefine.cs
func analysisArgs(args []string) bool {
	var curFunc dealCommand
	for i := 0; i < len(args); i++ {
		parm := args[i]
		if parm[0] == '-' {
			if curFunc != nil {
				if curFunc("") == false {
					return false
				}
			}
			curFunc = checkCommand(parm)
			if curFunc == nil {
				logErr("illegal command:" + parm)
				return false
			}
		} else {
			if curFunc == nil {
				logErr("illegal command:" + parm)
				return false
			}
			if curFunc(parm) == false {
				return false
			}
			curFunc = nil
		}
	}

	if curFunc != nil {
		if curFunc("") == false {
			return false
		}
	}

	return true
}

func checkCommand(aCommand string) dealCommand {
	switch aCommand {
	case "-s":
		return dealCommandS
	case "-i":
		return dealCommandI
	case "-go":
		return dealCommandGO
	case "-cs":
		return dealCommandCS
	case "-laya":
		return dealCommandLaya
	case "-cocos":
		return dealCommandCocos
	case "-h":
		return dealCommandHelp
	}
	return nil
}

type dealCommand func(arg string) bool

func dealCommandS(arg string) bool {
	if len(arg) <= 0 {
		logErr("the command -s lack of arg.")
		return false
	}
	gCommand.StructPath = arg

	return true
}

func dealCommandI(arg string) bool {
	if len(arg) <= 0 {
		logErr("the command -i lack of arg.")
		return false
	}
	gCommand.IDPath = arg

	return true
}

func dealCommandGO(arg string) bool {
	if len(arg) <= 0 {
		logErr("the command -go lack of arg.")
		return false
	}
	if len(arg) < 3 || arg[len(arg)-3:] != ".go" {
		logErr("the arg of -go must be a file name. like ./Protocol/MsgDefine.go")
		return false
	}
	gCommand.GoFileName = arg
	return true
}

func dealCommandCS(arg string) bool {
	if len(arg) <= 0 {
		logErr("the command -cs lack of arg.")
		return false
	}
	if len(arg) < 3 || arg[len(arg)-3:] != ".cs" {
		logErr("the arg of -cs must be a file name. like ./Protocol/MsgDefine.cs")
		return false
	}
	gCommand.CsFileName = arg
	return true
}

func dealCommandLaya(arg string) bool {
	if len(arg) <= 0 {
		logErr("the command -ts lack of arg.")
		return false
	}
	if len(arg) < 3 || arg[len(arg)-3:] != ".ts" {
		logErr("the arg of -ts must be a file name. like ./Protocol/MsgDefine.ts")
		return false
	}
	gCommand.LayaFileName = arg
	return true
}

func dealCommandCocos(arg string) bool {
	if len(arg) <= 0 {
		logErr("the command -ts lack of arg.")
		return false
	}
	if len(arg) < 3 || arg[len(arg)-3:] != ".ts" {
		logErr("the arg of -ts must be a file name. like ./Protocol/MsgDefine.ts")
		return false
	}
	gCommand.CocosFileName = arg
	return true
}

func dealCommandHelp(arg string) bool {
	log("-s : [path] input message structs file path eg. ./MsgStruct")
	log("-i : [path] input message ID file path eg. ./MsgID")
	log("-go : optional command. [path] go file path and file name. eg. ./Protocol/MsgDefine.go")
	log("-ts : optional command. [path] ts file path and file name. eg. ./Protocol/MsgDefine.ts")
	log("-cs : optional command. [path] cs file path and file name. eg. ./Protocol/MsgDefine.cs")
	return false
}
