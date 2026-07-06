package payment

import "github.com/google/wire"

var PaymentSet = wire.NewSet(
	NewRepositoryPayment,
)
