package step

import (
	gd "github.com/Mooxe000/GoDash"
	"github.com/clbanning/mxj"
)

func confJetty(jettyConfPath string, portValue string) string {
	xmlStr := gd.StrFromFilePath(jettyConfPath)
	// pln(xmlStr)
	xmlMap, _ := mxj.NewMapXml([]byte(xmlStr))
	// dd(xmlMap)
	// pln(typeof(xmlMap))

	step := Step{&xmlMap}
	(&step).
		SetJettyPort(portValue).
		FixColonKey()
	// dd(xmlMap)

	resultBytes, _ := mxj.AnyXmlIndent(xmlMap["beans"], "", "  ", "beans")
	r := gd.BytesToString(resultBytes)
	return r
}

func SavaJettyConf(ConfMap mxj.Map) {
	JettyFromPath := ConfMap["Jetty"].(map[string]interface{})["From"].(string)
	JettyToPath := ConfMap["Jetty"].(map[string]interface{})["To"].(map[string]interface{})

	for port, path := range JettyToPath {
		confJetty := confJetty(JettyFromPath, port)
		prf("write jettyconf to %s...\n", JettyFromPath)
		gd.StringToFilePath(path.(string), confJetty)
	}
}
