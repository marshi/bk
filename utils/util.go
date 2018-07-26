package utils

import "github.com/mitchellh/go-homedir"

func HistoryFile() (string, error) {
	home, e := homedir.Dir()
	if e != nil {
		return "", e
	}
	historyFileName := home + "/.jd_history"
	return historyFileName, e
}
