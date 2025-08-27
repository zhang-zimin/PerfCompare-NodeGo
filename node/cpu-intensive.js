// CPU 密集型计算测试 - 斐波那契数列

function fibonacci(n) {
    if (n <= 1) return n;
    return fibonacci(n - 1) + fibonacci(n - 2);
}

function fibonacciIterative(n) {
    if (n <= 1) return n;
    let a = 0, b = 1;
    for (let i = 2; i <= n; i++) {
        let temp = a + b;
        a = b;
        b = temp;
    }
    return b;
}

function runTest(name, fn, n, iterations = 5) {
    console.log(`\n=== ${name} (n=${n}) ===`);
    const times = [];

    for (let i = 0; i < iterations; i++) {
        const start = process.hrtime();
        const result = fn(n);
        const end = process.hrtime(start);
        const timeMs = end[0] * 1000 + end[1] / 1000000;
        times.push(timeMs);
        console.log(`Run ${i + 1}: ${timeMs.toFixed(3)}ms - Result: ${result}`);
    }

    const avgTime = times.reduce((a, b) => a + b, 0) / times.length;
    const minTime = Math.min(...times);
    const maxTime = Math.max(...times);

    console.log(`Average: ${avgTime.toFixed(3)}ms`);
    console.log(`Min: ${minTime.toFixed(3)}ms`);
    console.log(`Max: ${maxTime.toFixed(3)}ms`);

    return { average: avgTime, min: minTime, max: maxTime };
}

function main() {
    console.log('Node.js CPU Intensive Tests');
    console.log('===========================');

    // 测试不同规模的斐波那契数列
    const testCases = [35, 40, 42];

    testCases.forEach(n => {
        runTest(`Fibonacci Recursive`, fibonacci, n);
        runTest(`Fibonacci Iterative`, fibonacciIterative, n);
    });

    // 大数计算测试
    console.log('\n=== Large Number Fibonacci (Iterative) ===');
    const largeNumbers = [1000, 5000, 10000];
    largeNumbers.forEach(n => {
        runTest(`Large Fibonacci`, fibonacciIterative, n, 3);
    });
}

if (require.main === module) {
    main();
}
