









var PASS = true;


var u1 = g.Uint(1), u2 = g.Uint(2);

var u = g.Uint(1);
var u_ = g.Uint(1);
var u8 = g.Uint8(8);
var u16 = g.Uint16(16);
var u32 = g.Uint32(32);

var i = g.Int(1);
var i_ = g.Int(1);
var i8 = g.Int8(8);
var i16 = g.Int16(16);
var i32 = g.Int32(32);

var f32 = g.Float32(3.2);
var f32_ = g.Float32(3.2);

var b = g.Byte(8);
var b_ = g.Byte(8);

var r = g.Rune(32);
var r_ = g.Rune(32);







function value() {
	var pass = true;

	if (u !== 1 || u_ !== 1 || u8 !== 8 || u16 !== 16 || u32 !== 32) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: uint<br>");
		pass = false, PASS = false;
	}
	if (i !== 1 || i_ !== 1 || i8 !== 8 || i16 !== 16 || i32 !== 32) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: int<br>");
		pass = false, PASS = false;
	}
	if (f32 !== 3.2 || f32_ !== 3.2) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: float<br>");
		pass = false, PASS = false;
	}
	if (b !== 8 || b_ !== 8) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: byte<br>");
		pass = false, PASS = false;
	}
	if (r !== 32 || r_ !== 32) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: rune<br>");
		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function add() {
	var pass = true;

	if (u + 1 !== 2 || u_ + 1 !== 2 || u8 + 1 !== 9 || u16 + 1 !== 17 || u32 + 1 !== 33) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: uint<br>");
		pass = false, PASS = false;
	}
	if (i + 1 !== 2 || i_ + 1 !== 2 || i8 + 1 !== 9 || i16 + 1 !== 17 || i32 + 1 !== 33) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: int<br>");
		pass = false, PASS = false;
	}
	if (f32 + 1 !== 4.2 || f32_ + 1 !== 4.2) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: float<br>");
		pass = false, PASS = false;
	}
	if (b + 1 !== 9 || b_ + 1 !== 9) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: byte<br>");
		pass = false, PASS = false;
	}
	if (r + 1 !== 33 || r_ + 1 !== 33) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: rune<br>");
		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function subtract() {
	var pass = true;

	if (u - 1 !== 0 || u_ - 1 !== 0 || u8 - 1 !== 7 || u16 - 1 !== 15 || u32 - 1 !== 31) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: uint<br>");
		pass = false, PASS = false;
	}
	if (i - 1 !== 0 || i_ - 1 !== 0 || i8 - 1 !== 7 || i16 - 1 !== 15 || i32 - 1 !== 31) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: int<br>");
		pass = false, PASS = false;
	}
	if (f32 - 1 !== 2.2 || f32_ - 1 !== 2.2) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: float<br>");
		pass = false, PASS = false;
	}
	if (b - 1 !== 7 || b_ - 1 !== 7) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: byte<br>");
		pass = false, PASS = false;
	}
	if (r - 1 !== 31 || r_ - 1 !== 31) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: rune<br>");
		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function multiplication() {
	var pass = true;

	if (u * 2 !== 2 || u_ * 2 !== 2 || u8 * 2 !== 16 || u16 * 2 !== 32 || u32 * 2 !== 64) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: uint<br>");
		pass = false, PASS = false;
	}
	if (i * 2 !== 2 || i_ * 2 !== 2 || i8 * 2 !== 16 || i16 * 2 !== 32 || i32 * 2 !== 64) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: int<br>");
		pass = false, PASS = false;
	}
	if (f32 * 2 !== 6.4 || f32_ * 2 !== 6.4) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: float<br>");
		pass = false, PASS = false;
	}
	if (b * 2 !== 16 || b_ * 2 !== 16) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: byte<br>");
		pass = false, PASS = false;
	}
	if (r * 2 !== 64 || r_ * 2 !== 64) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: rune<br>");
		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function division() {
	var pass = true;

	if (u / 1 !== 1 || u_ / 1 !== 1 || u8 / 2 !== 4 || u16 / 2 !== 8 || u32 / 2 !== 16) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: uint<br>");
		pass = false, PASS = false;
	}
	if (i / 1 !== 1 || i_ / 1 !== 1 || i8 / 2 !== 4 || i16 / 2 !== 8 || i32 / 2 !== 16) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: int<br>");
		pass = false, PASS = false;
	}
	if (f32 / 2 !== 1.6 || f32_ / 2 !== 1.6) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: float<br>");
		pass = false, PASS = false;
	}
	if (b / 2 !== 4 || b_ / 2 !== 4) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: byte<br>");
		pass = false, PASS = false;
	}
	if (r / 2 !== 16 || r_ / 2 !== 16) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: rune<br>");
		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function main() {
	document.write("<br><br>== Numeric<br><br>");

	document.write("=== RUN value<br>");
	value();
	document.write("=== RUN add<br>");
	add();
	document.write("=== RUN subtract<br>");
	subtract();
	document.write("=== RUN multiplication<br>");
	multiplication();
	document.write("=== RUN division<br>");
	division();

	if (PASS) {
		document.write("PASS<br>");
	} else {
		document.write("FAIL<br>");
		alert("Fail: Numeric");
	}
} main();


var x = 14, y = 9;

var and = x & y;
var or = x | y;
var xor = x ^ y;
var not = !true;

var lShift = 9 << 2;
var rShift = 9 >> 2;
var lShiftNeg = -9 << 2;
var rShiftNeg = -9 >> 2;
/* Generated by GoScript (github.com/kless/goscript) */
