const canvas = document.getElementById('hexgrid');
const ctx = canvas.getContext('2d');

// Make this realtime adaptive (need to add in rerender function?)
var winwidth = window.innerWidth;
var winheight = window.innerHeight;
ctx.canvas.width = winwidth-20;
ctx.canvas.height = winheight-20;
var r = Math.min((winwidth-30)/15/Math.sqrt(3), (winheight-30)/11.5/2);

// CENTER TO CORNER
var hexwidth = Math.sqrt(3) * r;
var hexheight = 2 * r;

const x_offset = -5;
const y_offset = 1;

class hexpoint {
	constructor(x, y) {
		this.x = x; // screen x
		this.y = y; // screen y
	}
}
class hexcoord { // AXIAL - PRIMARY COORDINATE SYSTEM
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
class hexcube { // CUBE
	constructor(v, u, t) {
		this.v = v; // same as axial
		this.u = u; // same as axial
		this.t = t; // should be -v-u
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
// CONVERSIONS (axial <-> offset / cube)
function cube_to_axial(cube) {
	var v = cube.v;
	var u = cube.u;
	return new hexcoord(v, u);
}
function axial_to_cube(hex) {
	var v = hex.v;
	var u = hex.u;
	var t = -hex.v-hex.u;
	return new hexcube(v, u, t);
}
function axial_to_offset(hex) {
	var col = hex.u + (hex.v + (hex.v&1)) / 2;
	var row = hex.v;
	return new hexoffset(col, row);
}
function offset_to_axial(hex) {
	var u = hex.col - (hex.row + (hex.row&1)) / 2;
	var v = hex.row;
	return new hexcoord(v, u);
}

// DIRECTIONS
var axial_unit_vectors = [
	new hexcoord(1, 0), new hexcoord(1, -1), new hexcoord(0, -1),
	new hexcoord(-1, 0), new hexcoord(-1, 1), new hexcoord(0, 1)
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

// DISTANCES
function cube_add(a, b) {
	return new hexcube(a.v + b.v, a.u + b.u, a.t + b.t);
}
function cube_subtract(a, b) {
	return new hexcube(a.v - b.v, a.u - b.u, a.t - b.t);
}
function cube_distance(a, b) {
	var vec = cube_subtract(a, b);
	return (Math.abs(vec.q) + abs(vec.r) + abs(vec.s)) / 2;
}
// formula for Euclidean distance: sqrt(dq^2 + dr^2 + dq*dr)
function axial_distance(a, b) {
	var ac = axial_to_cube(a);
	var bc = axial_to_cube(b);
	return cube_distance(ac, bc);
}

// WIP SECTIONS BELOW, ABOVE TO BE REVISED AT A LATER DATE

// LINE DRAWING
function lerp(a, b, t) {

}
function cube_lerp(a, b, t) {

}
function cube_linedraw(a, b) {

}

// MOVEMENT RANGE
// [COORDINATE RANGES]
// var results = [];
// for each -N <= q <= N {
// 	var s = -q-r;
// 	results.append(cube_add(center, Cube(q, r, s)));
//
// for each -N <= q <= N {
// 	for each max(-N, -q-N) <= r <= min(N, -q+N) {
//		results.append(axial_add(center, Hex(q, r)))
//	}
// }
//
// [INTERSECTING RANGES]
// for each q_min <= q <= q_max {
// 	for each max(r_min, -q-s_max) <= r <= min(r_max, -q-s_min) {
//		results.append(Hex(q, r));
// 	}
// }
//
// [OBSTACLES]
// function hex_reachable(start, movement) {
// 	var visited = set(); // set of hexes
// 	add start to visited
// 	var fringes = []; // array of arrays of hexes?
// 	fringes.append([start]);
// 	for each 1 < k <= movement {
// 		fringes,.append([]);
// 		for each hex in fringes[k-1] {
// 			var neighbor = hex_neighbor(hex, dir);
// 			if neighbor not in visited and not blocked {
// 				add neighbor to visited
// 				fringes[k].append(neighbor);
// 			}
// 		}
// 	return visited;
// }

// ROTATION - confirm direction of rotation
function axial_rotate_clockwise(hex, center) {
	var vec = cube_subtract(axial_to_cube(hex), axial_to_center(hex));
	var rot = new hexcube(-vec.u, -vec.v, -vec.t);
	rot = cube_add(rot, center);
	return cube_to_axial(rot);
}

// REFLECTION - make another if you want a line that's not at 0.
function axial_reflectT(hex) {
	return new hexcoord(hex.u, hex.v);
}
function axial_reflectU(hex) {
	return new hexcoord(-hex.u-hex.v, hex.u);
}
function axial_reflectV(hex) {
	return new hexcoord(hex.v, -hex.u-hex.v);
}
// To reflect over a line that's not at 0, pick a reference point on that line.
// Subtract the reference point, perform the reflection, then add the reference point back.
// RINGS
//

// FIELD OF VIEW

// HEX TO PIXEL

// PIXEL TO HEX

// ROUNDING TO NEAREST HEX

// MAP STORAGE

// WRAPAROUND MAPS

// PATHFINDING



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
