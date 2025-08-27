// JSON 序列化和反序列化性能测试

function generateTestData(complexity = 'medium') {
    const complexities = {
        simple: () => ({
            id: 1,
            name: 'Test User',
            email: 'test@example.com',
            active: true
        }),

        medium: () => ({
            id: 1,
            name: 'Test User',
            email: 'test@example.com',
            profile: {
                age: 30,
                city: 'New York',
                interests: ['coding', 'music', 'travel'],
                settings: {
                    theme: 'dark',
                    notifications: true,
                    privacy: {
                        public: false,
                        friends: true
                    }
                }
            },
            orders: Array.from({ length: 100 }, (_, i) => ({
                id: i,
                product: `Product ${i}`,
                price: Math.random() * 100,
                date: new Date().toISOString(),
                items: Array.from({ length: Math.floor(Math.random() * 5) + 1 }, (_, j) => ({
                    id: j,
                    name: `Item ${j}`,
                    quantity: Math.floor(Math.random() * 10) + 1
                }))
            }))
        }),

        complex: () => ({
            metadata: {
                version: '1.0.0',
                generated: new Date().toISOString(),
                schema: 'user-data-v1'
            },
            users: Array.from({ length: 1000 }, (_, i) => ({
                id: i,
                username: `user${i}`,
                email: `user${i}@example.com`,
                profile: {
                    firstName: `First${i}`,
                    lastName: `Last${i}`,
                    bio: `This is a bio for user ${i}. `.repeat(10),
                    avatar: `https://example.com/avatar/${i}.jpg`,
                    social: {
                        twitter: `@user${i}`,
                        github: `user${i}`,
                        linkedin: `user-${i}`
                    }
                },
                preferences: {
                    theme: ['light', 'dark'][i % 2],
                    language: ['en', 'zh', 'es', 'fr'][i % 4],
                    notifications: {
                        email: i % 2 === 0,
                        push: i % 3 === 0,
                        sms: i % 5 === 0
                    }
                },
                activity: Array.from({ length: 50 }, (_, j) => ({
                    id: j,
                    type: ['login', 'logout', 'purchase', 'view'][j % 4],
                    timestamp: new Date(Date.now() - j * 3600000).toISOString(),
                    metadata: {
                        ip: `192.168.1.${j % 255}`,
                        userAgent: `Browser ${j % 10}`,
                        sessionId: `session-${i}-${j}`
                    }
                }))
            }))
        })
    };

    return complexities[complexity]();
}

function testSerialization(data, iterations = 5) {
    console.log('Testing JSON serialization...');
    const times = [];
    let jsonString = '';

    for (let i = 0; i < iterations; i++) {
        const start = process.hrtime();
        jsonString = JSON.stringify(data);
        const end = process.hrtime(start);
        const timeMs = end[0] * 1000 + end[1] / 1000000;
        times.push(timeMs);
        console.log(`Serialization ${i + 1}: ${timeMs.toFixed(3)}ms`);
    }

    const avgTime = times.reduce((a, b) => a + b, 0) / times.length;
    console.log(`Average serialization time: ${avgTime.toFixed(3)}ms`);
    console.log(`JSON size: ${jsonString.length} characters\n`);

    return { avgTime, jsonString };
}

function testDeserialization(jsonString, iterations = 5) {
    console.log('Testing JSON deserialization...');
    const times = [];

    for (let i = 0; i < iterations; i++) {
        const start = process.hrtime();
        const data = JSON.parse(jsonString);
        const end = process.hrtime(start);
        const timeMs = end[0] * 1000 + end[1] / 1000000;
        times.push(timeMs);
        console.log(`Deserialization ${i + 1}: ${timeMs.toFixed(3)}ms`);
    }

    const avgTime = times.reduce((a, b) => a + b, 0) / times.length;
    console.log(`Average deserialization time: ${avgTime.toFixed(3)}ms\n`);

    return avgTime;
}

function testRoundTrip(data, iterations = 5) {
    console.log('Testing round-trip (serialize + deserialize)...');
    const times = [];

    for (let i = 0; i < iterations; i++) {
        const start = process.hrtime();
        const jsonString = JSON.stringify(data);
        const parsedData = JSON.parse(jsonString);
        const end = process.hrtime(start);
        const timeMs = end[0] * 1000 + end[1] / 1000000;
        times.push(timeMs);
        console.log(`Round-trip ${i + 1}: ${timeMs.toFixed(3)}ms`);
    }

    const avgTime = times.reduce((a, b) => a + b, 0) / times.length;
    console.log(`Average round-trip time: ${avgTime.toFixed(3)}ms\n`);

    return avgTime;
}

function main() {
    console.log('Node.js JSON Processing Tests');
    console.log('=============================');

    const complexities = ['simple', 'medium', 'complex'];

    complexities.forEach(complexity => {
        console.log(`\n=== ${complexity.toUpperCase()} Data Structure ===`);
        const data = generateTestData(complexity);

        const { avgTime: serializationTime, jsonString } = testSerialization(data);
        const deserializationTime = testDeserialization(jsonString);
        const roundTripTime = testRoundTrip(data);

        console.log(`Summary for ${complexity} data:`);
        console.log(`- Serialization: ${serializationTime.toFixed(3)}ms`);
        console.log(`- Deserialization: ${deserializationTime.toFixed(3)}ms`);
        console.log(`- Round-trip: ${roundTripTime.toFixed(3)}ms`);
        console.log('---');
    });
}

if (require.main === module) {
    main();
}
