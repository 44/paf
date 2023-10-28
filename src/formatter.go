package paf

import (
	"strings"
	"path"
	"github.com/charmbracelet/lipgloss"
	_ "github.com/junegunn/fzf/src/util"
	ansi_parser "github.com/44/go-ansi-parser"
	"strconv"
	// "github.com/fatih/color"
	"github.com/gookit/color"
	"fmt"
)

var (
	// styleFname = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("6"))
	styleFname = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	styleDir = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
)

func ShortenDir(dir string) string {
	parts := strings.Split(dir, "/")
	result := ""
	for _, part := range parts {
		if len(part) <= 2 {
			continue
		}
		result += string(part[0:2]) + "/"
	}
	return result
}

func OptimizeFile(fpath string) (string, string) {
	trimmed, _, _ := extractColor(fpath, nil, nil)
	//chars := util.ToChars([]byte(trimmed))
	//str := string(chars.Bytes())
	name := path.Base(trimmed)
	dir := ShortenDir(path.Dir(trimmed))
	return dir, name //styleDir.Render(dir), styleFname.Render(name)
}

func FormatFile(fname string) string {
	d, f := OptimizeFile(fname)
	fmt.Println("FormatFile", d, f)
	return f + " " + d
}

func formatStyle(seg *ansi_parser.StyledText) string {
	result := "["
	result += "style=" + strconv.Itoa(int(seg.Style))
	result += " mode=" + strconv.Itoa(int(seg.ColourMode))
	if seg.FgCol != nil {
		result += " fg=some"
	}
	if seg.BgCol != nil {
		result += " bg=some"
	}
	return result + "]{" + seg.Label + "}"
}

func toPlainText(text []*ansi_parser.StyledText) string {
	result := ""
	for _, t := range text {
		result += t.Label
	}
	return result
}

func stripColor(text string) string {
	result, err := ansi_parser.Parse(text, ansi_parser.WithIgnoreInvalidCodes(), ansi_parser.ParseOption{NonColorCodes: ansi_parser.Remove})
	if err != nil {
		return text
	}
	return toPlainText(result)
}

func formatFname(fname string) string {
	fnameStyle := color.HEX("9da").Sprint
	dimStyle := color.HEX("777").Sprint

	name := path.Base(fname)
	dir := path.Dir(fname) // ShortenDir(path.Dir(fname))
	return fnameStyle(name) + " " + dimStyle(dir)
}

func formatPrefix(filepath string, line string, column string) string {
	fnameStyle := color.HEX("9da").Sprint
	dirStyle := color.HEX("888").Sprint
	sepStyle := color.HEX("666").Sprint
	lineStyle := color.HEX("fed").Sprint //color.New(color.FgHiGreen).SprintFunc()
	columnStyle := color.HEX("fed").Sprint //color.New(color.FgHiGreen).SprintFunc()

	fname := path.Base(filepath)
	dir := path.Dir(filepath)

	return fnameStyle(fname) + sepStyle(":") + lineStyle(line) + sepStyle(":") + columnStyle(column) + " " + dirStyle(dir) + sepStyle(":")
}

func splitLine(line string) (string, string, string, string) {
	parts := strings.SplitN(line, ":", 4)
	if len(parts) > 3 {
		// vimgrep - file:line:column:text
		lnum := stripColor(parts[1])
		_, err := strconv.Atoi(lnum)
		if err == nil {
			cnum := stripColor(parts[2])
			_, err := strconv.Atoi(cnum)
			if err == nil {
				return stripColor(parts[0]), stripColor(parts[1]), stripColor(parts[2]), parts[3]
			}
			return stripColor(parts[0]), stripColor(parts[1]), "", strings.Join(parts[2:], "")
		}
	}

	if len(parts) > 2 {
		// grep - file:line:text
		lnum := stripColor(parts[1])
		_, err := strconv.Atoi(lnum)
		if err == nil {
			return stripColor(parts[0]), stripColor(parts[1]), "", strings.Join(parts[2:], "")
		}
	}

	if len(parts) > 1 {
		return stripColor(parts[0]), "", "", strings.Join(parts[1:], "")
	}
	return stripColor(line), "", "", ""
}

func formatLine(filepath string, lnum string, cnum string, text string) string {
	fnameStyle := color.HEX("9da").Sprint
	dirStyle := color.HEX("888").Sprint
	sepStyle := color.HEX("666").Sprint
	lineStyle := color.HEX("fed").Sprint //color.New(color.FgHiGreen).SprintFunc()
	columnStyle := color.HEX("fed").Sprint //color.New(color.FgHiGreen).SprintFunc()

	fname := path.Base(filepath)
	dir := path.Dir(filepath)

	result := fnameStyle(fname)
	if lnum != "" {
		result += sepStyle(":") + lineStyle(lnum)
	}
	if cnum != "" {
		result += sepStyle(":") + columnStyle(cnum)
	}
	result += " " + dirStyle(dir)
	if text != "" {
		result += sepStyle(":") + text
	}

	return result
}

func FormatGrep(line string) string {
	filepath, lnum, cnum, text := splitLine(line)
	return formatLine(filepath, lnum, cnum, text)

	// fnameStyle := color.HEX("9da").Sprint
	// lineStyle := color.HEX("fed").Sprint //color.New(color.FgHiGreen).SprintFunc()
	// columnStyle := color.HEX("fed").Sprint //color.New(color.FgHiGreen).SprintFunc()
	// dimStyle := color.HEX("777").Sprint
	// parts := strings.SplitN(line, ":", 4)
	//
	// filepath := ""
	// lnum := ""
	// column := ""
	// text := ""
	//
	// if len(parts) > 0 {
	// 	filepath = stripColor(parts[0])
	// }
	// if len(parts) > 1 {
	// 	lnum = stripColor(parts[1])
	// }
	// if len(parts) > 2 {
	// 	column = stripColor(parts[2])
	// }
	// if len(parts) > 3 {
	// 	text = parts[3]
	// }
	// return formatPrefix(filepath, lnum, column) + text

	// text, err := ansi_parser.Parse(line, ansi_parser.WithIgnoreInvalidCodes(), ansi_parser.ParseOption{NonColorCodes: ansi_parser.Remove})
	// result := ""
	// if err == nil {
	// 	for _, t := range text {
	// 		result += formatStyle(t)
	//
	// 	}
	// } else {
	// 	result = err.Error()
	// }
	// return result

	// parts := strings.SplitN(line, ":", 3)
	// if len(parts) > 2 {
	// 	fname := parts[0]
	// 	line := parts[1]
	// 	text := parts[2]
	// 	d, f := OptimizeFile(fname)
	// 	return f + ":" + line + " " + d + " " + text
	// }
	// return line
}
