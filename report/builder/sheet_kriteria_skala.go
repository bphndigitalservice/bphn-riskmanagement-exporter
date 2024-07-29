package builder

import (
	"github.com/xuri/excelize/v2"
)

func (ex *ExcelBuilder) fillKriteriaDanSkala(f *excelize.File) {
	f.SetCellValue(SheetKriteriaDanSkala, "B10", "Powered by MR Report Generator.")
}
