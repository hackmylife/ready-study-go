package koan

func Tax(amount, percent int) int          { return amount * percent / 100 } // TODO: 内部関数を非公開にする
func InvoiceTotal(amount, percent int) int { return amount + Tax(amount, percent) }
