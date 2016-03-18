package step

import (
	"strings"
)

func (step *Step) FixAmpersand() *Step {
	xmlMap := step.Value
	Values, _ := xmlMap.ValuesForKey("transportConnector")
	for _, v := range Values {
		url := v.(map[string]interface{})["-uri"]
		var s string = url.(string)
		s = strings.Replace(s, "&", "&amp;", len(s))
		v.(map[string]interface{})["-uri"] = s
	}

	return step
}
