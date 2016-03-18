package step

func (step *Step) RemoveKahadb() *Step {
	xmlMap := step.Value

	paPath := xmlMap.PathForKeyShortest("persistenceAdapter")
	xmlMap.SetValueForPath("", paPath)

	return step
}
