// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"game/db/dao/internal"
)

// internalSmsVerificationDao is internal type for wrapping internal DAO implements.
type internalSmsVerificationDao = *internal.SmsVerificationDao

// smsVerificationDao is the data access object for table u_sms_verification.
// You can define custom methods on it to extend its functionality as you wish.
type smsVerificationDao struct {
	internalSmsVerificationDao
}

var (
	// SmsVerification is globally public accessible object for table u_sms_verification operations.
	SmsVerification = smsVerificationDao{
		internal.NewSmsVerificationDao(),
	}
)

// Fill with you ideas below.
