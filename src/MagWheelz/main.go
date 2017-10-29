package main

import (
	 "github.com/sirupsen/logrus"
	//"github.com/gopherjs/gopherjs/js"
)

//TODO: check go and gopherJS/ReactJS can co-exist

var LOGGER *logrus.Logger = logrus.New()

func main() {

	//js.Global.Call("alert", "Hello, JavaScript")
	//println("Hello, JS console")
	//js.Global.Get("document").Call("write", "Hello world!")
	LOGGER.Debug("Getting File Name..")
	var fileName string = GetFileName()
	var fileContents string = ReadMyFile(fileName)
	LOGGER.Info("Printing File Contents\n", fileContents)

}

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	LOGGER = logrus.StandardLogger()
}