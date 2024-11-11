package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	passwordOptions []string
	cursor          int
	selected        int
	inputMode       bool
	passwordLength  string
	completed       bool
	errorMsg        string
}

func initialModel() model {
	return model{
		passwordOptions: []string{"Easy to say", "Easy to read", "All characters"},
		selected:        -1,
		inputMode:       false,
		passwordLength:  "",
		errorMsg:        "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.inputMode {
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "enter":
				if len(m.passwordLength) > 0 {
					length := 0
					fmt.Sscanf(m.passwordLength, "%d", &length)
					if length < 8 {
						m.errorMsg = "Password length must be at least 8 characters"
						return m, nil
					}
					m.completed = true
				}
				return m, tea.Quit
			case "backspace":
				if len(m.passwordLength) > 0 {
					m.passwordLength = m.passwordLength[:len(m.passwordLength)-1]
				}
			default:
				if len(msg.String()) == 1 && msg.String() >= "0" && msg.String() <= "9" {
					m.passwordLength += msg.String()
				}
			}
		} else {
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "up":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down":
				if m.cursor < len(m.passwordOptions)-1 {
					m.cursor++
				}
			case "enter", " ":
				m.selected = m.cursor
				m.inputMode = true
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.completed {
		option := m.passwordOptions[m.selected]
		return fmt.Sprintf("\nSelected option: %s\nPassword length: %s\n\n", option, m.passwordLength)
	}

	if m.inputMode {
		s := "\n🔐 Enter desired password length (minimum 8 characters):\n\n"
		s += m.passwordLength
		if m.errorMsg != "" {
			s += "\n\n❌ " + m.errorMsg
		}
		s += "\n\nPress enter to confirm or ctrl+c to quit.\n"
		return s
	}

	s := "\n🔐 Effortlessly Generate Robust Passwords!\n\n"
	for i, choice := range m.passwordOptions {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if m.selected == i {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\nPress space or enter to select an option.\n"
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
