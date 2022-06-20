package tui

import (
	"cophee.team/project-manager/entry"
	"cophee.team/project-manager/project"
	tea "github.com/charmbracelet/bubbletea"
)

var Prog *tea.Program

type sessionState int

const (
	projectView sessionState = iota
	entryView
)

type MainModel struct {
	state sessionState
	project tea.Model
	entry tea.Model
	windowSize tea.WindowSizeMsg
	path string
}

func NewMainModel() MainModel {
	return MainModel{
		state: projectView,
		project: project.NewProjectModel(),
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowSize = msg
	case project.SelectMsg:
		m.path = msg.Path
		m.state = entryView
	case entry.BackMsg:
		m.state = projectView
	}

	switch m.state {
	case projectView:
		newProject, newCmd := m.project.Update(msg)
		projectModel, ok := newProject.(project.ProjectModel)
		if !ok {
			panic("could not perform assertion on project model")
		}
		m.project = projectModel
		cmd = newCmd
	case entryView:
		m.entry = entry.NewEntryModel(m.path)
		newEntry, newCmd := m.entry.Update(msg)
		entryModel, ok := newEntry.(entry.EntryModel)
		if !ok {
			panic("could not perform assertion on entry model")
		}
		m.entry = entryModel
		cmd = newCmd
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	switch m.state {
	case entryView:
		return m.entry.View()
	default:
		return m.project.View()
	}
}
