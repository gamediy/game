// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"game/db/dao/internal"
)

// internalPaymentDao is internal type for wrapping internal DAO implements.
type internalPaymentDao = *internal.PaymentDao

// paymentDao is the data access object for table p_payment.
// You can define custom methods on it to extend its functionality as you wish.
type paymentDao struct {
	internalPaymentDao
}

var (
	// Payment is globally public accessible object for table p_payment operations.
	Payment = paymentDao{
		internal.NewPaymentDao(),
	}
)

// Fill with you ideas below.
