/**
 * Tiny calculator module used to exercise CI.
 */

export function add(a, b) {
  return a + b;
}

export function sub(a, b) {
  return a - b;
}

export function mul(a, b) {
  return a * b;
}

export function div(a, b) {
  if (b === 0) {
    throw new RangeError("divide by zero");
  }
  return a / b;
}
