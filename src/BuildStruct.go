package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var gStructFileLimitCharacter = [255]int8{58: 1, 91: 1, 93: 1, 95: 1,
	48: 1, 49: 1, 50: 1, 51: 1, 52: 1, 53: 1, 54: 1, 55: 1, 56: 1, 57: 1,
	65: 1, 66: 1, 67: 1, 68: 1, 69: 1, 70: 1, 71: 1, 72: 1, 73: 1, 74: 1, 75: 1, 76: 1, 77: 1, 78: 1, 79: 1, 80: 1, 81: 1, 82: 1, 83: 1, 84: 1, 85: 1, 86: 1, 87: 1, 88: 1, 89: 1, 90: 1,
	97: 1, 98: 1, 99: 1, 100: 1, 101: 1, 102: 1, 103: 1, 104: 1, 105: 1, 106: 1, 107: 1, 108: 1, 109: 1, 110: 1, 111: 1, 112: 1, 113: 1, 114: 1, 115: 1, 116: 1, 117: 1, 118: 1, 119: 1, 120: 1, 121: 1, 122: 1}

func formatStructFileLine(aData []byte) (string, bool) {
	nLen := len(aData)
	strValidLine := ""
	for i := 0; i < nLen; i++ {
		c := aData[i]
		if c == ' ' || c == '\r' || c == '\n' || c == '\t' {
			continue
		}
		if c == '/' {
			break
		}
		if gStructFileLimitCharacter[c] == 0 {
			logErr("illegal character in this line: " + string(aData))
			return "", false
		}
		strValidLine += string(c)
	}
	return strValidLine, true
}

type sNode struct {
	Name string
	Type string
}

func (pOwn *sNode) isArray() bool {
	nPos := strings.LastIndex(pOwn.Type, "[]")
	if nPos != 0 {
		return false
	}
	return true
}

func (pOwn *sNode) checkType() bool {
	nPos := strings.LastIndex(pOwn.Type, "[]")
	if nPos > 0 {
		return false
	}
	return true
}

type sMessage struct {
	Name  string
	Nodes []*sNode
}

func (pOwn *sMessage) init(aName string) {
	pOwn.Name = aName
	pOwn.Nodes = make([]*sNode, 0, 8)
}

func (pOwn *sMessage) appendNode() *sNode {
	pTemp := new(sNode)
	pOwn.Nodes = append(pOwn.Nodes, pTemp)
	return pTemp
}

func buildStruct() bool {
	if len(gCommand.StructPath) <= 0 {
		return true
	}
	pFile, err := os.Open(gCommand.StructPath)
	if err != nil {
		logErr(err.Error())
		return false
	}
	defer pFile.Close()

	pReader := bufio.NewReader(pFile)
	var pNewMsg *sMessage
	msgList := make([]*sMessage, 0, 256)
	for true {
		bytes, bPrefix, err := pReader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			logErr(err.Error())
			return false
		}
		if bPrefix == true {
			logErr("the line is too long")
			return false
		}
		line, bOK := formatStructFileLine(bytes)
		if bOK == false {
			return false
		}
		if len(line) == 0 {
			continue
		}
		args := strings.Split(line, ":")
		if len(args) != 2 {
			logErr("wrong format in line: " + line)
			return false
		}

		if args[0] == "TMSG_BEGIN" {
			if pNewMsg != nil {
				logErr("incompleted message define: " + pNewMsg.Name)
				return false
			}
			pNewMsg = new(sMessage)
			pNewMsg.init(args[1])
		} else if args[0] == "TMSG_END" {
			if pNewMsg == nil {
				logErr("illegal message end: " + line)
				return false
			}
			msgList = append(msgList, pNewMsg)
			pNewMsg = nil
		} else {
			if pNewMsg == nil {
				logErr("illegal message item: " + line)
				return false
			}
			pNode := pNewMsg.appendNode()
			pNode.Name = args[0]
			pNode.Type = args[1]
			//判断数组语法格式
			if pNode.checkType() == false {
				logErr("illegal array define: " + line)
				return false
			}
		}
	}
	//开始转换代码
	doBuild(msgList)
	return true
}

func doBuild(aList []*sMessage) {
	for _, msg := range aList {
		gBuildGo.buildStruct(msg)
		gBuildLaya.buildStruct(msg)
		gBuildCs.buildStruct(msg)
		gBuildCocos.buildStruct(msg)
	}
}
