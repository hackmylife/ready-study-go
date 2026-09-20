//go:build ignore

package koan

func tax(amount, percent int) int          { return amount * percent / 100 }
func InvoiceTotal(amount, percent int) int { return amount + tax(amount, percent) }
