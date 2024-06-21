const a = ["1", "2", "3"];
const b = ["1", "2", "3"];

console.log(a == b); // false
console.log(a === b); // false
console.log(JSON.stringify(a) === JSON.stringify(b)); // true

const obj1 = { a: 1, b: 2 };
const obj2 = { a: 1, b: 2 };

console.log(obj1 == obj2); // false
console.log(JSON.stringify(obj1) === JSON.stringify(obj2)); // true


const n = null;
const n2 = null;

console.log(n == n2); // true