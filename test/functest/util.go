package functest

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/tableauio/tableau"
	"github.com/tableauio/tableau/format"
	"github.com/tableauio/tableau/internal/importer"
	"github.com/tableauio/tableau/log"
	"github.com/tableauio/tableau/options"
)

func genProto(t *testing.T) {
	err := tableau.GenProto(
		"protoconf",
		"./testdata",
		"./_proto",
		options.Input(
			&options.InputOption{
				Proto: &options.InputProtoOption{
					ProtoPaths: []string{"./_proto"},
					// ImportedProtoFiles: []string{
					// 	"common/cs_dbkeyword.proto",
					// 	"common/common.proto",
					// 	"common/time.proto",
					// },
					Formats: []format.Format{
						// format.Excel,
						format.CSV,
						format.XML,
					},
					Header: &options.HeaderOption{
						Namerow: 1,
						Typerow: 2,
						Noterow: 3,
						Datarow: 4,
					},
				},
			},
		),
		options.Output(
			&options.OutputOption{
				Proto: &options.OutputProtoOption{
					FilenameWithSubdirPrefix: false,
					FileOptions: map[string]string{
						"go_package": "github.com/tableauio/tableau/test/functest/protoconf",
					},
				},
			},
		),
		options.Log(
			&log.Options{
				Level: "DEBUG",
				Mode:  "FULL",
			},
		),
	)
	if err != nil {
		t.Errorf("%+v", err)
	}
}

func genConf(t *testing.T) {
	err := tableau.GenConf(
		"protoconf",
		"./testdata",
		"./_conf",
		options.Input(
			&options.InputOption{
				Conf: &options.InputConfOption{
					ProtoPaths: []string{"./_proto", "."},
					ProtoFiles: []string{"./_proto/*.proto"},
					Formats: []format.Format{
						// format.Excel,
						format.CSV,
						format.XML,
					},
				},
			},
		),
		options.Output(
			&options.OutputOption{
				Conf: &options.OutputConfOption{
					Pretty:          true,
					Formats:         []format.Format{format.JSON},
					EmitUnpopulated: true,
				},
			},
		),
	)
	if err != nil {
		t.Errorf("%+v", err)
	}
}

func rangeFilesByFormat(dir string, fmt format.Format, callback func(bookPath string) error) error {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	// book name -> existence(bool)
	csvBooks := map[string]bool{}
	for _, entry := range dirEntries {
		if entry.IsDir() {
			// scan and generate subdir recursively
			subdir := filepath.Join(dir, entry.Name())
			err = rangeFilesByFormat(subdir, fmt, callback)
			if err != nil {
				return err
			}
			continue
		}
		fileFmt := format.Ext2Format(filepath.Ext(entry.Name()))
		if fileFmt != fmt {
			continue
		}
		switch fmt {
		case format.Excel:
			bookPath := filepath.Join(dir, entry.Name())
			if err := callback(bookPath); err != nil {
				return err
			}
		case format.CSV:
			bookName, _, err := importer.ParseCSVFilenamePattern(entry.Name())
			if err != nil {
				return err
			}
			if _, ok := csvBooks[bookName]; ok {
				// NOTE: multiple CSV files construct the same book.
				continue
			}
			csvBooks[bookName] = true
			if err := callback(importer.GenCSVBookFilenamePattern(dir, bookName)); err != nil {
				return err
			}
		default:
			return errors.New("unknown fommat: " + string(fmt))
		}
	}
	return nil
}