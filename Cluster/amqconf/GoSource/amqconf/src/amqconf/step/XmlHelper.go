package step

import (
	gd "github.com/Mooxe000/GoDash"
	"strings"
)

func ParentPath(path string) string {
	var gl gd.GdList
	gi := gd.GdInterface{strings.Split(path, ".")}
	gl.Value = (&gi).ToInterfaces()
	Pgl := &gl

	Pgl.New()
	Pgl.Pop()
	Pgl.Sync()

	ss := (&(gd.GdInterfaces{Pgl.Value})).ToStrings()
	r := strings.Join(ss, ".")

	return r
}

/*
 * 此处 有坑
 * mxj 的 path 系统 并不遵循 xpath
 * slice 是无序的，每次都不一样，无法定位到
 */
/*
func PathForKeyValue(xmlMap *mxj.Map, key string, value string) string {
  paths := xmlMap.PathsForKey(key)

  rPath := ""
  var ri, rii int
  for i, path := range paths {
    parentPath := ParentPath(path)
    vs, _ := xmlMap.ValuesForPath(parentPath)
    for ii, v := range vs {
      if v.(map[string]interface{})["-value"] == value {
        ri = i
        rii = ii
        rPath = parentPath
        break
      }
    }
  }

  pln(ri)
  pln(rii)
  rPath = rPath +
    "[" + gd.IntToString(ri) +  "]" +
    "[" + gd.IntToString(rii) + "]"

  pln(rPath)
  rValue, _ := xmlMap.ValuesForPath(rPath)
  dd(typeof(rValue))
  // dd(rValue)

  return ""
}
*/
