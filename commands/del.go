package commands

import (
	"bk/utils"
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func Delete(c *cli.Context) error {
	deletePath := c.String("path")
	historyFileName, e := utils.HistoryFile()
	historyFile, e := os.OpenFile(historyFileName, os.O_RDONLY, 0600)
	if e != nil {
		return e
	}
	var texts string
	scanner := bufio.NewScanner(historyFile)
	for scanner.Scan() {
		t := scanner.Text()
		if deletePath != t {
			texts = texts + t + "\n"
		}
	}
	texts = strings.TrimRight(texts, "\n")
	historyFile.Close()

	historyFileW, e := os.OpenFile(historyFileName, os.O_WRONLY|os.O_TRUNC, 0600)
	fmt.Fprintln(historyFileW, texts)
	historyFileW.Close()
	return nil
}
