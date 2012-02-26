/* Generated by GoScript <github.com/kless/GoScript> */















var g = {}; (function() {


function Export(pkg, exported) {
	var v; for (_ in exported) { v = exported[_];
		pkg.v = v;
	}
}





function S(f, len, cap) {
	this.f=f;
	this.len=len;
	this.cap=cap;
}


S.prototype.set = function(i, low, high) {
	this.len = high - low;

	if (i.f !== undefined) {
		this.f = i.f.slice(low, high);
		this.cap = i.cap - low;
	} else {
		this.f = i.slice(low, high);
		this.cap = i.length - low;
	}
}


S.prototype.make = function(len, cap) {
	this.len = len;

	for (var i = 0; i < len; i++) {
		this.f[i] = 0;
	}

	if (cap !== undefined) {
		this.cap = cap;
	} else {
		this.cap = len;
	}
}


S.prototype.append = function(elt) {
	if (JSON.stringify(this.len) === JSON.stringify(this.cap)) {
		this.cap = this.len * 2;
	}
	this.len++;
}


S.prototype.toString = function() {
	return this.f.join("");
}


S.prototype.isNil = function() {
	if (this.len !== 0) {
		return false;
	}
	return true;
}







function M(f, zero) {
	this.f=f;
	this.zero=zero;
}




M.prototype.get = function(k) {
	var v = this.f;


	for (var i = 0; i < arguments.length; i++) {
		v = v[arguments[i]];
	}

	if (v === undefined) {
		return [this.zero, false];
	}
	return [v, true];
}

g.Export(g, [Export, S, M]);
})();
