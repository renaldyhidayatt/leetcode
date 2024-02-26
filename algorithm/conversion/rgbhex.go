package conversion

func HEXToRGB(hex uint) (red, green, blue byte) {

	blue = byte(hex & 0xFF)
	green = byte((hex >> 8) & 0xFF)
	red = byte((hex >> 16) & 0xFF)
	return
}

func RGBToHex(red, green, blue byte) uint {
	return (uint(red) << 16) | (uint(green) << 8) | uint(blue)
}
