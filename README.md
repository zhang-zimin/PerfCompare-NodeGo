# Node.js vs Go 性能测试 ⚡

这个项目提供了全面的 Node.js 和 Go 性能对比测试，包含多个真实场景的基准测试。

## 🚀 快速开始

```bash
# 1. 安装依赖
npm run install-deps

# 2. 运行完整基准测试
npm run benchmark

# 3. 或运行单项测试
npm run test:cpu        # CPU 密集型测试
npm run test:file-io    # 文件 I/O 测试
npm run test:json       # JSON 处理测试
```

## 📊 测试场景

1. **HTTP 服务器性能测试** - 并发请求处理能力
2. **CPU 密集型计算** - 斐波那契数列递归和迭代算法
3. **文件 I/O 操作** - 大小不同文件的读写性能和并发操作
4. **JSON 序列化/反序列化** - 简单到复杂数据结构的 JSON 处理

## 🛠️ 环境准备

### 必需软件
- **Node.js** (v18+ 推荐) - [下载](https://nodejs.org/)
- **Go** (v1.21+ 推荐) - [下载](https://golang.org/dl/)

## 📈 性能测试命令

### 完整基准测试
```bash
npm run benchmark       # 运行所有对比测试
```

### 单项测试
```bash
# CPU 密集型计算对比
npm run test:cpu

# 文件 I/O 性能对比  
npm run test:file-io

# JSON 处理性能对比
npm run test:json

# HTTP 服务器测试 (需要手动启动服务器)
npm run test:http
```

### 单独运行 Node.js 测试
```bash
npm run node:cpu        # CPU 测试
npm run node:fileio     # 文件 I/O 测试
npm run node:json       # JSON 测试
```

### 单独运行 Go 测试  
```bash
npm run go:cpu          # CPU 测试
npm run go:fileio       # 文件 I/O 测试
npm run go:json         # JSON 测试
```

### HTTP 服务器测试
```bash
# 启动服务器 (选择一个)
npm run start:node-server  # Node.js 服务器 (端口 3000)
npm run start:go-server    # Go 服务器 (端口 3001)

# 测试端点
curl http://localhost:3000/health  # Node.js
curl http://localhost:3001/health  # Go
```

## 🎯 初步性能结果

基于开发环境的初步测试结果：

### CPU 密集型计算 (斐波那契 n=40, 递归)
- **Go**: ~510ms
- **Node.js**: ~1067ms  
- **结果**: Go 比 Node.js 快 **2.1x**

### CPU 密集型计算 (斐波那契 n=40, 迭代)
- **Go**: <1ms
- **Node.js**: <1ms
- **结果**: 两者都极快，性能相近

### 内存使用
- **Go**: 编译后二进制小，启动快，内存占用低
- **Node.js**: V8 引擎优化好，但启动内存占用较高


## 🔍 测试详情

### CPU 密集型测试
- **递归斐波那契**: 测试纯计算性能和函数调用开销
- **迭代斐波那契**: 测试循环优化和数值计算效率
- **大数计算**: 测试大整数处理能力 (1000, 5000, 10000)

### 文件 I/O 测试
- **顺序读写**: 不同大小文件的读写性能 (1K, 10K, 100K 行)
- **并发操作**: 多文件同时读写的并发处理能力
- **缓存效应**: 重复读取同一文件的缓存优化

### JSON 处理测试
- **简单结构**: 基础对象序列化/反序列化
- **中等复杂**: 嵌套对象和数组处理 (100 项)
- **复杂结构**: 深度嵌套和大数据量处理 (1000 用户, 50 活动记录)

### HTTP 服务器测试
- **简单响应**: 纯文本和 JSON 响应速度
- **CPU 任务**: 服务器端斐波那契计算
- **数据处理**: POST 数据的解析和处理
- **并发负载**: 多客户端并发请求处理
