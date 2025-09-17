package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type commitItem struct {
	title string
	desc  string
}

func (ct *commitItem) Title() string       { return ct.title }
func (ct *commitItem) Description() string { return ct.desc }
func (ct *commitItem) FilterValue() string { return ct.title }

type logModel struct {
	list list.Model
}

func NewLogList(commits []string) tea.Model {
	items := make([]list.Item, len(commits))

	for i, c := range commits {
		parts := strings.SplitN(c, " ", 2)
		if len(parts) == 2 {
			items[i] = &commitItem{
				title: parts[0],
				desc:  parts[1],
			}
		}
	}

	l := list.New(items, list.NewDefaultDelegate(), 50, 14)

	return logModel{list: l}
}

func (m logModel) Init() tea.Cmd {
	return nil
}

func (m logModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m logModel) View() string {
	return lipgloss.NewStyle().Margin(1, 2).Render(m.list.View())
}
