package timecalc

import (
	"fmt"
	stc "strconv"
	"time"

	"../../domain"
	"github.com/lib/pq"
)

func DifftoNow(previous pq.NullTime) (res domain.AchieveInfo) {

	res.Today = false

	if !previous.Valid {
		res.Last = "達成日はありません"
		return
	}

	duration := time.Now().YearDay() - previous.Time.YearDay()

	if duration != 0 {
		res.Last = fmt.Sprintf("%s日前", stc.Itoa(duration))
		return
	}

	res.Last = "今日"
	res.Today = true

	return

}

func PickDate(datetime time.Time) (date string) {
	const layout = "2006-01-02"
	date = fmt.Sprintf(datetime.Format(layout))
	return
}
