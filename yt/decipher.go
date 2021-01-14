package yt

import (
	"fmt"
	"github.com/kraxarn/website/common"
	"regexp"
	"strings"
)

func getFuncId(line string) string {
	// TODO
	return ""
}

func matches(funcId, expression string) bool  {
	var expr *regexp.Regexp
	if expr, err := regexp.Compile(fmt.Sprintf("%s%s", funcId, expression)); err != nil {
		return false
	}
	return expr.
}

func decipher(cipher string) (string, error) {
	if len(baseJs) == 0 {
		baseJs, err := common.Get(baseJsUrl)
		if err != nil {
			return "", err
		}
	}

	funcExpr := regexp.MustCompile(`(\w+)=function\(\w+\){(\w+)=\2\.split\(\x22{2}\);.*?return\s+\2\.join\(\x22{2}\)}`)
	funcName := funcExpr.FindString(baseJs)
	if strings.Contains(funcName, "$") {
		funcName = `\` + funcName
	}

	lines := strings.Split(funcName, ";")

	var idReverse, idSlice, idCharSwap, funcId, operations string

	for _, line := range lines[1:len(lines)-2] {
		if len(idReverse) > 0 && len(idSlice) > 0 && len(idCharSwap) > 0 {
			break
		}

		funcId = getFuncId(line)

		if expr, err := regexp.Compile(fmt.Sprintf("%s:\bfunction\b\(\w+\)", funcId))
	}
}