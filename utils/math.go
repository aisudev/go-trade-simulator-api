package utils

func CalculateOpenPosition(last_price, last_amount, open_price, amount float64) float64 {
	return (open_price * last_amount / last_price) + amount
}

func CalculateClosePosition(last_price, last_amount, close_price, amount float64) float64 {
	return (close_price * last_amount / last_price) - amount
}

func ClosePosition(last_price, last_amount, close_price float64) float64 {
	return (close_price * last_amount / last_price)
}
