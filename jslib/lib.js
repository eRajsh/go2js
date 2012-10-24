







var g = {}; (function() {


const 
invalidT = 0,
arrayT = 1,
mapT = 2,
sliceT = 3;





(function() {


	if (!Array.isArray) {
		Array.isArray = function(arg) {
			return Object.prototype.toString.call(arg) === "[object Array]";
		};
	}
}());





function arrayType(v, len_) {
	this.v=v;

	this.len_=len_
}


arrayType.prototype.len = function(dim) {
	if (dim === undefined) {
		return this.len_[0];
	}
	return this.len_[arguments.length];
}


arrayType.prototype.cap = function(dim) {
	if (dim === undefined) {
		return this.len_[0];
	}
	return this.len_[arguments.length];
}


arrayType.prototype.str = function() {
	return this.v.join("");
}


arrayType.prototype.typ = function() { return arrayT; }



function MkArray(dim, zero, data) {
	var a = new arrayType([], g.Map(0));

	if (data !== undefined) {
		if (!equalDim(dim, getDimArray(data))) {
			a.v = initArray(dim, zero);
			mergeArray(a.v, data);
		} else {
			a.v = data;
		}
	} else {
		a.v = initArray(dim, zero);
	}

	var v; for (var i in dim) { v = dim[i];
		a.len_[i] = v;
	}

	return a;
}




function equalDim(d1, d2) {
	if (d1.length !== d2.length) {
		return false;
	}
	var v; for (var i in d1) { v = d1[i];
		if (JSON.stringify(v) !== JSON.stringify(d2[i])) {
			return false;
		}
	}
	return true;
}


function getDimArray(a) { var dim = [];
	for (;;) {
		dim.push(a.length);

		if (Array.isArray(a[0])) {
			a = a[0];
		} else {
			break;
		}
	}
	return dim;
}


function initArray(dim, zero) { var a = [];
	if (dim.length === 0) {
		return zero;
	}
	var nextArray = initArray(dim.slice(1), zero);

	for (var i = 0; i < dim[0]; i++) {
		a[i] = nextArray;
	}
	return a;
}


function mergeArray(dst, src) {
	var srcVal; for (var i in src) { srcVal = src[i];
		if (Array.isArray(srcVal)) {
			mergeArray(dst[i], srcVal);
		} else {
			var isHashMap = false;


			if (typeof(srcVal) === "object") {
				var v; for (var k in srcVal) { v = srcVal[k];
					if (srcVal.hasOwnProperty(k)) {
						isHashMap = true;
						i = k;
						dst[i] = v;
					}
				}
			}
			if (!isHashMap) {
				dst[i] = srcVal;
			}
		}
	}
}





function sliceType(v, low, high, len, cap, nil_) {
	this.v=v;

	this.low=low;
	this.high=high;
	this.len=len;
	this.cap=cap;

	this.nil_=nil_
}


sliceType.prototype.typ = function() { return sliceT; }

sliceType.prototype.isNil = function() {
	if (this.len !== 0 || this.cap !== 0) {
		return false;
	}
	return this.nil_;
}


function MkSlice(zero, len, cap) {
	var s = new sliceType([], 0, 0, 0, 0, false);

	if (zero === undefined) {
		s.nil_ = true;
		return s;
	}


	s.v = Array(len);
	for (var i = 0; i < len; i++) {
		s.v[i] = zero;
	}

	if (cap !== undefined) {
		s.cap = cap;
	} else {
		s.cap = len;
	}

	s.len = len;
	return s;
}


function Slice(zero, data) {
	var s = new sliceType([], 0, 0, 0, 0, false);

	if (zero === undefined) {
		s.nil_ = true;
		return s;
	}

	var srcVal; for (var i in data) { srcVal = data[i];
		var isHashMap = false;


		if (typeof(srcVal) === "object") {
			var v; for (var k in srcVal) { v = srcVal[k];
				if (srcVal.hasOwnProperty(k)) {
					isHashMap = true;

					for (i; i < k; i++) {
						s.v[i] = zero;
					}
					s.v[i] = v;
				}
			}
		}
		if (!isHashMap) {
			s.v[i] = srcVal;
		}
	}

	s.len = s.v.length;
	s.cap = s.len;
	return s;
}


function SliceFrom(src, low, high) {
	var s = new sliceType([], 0, 0, 0, 0, false);

	if (low !== undefined) {
		s.low = low;
	} else {
		s.low = 0;
	}
	if (high !== undefined) {
		s.high = high;
	} else {
		if (src.arr !== undefined) {
			s.high = src.len;
		} else {
			s.high = src.v.length;
		}
	}

	s.len = s.high - s.low;

	if (src.nil_ !== undefined) {
		s.cap = src.cap - s.low;


	} else {
		s.cap = src.cap() - s.low;
	}

	s.v = src.v.slice(s.low, s.high);

	return s;
}


sliceType.prototype.str = function() {
	return this.v.join("");
}




function Append(src) { var elt = [].slice.call(arguments).slice(1); var dst = new sliceType([], 0, 0, 0, 0, false);

	dst.low = src.low;
	dst.high = src.high;
	dst.len = src.len;
	dst.cap = src.cap;
	dst.nil_ = src.nil_;

	dst.v = Array(src.len);
	var v; for (var i in src.v) { v = src.v[i];
		dst.v[i] = v;
	}





	var v; for (var _ in elt) { v = elt[_];
		if (Array.isArray(v)) {
			var vArr; for (var _ in v) { vArr = v[_];
				dst.v.push(vArr);
				if (JSON.stringify(dst.len) === JSON.stringify(dst.cap)) {
					dst.cap = dst.len * 2;
				}
				dst.len++;
			}
			break;
		}

		dst.v.push(v);
		if (JSON.stringify(dst.len) === JSON.stringify(dst.cap)) {
			dst.cap = dst.len * 2;
		}
		dst.len++;
	}
	return dst;
}


function Copy(dst, src) { var n = 0;

	if (src.nil_ !== undefined) {
		for (; n < src.len; n++) {
			if (JSON.stringify(n) === JSON.stringify(dst.len)) {
				return n;
			}
			dst.v[n] = src.v[n];
		}
		return n;
	}


	for (; n < src.length; n++) {
		if (JSON.stringify(n) === JSON.stringify(dst.len)) {
			break;
		}
		dst.v[n] = src[n];
	}
	return n;
}












function mapType(v, zero) {
	this.v=v;
	this.zero=zero
}


mapType.prototype.len = function() {
	var len = 0;
	var _; for (var key in this.v) { _ = this.v[key];
		if (this.v.hasOwnProperty(key)) {
			len++;
		}
	}
	return len;
}


mapType.prototype.typ = function() { return mapT; }


function Map(zero, v) {
	var m = new mapType(v, zero);
	return m;
}




mapType.prototype.get = function(k) {
	var v = this.v;


	for (var i = 0; i < arguments.length; i++) {
		v = v[arguments[i]];
	}

	if (v === undefined) {
		return [this.zero, false];
	}
	return [v, true];
}





function Export(pkg, exported) {
	var v; for (var _ in exported) { v = exported[_];
		pkg.v = v;
	}
}

g.MkArray = MkArray;
g.MkSlice = MkSlice;
g.Slice = Slice;
g.SliceFrom = SliceFrom;
g.Append = Append;
g.Copy = Copy;
g.Map = Map;
g.Export = Export;

})();
/* Generated by GoScript (github.com/kless/goscript) */
