package world

func wConvert(w uint8) (uint8, uint8) {
	return (w / 15, w % 15); // (v, u)
}
func vuConvert(v uint8, u uint8) uint8 {
	return v * 15 + u; // w
}

type boundType uint8;
const (
	onGrid boundType = iota
	offGrid
	xplusEdge
	yplusEdge
	zplusEdge
	xminusEdge
	yminusEdge
	zminusEdge
	xzCorner
	xyCorner
	yzCorner
	yxCorner
	zyCorner
	zxCorner
)
func wBounds(w uint8) boundType {
	return vuBounds(wConvert(w));
}
func vuBounds(v uint8, u uint8) boundType {
	switch {
	case v+u < 6, v+u > 22, v > 14, u > 14:
		return offGrid;
	case v == 14 && u == 0:
		return zxCorner;
	case v == 14 && u == 7:
		return yxCorner;
	case v == 7 && u == 14:
		return yzCorner;
	case v == 0 && u == 14:
		return xzCorner;
	case v == 0 && u == 7:
		return xyCorner;
	case v == 7 && u == 0:
		return zyCorner;
	case v == 0:
		return xplusEdge;
	case v == 14:
		return xminusEdge;
	case v+u == 21:
		return yplusEdge;
	case v+u == 7:
		return yminusEdge;
	case u == 0:
		return zplusEdge;
	case u == 14:
		return zminusEdge;
	default:
		return onGrid;
	}
}

// zone changes (x,y,z) have special return values [+] (0, 1, 2) & [-] (255, 254, 253):w
func vplusWalk(w uint8) uint8 {
	switch wBounds(w) {
	case xminusEdge, zxCorner, yxCorner:
		// advance one zone in the -x direction.
		return 255;
	case yplusEdge, yzCorner:
		// advance one zone in the +y direction.
		return 1;
	case offGrid:
		return w;
	default:
		return w + 15;
	}
}
func vminusWalk(w uint8) uint8 {
	switch wBounds(w) {
	case xplusEdge, xzCorner, xyCorner:
		// advance one zone in the +x direction.
		return 0;
	case yminusEdge, zyCorner:
		// advance one zone in the -y direction.
		return 254;
	case offGrid:
		return w;
	default:
		return w - 15;
	}
}
func uplusWalk(w uint8) uint8 {
	switch wBounds(w) {
	case yplusEdge, yxCorner, yzCorner:
		// advance one zone in the +y direction.
		return 1;
	case zminusEdge, xzCorner:
		// advance one zone in the -z direction.
		return 253;
	case offGrid:
		return w;
	default:
		return u + 1;
	}
}
func uminusWalk(w uint8) uint8 {
	switch wBounds(w) {
	case yminusEdge, xyCorner, zyCorner:
		// advance one zone in the -y direction.
		return 254;
	case zplusEdge, zxCorner:
		// advance one zone in the +z direction.
		return 2;
	case offGrid:
		return w;
	default:
		return w + 1;
	}
}
func vuplusWalk(w uint8) uint8 {
	swtich wBounds(w) {
	case zplusEdge, zyCorner, zxCorner:
		// advance one zone in the +z direction.
		return 2;
	case xminusEdge, yxCorner:
		// advance one zone in the -x direction.
		return 255;
	case offGrid:
		return w;
	default:
		return w + 14;
	}
}
func vuminusWalk(w uint8) uint8 {
	switch wBounds(w) {
	case zminusEdge, yzCorner, xzCorner:
		// advance one zone in the -z direction.
		return 253;
	case xplusEdge, xyCorner:
		// advance one zone in the +x direction.
		return 0;
	case offGrid:
		return w;
	default:
		return w - 14;
	}
}
