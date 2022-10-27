package main

import logger "mylogger/component"

func main() {
	author := "Dc"
	logger.Debugf("hello,%s", author)
	logger.Infof("hello,%s", author)
	logger.Warningf("hello,%s", author)
	logger.Errorf("hello,%s", author)

}
