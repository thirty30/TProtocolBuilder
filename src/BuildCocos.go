package main

import (
	"fmt"
	"os"
)

type sBuildCocosFile struct {
	mFilePath        string
	mFile            *os.File
	mMapTypeRelation map[string]string
}

func (pOwn *sBuildCocosFile) getRealType(aItem *sMessageItem) string {
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
	//非基础类型
	return "TMsg" + strTempType
}

func (pOwn *sBuildCocosFile) init() bool {
	pOwn.mMapTypeRelation = make(map[string]string)
	pOwn.mMapTypeRelation["BOOL"] = "boolean"
	pOwn.mMapTypeRelation["N8"] = "number"
	pOwn.mMapTypeRelation["N16"] = "number"
	pOwn.mMapTypeRelation["N32"] = "number"
	pOwn.mMapTypeRelation["N64"] = "string"
	pOwn.mMapTypeRelation["U8"] = "number"
	pOwn.mMapTypeRelation["U16"] = "number"
	pOwn.mMapTypeRelation["U32"] = "number"
	pOwn.mMapTypeRelation["U64"] = "string"
	pOwn.mMapTypeRelation["F32"] = "number"
	pOwn.mMapTypeRelation["F64"] = "number"
	pOwn.mMapTypeRelation["STR"] = "string"

	var err error
	pOwn.mFile, err = os.Create(pOwn.mFilePath)
	if err != nil {
		logErr("can not create go file")
		return false
	}

	pOwn.mFile.WriteString("import { ByteBuffer } from '../../Core/Assist/ByteBuffer';\n")
	pOwn.mFile.WriteString("import { ISerializable, Serialize } from '../../Core/Assist/Serialize';\n")

	return true
}

func (pOwn *sBuildCocosFile) clear() {
	pOwn.mFile.Close()
}

func (pOwn *sBuildCocosFile) getCommandDesc() string {
	return "-cocos [file]: optional command. Generate go message file. eg. -cocos ./MsgDefine.ts"
}

func (pOwn *sBuildCocosFile) verifyCommandParm(aParm []string) bool {
	if len(aParm) != 1 {
		logErr("the command -cocos needs 1 (only 1) argument.")
		return false
	}
	pOwn.mFilePath = aParm[0]
	return true
}

func (pOwn *sBuildCocosFile) buildMessageStruct() bool {
	for _, msg := range gMessageStructList {
		pOwn.doBuildMessageStruct(msg)
	}
	return true
}

func (pOwn *sBuildCocosFile) doBuildMessageStruct(aMsg *sMessage) {
	strMsgName := aMsg.Name
	strContent := ""

	strContent += fmt.Sprintf("export class TMsg%s implements ISerializable\n{\n", strMsgName)

	//消息字段定义
	for _, node := range aMsg.Nodes {
		strRealType := pOwn.getRealType(node)
		if isArray(node.Type) == false {
			strContent += fmt.Sprintf("public %s : %s;\n", node.Name, strRealType)
		} else {
			strContent += fmt.Sprintf("private m%s : %s[];\n", node.Name, strRealType)
		}
	}

	//构造函数
	strContent += "constructor()\n{\n"
	for _, node := range aMsg.Nodes {
		strRealType := pOwn.getRealType(node)
		strBaseType, bOK := pOwn.mMapTypeRelation[node.Type]
		//基础类型
		if bOK == true {
			if strBaseType == "boolean" {
				strContent += fmt.Sprintf("this.%s = false;", node.Name)
			} else if strBaseType == "number" {
				strContent += fmt.Sprintf("this.%s = 0;", node.Name)
			} else {
				strContent += fmt.Sprintf("this.%s = \"\";", node.Name)
			}
			continue
		}
		//对象
		if isArray(node.Type) == false {
			strContent += fmt.Sprintf("this.%s = new %s();", node.Name, strRealType)
			continue
		}
		strContent += fmt.Sprintf("this.m%s = [];", node.Name)
	}
	strContent += "}\n"

	//数组方法
	for _, node := range aMsg.Nodes {
		strType := pOwn.getRealType(node)
		if isArray(node.Type) == false {
			continue
		}
		strContent += fmt.Sprintf("public Get%sCount() : number { return this.m%s.length; }\n", node.Name, node.Name)
		strContent += fmt.Sprintf("public Get%sAt(aIdx : number) : %s { return this.m%s[aIdx]; }\n", node.Name, strType, node.Name)
		strContent += fmt.Sprintf("public Append%s(rData : %s) : void { this.m%s.push(rData); }\n ", node.Name, strType, node.Name)
	}

	//序列化
	strContent += "public Serialize(rBuffer: ByteBuffer): void\n{\n"
	for _, node := range aMsg.Nodes {
		_, bOK := pOwn.mMapTypeRelation[node.Type]
		//基础类型
		if bOK == true {
			strContent += fmt.Sprintf("Serialize.Serialize%s(rBuffer, this.%s);\n", node.Type, node.Name)
			continue
		}
		//非基础类型
		if isArray(node.Type) == false {
			strContent += fmt.Sprintf("this.%s.Serialize(rBuffer);\n", node.Name)
			continue
		}
		//数组
		arrayType := node.Type[2:]
		_, bOK = pOwn.mMapTypeRelation[arrayType]
		strContent += fmt.Sprintf("let n%sCount = this.m%s.length;\nSerialize.SerializeN32(rBuffer, n%sCount);\n", node.Name, node.Name, node.Name)
		//基础类型数组
		if bOK == true {
			strContent += fmt.Sprintf("for (let i = 0; i < n%sCount; ++i){ Serialize.Serialize%s(rBuffer, this.m%s[i]); }\n", node.Name, arrayType, node.Name)
			continue
		}
		strContent += fmt.Sprintf("for (let i = 0; i < n%sCount; ++i){ this.m%s[i].Serialize(rBuffer); }\n", node.Name, node.Name)
	}
	strContent += "}\n"

	//反序列化
	strContent += "public Deserialize(rBuffer: ByteBuffer): void\n{\n"
	for _, node := range aMsg.Nodes {
		_, bOK := pOwn.mMapTypeRelation[node.Type]
		//基础类型
		if bOK == true {
			strContent += fmt.Sprintf("this.%s = Serialize.Deserialize%s(rBuffer);\n", node.Name, node.Type)
			continue
		}
		//非基础类型
		if isArray(node.Type) == false {
			strContent += fmt.Sprintf("this.%s.Deserialize(rBuffer);\n", node.Name)
			continue
		}
		//数组
		arrayType := node.Type[2:]
		strRealType := pOwn.getRealType(node)
		_, bOK = pOwn.mMapTypeRelation[arrayType]
		strContent += fmt.Sprintf("let n%sCount = Serialize.DeserializeN32(rBuffer);\n", node.Name)
		//基础类型数组
		if bOK == true {
			strContent += fmt.Sprintf("for (let i = 0; i < n%sCount; ++i){ this.m%s[i] = Serialize.Deserialize%s(rBuffer); }\n", node.Name, node.Name, arrayType)
			continue
		}
		strContent += fmt.Sprintf("for (let i = 0; i < n%sCount; ++i){ this.m%s[i] = new %s();\nthis.m%s[i].Deserialize(rBuffer);\n }\n", node.Name, node.Name, strRealType, node.Name)
	}

	strContent += "}\n"

	strContent += "}\n"
	pOwn.mFile.WriteString(strContent)
}

func (pOwn *sBuildCocosFile) buildMessageID() bool {
	strContent := "export enum TMsgID {\n"
	for _, node := range gMessageIDList {
		strContent += fmt.Sprintf("%s = %d,\n", node.Name, node.Num)
	}
	strContent += "}\n"
	pOwn.mFile.WriteString(strContent)
	return true
}
