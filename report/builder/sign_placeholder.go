package builder

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func (ex *ExcelBuilder) signPlaceholder(f *excelize.File, sheet string, rowNumberStart int) {
	dateCell := fmt.Sprintf("E%d", rowNumberStart)

	// Date cell
	f.SetCellValue(sheet, dateCell, "Jakarta, <Date here>")
	// Signer Role cell
	signerRoleCell := fmt.Sprintf("E%d", rowNumberStart+1)
	f.SetCellValue(sheet, signerRoleCell, "Kepala Badan Pembinaan Hukum Nasional")
	// Signer Name cell
	signerNameCell := fmt.Sprintf("E%d", rowNumberStart+6)
	f.SetCellValue(sheet, signerNameCell, "<Name here>")
	// Employee ID cell
	signerIDCell := fmt.Sprintf("E%d", rowNumberStart+7)
	f.SetCellValue(sheet, signerIDCell, "NIP.<NIP here>")
}
