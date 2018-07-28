package commands

import (
	"bk/utils"
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
	"os/exec"
	"io"
	"bytes"
	)

func Delete(c *cli.Context) error {
	show := exec.Command("bk", "show")
	peco := exec.Command("peco")
	r, w := io.Pipe()
	show.Stdout = w
	peco.Stdin = r
	var out bytes.Buffer
	peco.Stdout = &out

	show.Start()
	peco.Start()
	show.Wait()
	w.Close()
	peco.Wait()
	deletePath := strings.TrimRight(out.String(), "\n")

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
	historyFile.Close()
	texts = strings.TrimRight(texts, "\n")
	if len(texts) == 0 {
		exec.Command("cp", "/dev/null", historyFileName).Start()
	} else {
		historyFileW, e := os.OpenFile(historyFileName, os.O_WRONLY|os.O_TRUNC, 0600)
		if e != nil {
			return e
		}
		fmt.Fprintln(historyFileW, texts)
		historyFileW.Close()
	}
	return nil
}
