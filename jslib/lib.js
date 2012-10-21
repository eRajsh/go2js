







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













function Copy(dst, src) { var n = 0;

	if (src.arr !== undefined) {
		for (var i = src.low; i < src.high; i++) {
			if (JSON.stringify(n) === JSON.stringify(dst.len)) {
				return n;
			}
			dst.arr.v[n] = src.arr.v[i];
			n++;
		}
		var v; for (var i in src.v) { v = src.v[i];
			if (JSON.stringify(n) === JSON.stringify(dst.len)) {
				return n;
			}
			dst.v[i] = v;
			n++;
		}
		return n;
	}


	for (; n < src.length; n++) {
		if (JSON.stringify(n) === JSON.stringify(dst.len)) {
			break;
		}
		dst.arr.v[n] = src[n];
	}
	return n;
}


function sliceType(arr, v, low, high, len, cap, nil_) {
	this.arr=arr;
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
	var s = new sliceType(undefined, [], 0, 0, 0, 0, false);

	if (zero === undefined) {
		s.nil_ = true;
		return s;
	}

	var a = new arrayType([], g.Map(0));
	a.len_[0] = len;
	for (var i = 0; i < len; i++) {
		a.v[i] = zero;
	}
	s.arr = a;

	if (cap !== undefined) {
		s.cap = cap;
	} else {
		s.cap = len;
	}

	s.len = len;
	s.high = len;
	return s;
}


function Slice(zero, data) {
	var s = new sliceType(undefined, [], 0, 0, 0, 0, false);

	if (zero === undefined) {
		s.nil_ = true;
		return s;
	}

	var a = new arrayType([], g.Map(0));
	var srcVal; for (var i in data) { srcVal = data[i];
		var isHashMap = false;


		if (typeof(srcVal) === "object") {
			var v; for (var k in srcVal) { v = srcVal[k];
				if (srcVal.hasOwnProperty(k)) {
					isHashMap = true;

					for (i; i < k; i++) {
						a.v[i] = zero;
					}
					a.v[i] = v;
				}
			}
		}
		if (!isHashMap) {
			a.v[i] = srcVal;
		}
	}
	s.len = a.v.length;
	a.len_[0] = s.len;
	s.arr = a;

	s.cap = s.len;
	s.high = s.len;
	return s;
}


function SliceFrom(src, low, high) {
	var s = new sliceType(undefined, [], 0, 0, 0, 0, false);
	s.set(src, low, high);
	return s;
}


sliceType.prototype.set = function(src, low, high) {
	if (low !== undefined) {
		this.low = low;
	} else {
		this.low = 0;
	}
	if (high !== undefined) {
		this.high = high;
	} else {
		if (src.arr !== undefined) {
			this.high = src.len;
		} else {
			this.high = src.v.length;
		}
	}

	this.len = this.high - this.low;

	if (src.arr !== undefined) {
		this.arr = src.arr;
		this.cap = src.cap - this.low;
		this.low += src.low;
		this.high += src.low;
	} else {
		this.arr = src;
		this.cap = src.cap() - this.low;
	}
}







sliceType.prototype.get = function() {
	if (this.arr !== undefined) {
		return this.arr.v.slice(this.low, this.high);
	}
	return this.v;
}


sliceType.prototype.str = function() {
	var _s = this.get();
	return _s.join("");
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
g.Copy = Copy;
g.MkSlice = MkSlice;
g.Slice = Slice;
g.SliceFrom = SliceFrom;
g.Map = Map;
g.Export = Export;

})();
/* Generated by GoScript (github.com/kless/goscript) */
