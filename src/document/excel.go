package document

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"strconv"
)

type sheetDoc struct {
	xlsxFile *excelize.File
	sheetName string
	fileName string
	rows [][]string
}

var label [26]string

func NewExcel(fileName string, sheetName string) *sheetDoc {
	return new(sheetDoc).readIntoMemeory(fileName, sheetName)
}

func (sheetdoc *sheetDoc) readIntoMemeory(name string, sheet string) *sheetDoc {
	xlsx, err := excelize.OpenFile(name)
	if len(label[0]) == 0 {
		initLabel()
	}
	if err != nil {
		log.Println("File [" + name + "] doesn't exsit, new xlsx file created")
		sheetdoc.xlsxFile = excelize.NewFile()
		index := sheetdoc.xlsxFile.NewSheet(sheet)
		sheetdoc.xlsxFile.SetActiveSheet(index)
		sheetdoc.sheetName = sheet
		sheetdoc.fileName = name
	} else {
		sheetdoc.fileName = name
		sheetdoc.sheetName = sheet
		sheetdoc.xlsxFile = xlsx
		sheetdoc.rows, err= xlsx.GetRows(sheet)
	}
	return sheetdoc
}

func (sheetdoc *sheetDoc) GetRows() [][]string  {
	return sheetdoc.rows
}

func (sheetdoc *sheetDoc) GetRow(rowNum int) []string {
	return sheetdoc.rows[rowNum]
}

func (sheetdoc *sheetDoc) GetCellValue(row int, column int) string {
	return sheetdoc.rows[row][column]
}

func (sheetdoc *sheetDoc) GetCell(coordinate string) string {
	value, _ := sheetdoc.xlsxFile.GetCellValue(sheetdoc.sheetName, coordinate)
	return value
}

func (sheetdoc *sheetDoc) SetCellValue(row int, column int, value string) {
	sheetdoc.SetCell(convertCoordinate(row, column), value)
}

func (sheetdoc *sheetDoc) SetCell(coordinate string, value string) {
	style, _ := sheetdoc.xlsxFile.NewStyle(getDeafaultStyle())
	err1 := sheetdoc.xlsxFile.SetCellStyle(sheetdoc.sheetName, coordinate, coordinate, style)
	err2 := sheetdoc.xlsxFile.SetCellValue(sheetdoc.sheetName, coordinate, value)
	if err1 != nil {

	}
	if err2 != nil {

	}
}

func (sheetdoc *sheetDoc) SetRowValue(rowNum int, columnNum int, values *[]string) {
	sheetdoc.SetRowStyle(rowNum, columnNum, len(*values), getDeafaultStyle())
	err := sheetdoc.xlsxFile.SetSheetRow(sheetdoc.sheetName, convertCoordinate(rowNum, columnNum), values)
	if err!= nil {
		log.Fatal(err)
	}
}

func (sheetdoc *sheetDoc) Save()  {
	err := sheetdoc.xlsxFile.SaveAs(sheetdoc.fileName)
	if err != nil {
		log.Fatal(err)
	}
}

func getDeafaultStyle() string {
	return `{"font":{"bold":false,"italic":false,"family":"微软雅黑","size":11,"color":"#000000"},
			"alignment":{"horizontal":"center","vertical":"center","wrap_text":true}}`
}

func (sheetdoc *sheetDoc) SetRowStyle(rowNum int, columnNum int, length int, styleStr string)  {
	style, _ := sheetdoc.xlsxFile.NewStyle(styleStr)
	for i := columnNum; i < length; i++ {
		coordinate := convertCoordinate(rowNum, i)
		sheetdoc.xlsxFile.SetCellStyle(sheetdoc.sheetName, coordinate, coordinate, style)
	}
}

func convertCoordinate(row int, column int) string {
	return label[column] + strconv.Itoa(row + 1)
}

func initLabel()  {
	idx := 0
	for lab := 'A'; lab <= 'Z'; lab ++ {
		label[idx] = string(lab)
		idx ++
	}
}