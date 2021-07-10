package main

import (
	"fmt"
	"os"
)

type sBuildGoFile struct {
	mFilePath        string
	mFile            *os.File
	mMapTypeRelation map[string]string
}

func (pOwn *sBuildGoFile) getRealType(aItem *sMessageItem) string {
	strTempType := ""
	//得到真实的类型
	if aItem.isArray() == true {
		strTempType = aItem.Type[2:]
	} else {
		strTempType = aItem.Type
	}

	strType, bOK := pOwn.mMapTypeRelation[strTempType]
	//基础类型
	if bOK == true {
		return strType
	}
	//非基础类型的数组用指针
	if aItem.isArray() == true {
		return "*" + strTempType
	}
	//非基础类型的对象不用指针
	return strTempType
}

func (pOwn *sBuildGoFile) init() bool {
	pOwn.mMapTypeRelation = make(map[string]string)
	pOwn.mMapTypeRelation["BOOL"] = "bool"
	pOwn.mMapTypeRelation["N8"] = "int8"
	pOwn.mMapTypeRelation["N16"] = "int16"
	pOwn.mMapTypeRelation["N32"] = "int32"
	pOwn.mMapTypeRelation["N64"] = "int64"
	pOwn.mMapTypeRelation["U8"] = "uint8"
	pOwn.mMapTypeRelation["U16"] = "uint16"
	pOwn.mMapTypeRelation["U32"] = "uint32"
	pOwn.mMapTypeRelation["U64"] = "uint64"
	pOwn.mMapTypeRelation["F32"] = "float32"
	pOwn.mMapTypeRelation["F64"] = "float64"
	pOwn.mMapTypeRelation["STR"] = "string"

	var err error
	pOwn.mFile, err = os.Create(pOwn.mFilePath)
	if err != nil {
		logErr("can not create go file")
		return false
	}

	pOwn.mFile.WriteString("package tprotocol\n")
	return true
}

func (pOwn *sBuildGoFile) clear() {
	pOwn.mFile.Close()
}

func (pOwn *sBuildGoFile) getCommandDesc() string {
	return "-go [file]: optional command. Generate go message file. eg. -go ./MsgDefine.go"
}

func (pOwn *sBuildGoFile) verifyCommandParm(aParm []string) bool {
	if len(aParm) != 1 {
		logErr("the command -go needs 1 (only 1) argument.")
		return false
	}
	pOwn.mFilePath = aParm[0]
	return true
}

func (pOwn *sBuildGoFile) buildMessageStruct() bool {
	for _, msg := range gMessageStructList {
		pOwn.doBuildMessageStruct(msg)
	}
	return true
}

func (pOwn *sBuildGoFile) doBuildMessageStruct(aMsg *sMessage) {
	strMsgName := aMsg.Name
	strContent := ""
	//消息定义
	strContent += fmt.Sprintf("type %s struct{\n", strMsgName)
	for _, node := range aMsg.Nodes {
		strType := pOwn.getRealType(node)
		if node.isArray() == true {
			strContent += fmt.Sprintf("m%s []%s\n", node.Name, strType)
		} else {
			strContent += fmt.Sprintf("%s %s\n", node.Name, strType)
		}
	}
	strContent += "}\n"

	//数组方法
	for _, node := range aMsg.Nodes {
		strType := pOwn.getRealType(node)
		if node.isArray() == false {
			continue
		}
		strContent += fmt.Sprintf("func (pOwn *%s) Get%sCount() int32 { return int32(len(pOwn.m%s)) }\n", strMsgName, node.Name, node.Name)
		strContent += fmt.Sprintf("func (pOwn *%s) Get%sAt(aIdx int32) %s { return pOwn.m%s[aIdx] }\n", strMsgName, node.Name, strType, node.Name)
		strContent += fmt.Sprintf("func (pOwn *%s) Append%s(aData %s) {\n ", strMsgName, node.Name, strType)
		strContent += fmt.Sprintf("if pOwn.m%s == nil { pOwn.m%s = make([]%s, 0, 8) }\n ", node.Name, node.Name, strType)
		strContent += fmt.Sprintf("pOwn.m%s = append(pOwn.m%s, aData)\n }\n", node.Name, node.Name)
	}

	//序列化
	strContent += fmt.Sprintf("func (pOwn *%s) Serialize(aBuffer []byte, aSize uint32) uint32{\n", strMsgName)
	strContent += "var nOffset uint32\n"
	for _, node := range aMsg.Nodes {
		_, bOK := pOwn.mMapTypeRelation[node.Type]
		//基础类型
		if bOK == true {
			strContent += fmt.Sprintf("nOffset += serialize%s(aBuffer[nOffset:], aSize-nOffset, pOwn.%s)\n", node.Type, node.Name)
			continue
		}
		//非基础类型对象
		if node.isArray() == false {
			strContent += fmt.Sprintf("nOffset += pOwn.%s.Serialize(aBuffer[nOffset:], aSize-nOffset)\n", node.Name)
			continue
		}
		//数组
		arrayType := node.Type[2:]
		_, bOK = pOwn.mMapTypeRelation[arrayType]
		strContent += fmt.Sprintf("n%sCount := len(pOwn.m%s)\n", node.Name, node.Name)
		strContent += fmt.Sprintf("nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, int32(n%sCount))\n", node.Name)
		//基础类型数组
		if bOK == true {
			strContent += fmt.Sprintf("for i := 0; i < n%sCount; i++ { nOffset += serialize%s(aBuffer[nOffset:], aSize-nOffset, pOwn.m%s[i]) }\n", node.Name, arrayType, node.Name)
			continue
		}
		//非基础类型数组
		strContent += fmt.Sprintf("for i := 0; i < n%sCount; i++ { nOffset += pOwn.m%s[i].Serialize(aBuffer[nOffset:], aSize-nOffset) }\n", node.Name, node.Name)
	}
	strContent += "return nOffset\n}\n"

	//反序列化
	strContent += fmt.Sprintf("func (pOwn *%s) Deserialize(aBuffer []byte, aSize uint32) uint32{\n", strMsgName)
	strContent += "var nOffset uint32\n"
	strContent += "var nTemp uint32\n"
	for _, node := range aMsg.Nodes {
		_, bOK := pOwn.mMapTypeRelation[node.Type]
		//基础类型
		if bOK == true {
			strContent += fmt.Sprintf("pOwn.%s, nTemp = deserialize%s(aBuffer[nOffset:], aSize-nOffset)\nnOffset += nTemp\n", node.Name, node.Type)
			continue
		}
		//非基础类型对象
		if node.isArray() == false {
			strContent += fmt.Sprintf("nOffset += pOwn.%s.Deserialize(aBuffer[nOffset:], aSize-nOffset)\n", node.Name)
			continue
		}
		//数组
		arrayType := node.Type[2:]
		strRealType, bOK := pOwn.mMapTypeRelation[arrayType]
		strContent += fmt.Sprintf("n%sCount, nTemp := deserializeN32(aBuffer[nOffset:], aSize-nOffset)\nnOffset += nTemp\n", node.Name)
		strContent += fmt.Sprintf("for i := 0; i < int(n%sCount); i++ {\n", node.Name)
		//基础类型数组
		if bOK == true {
			strContent += fmt.Sprintf("obj, nTemp := deserialize%s(aBuffer[nOffset:], aSize-nOffset)\nnOffset += nTemp\n", arrayType)
			strContent += fmt.Sprintf("if pOwn.m%s == nil { pOwn.m%s = make([]%s, 0, 8) }\n", node.Name, node.Name, strRealType)
			strContent += fmt.Sprintf("pOwn.m%s = append(pOwn.m%s, obj)\n", node.Name, node.Name)
		} else {
			//非基础类型数组
			strContent += fmt.Sprintf("var obj %s\n", arrayType)
			strContent += "nOffset += obj.Deserialize(aBuffer[nOffset:], aSize-nOffset)\n"
			strContent += fmt.Sprintf("if pOwn.m%s == nil { pOwn.m%s = make([]*%s, 0, 8) }\n", node.Name, node.Name, arrayType)
			strContent += fmt.Sprintf("pOwn.m%s = append(pOwn.m%s, &obj)\n", node.Name, node.Name)
		}
		strContent += "}\n"
	}

	strContent += "return nOffset\n}\n"

	pOwn.mFile.WriteString(strContent)
}

func (pOwn *sBuildGoFile) buildMessageID() bool {
	strContent := "const (\n"
	for _, node := range gMessageIDList {
		strContent += fmt.Sprintf("%s = %d\n", node.Name, node.Num)
	}
	strContent += ")\n"
	pOwn.mFile.WriteString(strContent)
	return true
}
