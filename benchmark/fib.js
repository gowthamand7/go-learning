function fibonacci() {
    let a = 0, b = 1;
    return function() {
        const result = a;
        const c = a + b;
        a = b;
        b = c;
        return result;
    };
}

console.time("nodejs");

const fib = fibonacci();
for (let i = 0; i < 100000000; i++) {
    fib();
}

console.timeEnd("nodejs");
