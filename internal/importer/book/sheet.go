package book

import (
	"bytes"
	"encoding/csv"
	"io"
	"math"

	"github.com/pkg/errors"
	"github.com/tableauio/tableau/internal/atom"
	"github.com/tableauio/tableau/proto/tableaupb"
	"github.com/xuri/excelize/v2"
	"google.golang.org/protobuf/proto"
)

// MetaSheetName defines the meta data of each worksheet.
const MetaSheetName = "@TABLEAU"

type SheetParser interface {
	Parse(protomsg proto.Message, sheet *Sheet) error
}

type Sheet struct {
	Name   string
	MaxRow int
	MaxCol int

	Rows [][]string // 2D array of string.

	Meta *tableaupb.SheetMeta
}

// NewSheet creats a new Sheet.
func NewSheet(name string, rows [][]string) *Sheet {
	maxRow := len(rows)
	maxCol := 0
	// MOTE: different row may have different length.
	// We need to find the max col.
	for _, row := range rows {
		n := len(row)
		if n > maxCol {
			maxCol = n
		}
	}
	return &Sheet{
		Name:   name,
		MaxRow: maxRow,
		MaxCol: maxCol,
		Rows:   rows,
	}
}

// ExtendSheet extends an existing Sheet.
func ExtendSheet(sheet *Sheet, rows [][]string) {
	maxRow := len(rows)
	maxCol := 0
	// MOTE: different row may have different length.
	// We need to find the max col.
	for _, row := range rows {
		n := len(row)
		if n > maxCol {
			maxCol = n
		}
	}
	sheet = &Sheet{
		Name:   sheet.Name,
		MaxRow: sheet.MaxRow + maxRow,
		MaxCol: int(math.Max(float64(sheet.MaxCol), float64(maxCol))),
		Rows:   append(sheet.Rows, rows...),
		Meta:   sheet.Meta,
	}
}

// Cell returns the cell at (row, col).
func (s *Sheet) Cell(row, col int) (string, error) {
	if row < 0 || row >= s.MaxRow {
		return "", errors.Errorf("row %d out of range", row)
	}
	if col < 0 || col >= s.MaxCol {
		return "", errors.Errorf("col %d out of range", col)
	}
	// MOTE: different row may have different length.
	if col >= len(s.Rows[row]) {
		return "", nil
	}
	return s.Rows[row][col], nil
}

// String returns the string representation (CSV) of the sheet.
func (s *Sheet) String() string {
	var buffer bytes.Buffer
	w := csv.NewWriter(&buffer)
	err := w.WriteAll(s.Rows) // calls Flush internally
	if err != nil {
		atom.Log.Panicf("write csv failed: %v", err)
	}
	return buffer.String()
}

func (s *Sheet) ExportCSV(writer io.Writer) error {
	w := csv.NewWriter(writer)
	// FIXME(wenchy): will be something wrong if we add the empty cell?
	// TODO: deepcopy a new rows!
	for nrow, row := range s.Rows {
		for i := len(row); i < s.MaxCol; i++ {
			// atom.Log.Debugf("add empty cell: %s", s.Name)
			row = append(row, "")
		}
		s.Rows[nrow] = row
	}
	// TODO: escape the cell value with `,` and `"`.
	return w.WriteAll(s.Rows) // calls Flush internally
}

func (s *Sheet) ExportExcel(file *excelize.File) error {
	file.NewSheet(s.Name)
	// TODO: clean up the sheet by using RemoveRow API.

	for nrow, row := range s.Rows {
		// file.SetRowHeight(s.Name, nrow, 20)
		for ncol, cell := range row {
			colname, err := excelize.ColumnNumberToName(ncol + 1)
			if err != nil {
				return errors.Wrapf(err, "failed to convert column number %d to name", ncol+1)
			}
			file.SetColWidth(s.Name, colname, colname, 20)

			axis, err := excelize.CoordinatesToCellName(ncol+1, nrow+1)
			if err != nil {
				return errors.Wrapf(err, "failed to convert coordinates (%d,%d) to cell name", ncol+1, nrow+1)
			}
			err = file.SetCellValue(s.Name, axis, cell)
			if err != nil {
				return errors.Wrapf(err, "failed to set cell value %s", axis)
			}
		}
	}
	return nil
}

func MetsasheetOptions() *tableaupb.WorksheetOptions {
	return &tableaupb.WorksheetOptions{
		Name:    MetaSheetName,
		Namerow: 1,
		Datarow: 2,
	}
}