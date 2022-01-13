const canvas = document.getElementById('hexgrid');
const ctx = canvas.getContext('2d');

// make radius programmatic
const r = 25;
const a = Math.PI / 3;


function init() {
	for (var v = 0; v < 15; v++) {
		for (var u = 0; u < 15; u++) {
			if (v + u > 6 && v + u < 22) {
				drawHex(v, u);
			}
		}
	}

}
init();

// swap sin and cos for flat top, this makes pointy tops.
function drawHex(v, u) { 
	ctx.beginPath();
	for (var i = 0; i < 6; i++) {
		ctx.lineTo(u * r * Math.sin(a) + v * r * Math.sin(a) * 2 + r * Math.sin(a * i) - 5*r, r + u * r * (1+Math.cos(a)) + r * Math.cos(a * i));
	}
	ctx.closePath();
	ctx.stroke();
}

