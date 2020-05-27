package gofakeit

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

// CSVFileOptions defines values needed for csv generation
type CSVFileOptions struct {
	*CSVOptions
	FilePath string `json:"file_path" xml:"file_path"`
}

// CSVFile generates an object or an array of objects in json format
func CSVFile(cfo *CSVFileOptions) error {
	// Check delimiter
	if cfo.Delimiter == "" {
		cfo.Delimiter = ","
	}
	if strings.ToLower(cfo.Delimiter) == "tab" {
		cfo.Delimiter = "\t"
	}
	if cfo.Delimiter != "," && cfo.Delimiter != "\t" {
		return errors.New("Invalid delimiter type")
	}

	// Check fields
	if cfo.Fields == nil || len(cfo.Fields) <= 0 {
		return errors.New("Must pass fields in order to build json object(s)")
	}

	// Make sure you set a row count
	if cfo.RowCount <= 0 {
		return errors.New("Must have row count")
	}

	outputFile, outputError := os.OpenFile(cfo.FilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Println("An error occurred with file opening or creation\n", cfo.FilePath, outputError)
		return outputError
	}
	defer outputFile.Close()
	w := csv.NewWriter(outputFile)
	w.Comma = []rune(cfo.Delimiter)[0]

	// Add header row
	header := make([]string, len(cfo.Fields))
	for i, field := range cfo.Fields {
		header[i] = field.Name
	}
	w.Write(header)

	// Loop through row count and add fields
	for i := 1; i < int(cfo.RowCount); i++ {
		vr := make([]string, len(cfo.Fields))

		// Loop through fields and add to them to map[string]interface{}
		for ii, field := range cfo.Fields {
			if field.Function == "autoincrement" {
				vr[ii] = fmt.Sprintf("%d", i)
				continue
			}

			// Get function info
			funcInfo := GetFuncLookup(field.Function)
			if funcInfo == nil {
				return errors.New("Invalid function, " + field.Function + " does not exist")
			}

			value, err := funcInfo.Call(&field.Params, funcInfo)
			if err != nil {
				return err
			}

			vr[ii] = fmt.Sprintf("%v", value)
		}

		w.Write(vr)
	}

	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}

func addFileCSVFileLookup() {
	AddFuncLookup("csvfile", Info{
		Display:     "CSV file",
		Category:    "csv file",
		Description: "Generates array of rows in csv format",
		Example: `
			id,first_name,last_name,password
			1,Markus,Moen,Dc0VYXjkWABx
			2,Osborne,Hilll,XPJ9OVNbs5lm
		`,
		Output: "[]byte",
		Params: []Param{
			{Field: "rowcount", Display: "Row Count", Type: "int", Default: "100", Description: "Number of rows in JSON array"},
			{Field: "fields", Display: "Fields", Type: "[]Field", Description: "Fields containing key name and function to run in json format"},
			{Field: "delimiter", Display: "Delimiter", Type: "string", Default: ",", Description: "Separator in between row values"},
			{Field: "filepath", Display: "FilePath", Type: "string", Default: "./", Description: "file path"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			co := CSVOptions{}

			rowcount, err := info.GetInt(m, "rowcount")
			if err != nil {
				return "", err
			}
			co.RowCount = rowcount

			fieldsStr, err := info.GetStringArray(m, "fields")
			if err != nil {
				return "", err
			}

			// Check to make sure fields has length
			if len(fieldsStr) > 0 {
				co.Fields = make([]Field, len(fieldsStr))

				for i, f := range fieldsStr {
					// Unmarshal fields string into fields array
					err = json.Unmarshal([]byte(f), &co.Fields[i])
					if err != nil {
						return "", errors.New("Unable to decode json string")
					}
				}
			}

			delimiter, err := info.GetString(m, "delimiter")
			if err != nil {
				return "", err
			}
			co.Delimiter = delimiter

			filePath, err := info.GetString(m, "filepath")
			if err != nil {
				return "", err
			}
			cfo := CSVFileOptions{&co, filePath}
			err = CSVFile(&cfo)
			if err != nil {
				return "", err
			}
			return "", nil
		},
	})
}
