package constants

import (
	"cophee.team/project-manager/project"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var StatusMessageStyle = lipgloss.NewStyle().
						Foreground(lipgloss.AdaptiveColor{
							Light: "#04B575",
							Dark: "#04B575"}).Render

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

func NewItemDelegate(keys *KeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func (msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(project.Project); ok { 
			title = i.Name
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.Select):
				return m.NewStatusMessage(StatusMessageStyle("You chose " + title))
			case key.Matches(msg, keys.Delete):
				index := m.Index()
				m.RemoveItem(index)
				if len(m.Items()) == 0 {
					keys.Delete.SetEnabled(false)
				}
				return m.NewStatusMessage(StatusMessageStyle("Deleted " + title))
			}
		}
		return nil
	}

	help := []key.Binding{keys.Select, keys.Delete, keys.Quit, keys.Up, keys.Down}

	d.ShortHelpFunc = func () []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}