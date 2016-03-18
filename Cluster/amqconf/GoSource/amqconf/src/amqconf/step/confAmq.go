package step

import (
	gd "github.com/Mooxe000/GoDash"
	"github.com/clbanning/mxj"
)

func confAmq(
	amqConfPath string,
	confPolicyEntry string,
	confLevelDB string) string {

	xmlStr := gd.StrFromFilePath(amqConfPath)
	xmlMap, _ := mxj.NewMapXml([]byte(xmlStr))

	step := Step{&xmlMap}

	// flow
	(&step).
		SetJvmHeap("80").
		RemoveLogQuery().
		InsertPolicyEntry(confPolicyEntry).
		RemoveKahadb().
		AddLevelDB(confLevelDB).
		FixAmpersand().
		FixColonKey()
	// dd(xmlMap)

	resultBytes, _ := mxj.AnyXmlIndent(
		xmlMap["beans"], "", "  ", "beans")
	r := gd.BytesToString(resultBytes)

	return r
}

func SavaAmqConf(ConfMap mxj.Map) {
	AmqFromPath := ConfMap["Amq"].
		(map[string]interface{})["From"].
		(string)
	AmqToPath := ConfMap["Amq"].
		(map[string]interface{})["To"].
		(string)
	confPolicyEntryPath := ConfMap["Amq"].
		(map[string]interface{})["PolicyEntry"].
		(string)
	confLevelDBPath := ConfMap["Amq"].
		(map[string]interface{})["LevelDB"].
		(string)

	confAmq := confAmq(AmqFromPath, confPolicyEntryPath, confLevelDBPath)

	prf("write amqconf to %s...", AmqFromPath)
	gd.StringToFilePath(AmqToPath, confAmq)
}
