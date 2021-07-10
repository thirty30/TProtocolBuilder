package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type sBuildMessage struct {
	mMessageStructFile string
	mMessageIDFile     string
}

func (pOwn *sBuildMessage) init() bool {
	if len(pOwn.mMessageStructFile) == 0 || len(pOwn.mMessageIDFile) == 0 {
		logErr("need 2 files' path: message struct and message id definition file path")
		return false
	}
	return true
}

func (pOwn *sBuildMessage) clear() {

}

func (pOwn *sBuildMessage) getCommandDesc() string {
	return "-md [MessageStructFile] [MessageIDFile] : necessary command, the parms are message struct and message id definition file path eg. -md ./MsgStruct ./MsgID"
}

func (pOwn *sBuildMessage) verifyCommandParm(aParm []string) bool {
	if len(aParm) != 2 {
		logErr("need 2 files' path: message struct and message id definition file path")
		return false
	}
	pOwn.mMessageStructFile = aParm[0]
	pOwn.mMessageIDFile = aParm[1]
	return true
}

func (pOwn *sBuildMessage) formatStructFileLine(aData []byte) (string, bool) {
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

func (pOwn *sBuildMessage) buildMessageStruct() bool {
	pFile, err := os.Open(pOwn.mMessageStructFile)
	if err != nil {
		logErr(err.Error())
		return false
	}
	defer pFile.Close()

	pReader := bufio.NewReader(pFile)
	var pNewMsg *sMessage
	for {
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
		line, bOK := pOwn.formatStructFileLine(bytes)
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
			gMessageStructList = append(gMessageStructList, pNewMsg)
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
	return true
}

func (pOwn *sBuildMessage) formatIDFileLine(aData []byte) (string, bool) {
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
		if gIDFileLimitCharacter[c] == 0 {
			logErr("illegal character in this line: " + string(aData))
			return "", false
		}
		strValidLine += string(c)
	}
	return strValidLine, true
}

func (pOwn *sBuildMessage) buildMessageID() bool {
	pFile, err := os.Open(pOwn.mMessageIDFile)
	if err != nil {
		logErr(err.Error())
		return false
	}
	defer pFile.Close()

	pReader := bufio.NewReader(pFile)
	num := 0
	for {
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

		line, bOK := pOwn.formatIDFileLine(bytes)
		if bOK == false {
			return false
		}
		if len(line) == 0 {
			continue
		}
		args := strings.Split(line, ":")
		if len(args) == 2 {
			if args[0] == "TMSG_MARK" {
				num, err = strconv.Atoi(args[1])
				if err != nil {
					logErr("illegal mark num in line: " + line)
					return false
				}
			} else {
				logErr("illegal mark in line: " + line)
				return false
			}
		} else if len(args) == 1 {
			pNode := new(sMessageIDNode)
			pNode.Name = args[0]
			pNode.Num = num
			gMessageIDList = append(gMessageIDList, pNode)
		} else {
			logErr("illegal sentence in line: " + line)
			return false
		}
		num++
	}
	return true
}
