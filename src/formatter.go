package paf

import (
	"strings"
	"path"
	"github.com/charmbracelet/lipgloss"
	_ "github.com/junegunn/fzf/src/util"
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

func FormatGrep(line string) string {
	parts := strings.SplitN(line, ":", 3)
	if len(parts) > 2 {
		fname := parts[0]
		line := parts[1]
		text := parts[2]
		d, f := OptimizeFile(fname)
		return f + ":" + line + " " + d + " " + text
	}
	return line
}
