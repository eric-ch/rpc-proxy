package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Rule for rpc-proxy Firewall
type Rule struct {
	allow       bool
	Direction   Direction
	Subject     Subject
	Destination string
	Interface   string
	Member      string
	DomUUID     string
	DomID       string
	DomType     string
	Sender      string
	specStubdom bool
	Stubdom     bool
	ifBool      ifBool
	Int         int
}

type ifBool struct {
	condition  bool
	identifier string
}

//Direction ...
type Direction int

//Directions
const (
	Incoming Direction = iota
	Outgoing
)

//Subject of the message to be allowed
type Subject int

//Subjects
const (
	Any Subject = iota
	Call
	Signal
	Return
	Error
)

func tempReadConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ruleStr := scanner.Text()
		fmt.Println("Read line: ", ruleStr)
		if ruleStr == "" || ruleStr[0] == '#' {
			fmt.Println("TODO: SKIP")
		} else {
			ruleSlc := strings.Fields(ruleStr)

			r := createRule(ruleSlc)
			fmt.Println("Rule: ", r)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func createRule(ruleSlc []string) Rule {
	l := lex("Testing", ruleSlc)
	var newRule Rule
	for aItem := range l.items {
		//fmt.Println("Item: ", aItem)
		switch aItem.Ityp {
		case itemRule:
			if aItem.val == "allow" {
				newRule.allow = true
			} else {
				newRule.allow = false
			}
		case itemDirSub:
			dirsub, _ := strconv.Atoi(aItem.val)
			newRule.Direction = Direction(dirsub / 10)
			newRule.Subject = Subject(dirsub % 10)
		case itemDest:
			newRule.Destination = aItem.val
		case itemInter:
			newRule.Interface = aItem.val
		case itemMember:
			newRule.Member = aItem.val
		case itemDomUUID:
			newRule.DomUUID = aItem.val
		case itemDomID:
			newRule.DomID = aItem.val
		case itemDomType:
			newRule.DomType = aItem.val
		case itemSender:
			newRule.Sender = aItem.val
		case itemStubdom:
			newRule.specStubdom = true
			if aItem.val == "true" {
				newRule.Stubdom = true
			} else {
				newRule.Stubdom = false
			}
		case itemIfBoolTrue:
			newRule.ifBool.condition = true
			newRule.ifBool.identifier = aItem.val
		case itemIfBoolFalse:
			newRule.ifBool.condition = false
			newRule.ifBool.identifier = aItem.val

		}

	}
	//fmt.Println("The thread got closed!")

	return newRule
}
