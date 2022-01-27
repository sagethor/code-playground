package world

// rewritten using "hexgrid positions & movement tables.pdf"
// not sure if v & u are strictly needed anymore, but the functions are still available.
func vuCoordinates(w uint8) (uint8, uint8) {
	return (w >> 4, w & 0x0f); // (v, u)
}
func vuCoordinate(v, u uint8) uint8 {
	return (v << 4) + u; // w 
}

// this will likely need rewrite with other sections once integrated with client
func vStepPlus(w uint8) uint8 {
	return w + 16;
}
func vStepMinus(w uint8) uint8 {
	return w - 16;
}
func uStepPlus(w uint8) uint8 {
	return w + 1;
}
func uStepMinus(w uint8) uint8 {
	return w - 1;
}
func vuStepPlus(w uint8) uint8 {
	return w + 15;
}
func vuStepMinus(w uint8) uint8 {
	return w - 15;
}

type boundType uint8;
const (
	onGrid boundType = 1 << iota;
	offGrid
	xPlus
	yPlus
	zPlus
	xMinus
	yMinus
	zMinus
)

func wBounds(w uint8) boundType {
	switch w {
	case 0, 1, 2, 3, 4, 5, 6,
	16, 17, 18, 19, 20, 21, 31
	32, 33, 36, 35, 36, 47,
	48, 49, 50, 51, 63,
	64, 65, 66, 79,
	80, 81, 95,
	96, 111,
	127,
	142, 143,
	157, 158, 159,
	172, 173, 174, 175,
	187, 188, 189, 190, 191,
	202, 203, 204, 205, 206, 207,
	217, 218, 219, 220, 221, 222, 223,
	232, 233, 234, 235, 236, 237, 238, 239,
	240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 253, 254, 255:
		return offGrid;
	case 7, 8, 9, 10, 11, 12, 13:
		return xPlus;
	case 225, 226, 227, 228, 229, 230, 231:
		return xMinus;
	case 126, 141, 156, 171, 186, 201, 216:
		return yPlus;
	case 22, 37, 52, 67, 82, 97, 112:
		return yMinus;
	case 128, 144, 160, 176, 192, 208, 224:
		return zPlus;
	case 14, 30, 46, 62, 78, 94, 110:
		return zMinus;
	case default:
		return onGrid;
	}
}
func vuBounds(v, u uint8) boundType {
	return wBounds(vuCoordinate(v, u));
}

// refactor/rename when process flow is decided upon
// ALSO TEST THIS IN EXCEL?
func zoneChange(w uint8) uint8 {
	eval := wBounds(w);
	switch eval {
	case xPlus:
		return w + 202;
	case xMinus:
		return w - 202;
	case yPlus:
		return w - 103;
	case yMinus:
		return w + 103;
	case zPlus:
		return w - 115;
	case zMinus:
		return w + 115;
	default:
		return w;
}
