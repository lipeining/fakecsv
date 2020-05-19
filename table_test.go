package fakecsv

import (
	"fmt"
	"testing"

	"github.com/lipeining/fakecsv/model"
	"github.com/stretchr/testify/assert"
)

func TestMakeColumnFuncFactory(t *testing.T) {
	cols := []model.Column{
		{"id", "int32", "1000000", "1", "", true, false, true},
		{"name", "string", "80", "0", "", true, false, false},
		{"card", "string", "80", "0", "", true, false, false},
		{"create_time", "datetime", "2026-01-02 15:04:05", "2006-01-02 15:04:05", "", true, false, false},
	}
	generator := MakeColumnFuncFactory(cols)
	line := generator(1)
	fmt.Println(line)
}
func TestWritetxt(t *testing.T) {
	cols := []model.Column{
		{"id", "int32", "1000000", "1", "", true, false, true},
		{"name", "string", "80", "0", "", true, false, false},
		{"card", "string", "80", "0", "", true, false, false},
		{"create_time", "datetime", "2026-01-02 15:04:05", "2006-01-02 15:04:05", "", true, false, false},
	}
	generator := MakeColumnFuncFactory(cols)
	dir := "C:/ProgramData/MySQL/MySQL Server 5.7/Uploads"
	basename := "card"
	err := Writetxt(dir, basename, 1, 10000, generator)
	assert.Equal(t, nil, err)
	// go Writetxt(basename, 1, 10000, generator)
	// go Writetxt(basename, 10001, 20000, generator)
}
func TestWriteCSV(t *testing.T) {
	cols := []model.Column{
		{"id", "int32", "1000000", "1", "", true, false, true},
		{"name", "string", "80", "0", "", true, false, false},
		{"card", "string", "80", "0", "", true, false, false},
		{"create_time", "datetime", "2026-01-02 15:04:05", "2006-01-02 15:04:05", "", true, false, false},
	}
	generator := MakeColumnFuncFactory(cols)
	dir := "C:/ProgramData/MySQL/MySQL Server 5.7/Uploads"
	basename := "card"
	err := WriteCSV(dir, basename, 1, 10000, generator)
	assert.Equal(t, nil, err)
	// go Writetxt(basename, 1, 10000, generator)
	// go Writetxt(basename, 10001, 20000, generator)
}
