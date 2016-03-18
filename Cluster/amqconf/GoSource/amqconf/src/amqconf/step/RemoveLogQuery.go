package step

import (
	gd "github.com/Mooxe000/GoDash"
)

func (step *Step) RemoveLogQuery() *Step {
	xmlMap := step.Value

	keyPath := xmlMap.PathForKeyShortest("-id")
	parentPath := ParentPath(keyPath)
	pathValues, _ := xmlMap.ValuesForPath(parentPath)

	var key int
	if gd.IsSlice(pathValues) {
		for i, pathValue := range pathValues {
			if pathValue.(map[string]interface{})["-id"] == "logQuery" {
				key = i
			}
		}
	}

	var gl gd.GdList
	gl.Value = pathValues
	Pgl := &gl
	Pgl.New()
	Pgl.Remove(key).Sync()

	xmlMap.SetValueForPath((&(gd.GdInterface{Pgl.Value})).Value, parentPath)

	return step
}
