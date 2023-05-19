package main

import (
	"fmt"
	"log"

	"github.com/antonioiubatti93/ophelos/interest"
	termstructuremodel "github.com/antonioiubatti93/ophelos/termstructure/model"
)

func main() {
	fmt.Println("Example with coupon-bearing bonds")

	ts, err := termstructuremodel.NewDiscrete(
		termstructuremodel.NewTenorValue(termstructuremodel.ON, 0.01),
		termstructuremodel.NewTenorValue(termstructuremodel.W1, 0.02),
		termstructuremodel.NewTenorValue(termstructuremodel.M1, 0.03),
		termstructuremodel.NewTenorValue(termstructuremodel.M3, 0.04),
		termstructuremodel.NewTenorValue(termstructuremodel.M6, 0.045),
		termstructuremodel.NewTenorValue(termstructuremodel.M12, 0.04),
	)
	if err != nil {
		log.Fatal(err)
	}

	const (
		fixed    = 0.02
		spread   = 0.005
		maturity = 1.0 / 3.0
		notional = 1.0
	)

	floater := interest.NewFloatingRate(ts, spread)

	bond := NewBond(notional, maturity,
		NewFixedCoupon(fixed, 1.0/12.0),
		NewFixedCoupon(fixed, 1.0/6.0),
		NewFloatingCoupon(floater, 1.0/4.0),
		NewFloatingCoupon(floater, 1.0/3.0),
	)

	rate := interest.NewRate(ts, interest.NewSimpleDiscount())

	npv := bond.Value(rate)
	fmt.Println("bond value:", npv)

	const shift = 0.0001 // 1 basis point

	shifted := rate.Shift(shift)

	duration := -(bond.Value(shifted) - npv) / (shift * npv)

	fmt.Println("duration [y]:", duration, "vs maturity:", maturity)
}
