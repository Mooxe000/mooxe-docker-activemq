package step

func (step *Step) FixColonKey() *Step {
	xmlMap := step.Value

	xmlMap.RenameKey(xmlMap.PathForKeyShortest("-xsi"), "-xmlns:xsi")
	xmlMap.RenameKey(xmlMap.PathForKeyShortest("-schemaLocation"), "-xsi:schemaLocation")

	return step
}
