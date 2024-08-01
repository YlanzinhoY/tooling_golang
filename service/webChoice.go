package service

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ylanzinhoy/sollievo/model"
)

func (s *CommandsStruct) WebChoice() {

	firstQuestion()

}

func firstQuestion() {

	types := []string{"Htmx", "react"}

	prompt := &survey.Select{
		Message: "First Choose",
		Options: types,
	}

	var frameworks string
	err := survey.AskOne(prompt, &frameworks, nil)

	cs := CommandsStruct{}

	switch frameworks {
	case "react":
		var pn string
		fmt.Println("Nome do projeto?")
		_, err := fmt.Scan(&pn)
		if err != nil {
			return
		}
		command := fmt.Sprintf("pnpm create vite@latest %s  --template react-ts", pn)
		err = cs.CommandRunnerNodeJS("react", command)
		if err != nil {
			log.Fatal(err)
		}

		acceptTailwind(&cs, pn)
		break
	}

	if err != nil {
		log.Fatal(err)
	}
}

func acceptTailwind(cs *CommandsStruct, path string) {
	var choice string
	fmt.Println("Deseja Tailwind? s/n")
	_, err := fmt.Scan(&choice)

	if err != nil {
		log.Fatal(err)
	}

	if choice == strings.ToUpper("s") || choice == strings.ToLower("s") {
		err = cs.CommandRunnerNodeJS("tailwind", fmt.Sprintf("cd %s && npm install -D tailwindcss postcss autoprefixer && npx tailwindcss init -p", path))
		if err != nil {
			log.Panic(err)
		}
	} else {
		os.Exit(1)
	}

}

func secondQuestion() {
	modeltools := model.Tools{}

	fm := model.FrameworkModel{}

	modeltools.Tools = fm.ToolsMap()

	res, err := modeltools.ToolsChoice()

	if err != nil {
		log.Panic(err)
		return
	}

	prompt := &survey.Select{
		Renderer: survey.Renderer{},
		Message:  fmt.Sprintf("qual framework voce quer? %s", res),
		Options:  res,
	}

	var choice string

	err = survey.AskOne(prompt, &choice, nil)

	if err != nil {
		log.Fatal(err)
	}

}

func processAwnser(awnsers []string, choice string) {

	length := len(awnsers) - 1

	for i := 0; i < length; i++ {
		if awnsers[i] == choice {
			// generateAwnser
			fmt.Printf("choice %s\n", choice)
			cs := CommandsStruct{}
			cs.CommandRunner(choice, "web")
		}
	}

}
