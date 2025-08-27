const { execSync } = require('child_process');
const path = require('path');

async function runTest(name, command) {
    console.log(`\n=== ${name} ===`);
    console.log(`Command: ${command}`);
    console.log('----------------------------------------');

    try {
        const result = execSync(command, {
            encoding: 'utf8',
            timeout: 120000, // 2 minutes timeout
            stdio: 'inherit'
        });
    } catch (error) {
        console.error(`Error running ${name}: ${error.message}`);
    }
}

async function main() {
    console.log('File I/O Performance Comparison');
    console.log('===============================');

    const nodeDir = path.join(__dirname, '..', 'node');
    const goDir = path.join(__dirname, '..', 'go');

    // Node.js tests
    await runTest('Node.js File I/O Test', `cd "${nodeDir}" && node file-io.js`);

    // Go tests
    await runTest('Go File I/O Test', `cd "${goDir}/fileio" && go run main.go`);

    console.log('\n=== File I/O Performance Test Complete ===');
}

if (require.main === module) {
    main().catch(console.error);
}
