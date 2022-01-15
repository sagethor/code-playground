const canvas = document.getElementById('hexgrid');
const ctx = canvas.getContext('2d');

// make radius programmatic
const r = 25;

const hexwidth = Math.sqrt(3) * r;
const hexheight = 2 * r;

const x_offset = -5;
const y_offset = 1;

class hexpoint {
	constructor(x, y) {
		this.x = x; // screen x
		this.y = y; // screen y
	}
}
class hexcoord { // AXIAL
	constructor(v, u) {
		this.v = v; // rows
		this.u = u; // columns
	}
}
class hexoffset { // OFFSET
	constructor(col, row) {
		this.col = col;
		this.row = row;
	}
}

function hexcorner(center, size, i) {
	var angle_deg = 60 * i - 30;
	var angle_rad = Math.PI / 180 * angle_deg;
	return new hexpoint(center.x + size * Math.cos(angle_rad),
		center.y + size * Math.sin(angle_rad));
}

function hexdraw(center, size) {
	ctx.beginPath();
	for (var i = 0; i < 6; i++) {
		corner = hexcorner(center, size, i);
		ctx.lineTo(corner.x, corner.y);
	}
	ctx.closePath();
	ctx.stroke();
}
// not 100% sure if this is neccessary.
function axial_to_offset(hex) {
	var col = hex.u + (hex.v + (hex.v&1)) / 2;
	var row = hex.v;
	return new hexoffset(col, row);
}
function offset_to_axial(hex) {
	var u = hex.col - (hex.row + hex.row&1)) / 2;
	var v = hex.row;
	return new hexcoord(v, u);
}

// again, not 100% sure on all of this, but I guess this boilerplate can't hurt
var axial_unit_vectors = [
	hexcoord(1, 0), hexcoord(1, -1), hexcoord(0, -1),
	hexcoord(-1, 0), hexcoord(-1, 1), hexcoord(0, 1)
];
function axial_direction(direction) {
	return axial_direction_vectors[direction];
}
function axial_add(hex, vec) {
	return Hex(hex.v + vec.v, hex.u + vec.u);
}
function axial_neighbor(hex, direction) {
	return axial_add(hex, axial_direction(direction));
}

// stopping here - https://www.redblobgames.com/grids/hexagons/#neighbors-offset


function init() {
	for (var v = 0; v < 15; v++) {
		for (var u = 0; u < 15; u++) {
			if (v + u > 6 && v + u < 22) {
				center = new hexpoint(u * hexwidth * .5 + v * hexwidth + r * x_offset, u * hexheight * .75 + r * y_offset);
				hexdraw(center, r);
			}
		}
	}

}
init();
