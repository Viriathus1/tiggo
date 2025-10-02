package tui

import (
	"fmt"
	"time"

	"github.com/Viriathus1/tiggo/internal/gitclient"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	hashStyle = lipgloss.NewStyle().Bold(true)

	selectedStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#444444")).
			Foreground(lipgloss.Color("#FFFFFF"))

	normalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#808080"))
)

type commitItem struct {
	title string
	desc  string
}

func (ct *commitItem) Title() string       { return hashStyle.Render(ct.title) }
func (ct *commitItem) Description() string { return ct.desc }
func (ct *commitItem) FilterValue() string { return ct.desc }

type logModel struct {
	gitClient *gitclient.GitClient
	list      list.Model
}

func NewLogList(gitClient *gitclient.GitClient) tea.Model {
	commits, _ := gitClient.GetCommitHistory()
	items := make([]list.Item, len(commits))

	for i, commit := range commits {
		itemMetaDesc := fmt.Sprintf("Meta:\tcommited by %s on %s\n",
			commit.Author.Name,
			commit.Author.When.Format(time.DateOnly),
		)
		itemCommitMessage := fmt.Sprintf("Message: %s\n", commit.Message)
		items[i] = &commitItem{
			title: commit.Hash.String(),
			desc:  itemMetaDesc + itemCommitMessage,
		}
	}

	d := list.NewDefaultDelegate()
	d.SetHeight(3)
	d.Styles.NormalTitle = normalStyle
	d.Styles.NormalDesc = normalStyle
	d.Styles.SelectedTitle = selectedStyle
	d.Styles.SelectedDesc = selectedStyle

	l := list.New(items, d, 0, 0)
	l.Title = "Git Log"
	l.Styles.Title = selectedStyle.PaddingLeft(2).PaddingRight(2)
	l.SetShowStatusBar(false)

	return logModel{
		gitClient: gitClient,
		list:      l,
	}
}

func (m logModel) Init() tea.Cmd {
	return nil
}

func (m logModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m logModel) View() string {
	return m.list.View()
}
