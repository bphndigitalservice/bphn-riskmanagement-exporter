package builder

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"time"
)

func (ex *ExcelBuilder) signPlaceholder(f *excelize.File, sheet string, rowNumberStart int, col string) {
	dateCell := fmt.Sprintf("%s%d", col, rowNumberStart)

	now := time.Now()
	timeValue := now.Format("02 January 2006")
	// Date cell
	f.SetCellValue(sheet, dateCell, fmt.Sprintf("Jakarta , %s", timeValue))
	// Signer Role cell
	signerRoleCell := fmt.Sprintf("%s%d", col, rowNumberStart+1)
	role := os.Getenv("SIGN_ROLE")
	f.SetCellValue(sheet, signerRoleCell, role)
	// Signer Name cell
	name := os.Getenv("SIGN_NAME")
	signerNameCell := fmt.Sprintf("%s%d", col, rowNumberStart+6)
	f.SetCellValue(sheet, signerNameCell, name)
	// Employee ID cell
	signerIDCell := fmt.Sprintf("%s%d", col, rowNumberStart+7)
	f.SetCellValue(sheet, signerIDCell, fmt.Sprintf("NIP.%s", os.Getenv("SIGN_NIP")))
}
