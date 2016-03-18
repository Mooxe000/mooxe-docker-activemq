package step

func (step *Step) SetJvmHeap(heapValue string) *Step {
	xmlMap := step.Value

	keyPath := xmlMap.PathForKeyShortest("-percentOfJvmHeap")
	xmlMap.SetValueForPath(heapValue, keyPath)

	return step
}
