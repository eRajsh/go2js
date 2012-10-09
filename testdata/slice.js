/* Generated by GoScript (github.com/kless/goscript) */









var PASS = true;

function builtIn() {
	var pass = true;

	var s1 = g.NilSlice();
	var s2 = g.Slice();
	var s3 = g.MkSlice(0, 0);

	var _ = function(msg, in_, out) { return {
		msg:msg,
		in_:in_,
		out:out
	};}; tests = [
		_("nil s1", s1.isNil, true),
		_("nil s2", s2.isNil, false),
		_("nil s3", s3.isNil, false),
		_("len s1", s1.len === 0, true),
		_("len s2", s2.len === 0, true),
		_("len s3", s3.len === 0, true),
		_("cap s1", s1.cap === 0, true),
		_("cap s2", s2.cap === 0, true),
		_("cap s3", s3.cap === 0, true)
	];

	var t; for (var _ in tests) { t = tests[_];
		if (JSON.stringify(t.in_) !== JSON.stringify(t.out)) {
			document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: " + t.msg + " => got " + t.in_ + ", want " + t.out + "<br>");
			pass = false, PASS = false;
		}
	}
	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function shortHand() {
	var pass = true;

	var array = g.MkArray([10], 0, ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j']);
	var a_slice = g.NilSlice(), b_slice = g.NilSlice();



	a_slice.set(array, 4, 8);
	if (a_slice.toString() === "efgh" && a_slice.len === 4 && a_slice.cap === 6) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. [4:8] => got " + a_slice.v + ", len=" + a_slice.len + ", cap=" + a_slice.cap + "<br>");

		pass = false, PASS = false;
	}

	a_slice.set(array, 6, 7);
	if (a_slice.toString() !== "g") {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. [6:7] => got " + a_slice.v + "<br>");
		pass = false, PASS = false;
	}

	a_slice.set(array, 0, 3);
	if (a_slice.toString() === "abc" && a_slice.len === 3 && a_slice.cap === 10) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. [:3] => got " + a_slice.v + ", len=" + a_slice.len + ", cap=" + a_slice.cap + "<br>");

		pass = false, PASS = false;
	}

	a_slice.set(array, 5);
	if (a_slice.toString() !== "fghij") {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. [5:] => got " + a_slice.v + "<br>");
		pass = false, PASS = false;
	}

	a_slice.set(array, 0);
	if (a_slice.toString() !== "abcdefghij") {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. [:] => got " + a_slice.v + "<br>");
		pass = false, PASS = false;
	}

	a_slice.set(array, 3, 7);
	if (a_slice.toString() === "defg" && a_slice.len === 4 && a_slice.cap === 7) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. [3:7] => got " + a_slice.v + ", len=" + a_slice.len + ", cap=" + a_slice.cap + "<br>");

		pass = false, PASS = false;
	}



	b_slice.set(a_slice, 1, 3);
	if (b_slice.toString() === "ef" && b_slice.len === 2 && b_slice.cap === 6) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 2. [1:3] => got " + b_slice.v + ", len=" + b_slice.len + ", cap=" + b_slice.cap + "<br>");

		pass = false, PASS = false;
	}

	b_slice.set(a_slice, 0, 3);
	if (b_slice.toString() !== "def") {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 2. [:3] => got " + b_slice.v + "<br>");
		pass = false, PASS = false;
	}

	b_slice.set(a_slice, 0);
	if (b_slice.toString() !== "defg") {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 2. [:] => got " + b_slice.v + "<br>");
		pass = false, PASS = false;
	}



	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function useFunc() {
	var pass = true;


	var Max = function(slice) {
		var max = slice[0];
		for (var index = 1; index < slice.length; index++) {
			if (slice[index] > max) {
				max = slice[index];
			}
		}
		return max;
	};

	var A1 = g.MkArray([10], 0, [1, 2, 3, 4, 5, 6, 7, 8, 9]);
	var A2 = g.MkArray([4], 0, [1, 2, 3, 4]);
	var A3 = g.MkArray([1], 0, [1]);

	var slice = g.NilSlice();

	slice.set(A1, 0);
	if (Max(slice.v) !== 9) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: A1 => got " + Max(slice.v) + ", want 9<br>");
		pass = false, PASS = false;
	}

	slice.set(A2, 0);
	if (Max(slice.v) !== 4) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: A2 => got " + Max(slice.v) + ", want 4<br>");
		pass = false, PASS = false;
	}

	slice.set(A3, 0);
	if (Max(slice.v) !== 1) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: A3 => got " + Max(slice.v) + ", want 1<br>");
		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function reference() {
	var pass = true;

	var fmtSlice = function(slice) {
		var s = "[";
		for (var index = 0; index < slice.len - 1; index++) {
			s += "" + slice.v[index] + ",";
		}
		s += "" + slice.v[slice.len - 1] + "]";

		return s;
	};

	var A = g.MkArray([10], 0, ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j']);
	var slice1 = g.Slice(A, 3, 7);
	var slice2 = g.Slice(A, 5);
	var slice3 = g.Slice(slice1, 0, 2);



	var _ = function(msg, in_, out) { return {
		msg:msg,
		in_:in_,
		out:out
	};}; tests = [
		_("A", fmtSlice(A, 0), "['a','b','c','d','e','f','g','h','i','j']"),
		_("slice1", fmtSlice(slice1), "['d','e','f','g']"),
		_("slice2", fmtSlice(slice2), "['f','g','h','i','j']"),
		_("slice3", fmtSlice(slice3), "['d','e']")
	];

	var t; for (var _ in tests) { t = tests[_];
		if (JSON.stringify(t.in_) !== JSON.stringify(t.out)) {
			document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. " + t.msg + " => got " + t.in_ + ", want " + t.out + "<br>");
			pass = false, PASS = false;
		}
	}


	A.v[4] = 'E';

	_ = function(msg, in_, out) { return {
		msg:msg,
		in_:in_,
		out:out
	};}; tests = [
		_("A", fmtSlice(A, 0), "['a','b','c','d','E','f','g','h','i','j']"),
		_("slice1", fmtSlice(slice1), "['d','E','f','g']"),
		_("slice2", fmtSlice(slice2), "['f','g','h','i','j']"),
		_("slice3", fmtSlice(slice3), "['d','E']")
	];

	var t; for (var _ in tests) { t = tests[_];
		if (JSON.stringify(t.in_) !== JSON.stringify(t.out)) {
			document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 2. " + t.msg + " => got " + t.in_ + ", want " + t.out + "<br>");
			pass = false, PASS = false;
		}
	}


	slice2[1] = 'G';

	_ = function(msg, in_, out) { return {
		msg:msg,
		in_:in_,
		out:out
	};}; tests = [
		_("A", fmtSlice(A, 0), "['a','b','c','d','E','f','G','h','i','j']"),
		_("slice1", fmtSlice(slice1), "['d','E','f','G']"),
		_("slice2", fmtSlice(slice2), "['f','G','h','i','j']"),
		_("slice3", fmtSlice(slice3), "['d','E']")
	];

	var t; for (var _ in tests) { t = tests[_];
		if (JSON.stringify(t.in_) !== JSON.stringify(t.out)) {
			document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 3. " + t.msg + " => got " + t.in_ + ", want " + t.out + "<br>");
			pass = false, PASS = false;
		}
	}



	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function resize() {
	var pass = true;

	var slice = g.NilSlice();


	slice = g.MkSlice(0, 4, 5);

	if (slice.len === 4 && slice.cap === 5 && slice.v[0] === 0 && slice.v[1] === 0 && slice.v[2] === 0 && slice.v[3] === 0) {


	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. got " + slice.v + ", want [0 0 0 0])<br>");
		pass = false, PASS = false;
	}


	slice.v[1] = 2, slice.v[3] = 3;

	if (slice.v[0] === 0 && slice.v[1] === 2 && slice.v[2] === 0 && slice.v[3] === 3) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 2. got " + slice.v + ", want [0 2 0 3])<br>");
		pass = false, PASS = false;
	}


	slice = g.MkSlice(0, 2);

	if (slice.len === 2 && slice.cap === 2 && slice.v[0] === 0 && slice.v[1] === 0) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 3. got " + slice.v + ", want [0 0])<br>");
		pass = false, PASS = false;
	}


	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function main() {
	document.write("<br><br>== Slices<br><br>");

	document.write("=== RUN builtIn<br>");
	builtIn();
	document.write("=== RUN shortHand<br>");
	shortHand();
	document.write("=== RUN useFunc<br>");
	useFunc();
	document.write("=== RUN reference<br>");
	reference();
	document.write("=== RUN resize<br>");
	resize();

	if (PASS) {
		document.write("PASS<br>");
	} else {
		document.write("FAIL<br>");
		alert("Fail: Slices");
	}
} main();
