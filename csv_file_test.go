package gofakeit

import (
	"fmt"
	"testing"
)

func TestCSVFile(t *testing.T) {
	Seed(11)

	err := CSVFile(&CSVFileOptions{
		&CSVOptions{
			RowCount: 3,
			Fields: []Field{
				{Name: "id", Function: "autoincrement"},
				{Name: "first_name", Function: "firstname"},
				{Name: "last_name", Function: "lastname"},
				{Name: "password", Function: "password", Params: map[string][]string{"special": {"false"}}},
			},
		},
		"./csv_file.csv",
	})
	if err != nil {
		fmt.Println(err)
	}
}

func TestCSVFileLookup(t *testing.T) {
	addFileCSVFileLookup()
	info := GetFuncLookup("csvfile")

	m := map[string][]string{
		"filepath": {"./csv_file_lookup.csv"},
		"rowcount": {"10"},
		"fields": {
			`{"name":"id","function":"autoincrement"}`,
			`{"name":"first_name","function":"firstname"}`,
			`{"name":"password","function":"password","params":{"special":["false"],"length":["20"]}}`,
		},
	}
	_, err := info.Call(&m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	// t.Fatal(fmt.Sprintf("%s", value.([]byte)))
}
