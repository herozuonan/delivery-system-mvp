# 开发基线文档

## 文档版本

| 版本 | 日期 | 作者 | 变更说明 | 状态 |
|------|------|------|---------|------|
| v1.0 | 2024-01-XX | Engineer | 初稿 | 待评审 |

## 基线冻结声明

本文档定义了自动化交付系统的开发基线，包括：
- 需求规格书（冻结版）
- 技术栈选定
- 系统架构设计
- 测试计划
- 部署流程
- 项目计划

**冻结时间：** 2024-01-XX  
**冻结人：** ___________  
**评审通过：** [ ] 是 [ ] 否  

## 开发启动条件检查表

### 需求确认
- [x] 需求规格书完成
- [x] MVP功能范围明确
- [x] 优先级划分清晰
- [x] 非功能需求定义
- [x] 约束条件明确
- [ ] 业务方签字确认

### 技术方案
- [x] 技术栈选定（Go + PostgreSQL + Docker）
- [x] 基础设施方案确定（单VPS + Docker Compose）
- [x] CI/CD方案确定（GitHub Actions）
- [x] 成本估算完成
- [ ] TechLead签字确认

### 架构设计
- [x] 系统架构图完成
- [x] 数据流设计完成
- [x] 部署架构设计完成
- [x] 核心模块设计完成
- [x] API接口设计完成
- [ ] 架构评审通过

### 测试计划
- [x] 测试策略定义
- [x] 测试范围明确
- [x] 测试工具选定
- [x] 测试环境设计
- [x] 测试时间预估
- [ ] 测试计划评审通过

### 部署流程
- [x] 部署架构设计
- [x] 部署流程文档
- [x] 部署配置文件模板
- [x] 回滚流程设计
- [x] 部署检查清单
- [ ] 部署流程评审通过

### 项目计划
- [x] 工作量评估
- [x] 时间表制定
- [x] 里程碑定义
- [x] 风险评估
- [x] 资源需求
- [ ] 项目计划评审通过

### 开发环境
- [ ] VPS购买完成
- [ ] Docker环境搭建
- [ ] PostgreSQL安装配置
- [ ] Docker Registry搭建
- [ ] Git仓库创建
- [ ] CI/CD流程配置
- [ ] 开发工具安装

### 文档完善
- [x] 需求规格书
- [x] 技术栈决策文档
- [x] 系统架构设计文档
- [x] 测试计划
- [x] 部署流程文档
- [x] 项目计划书
- [x] 开发基线文档
- [ ] API文档（待开发）
- [ ] 运维手册（待开发）
- [ ] 用户手册（待开发）

### 评审和批准
- [ ] 需求评审通过
- [ ] 架构评审通过
- [ ] 测试计划评审通过
- [ ] 部署流程评审通过
- [ ] 项目计划评审通过
- [ ] 所有利益相关方确认
- [ ] 开发启动批准

## 版本控制和变更管理

### 版本控制策略
- **主分支（main）**：生产环境代码，受保护
- **开发分支（develop）**：开发环境代码，集成分支
- **功能分支（feature/xxx）**：单个功能开发
- **修复分支（bugfix/xxx）**：bug修复
- **发布分支（release/xxx）**：发布准备

### 分支管理规则
```
main (生产)
  ↑
  └─ release/v1.0.0 (发布准备)
       ↑
       └─ develop (开发集成)
            ↑
            ├─ feature/webhook (功能开发)
            ├─ feature/workflow (功能开发)
            ├─ feature/deploy (功能开发)
            └─ bugfix/xxx (bug修复)
```

### 提交规范
```
<type>(<scope>): <subject>

<body>

<footer>

type: feat|fix|docs|style|refactor|test|chore
scope: 模块名称
subject: 简短描述（50字以内）
body: 详细描述（可选）
footer: 关联issue（可选）

示例：
feat(webhook): add GitHub webhook receiver

- Implement webhook signature verification
- Add event parsing and validation
- Support push and pull_request events

Closes #123
```

### 变更控制流程

```
需求变更申请
    ↓
评估工作量和风险
    ↓
TechLead审批
    ↓
[批准] → 更新基线文档 → 调整计划
[拒绝] → 记录原因 → 后续迭代
```

### 变更申请模板
```
变更标题：[功能名称]
变更类型：[新增/修改/删除]
优先级：[P0/P1/P2]
工作量：[X小时]
风险：[低/中/高]
理由：[变更原因]
影响范围：[受影响的模块]
```

## 代码规范

### Go代码规范
- 遵循 [Effective Go](https://golang.org/doc/effective_go)
- 使用 `gofmt` 格式化代码
- 使用 `golint` 检查代码风格
- 使用 `go vet` 检查代码问题
- 函数注释必须以函数名开头
- 包注释必须以包名开头

### 代码审查规则
- 所有代码必须通过代码审查
- 至少1个reviewer批准
- 所有测试必须通过
- 代码覆盖率不能下降
- 无lint警告

### 提交前检查
```bash
# 格式化代码
go fmt ./...

# 检查代码风格
golint ./...

# 检查代码问题
go vet ./...

# 运行测试
go test -v -race ./...

# 生成覆盖率报告
go test -coverprofile=coverage.txt ./...
```

## 文件结构

```
delivery-system/
├── docs/                          # 文档
│   ├── 01-REQUIREMENTS.md        # 需求规格书
│   ├── 02-TECH_STACK.md          # 技术栈决策
│   ├── 03-ARCHITECTURE.md        # 系统架构
│   ├── 04-TEST_PLAN.md           # 测试计划
│   ├── 05-DEPLOYMENT.md          # 部署流程
│   ├── 06-PROJECT_PLAN.md        # 项目计划
│   └── 07-BASELINE.md            # 开发基线（本文件）
├── cmd/                           # 可执行程序
│   └── delivery/
│       └── main.go               # 主程序入口
├── internal/                      # 内部包
│   ├── api/                      # API层
│   │   ├── handler.go
│   │   └── middleware.go
│   ├── service/                  # 业务逻辑层
│   │   ├── workflow.go
│   │   ├── task.go
│   │   ├── deploy.go
│   │   └── monitor.go
│   ├── repository/               # 数据访问层
│   │   ├── workflow.go
│   │   ├── task.go
│   │   └── deployment.go
│   ├── executor/                 # 执行层
│   │   ├── builder.go
│   │   ├── tester.go
│   │   └── deployer.go
│   ├── config/                   # 配置管理
│   │   └── config.go
│   └── logger/                   # 日志管理
│       └── logger.go
├── pkg/                           # 公共包
│   ├── models/                   # 数据模型
│   ├── utils/                    # 工具函数
│   └── errors/                   # 错误定义
├── test/                          # 测试文件
│   ├── unit/
│   ├── integration/
│   └── e2e/
├── scripts/                       # 脚本
│   ├── setup.sh                  # 环境搭建
│   ├── deploy.sh                 # 部署脚本
│   └── rollback.sh               # 回滚脚本
├── config/                        # 配置文件
│   ├── docker-compose.yml
│   ├── docker-compose.test.yml
│   └── nginx.conf
├── .github/                       # GitHub配置
│   └── workflows/
│       └── deploy.yml            # CI/CD流程
├── Dockerfile                     # Docker镜像定义
├── go.mod                         # Go模块定义
├── go.sum                         # Go依赖锁定
├── Makefile                       # 构建脚本
├── README.md                      # 项目说明
└── .gitignore                     # Git忽略规则
```

## 开发工作流

### 日常开发流程

```
1. 从develop分支创建功能分支
   git checkout develop
   git pull origin develop
   git checkout -b feature/xxx

2. 开发功能
   - 编写代码
   - 编写测试
   - 运行测试
   - 代码审查

3. 提交代码
   git add .
   git commit -m "feat(scope): description"
   git push origin feature/xxx

4. 创建Pull Request
   - 填写PR描述
   - 关联issue
   - 等待审查

5. 代码审查
   - 至少1个reviewer批准
   - 所有测试通过
   - 无lint警告

6. 合并到develop
   git checkout develop
   git pull origin develop
   git merge --no-ff feature/xxx
   git push origin develop

7. 删除功能分支
   git branch -d feature/xxx
   git push origin --delete feature/xxx
```

### 发布流程

```
1. 从develop创建发布分支
   git checkout develop
   git pull origin develop
   git checkout -b release/v1.0.0

2. 准备发布
   - 更新版本号
   - 更新变更日志
   - 运行全量测试
   - 性能测试

3. 发布到main
   git checkout main
   git pull origin main
   git merge --no-ff release/v1.0.0
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin main
   git push origin v1.0.0

4. 合并回develop
   git checkout develop
   git pull origin develop
   git merge --no-ff release/v1.0.0
   git push origin develop

5. 删除发布分支
   git branch -d release/v1.0.0
   git push origin --delete release/v1.0.0
```

## 开发启动清单

### 第一天
- [ ] 克隆代码仓库
- [ ] 安装开发工具（Go、Docker、PostgreSQL）
- [ ] 运行本地开发环境
- [ ] 运行测试验证环境
- [ ] 阅读项目文档

### 第一周
- [ ] 完成基础设施搭建
- [ ] 创建数据库和表结构
- [ ] 实现Webhook接收器
- [ ] 实现工作流引擎基础框架
- [ ] 编写单元测试

### 第二周
- [ ] 实现任务调度器
- [ ] 实现状态机管理
- [ ] 实现代码拉取模块
- [ ] 实现构建执行器
- [ ] 编写集成测试

### 第三周
- [ ] 实现测试执行器
- [ ] 实现Docker镜像构建
- [ ] 实现部署管理器
- [ ] 实现回滚机制
- [ ] 编写E2E测试

### 第四周
- [ ] 实现日志收集
- [ ] 实现告警规则引擎
- [ ] 实现通知发送器
- [ ] 实现RESTful API
- [ ] 完成测试和优化

### 第五周
- [ ] 生产环境配置
- [ ] 部署脚本准备
- [ ] 监控配置
- [ ] 备份验证
- [ ] 回滚演练

### 第六周
- [ ] 上线执行
- [ ] 上线验证
- [ ] 文档完善
- [ ] 知识转移

## 沟通和协作

### 日常沟通
- **站会**：每天10:00（15分钟）
  - 昨天完成了什么
  - 今天计划做什么
  - 遇到什么阻塞

- **周会**：每周五16:00（30分钟）
  - 周进度总结
  - 下周计划
  - 风险评估

### 文档更新
- 每周更新项目进度
- 每个里程碑完成后更新基线文档
- 重大决策记录在案

### 问题上报
- 遇到技术难题立即上报
- 工作量评估有误立即调整
- 风险识别立即通知

## 成功标准

### 开发完成标准
- [ ] 所有P0功能开发完成
- [ ] 单元测试覆盖率>80%
- [ ] 集成测试覆盖率>60%
- [ ] 无P0/P1级别bug
- [ ] 代码审查通过
- [ ] 文档完善

### 上线成功标准
- [ ] 生产环境部署成功
- [ ] 系统稳定运行24小时
- [ ] 监控告警正常工作
- [ ] 备份和回滚机制验证通过
- [ ] 用户反馈良好

## 附录

### 常用命令

```bash
# 本地开发
make dev              # 启动本地开发环境
make test             # 运行所有测试
make test-unit        # 运行单元测试
make test-integration # 运行集成测试
make test-e2e         # 运行E2E测试
make coverage         # 生成覆盖率报告
make lint             # 代码检查
make build            # 构建二进制
make docker-build     # 构建Docker镜像
make docker-push      # 推送Docker镜像
make deploy           # 部署到生产
make rollback         # 回滚到前一版本
```

### 常用链接
- GitHub仓库：https://github.com/xxx/delivery-system
- Docker Registry：https://registry.example.com
- 项目文档：https://docs.example.com
- 监控面板：https://monitor.example.com

## 签字确认

**TechLead：** ___________  日期：___________

**业务方：** ___________  日期：___________

**Engineer：** ___________  日期：___________

---

**文档状态：** 待评审  
**最后更新：** 2024-01-XX  
**下次评审：** 开发启动前
