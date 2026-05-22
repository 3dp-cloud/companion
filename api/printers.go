package api

type PrinterDeclaration struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Endpoint string `json:"endpoint"`
}

// Make http request to get assigned printers
func (c Client) GetAssignedPrinters() []PrinterDeclaration {
	// For now, return a hardcoded list of printers
	return []PrinterDeclaration{
		{ID: "printer1", Type: "generic_klipper", Endpoint: "http://printer1.local"},
		{ID: "printer2", Type: "generic_klipper", Endpoint: "http://printer2.local"},
	}
}
