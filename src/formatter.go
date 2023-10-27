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

func FormatGrep(line string) string {
	fnameStyle := color.HEX("9da").Sprint
	lineStyle := color.HEX("fed").Sprint //color.New(color.FgHiGreen).SprintFunc()
	columnStyle := color.HEX("fed").Sprint //color.New(color.FgHiGreen).SprintFunc()
	dimStyle := color.HEX("777").Sprint
	parts := strings.SplitN(line, ":", 4)
	result := ""
	if len(parts) > 0 {
		result += fnameStyle(stripColor(parts[0]))
		result += dimStyle(":")
	}
	if len(parts) > 1 {
		result += lineStyle(stripColor(parts[1]))
		result += dimStyle(":")
	}
	if len(parts) > 2 {
		result += columnStyle(stripColor(parts[2]))
		result += dimStyle(":")
	}
	if len(parts) > 3 {
		result += parts[3]
		// colored, err := ansi_parser.Parse(parts[3], ansi_parser.WithIgnoreInvalidCodes(), ansi_parser.ParseOption{NonColorCodes: ansi_parser.Remove})
		// if err != nil {
		// 	result += "ERR:" + parts[3]
		// }
		// for _, t := range colored {
		// 	result += formatStyle(t)
		// }
		// result += parts[3]
	}
	return result

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
