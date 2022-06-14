package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"cophee.team/project-manager/config"
	"cophee.team/project-manager/constants"
	"cophee.team/project-manager/project"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var Configuration *config.Config
var AppStyle = lipgloss.NewStyle().Padding(1, 2)
var TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

func main() {
	Configuration = config.ReadConfig()
	if len(Configuration.ImportDirs) == 0 {
		fmt.Println("No import directories configured")
		os.Exit(1)
	}

	if err := tea.NewProgram(newModel()).Start(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
}

type model struct {
	list list.Model
	keys *listKeyMap
	delegateKeys *constants.KeyMap
}

type listKeyMap struct {
	toggleSpinner    key.Binding
	toggleTitleBar   key.Binding
	toggleStatusBar  key.Binding
	togglePagination key.Binding
	toggleHelpMenu   key.Binding
	insertItem       key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		insertItem: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add item"),
		),
		toggleSpinner: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "toggle spinner"),
		),
		toggleTitleBar: key.NewBinding(
			key.WithKeys("T"),
			key.WithHelp("T", "toggle title"),
		),
		toggleStatusBar: key.NewBinding(
			key.WithKeys("S"),
			key.WithHelp("S", "toggle status"),
		),
		togglePagination: key.NewBinding(
			key.WithKeys("P"),
			key.WithHelp("P", "toggle pagination"),
		),
		toggleHelpMenu: key.NewBinding(
			key.WithKeys("H"),
			key.WithHelp("H", "toggle help"),
		),
	}
}


func getProjects() []project.Project {
	var projects []project.Project
	
	items, _ := ioutil.ReadDir(Configuration.ImportDirs[0])
	for _, item := range items {
		if item.IsDir() {
			projects = append(projects, project.Project{
				Name: item.Name(),
				Path: Configuration.ImportDirs[0] + "/" + item.Name(),
			})
		}
	}
	
	return projects
}

func newModel() model {

	var (
		delegateKeys = constants.DefaultKeyMap
		listKeys = newListKeyMap()
	)

	delegate := constants.NewItemDelegate(&delegateKeys)
	projects := getProjects()
	items := make([]list.Item, len(projects))
	for i, project := range projects {
		items[i] = project
	}
	projectList := list.New(items, delegate, 0, 0)
	projectList.Title = "Projects"
	projectList.Styles.Title = TitleStyle
	projectList.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.toggleSpinner,
			listKeys.insertItem,
			listKeys.toggleTitleBar,
			listKeys.toggleStatusBar,
			listKeys.togglePagination,
			listKeys.toggleHelpMenu,
		}
	}
	return model{
		list: projectList,
		keys: listKeys,
		delegateKeys: &delegateKeys,
	}
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) View() string {
	return AppStyle.Render(m.list.View())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := AppStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
			break
		}
		switch {
		case key.Matches(msg, m.keys.toggleSpinner):
			cmd := m.list.ToggleSpinner()
			return m, cmd
		case key.Matches(msg, m.keys.toggleTitleBar):
			v := !m.list.ShowTitle()
			m.list.SetShowTitle(v)
			m.list.SetShowFilter(v)
			m.list.SetFilteringEnabled(v)
			return m, nil
		case key.Matches(msg, m.keys.toggleStatusBar):
			m.list.SetShowStatusBar(!m.list.ShowStatusBar())
			return m, nil
		case key.Matches(msg, m.keys.togglePagination):
			m.list.SetShowPagination(!m.list.ShowPagination())
			return m, nil
		case key.Matches(msg, m.keys.toggleHelpMenu):
			m.list.SetShowHelp(!m.list.ShowHelp())
			return m, nil
		// TODO: add item
		/*
		case key.Matches(msg, m.keys.insertItem):
			m.delegateKeys.Delete.SetEnabled(true)
			newItem := m.itemGenerator.next()
			insCmd := m.list.InsertItem(0, newItem)
			statusCmd := m.list.NewStatusMessage(StatusMessageStyle("Added " + newItem.Title()))
			return m, tea.Batch(insCmd, statusCmd)
			*/
		}
	}

	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}