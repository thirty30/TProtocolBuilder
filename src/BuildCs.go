package main

import (
	"fmt"
	"os"
)

type sBuildCSFile struct {
	mFilePath        string
	mFile            *os.File
	mMapTypeRelation map[string]string
}

func (pOwn *sBuildCSFile) getRealType(aItem *sMessageItem) string {
	strTempType := ""
	//得到真实的类型
	if isArray(aItem.Type) == true {
		strTempType = aItem.Type[2:]
	} else {
		strTempType = aItem.Type
	}

	strType, bOK := pOwn.mMapTypeRelation[strTempType]
	//基础类型
	if bOK == true {
		return strType
	}
	return strTempType
}

func (pOwn *sBuildCSFile) init() bool {
	pOwn.mMapTypeRelation = make(map[string]string)
	pOwn.mMapTypeRelation["BOOL"] = "bool"
	pOwn.mMapTypeRelation["N8"] = "byte"
	pOwn.mMapTypeRelation["N16"] = "short"
	pOwn.mMapTypeRelation["N32"] = "int"
	pOwn.mMapTypeRelation["N64"] = "long"
	pOwn.mMapTypeRelation["U8"] = "byte"
	pOwn.mMapTypeRelation["U16"] = "ushort"
	pOwn.mMapTypeRelation["U32"] = "uint"
	pOwn.mMapTypeRelation["U64"] = "ulong"
	pOwn.mMapTypeRelation["F32"] = "float"
	pOwn.mMapTypeRelation["F64"] = "double"
	pOwn.mMapTypeRelation["STR"] = "string"

	var err error
	pOwn.mFile, err = os.Create(pOwn.mFilePath)
	if err != nil {
		logErr("can not create cs file")
		return false
	}

	pOwn.mFile.WriteString("using System;\nusing System.Collections.Generic;\nusing System.IO;\nusing System.Text;\nusing TNet;\n")
	pOwn.mFile.WriteString("namespace Game.Message\n{\n")

	return true
}

func (pOwn *sBuildCSFile) clear() {
	pOwn.mFile.WriteString("\n}\n")
	pOwn.mFile.Close()
}

func (pOwn *sBuildCSFile) getCommandDesc() string {
	return "-cs [file]: optional command. Generate c# message file. eg. -cs ./MsgDefine.cs"
}

func (pOwn *sBuildCSFile) initCommandParm(aParm []string) bool {
	if len(aParm) != 1 {
		logErr("the command -cs needs 1 (only 1) argument.")
		return false
	}
	pOwn.mFilePath = aParm[0]
	return true
}

func (pOwn *sBuildCSFile) buildMessageStruct() bool {
	for _, msg := range gMessageStructList {
		pOwn.doBuildMessageStruct(msg)
	}
	return true
}

func (pOwn *sBuildCSFile) doBuildMessageStruct(aMsg *sMessage) {

	strMsgName := aMsg.Name
	strContent := ""
	//消息定义
	strContent += fmt.Sprintf("public class TMsg%s : ITNetMessage\n{\n", strMsgName)
	for _, node := range aMsg.Nodes {
		strType := pOwn.getRealType(node)
		if isArray(node.Type) == true {
			strContent += fmt.Sprintf("public List<%s> %s = new List<%s>();\n", strType, node.Name, strType)
		} else {
			_, bOK := pOwn.mMapTypeRelation[node.Type]
			if bOK == true {
				strContent += fmt.Sprintf("public %s %s;\n", strType, node.Name)
			} else {
				strContent += fmt.Sprintf("public %s %s = new %s();\n", strType, node.Name, strType)
			}
		}
	}

	//序列化
	strContent += "public void Serialize(byte[] aBuffer, int aSize, ref int nOffset)\n{\n"
	for _, node := range aMsg.Nodes {
		_, bOK := pOwn.mMapTypeRelation[node.Type]
		//基础类型
		if bOK == true {
			strContent += fmt.Sprintf("TNetEncode.Serialize%s(this.%s, aBuffer, ref nOffset);\n", node.Type, node.Name)
			continue
		}
		//非基础类型对象
		if isArray(node.Type) == false {
			strContent += fmt.Sprintf("this.%s.Serialize(aBuffer, aSize - nOffset, ref nOffset);\n", node.Name)
			continue
		}
		//数组
		arrayType := node.Type[2:]
		_, bOK = pOwn.mMapTypeRelation[arrayType]
		strContent += fmt.Sprintf("int n%sCount = this.%s.Count;\n", node.Name, node.Name)
		strContent += fmt.Sprintf("TNetEncode.SerializeN32(n%sCount, aBuffer, ref nOffset);\n", node.Name)
		//基础类型数组
		if bOK == true {
			strContent += fmt.Sprintf("for (int i = 0; i < n%sCount; i++){TNetEncode.Serialize%s(this.%s[i], aBuffer, ref nOffset);}\n", node.Name, arrayType, node.Name)
			continue
		}
		//非基础类型数组
		strContent += fmt.Sprintf("for (int i = 0; i < n%sCount; i++){this.%s[i].Serialize(aBuffer, aSize - nOffset, ref nOffset);}\n", node.Name, node.Name)
	}
	strContent += "\n}\n"

	//反序列化
	strContent += "public void Deserialize(byte[] aBuffer, int aSize, ref int nOffset)\n{\n"
	for _, node := range aMsg.Nodes {
		_, bOK := pOwn.mMapTypeRelation[node.Type]
		//基础类型
		if bOK == true {
			strContent += fmt.Sprintf("this.%s = TNetEncode.Deserialize%s(aBuffer, ref nOffset);\n", node.Name, node.Type)
			continue
		}
		//非基础类型对象
		if isArray(node.Type) == false {
			strContent += fmt.Sprintf("this.%s.Deserialize(aBuffer, aSize - nOffset, ref nOffset);\n", node.Name)
			continue
		}
		//数组
		arrayType := node.Type[2:]
		_, bOK = pOwn.mMapTypeRelation[arrayType]
		strContent += fmt.Sprintf("int n%sCount = TNetEncode.DeserializeN32(aBuffer, ref nOffset);\n", node.Name)
		//基础类型数组
		if bOK == true {
			strContent += fmt.Sprintf("for(int i = 0; i < n%sCount; i++)\n{\n", node.Name)
			strContent += fmt.Sprintf("this.%s.Add(TNetEncode.Deserialize%s(aBuffer, ref nOffset));\n}\n", node.Name, arrayType)
			continue
		}
		//非基础类型数组
		strContent += fmt.Sprintf("for(int i = 0; i < n%sCount; i++)\n{\n", node.Name)
		strContent += fmt.Sprintf("%s temp = new %s();temp.Deserialize(aBuffer, aSize - nOffset, ref nOffset);this.%s.Add(temp);\n}\n", arrayType, arrayType, node.Name)
	}
	strContent += "\n}\n"

	strContent += "}\n"

	pOwn.mFile.WriteString(strContent)
}

func (pOwn *sBuildCSFile) buildMessageID() bool {
	strContent := "public class TMessageID\n{\n"
	for _, node := range gMessageIDList {
		strContent += fmt.Sprintf("public const int %s = %d;\n", node.Name, node.Num)
	}
	strContent += "}\n"
	pOwn.mFile.WriteString(strContent)
	return true
}
