package main

import "strings"

type sCommandItem struct {
	mCmd        string
	mParm       []string
	mBuilder    IBuilder
	mCanExecute bool
}

type IBuilder interface {
	init() bool
	clear()
	getCommandDesc() string
	initCommandParm(aParm []string) bool
	buildMessageStruct() bool
	buildMessageID() bool
}

type sMessageIDNode struct {
	Name string
	Num  int
}

type sMessageItem struct {
	Name string
	Type string
}

type sMessage struct {
	Name  string
	Nodes []*sMessageItem
}

func (pOwn *sMessage) init(aName string) {
	pOwn.Name = aName
	pOwn.Nodes = make([]*sMessageItem, 0, 8)
	gTypeKeyWord = append(gTypeKeyWord, aName)
}

func (pOwn *sMessage) appendNode() *sMessageItem {
	pTemp := new(sMessageItem)
	pOwn.Nodes = append(pOwn.Nodes, pTemp)
	return pTemp
}

var gIDFileLimitCharacter = [255]int8{58: 1, 95: 1,
	48: 1, 49: 1, 50: 1, 51: 1, 52: 1, 53: 1, 54: 1, 55: 1, 56: 1, 57: 1,
	65: 1, 66: 1, 67: 1, 68: 1, 69: 1, 70: 1, 71: 1, 72: 1, 73: 1, 74: 1, 75: 1, 76: 1, 77: 1, 78: 1, 79: 1, 80: 1, 81: 1, 82: 1, 83: 1, 84: 1, 85: 1, 86: 1, 87: 1, 88: 1, 89: 1, 90: 1,
	97: 1, 98: 1, 99: 1, 100: 1, 101: 1, 102: 1, 103: 1, 104: 1, 105: 1, 106: 1, 107: 1, 108: 1, 109: 1, 110: 1, 111: 1, 112: 1, 113: 1, 114: 1, 115: 1, 116: 1, 117: 1, 118: 1, 119: 1, 120: 1, 121: 1, 122: 1}

var gStructFileLimitCharacter = [255]int8{58: 1, 91: 1, 93: 1, 95: 1,
	48: 1, 49: 1, 50: 1, 51: 1, 52: 1, 53: 1, 54: 1, 55: 1, 56: 1, 57: 1,
	65: 1, 66: 1, 67: 1, 68: 1, 69: 1, 70: 1, 71: 1, 72: 1, 73: 1, 74: 1, 75: 1, 76: 1, 77: 1, 78: 1, 79: 1, 80: 1, 81: 1, 82: 1, 83: 1, 84: 1, 85: 1, 86: 1, 87: 1, 88: 1, 89: 1, 90: 1,
	97: 1, 98: 1, 99: 1, 100: 1, 101: 1, 102: 1, 103: 1, 104: 1, 105: 1, 106: 1, 107: 1, 108: 1, 109: 1, 110: 1, 111: 1, 112: 1, 113: 1, 114: 1, 115: 1, 116: 1, 117: 1, 118: 1, 119: 1, 120: 1, 121: 1, 122: 1}

func isArray(aType string) bool {
	return strings.Contains(aType, "[]")
}

func verifyType(aType string) bool {
	for _, v := range gTypeKeyWord {
		if aType == v {
			return true
		}

		if aType == ("[]" + v) {
			return true
		}
	}
	return false
}
