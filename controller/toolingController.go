package controller

import (
	"fmt"
	"github.ylanzinhoy.tooling_golang/enums"
	"github.ylanzinhoy.tooling_golang/service"
	"net/http"
)

type ToolsServiceInterface interface {
	Exec(commandsStruct *service.CommandsStruct, value, goPackage string)
}

type ToolingControllerUpper struct{}

// region functions
func (t *ToolingControllerUpper) Exec(commandsStruct *service.CommandsStruct, value, goPackage string) {
	err := commandsStruct.CommandRunner(value, goPackage)
	if err != nil {
		fmt.Println(http.StatusBadRequest)
	}
	fmt.Printf("%d\n", http.StatusOK)
}

func (t *ToolingControllerUpper) ToolingController() {
	toolsChoices := make([]string, 0)

	tools := commandsStruct.Choices(modelTools.ToolsChoice(append(toolsChoices,
		enums.Gorm, enums.Viper, enums.Wire, enums.Default)), "tools")

	//region switch case
	for _, choice := range tools {
		switch choice {
		case enums.Gorm:
			t.Exec(&commandsStruct, enums.Gorm, enums.GormPackage)
		case enums.Viper:
			t.Exec(&commandsStruct, enums.Viper, enums.ViperPackage)
		case enums.Wire:
			t.Exec(&commandsStruct, enums.Wire, enums.WirePackage)
		case enums.Prometheus:
			t.Exec(&commandsStruct, enums.Prometheus, enums.PrometheusPackage)
		case enums.Default:
			fmt.Println(enums.Purple + "Saindo...")
			break
		}
	}
	//endregion
}

//endregion
