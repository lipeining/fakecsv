package fakecsv

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/lipeining/fakecsv/model"
)

func init() {
	gofakeit.Seed(0)
}

// LOWPRORITY use by load data
const LOWPRORITY bool = false

// CONCURRENT use by load data
const CONCURRENT bool = false

// LOCAL use by load data
const LOCAL bool = false

// Chareset use by load data
const Chareset string = "utf8mb4"

// FieldsTerminatedBy use by load data
const FieldsTerminatedBy string = ","

// FieldsEnclosedBy use by load data
const FieldsEnclosedBy string = ""

// FieldsEscapedBy use by load data
const FieldsEscapedBy string = "\\"

// LinesTerminatedBy use by load data
const LinesTerminatedBy string = "\n"

// const LinesStartingBy string = '\n'

// PathExists use os.stat to check path
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// ParseJSONColumn 通过 json 文件来解析 column
func ParseJSONColumn(filePath string) ([]model.Column, error) {
	content, err := ioutil.ReadFile(filePath)
	var cols []model.Column
	if err != nil {
		fmt.Println("read file error:", err)
		return nil, err
	}
	err = json.Unmarshal(content, &cols)
	if err != nil {
		fmt.Println("unmarshal error:", err)
		return nil, err
	}
	fmt.Println(filePath, " cols length: ", len(cols))
	return cols, nil
}

func randIntRange(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn(max-min+1) + min
}
func randInt32Range(min, max int32) int32 {
	if min == max {
		return min
	}
	return rand.Int31n(max-min+1) + min
}
func randInt64Range(min, max int64) int64 {
	if min == max {
		return min
	}
	return rand.Int63n(max-min+1) + min
}

// MakeColumnFuncFactory 通用生成一行数据的回调函数
func MakeColumnFuncFactory(cols []model.Column) func(int) []string {
	// 对于每一个 column 添加 fake 字段，用于直接调用 gofakeit 的函数即可。
	// 如何结合对应的 Param Field 等参数进行完美切合 gofakeit 的功能？
	return func(current int) []string {
		insertCols := make([]string, 0)
		for _, column := range cols {
			if column.Type == "string" {
				// 使用 Word 会出新 let's 这种数据，需要格外小心
				var max int = 30
				var min int = 0
				if fix, err := strconv.Atoi(column.Max); err == nil {
					max = int(fix)
				}
				if fix, err := strconv.Atoi(column.Min); err != nil {
					min = int(fix)
				}
				if min > max {
					min, max = max, min
				}
				want := randIntRange(min, max)
				insertCols = append(insertCols, gofakeit.Sentence(want))
			} else if column.Type == "datetime" {
				layout := "2006-01-02 15:04:05"
				var max time.Time
				var min time.Time
				var err error
				if max, err = time.Parse(layout, column.Max); err != nil {
					max = time.Date(2030, time.January, 1, 0, 0, 0, 0, time.UTC)
				}
				if min, err = time.Parse(layout, column.Min); err != nil {
					min = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
				}
				if !min.Before(max) {
					min, max = max, min
				}
				d := gofakeit.DateRange(min, max).Format(layout)
				insertCols = append(insertCols, d)
			} else if column.Type == "int64" {
				if column.Autoincr {
					insertCols = append(insertCols, strconv.FormatInt(int64(current), 10))
				} else {
					var max int64
					var min int64
					var err error
					if max, err = strconv.ParseInt(column.Max, 10, 64); err != nil {
						max = 1000
					}
					if min, err = strconv.ParseInt(column.Min, 10, 64); err != nil {
						min = 0
					}
					if min > max {
						min, max = max, min
					}
					want := randInt64Range(min, max)
					insertCols = append(insertCols, strconv.FormatInt(int64(want), 10))
				}

			} else if column.Type == "int32" {
				if column.Autoincr {
					insertCols = append(insertCols, strconv.FormatInt(int64(current), 10))
				} else {
					var max int32 = 1000
					var min int32 = 0
					if fix, err := strconv.ParseInt(column.Max, 10, 32); err != nil {
						max = int32(fix)
					}
					if fix, err := strconv.ParseInt(column.Min, 10, 32); err != nil {
						min = int32(fix)
					}
					if min > max {
						min, max = max, min
					}
					want := randInt32Range(min, max)
					insertCols = append(insertCols, strconv.FormatInt(int64(want), 10))
				}
			} else if column.Type == "int" {
				if column.Autoincr {
					insertCols = append(insertCols, strconv.FormatInt(int64(current), 10))
				} else {
					var max int = 1000
					var min int = 0
					if fix, err := strconv.ParseInt(column.Max, 10, 0); err != nil {
						max = int(fix)
					}
					if fix, err := strconv.ParseInt(column.Min, 10, 0); err != nil {
						min = int(fix)
					}
					if min > max {
						min, max = max, min
					}
					want := randIntRange(min, max)
					insertCols = append(insertCols, strconv.FormatInt(int64(want), 10))
				}
			}
		}
		return insertCols
	}
}

// Writetxt write 1,000,000
func Writetxt(dir, basename string, start, end int, makeColumn func(int) []string) error {
	fileName := basename + "_" + strconv.Itoa(end) + ".txt"
	filePath := filepath.Join(dir, fileName)
	exists, err := PathExists(filePath)
	if err != nil {
		fmt.Println("An error stat with file \n", filePath, err)
		return err
	}
	if exists {
		return nil
	}
	outputFile, outputError := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Println("An error occurred with file opening or creation\n", filePath, outputError)
		return outputError
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	// 每次写 10000 行
	every := 10000
	num := end - start + 1
	total := num / every
	if num%every != 0 {
		total++
	}
	current := start
	for i := 0; i < total; i++ {
		for j := 0; j < every && current <= end; j++ {
			current++
			outputString := strings.Join(makeColumn(current), FieldsTerminatedBy)
			if current != 1 {
				outputString = LinesTerminatedBy + outputString
			}
			outputWriter.WriteString(outputString)
		}
		outputWriter.Flush()
	}
	return nil
}

// WriteCSV write 1,000,000
func WriteCSV(dir, basename string, start, end int, makeColumn func(int) []string) error {
	fileName := basename + "_" + strconv.Itoa(end) + ".csv"
	filePath := filepath.Join(dir, fileName)
	exists, err := PathExists(filePath)
	if err != nil {
		fmt.Println("An error stat with file \n", filePath, err)
		return err
	}
	if exists {
		return nil
	}
	outputFile, outputError := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Println("An error occurred with file opening or creation\n", filePath, outputError)
		return outputError
	}
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	// wirter.UseCRLF = false
	// 每次写 10000 行
	every := 10000
	num := end - start + 1
	total := num / every
	if num%every != 0 {
		total++
	}
	current := start
	for i := 0; i < total; i++ {
		for j := 0; j < every && current <= end; j++ {
			record := makeColumn(current)
			writer.Write(record)
			current++
		}
		writer.Flush()
	}
	return writer.Error()
}
