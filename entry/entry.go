package entry

import (
	"fmt"

	"cophee.team/project-manager/config"
	"cophee.team/project-manager/constants"
	"cophee.team/project-manager/styles"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type BackMsg bool

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
	box := fmt.Sprintf("%v %v by ", m.Config.Title, m.Config.Version)
	for _, author := range m.Config.Authors {
		box += fmt.Sprintf("%v ", author)
	}
	s := styles.EntryBox.Render(box)
	s += "\n"
	for _, desc := range m.Config.Desc {
		s += fmt.Sprintf("%v\n", desc)
	}
	for _, lang := range m.Config.Languages {
		s += fmt.Sprintf("%v ", lang)
	}
	return styles.AppStyle.Render(s)
}

func (m EntryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.DefaultKeyMap.Back):
			return m, func() tea.Msg {
				return BackMsg(true)
			}
		}
	}

	return m, tea.Batch(cmds...)
}