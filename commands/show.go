package commands

import (
	"bk/utils"
	"os"
	"bufio"
	"fmt"
	"github.com/urfave/cli"
)

func Show(c *cli.Context) error {
	historyFileName, e := utils.HistoryFile()
	historyFile, e := os.OpenFile(historyFileName, os.O_RDONLY, 0400)
	defer historyFile.Close()
	if e != nil {
		return e
	}
	scanner := bufio.NewScanner(historyFile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil
}
