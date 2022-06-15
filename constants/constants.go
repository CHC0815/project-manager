package constants

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

var StatusMessageStyle = lipgloss.NewStyle().
						Foreground(lipgloss.AdaptiveColor{
							Light: "#04B575",
							Dark: "#04B575"}).Render

var AppStyle = lipgloss.NewStyle().Padding(1, 2)
var TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

type KeyMap struct {
	Up key.Binding
	Down key.Binding
	Quit key.Binding
	Select key.Binding
	Delete key.Binding
}

var DefaultKeyMap = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "move down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl-c"),
		key.WithHelp("q/ctrl-c", "quit"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter", "space"),
		key.WithHelp("enter/space", "select"),
	),
	Delete: key.NewBinding(
		key.WithKeys("backspace", "delete"),
		key.WithHelp("backspace/delete", "delete"),
	),
}

func (kmap KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		kmap.Up, 
		kmap.Down, 
		kmap.Quit, 
		kmap.Select, 
		kmap.Delete,
	}
}

func (kmap KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			kmap.Up,
			kmap.Down,
			kmap.Quit,
			kmap.Select,
			kmap.Delete,
		},
	}
}

type ListKeyMap struct {
	ToggleSpinner    key.Binding
	ToggleTitleBar   key.Binding
	ToggleStatusBar  key.Binding
	TogglePagination key.Binding
	ToggleHelpMenu   key.Binding
	InsertItem       key.Binding
}