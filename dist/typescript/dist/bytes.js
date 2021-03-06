"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const sprintf_js_1 = require("sprintf-js");
exports.char = (c) => {
    return c.charCodeAt(0);
};
function write(buf, value, type) {
    const regex = /\[(\w+)\]/m;
    let m;
    if ((m = regex.exec(type)) !== null) {
        const subtype = type.slice(m[0].length);
        return value.map((v) => {
            return write(buf, v, subtype);
        });
    }
    switch (type) {
        case 'bool': return buf.uint8(value ? 1 : 0);
        case 'byte': return buf.uint8(value);
        case 'uint8': return buf.uint8(value);
        case 'uint16': return buf.uint16le(value);
        case 'uint32': return buf.uint32le(value);
        case 'uint64': return buf.uint64le(value);
        default: throw new Error('Invalid type : ' + type);
    }
}
exports.write = write;
// read fills the value with the appropriate number of bytes from the buffer.
//
// This is useful for fixed size types such as int, float etc.
function read(buf, type) {
    const regex = /\[(\w+)\]/m;
    let m;
    if ((m = regex.exec(type)) !== null) {
        const subtype = type.slice(m[0].length);
        return [...Array(parseInt(m[1], 10))].map(() => {
            return read(buf, subtype);
        });
    }
    switch (type) {
        case 'bool': return buf.uint8() ? true : false;
        case 'uint8': return buf.uint8();
        case 'byte': return buf.uint8();
        case 'uint16': return buf.uint16le();
        case 'uint32': return buf.uint32le();
        case 'uint64': return buf.uint64le();
        default: throw new Error('Invalid type : ' + type);
    }
}
exports.read = read;
exports.WriteFixedChar = (buf, value, size) => {
    const v = value ? Buffer.from(value, 'ascii') : Buffer.from([]);
    if (v.length > size) {
        throw new Error(`FixedChar too long ${value.length} > ${size}`);
    }
    buf.write(v);
    // Pad with zeroes
    if (v.length < size) {
        buf.write(Buffer.alloc(size - v.length));
    }
    return;
};
// WriteVarBin writes a variable binary value
exports.WriteVarBin = (buf, value, sizeBits) => {
    const v = Buffer.from(value || []);
    exports.WriteVariableSize(buf, v.length, sizeBits, 0);
    buf.write(v);
    return;
};
// WriteVariableSize writes a size/length to the buffer using the specified size unsigned integer.
// defaultSizeBits is used if sizeBits is zero.
exports.WriteVariableSize = (buf, size, sizeBits, defaultSizeBits) => {
    if (sizeBits === 0) {
        sizeBits = defaultSizeBits;
    }
    if (size >= Math.pow(2, sizeBits)) {
        throw new Error(sprintf_js_1.sprintf('Size beyond size bits limit (%d) : %d', (Math.pow(2, sizeBits)) - 1, size));
    }
    switch (sizeBits) {
        case 8:
            return write(buf, size, 'uint8');
        case 16:
            return write(buf, size, 'uint16');
        case 32:
            return write(buf, size, 'uint32');
        case 64:
            return write(buf, size, 'uint64');
        default:
            throw new Error('Invalid variable size bits : ' + sizeBits);
    }
};
// WriteVarChar writes a variable character string
exports.WriteVarChar = (buf, value, sizeBits) => {
    const v = value ? Buffer.from(value, 'ascii') : Buffer.from([]);
    exports.WriteVariableSize(buf, v.length, sizeBits, 0);
    buf.write(v);
};
// ReadFixedChar reads a fixed character string
exports.ReadFixedChar = (buf, size) => {
    return buf.read(size).toString('ascii');
};
// ReadVarBin reads a variable binary value
exports.ReadVarBin = (buf, sizeBits) => {
    const size = exports.ReadVariableSize(buf, sizeBits, 0);
    return buf.read(size);
};
// ReadVariableSize reads a size/length from the buffer using the specified size unsigned integer.
// defaultSizeBits is used if sizeBits is zero.
exports.ReadVariableSize = (buf, sizeBits, defaultSizeBits) => {
    if (sizeBits === 0) {
        sizeBits = defaultSizeBits;
    }
    switch (sizeBits) {
        case 8: return read(buf, 'uint8');
        case 16: return read(buf, 'uint16');
        case 32: return read(buf, 'uint32');
        case 64: return read(buf, 'uint64');
        default: throw new Error('Invalid variable size bits : ' + sizeBits);
    }
};
// ReadVarChar reads a variable character string
exports.ReadVarChar = (buf, sizeBits) => {
    const size = exports.ReadVariableSize(buf, sizeBits, 0);
    return buf.read(size).toString('ascii');
};
// { write, read, ReadVarChar, ReadVariableSize, ReadVarBin, ReadFixedChar, WriteVarChar, WriteVariableSize, WriteFixedChar }
