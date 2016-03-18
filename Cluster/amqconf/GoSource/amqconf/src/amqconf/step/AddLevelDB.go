package step

import (
	gd "github.com/Mooxe000/GoDash"
	"github.com/clbanning/mxj"
)

func (step *Step) AddLevelDB(confLevelDB string) *Step {
	xmlMap := step.Value

	prf("read LevelDBconfPath %s...\n", confLevelDB)
	nxmlStr := gd.StrFromFilePath(confLevelDB)
	nxmlMap, _ := mxj.NewMapXml([]byte(nxmlStr))

	m := make(map[string]interface{})
	m["replicatedLevelDB"] = nxmlMap["replicatedLevelDB"]

	paPath := xmlMap.PathForKeyShortest("persistenceAdapter")
	xmlMap.SetValueForPath(m, paPath)

	return step
}
