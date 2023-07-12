package model

type CsvRecord struct {
	Company          string
	Person           string
	Name             string
	DeviceType       string
	MacAddress       string
	Registered       string // string in csv (Yes/No)
	Status           string
	UUIDCreationDate string
	DownloadDate     string
	HotDesking       string // string in csv (Yes/No)
	HotDeskingID     string
	HotDeskingPhone  string
	Location         string
	Group            string
	Comment          string
	Firmware         string
}
