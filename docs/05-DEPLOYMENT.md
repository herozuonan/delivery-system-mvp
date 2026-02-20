# 部署流程文档

## 部署架构

```
┌──────────────────────────────────────────────────────────┐
│                    GitHub仓库                             │
│  (代码 + Dockerfile + docker-compose.yml)                │
└────────────────────┬─────────────────────────────────────┘
                     │ git push
┌────────────────────▼─────────────────────────────────────┐
│              GitHub Actions CI/CD                         │
│  1. 代码拉取                                              │
│  2. 构建 (go build)                                       │
│  3. 测试 (go test)                                        │
│  4. 构建Docker镜像                                        │
│  5. 推送到Docker Registry                                │
└────────────────────┬─────────────────────────────────────┘
                     │ docker push
┌────────────────────▼─────────────────────────────────────┐
│            Docker Registry (镜像存储)                     │
│  app:v1.0.0, app:v1.0.1, app:latest                      │
└────────────────────┬─────────────────────────────────────┘
                     │ docker pull
┌────────────────────▼─────────────────────────────────────┐
│         生产环境 (单VPS + Docker Compose)                 │
│  ┌──────────────────────────────────────────────────┐   │
│  │ docker-compose up -d                             │   │
│  │ - app:v1.0.1 (新版本)                            │   │
│  │ - PostgreSQL                                     │   │
│  │ - Nginx (反向代理)                               │   │
│  └──────────────────────────────────────────────────┘   │
└──────────────────────────────────────────────────────────┘
```

## 部署流程（详细步骤）

### 阶段1：代码提交到部署触发

```
1. 开发者提交代码到GitHub
   git push origin main

2. GitHub Actions自动触发
   - 事件：push to main branch
   - 工作流：.github/workflows/deploy.yml

3. CI/CD流程
   a. 代码拉取
      git clone <repo>
      git checkout <commit-sha>
   
   b. 依赖安装
      go mod download
   
   c. 代码构建
      go build -o app -ldflags "-X main.Version=v1.0.1"
   
   d. 单元测试
      go test -v -race ./...
   
   e. 集成测试
      docker-compose -f docker-compose.test.yml up
      go test -v -tags=integration ./...
   
   f. Docker镜像构建
      docker build -t app:v1.0.1 .
      docker tag app:v1.0.1 registry.example.com/app:v1.0.1
      docker tag app:v1.0.1 registry.example.com/app:latest
   
   g. 镜像推送
      docker push registry.example.com/app:v1.0.1
      docker push registry.example.com/app:latest
   
   h. 部署通知
      发送Webhook到自动化交付系统
      POST /api/webhooks/deploy
      {
        "image": "registry.example.com/app:v1.0.1",
        "version": "v1.0.1",
        "environment": "production"
      }
```

### 阶段2：部署执行

```
1. 自动化交付系统接收部署通知
   - 验证镜像签名
   - 检查部署权限
   - 创建部署任务

2. 部署前检查
   - 检查目标环境健康状态
   - 检查磁盘空间（>1GB）
   - 检查网络连接
   - 备份当前版本镜像

3. 部署执行
   a. 拉取新镜像
      docker pull registry.example.com/app:v1.0.1
   
   b. 停止旧容器
      docker-compose down
   
   c. 更新docker-compose.yml
      image: registry.example.com/app:v1.0.1
   
   d. 启动新容器
      docker-compose up -d
   
   e. 等待容器启动（最多30秒）
      docker ps | grep app
   
   f. 健康检查（最多10次，每次间隔2秒）
      curl -f http://localhost:8080/health || exit 1
   
   g. 记录部署日志
      - 部署开始时间
      - 部署结束时间
      - 部署状态（成功/失败）
      - 新旧版本号

4. 部署后验证
   - 检查应用日志（无ERROR）
   - 检查数据库连接
   - 检查外部服务连接
   - 运行冒烟测试

5. 部署完成
   - 发送成功通知（邮件/Slack/钉钉）
   - 更新部署历史
   - 清理旧镜像（保留最近3个版本）
```

### 阶段3：部署失败处理

```
1. 检测到部署失败
   - 健康检查失败
   - 容器启动失败
   - 数据库连接失败

2. 自动回滚
   a. 停止新容器
      docker-compose down
   
   b. 恢复旧镜像
      docker pull registry.example.com/app:v1.0.0
      docker-compose up -d
   
   c. 验证回滚成功
      curl -f http://localhost:8080/health
   
   d. 记录回滚日志
      - 回滚原因
      - 回滚时间
      - 回滚结果

3. 告警通知
   - 立即发送告警（邮件/Slack/钉钉）
   - 包含错误日志和回滚信息
   - 标记为P0事件

4. 人工介入
   - 分析失败原因
   - 修复问题
   - 重新部署
```

## 部署配置文件

### docker-compose.yml（生产环境）
```yaml
version: '3.8'

services:
  app:
    image: registry.example.com/app:latest
    container_name: delivery-app
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_NAME=delivery
      - DB_USER=postgres
      - DB_PASSWORD=${DB_PASSWORD}
      - LOG_LEVEL=info
    depends_on:
      - postgres
    restart: always
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/health"]
      interval: 10s
      timeout: 5s
      retries: 3
      start_period: 30s
    volumes:
      - ./logs:/app/logs
      - ./config:/app/config

  postgres:
    image: postgres:14-alpine
    container_name: delivery-db
    environment:
      - POSTGRES_DB=delivery
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=${DB_PASSWORD}
    volumes:
      - postgres_data:/var/lib/postgresql/data
    restart: always

  nginx:
    image: nginx:alpine
    container_name: delivery-nginx
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf:ro
      - ./ssl:/etc/nginx/ssl:ro
    depends_on:
      - app
    restart: always

volumes:
  postgres_data:
```

### .github/workflows/deploy.yml（CI/CD配置）
```yaml
name: Build and Deploy

on:
  push:
    branches: [main]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Run tests
        run: go test -v -race ./...
      
      - name: Build
        run: go build -o app -ldflags "-X main.Version=${{ github.sha }}"
      
      - name: Build Docker image
        run: |
          docker build -t app:${{ github.sha }} .
          docker tag app:${{ github.sha }} app:latest
      
      - name: Push to registry
        run: |
          echo ${{ secrets.REGISTRY_PASSWORD }} | docker login -u ${{ secrets.REGISTRY_USER }} --password-stdin
          docker push app:${{ github.sha }}
          docker push app:latest
      
      - name: Trigger deployment
        run: |
          curl -X POST http://delivery-system.example.com/api/webhooks/deploy \
            -H "Content-Type: application/json" \
            -d '{
              "image": "app:${{ github.sha }}",
              "version": "${{ github.sha }}",
              "environment": "production"
            }'
```

## 部署检查清单

### 部署前
- [ ] 所有测试通过
- [ ] 代码审查完成
- [ ] 版本号更新
- [ ] 变更日志更新
- [ ] 数据库迁移脚本准备
- [ ] 回滚方案确认

### 部署中
- [ ] 监控部署进度
- [ ] 检查日志输出
- [ ] 验证健康检查
- [ ] 确认无错误告警

### 部署后
- [ ] 验证功能正常
- [ ] 检查性能指标
- [ ] 检查错误日志
- [ ] 通知相关人员
- [ ] 更新部署记录

## 部署时间窗口

- **工作日部署**：09:00-17:00（避免夜间故障）
- **紧急修复**：24小时可部署
- **大版本发布**：周五下午前完成（便于周末监控）

## 回滚流程

### 自动回滚
```bash
# 触发条件：健康检查失败
# 执行时间：<2分钟
# 操作：恢复前一个版本镜像
docker-compose down
docker pull registry.example.com/app:v1.0.0
docker-compose up -d
```

### 手动回滚
```bash
# 命令：
curl -X POST http://delivery-system.example.com/api/tasks/{task-id}/rollback

# 参数：
{
  "target_version": "v1.0.0",
  "reason": "性能下降"
}
```

## 部署监控

### 关键指标
- 部署成功率：>99%
- 部署耗时：<10分钟
- 回滚耗时：<2分钟
- 健康检查通过率：>99%

### 告警规则
- 部署失败 → P0告警
- 部署超时（>15分钟） → P1告警
- 健康检查失败 → P0告警
- 回滚触发 → P0告警

## 状态
- [ ] TechLead确认
- [ ] 部署流程评审通过
- [ ] 签字确认
- [ ] 冻结版本

**评审日期：** ___________
**评审人：** ___________
