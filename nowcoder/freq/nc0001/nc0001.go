package nc0001

import "math/big"

func solve(s string, t string) string {
	// write code here
	var (
		a = big.NewInt(0)
		b = big.NewInt(0)
		z = big.NewInt(0)
	)
	a.SetString(s, 10)
	b.SetString(t, 10)

	z.Add(a, b)

	return z.String()
}
