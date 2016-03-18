package step

import (
	gd "github.com/Mooxe000/GoDash"
	"github.com/clbanning/mxj"
)

func (step *Step) InsertPolicyEntry(confPolicyEntry string) *Step {
	xmlMap := step.Value

	prf("read PolicyEntryPath %s...\n", confPolicyEntry)
	nxmlStr := gd.StrFromFilePath(confPolicyEntry)
	nxmlMap, _ := mxj.NewMapXml([]byte(nxmlStr))

	pePath := xmlMap.PathForKeyShortest("policyEntry")
	peValues, _ := xmlMap.ValuesForPath(pePath)
	peValues = append(peValues, nxmlMap["policyEntry"])

	xmlMap.SetValueForPath(peValues, pePath)

	return step
}
