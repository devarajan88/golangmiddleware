package csv

import (
	"encoding/csv"
	"errors"
	"os"
	"reflect"
)

// CSV represents the implementation of the CSV contract.
type CSV struct {
	filename string
	records  []CSVRow
}

// CSVRow represents a single row in the CSV file.
type CSVRow struct {
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

type RowIterator struct {
	csvFile  *CSV
	position int
}

func NewCSV(filename string) (*CSV, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var csvRows []CSVRow
	for _, record := range records {
		csvRow := CSVRow{}
		err := csvRow.Unmarshal(record)
		if err != nil {
			// Handle error
		}
		csvRows = append(csvRows, csvRow)
	}

	return &CSV{
		filename: filename,
		records:  csvRows,
	}, nil
}

func (c *CSV) RowIterator(pos int) RowIterator {
	if pos < 0 || pos >= len(c.records) {
		pos = 0
	}

	return RowIterator{
		csvFile:  c,
		position: pos,
	}
}

func (c *CSV) Incorporate(other *CSV) {
	columns := c.getColumnNames()
	otherColumns := other.getColumnNames()

	for _, colName := range otherColumns {
		if !contains(columns, colName) {
			columns = append(columns, colName)
		}
	}

	for _, otherRow := range other.records {
		row := CSVRow{}
		rowValues := make([]string, 0, len(columns))

		for _, colName := range columns {
			colValue := reflect.ValueOf(otherRow).FieldByName(colName).String()
			rowValues = append(rowValues, colValue)
		}

		err := row.Unmarshal(rowValues)
		if err != nil {
			// Handle error
		}

		c.records = append(c.records, row)
	}
}

func (c *CSV) ToStringRFC4180() string {
	output := ""
	writer := csv.NewWriter(os.Stdout)

	for _, row := range c.records {
		rowValues := row.Marshal()
		writer.Write(rowValues)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		// Handle error
	}

	return output
}

func (r *RowIterator) Get() *CSVRow {
	if r.position >= len(r.csvFile.records) {
		return nil
	}

	return &r.csvFile.records[r.position]
}

func (r *RowIterator) Next() bool {
	r.position++
	return r.position < len(r.csvFile.records)
}

func (r *RowIterator) Pos() int {
	return r.position
}

func (c *CSV) getColumnNames() []string {
	if len(c.records) == 0 {
		return nil
	}

	csvRow := c.records[0]
	columns := make([]string, 0, len(csvRow))

	t := reflect.TypeOf(csvRow)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("csv")
		columns = append(columns, tag)
	}

	return columns
}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func (r *CSVRow) Unmarshal(values []string) error {
	if len(values) != reflect.TypeOf(*r).NumField() {
		return errors.New("csv: number of values does not match number of fields")
	}

	for i := 0; i < len(values); i++ {
		field := reflect.ValueOf(r).Elem().Field(i)
		if field.IsValid() && field.CanSet() {
			field.SetString(values[i])
		}
	}

	return nil
}

-func (r *CSVRow) Marshal() []string {
	t := reflect.TypeOf(*r)
	values := make([]string, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := reflect.ValueOf(*r).Field(i)
		values[i] = field.String()
	}

	return values
}
