// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"game/db/dao/internal"
)

// internalReportWalletDayDao is internal implement for wrapping internal DAO implements.
type internalReportWalletDayDao = *internal.ReportWalletDayDao

// reportWalletDayDao is the data access object for table r_report_wallet_day.
// You can define custom methods on it to extend its functionality as you wish.
type reportWalletDayDao struct {
	internalReportWalletDayDao
}

var (
	// ReportWalletDay is globally public accessible object for table r_report_wallet_day operations.
	ReportWalletDay = reportWalletDayDao{
		internal.NewReportWalletDayDao(),
	}
)

// Fill with you ideas below.
