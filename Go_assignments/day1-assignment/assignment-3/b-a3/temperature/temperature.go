package temperature

func FahrenheitToCelsius(ft int) float32 {
	t := (ft - 32) * 5 / 9
	return float32(t)
}
