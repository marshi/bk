package commands

import (
	"bk/utils"
		"github.com/urfave/cli"
	"os"
		"fmt"
	"io/ioutil"
	"strings"
)

func Show(c *cli.Context) error {
	historyFileName, e := utils.HistoryFile()
	historyFile, e := os.OpenFile(historyFileName, os.O_RDONLY, 0400)
	defer historyFile.Close()
	if e != nil {
		return e
	}
	bytes, e := ioutil.ReadAll(historyFile)
	if e != nil {
		return e
	}
	texts := string(bytes)
	texts = strings.TrimRight(texts, "\n")
	fmt.Println(texts)
	return nil
}
