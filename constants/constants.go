package constants

import (
	"github.com/charmbracelet/bubbles/key"
)


type KeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Quit   key.Binding
	Select key.Binding
	Delete key.Binding
	Back   key.Binding
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
	Back: key.NewBinding(
		key.WithKeys("b", "esc"),
		key.WithHelp("b/esc", "back"),
	),
}

func (kmap KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		kmap.Up, 
		kmap.Down, 
		kmap.Quit, 
		kmap.Select, 
		kmap.Delete,
		kmap.Back,
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
			kmap.Back,
		},
	}
}

type ListKeyMap struct {
	ToggleTitleBar   key.Binding
	ToggleStatusBar  key.Binding
	TogglePagination key.Binding
	ToggleHelpMenu   key.Binding
	InsertItem       key.Binding
}