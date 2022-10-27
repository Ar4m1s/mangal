package icon

import (
	"github.com/metafates/mangal/color"
	"github.com/metafates/mangal/style"
)

type Icon int

const (
	Lua Icon = iota + 1
	Go
	Fail
	Success
	Question
	Mark
	Downloaded
	Progress
)

var icons = map[Icon]*iconDef{
	Lua: {
		emoji:   "🌙",
		nerd:    style.Fg(color.Blue)("\uE620"),
		plain:   style.Fg(color.Blue)("Lua"),
		kaomoji: style.Fg(color.Blue)("(=^･ω･^=)"),
		squares: style.Fg(color.Blue)("◧"),
	},
	Go: {
		emoji:   "🐹",
		nerd:    style.Fg(color.Cyan)("\uE627"),
		plain:   style.Fg(color.Cyan)("Go"),
		kaomoji: style.Fg(color.Cyan)("ʕ •ᴥ• ʔ"),
		squares: style.Fg(color.Cyan)("◨"),
	},
	Fail: {
		emoji:   "💀",
		nerd:    style.Fg(color.Red)("ﮊ"),
		plain:   style.Fg(color.Red)("X"),
		kaomoji: style.Fg(color.Red)("┐('～`;)┌"),
		squares: style.Fg(color.Red)("▨"),
	},
	Success: {
		emoji:   "🎉",
		nerd:    style.Fg(color.Green)("\uF65F "),
		plain:   style.Fg(color.Green)("✓"),
		kaomoji: style.Fg(color.Green)("(ᵔ◡ᵔ)"),
		squares: style.Fg(color.Green)("▣"),
	},
	Mark: {
		emoji:   "🦐",
		nerd:    style.Fg(color.Green)("\uF6D9"),
		plain:   style.New().Bold(true).Foreground(color.Green).Render("*"),
		kaomoji: style.New().Bold(true).Foreground(color.Red).Render("炎"),
		squares: style.New().Bold(true).Foreground(color.Green).Render("■"),
	},
	Question: {
		emoji:   "🤨",
		nerd:    style.Fg(color.Yellow)("\uF128"),
		plain:   style.Fg(color.Yellow)("?"),
		kaomoji: style.Fg(color.Yellow)("(￢ ￢)"),
		squares: style.Fg(color.Yellow)("◲"),
	},
	Progress: {
		emoji:   "👾",
		nerd:    style.Fg(color.Blue)("\uF0ED "),
		plain:   style.Fg(color.Blue)("@"),
		kaomoji: style.Fg(color.Blue)("┌( >_<)┘"),
		squares: style.Fg(color.Blue)("◫"),
	},
	Downloaded: {
		emoji:   "📦",
		nerd:    style.Fg(color.Cyan)("\uF0C5 "),
		plain:   style.Fg(color.Cyan)("⬇"),
		kaomoji: style.Fg(color.Cyan)("⊂(◉‿◉)つ"),
		squares: style.Fg(color.Cyan)("◬"),
	},
}
