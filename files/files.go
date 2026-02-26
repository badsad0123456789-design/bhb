package files

import (
	"fmt"
	"okak/okak"
	"os"

	"github.com/fatih/color"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		okak.PrintError(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {

		okak.PrintError(err)
		return
	}
	color.Green("ураа запись жазылды")
}
