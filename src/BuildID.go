package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

var gIDFileLimitCharacter = [255]int8{58: 1, 95: 1,
	48: 1, 49: 1, 50: 1, 51: 1, 52: 1, 53: 1, 54: 1, 55: 1, 56: 1, 57: 1,
	65: 1, 66: 1, 67: 1, 68: 1, 69: 1, 70: 1, 71: 1, 72: 1, 73: 1, 74: 1, 75: 1, 76: 1, 77: 1, 78: 1, 79: 1, 80: 1, 81: 1, 82: 1, 83: 1, 84: 1, 85: 1, 86: 1, 87: 1, 88: 1, 89: 1, 90: 1}

func formatIDFileLine(aData []byte) (string, bool) {
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

type sIDNode struct {
	Name string
	Num  int
}

func buildID() bool {
	if len(gCommand.IDPath) <= 0 {
		return true
	}
	pFile, err := os.Open(gCommand.IDPath)
	if err != nil {
		logErr(err.Error())
		return false
	}
	defer pFile.Close()

	pReader := bufio.NewReader(pFile)
	num := 0
	list := make([]*sIDNode, 0, 128)
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

		line, bOK := formatIDFileLine(bytes)
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
			pNode := new(sIDNode)
			pNode.Name = args[0]
			pNode.Num = num
			list = append(list, pNode)
		} else {
			logErr("illegal sentence in line: " + line)
			return false
		}
		num++
	}

	gBuildGo.buildID(list)
	gBuildLaya.buildID(list)
	gBuildCs.buildID(list)
	gBuildCocos.buildID(list)

	return true
}
