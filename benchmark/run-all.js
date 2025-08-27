const { execSync } = require('child_process');
const path = require('path');
const fs = require('fs').promises;

async function createResultsDirectory() {
    const resultsDir = path.join(__dirname, '..', 'results');
    try {
        await fs.mkdir(resultsDir, { recursive: true });
    } catch (error) {
        // Directory already exists
    }
    return resultsDir;
}

async function runAllTests() {
    console.log('Complete Node.js vs Go Performance Benchmark');
    console.log('===========================================');
    console.log(`Started at: ${new Date().toISOString()}\n`);

    const benchmarkDir = __dirname;
    const resultsDir = await createResultsDirectory();
    const timestamp = new Date().toISOString().replace(/[:.]/g, '-');

    const tests = [
        {
            name: 'CPU Intensive Tests',
            script: 'test-cpu.js'
        },
        {
            name: 'File I/O Tests',
            script: 'test-file-io.js'
        },
        {
            name: 'JSON Processing Tests',
            script: 'test-json.js'
        }
    ];

    for (const test of tests) {
        console.log(`\n${'='.repeat(60)}`);
        console.log(`Running ${test.name}`);
        console.log(`${'='.repeat(60)}`);

        try {
            const command = `cd "${benchmarkDir}" && node ${test.script}`;
            execSync(command, {
                encoding: 'utf8',
                stdio: 'inherit',
                timeout: 300000 // 5 minutes timeout
            });
        } catch (error) {
            console.error(`Error running ${test.name}: ${error.message}`);
        }

        console.log(`\n${test.name} completed.\n`);
    }

    console.log(`\n${'='.repeat(60)}`);
    console.log('HTTP Server Tests');
    console.log(`${'='.repeat(60)}`);
    console.log('\nTo run HTTP server tests:');
    console.log('1. Start Node.js server: cd node && npm run http-server');
    console.log('2. In another terminal: cd benchmark && node test-http.js node');
    console.log('3. Stop Node.js server (Ctrl+C)');
    console.log('4. Start Go server: cd go && go run http-server.go');
    console.log('5. In another terminal: cd benchmark && node test-http.js go');
    console.log('6. Stop Go server (Ctrl+C)');

    console.log(`\n${'='.repeat(60)}`);
    console.log('All Benchmark Tests Completed!');
    console.log(`Finished at: ${new Date().toISOString()}`);
    console.log(`${'='.repeat(60)}`);
}

if (require.main === module) {
    runAllTests().catch(console.error);
}
