package cmd

import (
	"fmt"
	"os"

	utils "pw/utils"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	passwordOptions   []string
	cursor            int
	selected          []bool
	passwordLength    int
	completed         bool
	errorMsg          string
	generateButton    string
	generatedPassword string
}

func initialModel() model {
	return model{
		passwordOptions: []string{"Include Digits?", "Include Symbols?"},
		selected:        make([]bool, 2),
		passwordLength:  8,
		errorMsg:        "",
		generateButton:  "Generate Password",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.passwordOptions)+1 {
				m.cursor++
			}
		case "enter", " ":
			if m.cursor < len(m.passwordOptions) {
				m.selected[m.cursor] = !m.selected[m.cursor]
			} else if m.cursor == len(m.passwordOptions)+1 {
				includeDigits := m.selected[0]
				includeSymbols := m.selected[1]

				m.generatedPassword = utils.GeneratePassword(m.passwordLength, includeDigits, includeSymbols)
				utils.WriteToClipboard(m.generatedPassword)

				m.completed = true
				return m, tea.Quit
			}
		case "left":
			if m.cursor == len(m.passwordOptions) && m.passwordLength > 8 {
				m.passwordLength--
			}
		case "right":
			if m.cursor == len(m.passwordOptions) {
				m.passwordLength++
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.completed {
		return "\n📋 Your password has been copied to your clipboard!\n\n"
	}

	s := "\n"

	for i, choice := range m.passwordOptions {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if m.selected[i] {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	cursor := " "
	if m.cursor == len(m.passwordOptions) {
		cursor = ">"
	}
	s += fmt.Sprintf("%s Password Length: < %d >\n", cursor, m.passwordLength)

	cursor = " "
	if m.cursor == len(m.passwordOptions)+1 {
		cursor = ">"
	}
	s += fmt.Sprintf("\n%s [ %s ]\n", cursor, m.generateButton)

	if m.errorMsg != "" {
		s += "\n❌ " + m.errorMsg + "\n"
	}

	s += "\nPress space/enter to toggle options.\n"
	s += "Use left/right arrows to adjust password length.\n"
	s += "\nPress ctrl+c to quit.\n"
	return s
}

func runInteractiveMode() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
