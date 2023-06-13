package csv

type Device struct {
	Company          string `csv:"Company"`
	Person           string `csv:"Person"`
	Name             string `csv:"Name"`
	DeviceType       string `csv:"Device type"`
	MACAddress       string `csv:"MAC address"`
	Registered       string `csv:"Registered"`
	Status           string `csv:"Status"`
	UUIDCreationDate string `csv:"UUID creation date"`
	DownloadDate     string `csv:"Download date"`
	HotDesking       string `csv:"Hot desking"`
	HotDeskingID     string `csv:"Hot desking ID"`
	HotDeskingPhone  string `csv:"Hot desking phone"`
	Location         string `csv:"Location"`
	Group            string `csv:"Group"`
	Comment          string `csv:"Comment"`
	Firmware         string `csv:"Firmware"`
}
