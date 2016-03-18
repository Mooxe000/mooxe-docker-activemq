package main

import (
	. "amqconf/step"
	gd "github.com/Mooxe000/GoDash"
	"github.com/clbanning/mxj"
	"github.com/codegangsta/cli"
	"os"
)

var (
	dd     = gd.Dd
	pln    = gd.Pln
	prf    = gd.Prf
	typeof = gd.TypeOf
)

func confAmqConf(confPath string, confName string) map[string]interface{}{
	prf("read amqconf from %s ...\n", confPath)
	jsonStr := gd.StrFromFilePath(confPath)
	jsonMap, _ := mxj.NewMapJson([]byte(jsonStr))

	ConfMap := jsonMap[confName].
		([]interface {})[0].
		(map[string]interface{})

	return ConfMap
}

func main() {
	app := cli.NewApp()
	app.Name = "amqconf"
	app.Usage = "config activemq"
	app.Flags = []cli.Flag {
	  cli.StringFlag{
	    Name: "conf",
	    Value: "../../conf/amqconf.json",
	    Usage: "config path",
	  },
	}
	app.Action = func(c *cli.Context) {
		confPath := c.String("conf")
		if len(c.Args()) == 0 || typeof(c.Args()[0]) != "string" {
			pln("Please provide a strig as config name.")
			return
		}
		ConfMap := confAmqConf(confPath, c.Args()[0])
		SavaAmqConf(ConfMap)
	}

	app.Run(os.Args)
}
