const express = require('express');

const app = express();
const port = 3000;

app.use(express.json());

// 简单的健康检查
app.get('/health', (req, res) => {
    res.json({ status: 'ok', timestamp: Date.now() });
});

// 返回简单文本
app.get('/hello', (req, res) => {
    res.text('Hello World!');
});

// JSON 响应
app.get('/json', (req, res) => {
    res.json({
        message: 'Hello World!',
        timestamp: Date.now(),
        data: Array.from({ length: 100 }, (_, i) => ({ id: i, value: `item-${i}` }))
    });
});

// CPU 密集型任务
app.get('/cpu/:n', (req, res) => {
    const n = parseInt(req.params.n) || 40;

    function fibonacci(n) {
        if (n <= 1) return n;
        return fibonacci(n - 1) + fibonacci(n - 2);
    }

    const start = process.hrtime();
    const result = fibonacci(n);
    const end = process.hrtime(start);

    res.json({
        input: n,
        result: result,
        time_ms: end[0] * 1000 + end[1] / 1000000
    });
});

// POST 数据处理
app.post('/data', (req, res) => {
    const data = req.body;

    // 简单的数据处理
    const processed = {
        received_count: Array.isArray(data) ? data.length : Object.keys(data).length,
        processed_at: Date.now(),
        summary: JSON.stringify(data).length
    };

    res.json(processed);
});

app.listen(port, () => {
    console.log(`Node.js HTTP server running on port ${port}`);
});
