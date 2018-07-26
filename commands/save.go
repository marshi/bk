package commands

import (
	"os"
	"github.com/urfave/cli"
	"bk/utils"
	"fmt"
	"bufio"
)

func Save(c *cli.Context) error {
	historyFileName, e := utils.HistoryFile()
	historyFile, e := os.OpenFile(historyFileName, os.O_RDWR|os.O_APPEND, 0600)
	defer historyFile.Close()
	if e != nil {
		return e
	}
	currentDirName, e := os.Getwd()
	if e != nil {
		return e
	}
	var isDuplicate = false
	scanner := bufio.NewScanner(historyFile)
	for scanner.Scan() {
		if currentDirName == scanner.Text() {
			isDuplicate = true
		}
	}
	if isDuplicate {
		return nil
	}
	fmt.Fprintln(historyFile, currentDirName)
	return nil
}