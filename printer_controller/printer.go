package printer_controller

import (
	"fmt"
	"time"

	"github.com/3dp-cloud/companion/api"
)

type PrinterController interface {
	testConnection() error
	stopPrint() error
	startPrint(fileUrl string) error
	getStatus() (string, error)
}

func Start(printerType string, apiClient api.Client) error {
	controller, err := GetPrinterController(printerType, apiClient)
	if err != nil {
		return err
	}

	// Core loop

	time.Sleep(10 * time.Second)

	return controller.testConnection()
}

func GetPrinterController(printerType string, apiClient api.Client) (PrinterController, error) {
	switch printerType {
	case "generic_klipper":
		return NewGenericKlipperController(), nil
	default:
		return nil, fmt.Errorf("unsupported printer type: %s", printerType)
	}
}
