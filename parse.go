// Package stringargv implements parse string cmd to argv list
package stringargv

import (
	"errors"
	"fmt"
)

const (
	singleQuote = '\''
	doubleQuote = '"'
	space       = ' '
)

// ErrUnClosedQuote represents a invalid cmd string with unclosed quote
var ErrUnClosedQuote = errors.New("invalid cmd string: unclosed quote")

// Parse cmd with space split return the []string
func Parse(cmd string) (argv []string, err error) {
	return ParseSplit(cmd, space)
}

// ParseSplit will parse cmd with custom split
func ParseSplit(cmd string, split byte) (argv []string, err error) {
	var (
		temp      []byte
		prevQuote byte
	)
	for i := 0; i < len(cmd); i++ {
		switch cmd[i] {
		case split:
			if prevQuote == 0 {
				if len(temp) != 0 {
					argv = append(argv, string(temp))
					temp = temp[:0]
				}
				continue // skip space
			}
		case singleQuote, doubleQuote:
			if prevQuote == 0 {
				if i == 0 || cmd[i-1] == split {
					prevQuote = cmd[i]
					continue
				}
			} else if cmd[i] == prevQuote {
				if i == len(cmd)-1 {
					if len(temp) != 0 {
						argv = append(argv, string(temp))
						temp = temp[:0]
					}
				} else if cmd[i+1] != split {
					argv = argv[:0]
					return nil, fmt.Errorf("invalid cmd string: %s", cmd)
				}
				prevQuote = 0
				if len(temp) != 0 {
					argv = append(argv, string(temp))
					temp = temp[:0]
				}
				continue
			}
		}
		temp = append(temp, cmd[i])
		if len(cmd)-1 == i {
			argv = append(argv, string(temp))
			temp = temp[:0]
		}
	}
	if prevQuote != 0 {
		err = ErrUnClosedQuote
		argv = argv[:0]
	}
	return
}
