import { test } from "node:test";
import assert from "node:assert/strict";
import { add, sub, mul, div } from "../src/calc.js";

test("add", () => {
  assert.equal(add(2, 3), 5);
});

test("sub", () => {
  assert.equal(sub(7, 4), 3);
});

test("mul", () => {
  assert.equal(mul(6, 7), 42);
});

test("div", () => {
  assert.equal(div(10, 2), 5);
});

test("div by zero throws", () => {
  assert.throws(() => div(1, 0), RangeError);
});
