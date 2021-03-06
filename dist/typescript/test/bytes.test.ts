import * as mocha from 'mocha';
import * as chai from 'chai';
import { read, write } from '../src/bytes';
import _ from '@keyring/util';
const expect = chai.expect;
[mocha]

describe('read', () => {
	it('should read array', () => {
		let size = 2;
		let buf = new _.Reader(Buffer.from([1, 2, 3, 4, 5, 6]));
		const type = '[][3]byte'.slice(2);
		let arr = [...Array(size)].map(() => read(buf, type));
		console.log('read arr:', arr);
		expect(arr).to.eql([[1,2,3],[4,5,6]])
	});
});

describe('write', () => {
	it('should write array', () => {
		let size = 2;
		let buf = new _.Writer();
		const type = '[][3]byte'.slice(2);
		[...Array(size)].map(() => write(buf, [1, 2, 3], type));
		console.log('write buf: ', buf.buf);
		expect(buf.buf).to.eql(Buffer.from([1,2,3,1,2,3]));
	});
});

