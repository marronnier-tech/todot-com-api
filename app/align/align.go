package align

import (
	"fmt"

	"gorm.io/gorm"
)

func ListOrder(tx *gorm.DB, table string, joined bool, param string) (ordered *gorm.DB) {
	if joined {
		ordered = tx.Order(fmt.Sprintf("%s.%s desc", table, param))
		return
	}

	ordered = tx.Order(fmt.Sprintf("%s desc", param))
	return

}
