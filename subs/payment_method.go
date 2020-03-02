package subs

type PaymentMethod string

const (
	CREDIT_CARD PaymentMethod = "CREDIT_CARD"
	BOLETO      PaymentMethod = "BOLETO"
	ALL         PaymentMethod = "ALL"
)
