package entry

import (
	"cophee.team/project-manager/config"
	"cophee.team/project-manager/constants"
	tea "github.com/charmbracelet/bubbletea"
)

type Entry struct {
	Name string
	Path string
}

func (e Entry) FilterValue() string {
	return e.Name
}

func (e Entry) Title() string {
	return e.Name
}

func (e Entry) Description () string {
	return e.Path
}

type EntryModel struct {
	Config *config.ProjectConfig
}

func NewEntryModel(path string) EntryModel {
	return EntryModel{
		Config: config.ReadProjectConfig(path),
	}
}

func (m EntryModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m EntryModel) View() string {
	s := m.Config.Title
	s += "\n" + m.Config.Author
	return constants.AppStyle.Render(s)
}

func (m EntryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	return m, tea.Batch(cmds...)
}