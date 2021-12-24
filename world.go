package main

// may need to adjust all below for modifying struct field

func wConvert(w uint8) (uint8, uint8) {
	return (w / 15, w % 15); // (v, u)
}
func vuConvert(v uint8, u uint8) uint8 {
	return v * 15 + u; // w
}

// re-evaluate for xy, yz, zx... do we just include those? in need of large rewrite...
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
)
func wBounds(w uint8) boundType {
	return vuBounds(wConvert(w));
}
func vuBounds(v uint8, u uint8) boundType {
	switch {
	case v+u < 6, v+u > 22:
		return offGrid;
	case v == 0 && u != 14 && u != 0:
		return xplusEdge;
	case v == 14 && u != 14 && u != 0:
		return xminusEdge;
	case v+u == 21 && v != 14 && v != 7:
		return yplusEdge;
	case v+u == 7 && v!= 7 && v != 0:
		return yminusEdge;
	case v != 7 && v != 14 && u == 0:
		return zplusEdge;
	case v!= 7 && v != 14 && u == 14:
		return zminusEdge;
	default:
		return onGrid;
	}
}

// rewrite for boundType usage
func vWalk(fwd bool, w uint8) uint8 {
	switch wBounds(w) {
	case offGrid:
		// TO-DO: define behavior for locations off-grid.
	case xplusEdge:
		if fwd {
			return w + 15;
		} else {
			// TO-DO: define moving one zone over in the x-direction
		}
	case xminusEdge:
		if fwd {
			// TO-DO: return
	case default:
		if fwd {
			return w + 15;
		} else {
			return w - 15;
		}
	}
}
func uWalk(fwd bool, w uint8) uint8 {
	var prime = w;
	if fwd {
		prime++;
	} else {
		prime--;
	}
	if wBounds(prime) {
		return prime;
	} else {
		return w;
	}
}
func vuWalk(fwd bool, w uint8) uint8 {
	var prime = w;
	if fwd {
		prime += 14;
	} else {
		prime -= 14;
	}
	if wBounds(prime) {
		return prime;
	} else {
		return w;
	}
}
