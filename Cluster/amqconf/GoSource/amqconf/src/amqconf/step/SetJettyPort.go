package step

func (step *Step) SetJettyPort(portValue string) *Step {
	xmlMap := step.Value

	paths := xmlMap.PathsForKey("-value")

	for _, path := range paths {
		parentPath := ParentPath(path)
		vs, _ := xmlMap.ValuesForPath(parentPath)
		for i, v := range vs {
			if v.(map[string]interface{})["-value"] == "8161" {
				vs[i].(map[string]interface{})["-value"] = portValue
				break
			}
		}
	}

	return step
}
