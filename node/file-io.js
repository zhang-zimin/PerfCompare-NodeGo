const fs = require('fs').promises;
const path = require('path');

// 文件 I/O 性能测试

async function createTestData(size) {
    // 创建指定大小的测试数据
    const data = Array.from({ length: size }, (_, i) => `Line ${i}: ${'x'.repeat(50)}\n`).join('');
    return data;
}

async function testFileWrite(filePath, data, iterations = 5) {
    console.log(`Testing file write (${data.length} bytes)...`);
    const times = [];

    for (let i = 0; i < iterations; i++) {
        const testFile = `${filePath}_${i}.txt`;
        const start = process.hrtime();
        await fs.writeFile(testFile, data);
        const end = process.hrtime(start);
        const timeMs = end[0] * 1000 + end[1] / 1000000;
        times.push(timeMs);

        // 清理文件
        await fs.unlink(testFile);
        console.log(`Write ${i + 1}: ${timeMs.toFixed(3)}ms`);
    }

    const avgTime = times.reduce((a, b) => a + b, 0) / times.length;
    console.log(`Average write time: ${avgTime.toFixed(3)}ms\n`);
    return avgTime;
}

async function testFileRead(filePath, iterations = 5) {
    console.log(`Testing file read...`);
    const times = [];

    for (let i = 0; i < iterations; i++) {
        const start = process.hrtime();
        const data = await fs.readFile(filePath, 'utf8');
        const end = process.hrtime(start);
        const timeMs = end[0] * 1000 + end[1] / 1000000;
        times.push(timeMs);
        console.log(`Read ${i + 1}: ${timeMs.toFixed(3)}ms (${data.length} bytes)`);
    }

    const avgTime = times.reduce((a, b) => a + b, 0) / times.length;
    console.log(`Average read time: ${avgTime.toFixed(3)}ms\n`);
    return avgTime;
}

async function testConcurrentFileOps(basePath, fileCount = 10) {
    console.log(`Testing concurrent file operations (${fileCount} files)...`);
    const data = await createTestData(1000);

    const start = process.hrtime();

    // 并发写入
    const writePromises = Array.from({ length: fileCount }, (_, i) =>
        fs.writeFile(`${basePath}_concurrent_${i}.txt`, data)
    );
    await Promise.all(writePromises);

    // 并发读取
    const readPromises = Array.from({ length: fileCount }, (_, i) =>
        fs.readFile(`${basePath}_concurrent_${i}.txt`, 'utf8')
    );
    const results = await Promise.all(readPromises);

    const end = process.hrtime(start);
    const timeMs = end[0] * 1000 + end[1] / 1000000;

    // 清理文件
    const cleanupPromises = Array.from({ length: fileCount }, (_, i) =>
        fs.unlink(`${basePath}_concurrent_${i}.txt`)
    );
    await Promise.all(cleanupPromises);

    console.log(`Concurrent operations completed: ${timeMs.toFixed(3)}ms`);
    console.log(`Average per file: ${(timeMs / fileCount).toFixed(3)}ms\n`);

    return timeMs;
}

async function main() {
    console.log('Node.js File I/O Tests');
    console.log('======================');

    const dataDir = path.join(__dirname, '..', 'data');
    try {
        await fs.mkdir(dataDir, { recursive: true });
    } catch (error) {
        // Directory already exists
    }

    const basePath = path.join(dataDir, 'test_node');

    // 测试不同大小的文件
    const sizes = [1000, 10000, 100000]; // 行数

    for (const size of sizes) {
        console.log(`\n=== Testing with ${size} lines ===`);
        const data = await createTestData(size);
        const filePath = `${basePath}_${size}.txt`;

        // 写入测试
        await testFileWrite(filePath, data);

        // 先创建文件用于读取测试
        await fs.writeFile(filePath, data);

        // 读取测试
        await testFileRead(filePath);

        // 清理测试文件
        await fs.unlink(filePath);
    }

    // 并发测试
    console.log('\n=== Concurrent File Operations ===');
    await testConcurrentFileOps(basePath, 10);
    await testConcurrentFileOps(basePath, 50);
}

if (require.main === module) {
    main().catch(console.error);
}
