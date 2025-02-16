package logic

import (
	"fmt"
	"time"
)

func LeaveRequest(joinDate, leaveDate time.Time, totalHoliday, leaveDuration float64) (bool, string) {
	startLeaveDate := joinDate.Add(180 * 24 * time.Hour)

	// if the employee does not meet the requirements to take personal leave (not yet 180 working days)
	if leaveDate.Before(startLeaveDate) {
		return false, "Karena belum 180 hari sejak tanggal join karyawan."
	}

	// calculate the remaining day from the date of leave (after 180 days) until the end of the year
	endOfYear := time.Date(leaveDate.Year(), time.December, 31, 23, 59, 59, 0, time.UTC)
	var remainingDays = int(endOfYear.Sub(startLeaveDate).Hours() / 24)

	// calculate the number of personal leave that can be taken (rounding down)
	var leaveAvailable = int(float64(remainingDays) / 365 * float64(totalHoliday))

	// check whether the duration of leave does not exceed the maximum personal leave
	if int(leaveDuration) >= leaveAvailable {
		return false, fmt.Sprintf("Karyawan hanya dapat mengambil %d hari cuti pribadi, durasi yang diminta terlalu banyak.", leaveAvailable)
	}

	// check if the leave duration is more than 3 days
	if leaveDuration > 3 {
		return false, "Durasi cuti pribadi maksimal 3 hari berturut-turut."
	}

	return true, "Cuti pribadi dapat diambil."
}
