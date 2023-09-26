package csvdb

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

type Database[T any] struct {
	data []T
}

func OpenCsvDatabase[T any](path string) (*Database[T], error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	header := lines[0]
	l := lines[1:]
	llen := len(l)
	data := make([]T, llen)
	for i := 0; i < llen; i++ {
		d, err := toStruct[T](header, l[i])
		if err != nil {
			return nil, err
		}
		data[i] = d
	}
	csvdb := &Database[T]{}
	return csvdb, nil
}

func toStruct[T any](header []string, line []string) (T, error) {
	var result T
	d := strings.Builder{}
	d.WriteString("{")
	for i, h := range header {
		if i > 0 {
			d.WriteString(",")
		}
		d.WriteString("\"" + h + "\"")
		d.WriteString(":")
		_, ok := strconv.Atoi(line[i])
		if ok != nil {
			d.WriteString("\"" + line[i] + "\"")
		} else {
			d.WriteString(line[i])
		}
	}
	d.WriteString("}")
	err := json.Unmarshal([]byte(d.String()), &result)
	return result, err
}
