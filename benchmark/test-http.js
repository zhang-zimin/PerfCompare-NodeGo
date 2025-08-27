const { execSync, spawn } = require('child_process');
const fs = require('fs').promises;
const path = require('path');

async function runHttpBenchmark(target) {
    console.log(`\n=== HTTP Benchmark - ${target.toUpperCase()} ===`);

    const port = target === 'node' ? 3000 : 3001;
    const baseUrl = `http://localhost:${port}`;

    // 测试端点
    const endpoints = [
        { name: 'Health Check', path: '/health' },
        { name: 'Hello World', path: '/hello' },
        { name: 'JSON Response', path: '/json' },
        { name: 'CPU Task (n=35)', path: '/cpu/35' }
    ];

    for (const endpoint of endpoints) {
        console.log(`\nTesting ${endpoint.name}...`);

        try {
            // 使用 curl 进行简单的性能测试
            const command = `curl -s -w "Time: %{time_total}s\\nStatus: %{http_code}\\n" -o /dev/null "${baseUrl}${endpoint.path}"`;

            for (let i = 0; i < 5; i++) {
                try {
                    const result = execSync(command, { encoding: 'utf8', timeout: 10000 });
                    console.log(`Run ${i + 1}: ${result.trim()}`);
                } catch (error) {
                    console.log(`Run ${i + 1}: Error - ${error.message}`);
                }
            }
        } catch (error) {
            console.log(`Error testing ${endpoint.name}: ${error.message}`);
        }
    }
}

async function main() {
    const target = process.argv[2];

    if (!target || !['node', 'go'].includes(target)) {
        console.log('Usage: node test-http.js <node|go>');
        process.exit(1);
    }

    console.log(`HTTP Performance Test for ${target.toUpperCase()}`);
    console.log('================================================');

    console.log(`\nNote: Make sure the ${target} HTTP server is running first!`);
    console.log(`Start it with:`);
    if (target === 'node') {
        console.log(`  cd node && npm run http-server`);
    } else {
        console.log(`  cd go && go run http-server.go`);
    }

    console.log('\nWaiting 3 seconds before starting tests...');
    await new Promise(resolve => setTimeout(resolve, 3000));

    await runHttpBenchmark(target);

    console.log('\n=== HTTP Benchmark Complete ===');
}

if (require.main === module) {
    main().catch(console.error);
}
