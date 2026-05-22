package main

import (
	"fmt"
	"time"

	"github.com/3dp-cloud/companion/api"
	"github.com/3dp-cloud/companion/printer_controller"
	"github.com/3dp-cloud/companion/utils"
)

func main() {
	cfg := utils.LoadConfig()

	apiClient := api.Connect(cfg)
	defer apiClient.Disconnect(250)

	printers := apiClient.GetAssignedPrinters()

	for _, printer := range printers {
		go func(printer api.PrinterDeclaration) {
			for {
				err := printer_controller.Start(printer.Type, apiClient)
				if err != nil {
					fmt.Printf("Error from printer controller %s: %v\n", printer.ID, err)
				}
				time.Sleep(1 * time.Minute)
			}
		}(printer)
	}

	select {}
}
