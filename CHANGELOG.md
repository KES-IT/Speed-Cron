# Changelog

## Release 0.4.3 - 2023-07-27

### <!-- 1 -->🐛 Bug Fixes

- 修复了一个在for中使用defer可能造成内存泄漏的问题 ([8c3c11f](https://github.com/hamster1963/Speed-Cron/commit/8c3c11fb2a5792a5737b0ae2a16041b039aeb627))
- 修复了更新与测速缓存锁没有被正确释放的问题 ([d0888a8](https://github.com/hamster1963/Speed-Cron/commit/d0888a8c0fd277d0d4ef304cd0439cac9ee4a559))

## Release 0.4.2 - 2023-07-24

### <!-- 1 -->🐛 Bug Fixes

- 修复多站点测速模块的内存泄露问题 ([bfeda95](https://github.com/hamster1963/Speed-Cron/commit/bfeda95568bef69850bf80ac3ab12c3b22d6201b))

## Release 0.4.1 - 2023-07-21

### <!-- 0 -->⛰️  Features

- 增加多站点测速模块 ([a8f4680](https://github.com/hamster1963/Speed-Cron/commit/a8f4680664c79b8dc8196344a9ed02cd12e95dae))
- 采用注入版本号的形式防止后端地址泄露 ([9733f92](https://github.com/hamster1963/Speed-Cron/commit/9733f92b7de5ec64cbd02ae8113f60591f3ca856))
- 增加更新通道，根据通道获取不同的更新 ([3e8d667](https://github.com/hamster1963/Speed-Cron/commit/3e8d6671820a3ea9c2c8d88e50d901e94d1d5404))

### <!-- 1 -->🐛 Bug Fixes

- 测速HTTP客户端换为G.Client并设置10s超时 ([a8f4680](https://github.com/hamster1963/Speed-Cron/commit/a8f4680664c79b8dc8196344a9ed02cd12e95dae))
- 修复推送延迟时错误为空情况下空指针的问题 ([024dd15](https://github.com/hamster1963/Speed-Cron/commit/024dd15fd487cd413e080bde32634ac761205ef6))
- 修复无法正确获取后端地址的问题 ([9733f92](https://github.com/hamster1963/Speed-Cron/commit/9733f92b7de5ec64cbd02ae8113f60591f3ca856))
- 新增下载文件名判断，避免下载到错误md5文件 ([3e8d667](https://github.com/hamster1963/Speed-Cron/commit/3e8d6671820a3ea9c2c8d88e50d901e94d1d5404))
- 修复更新通道获取失败导致的更新失败 ([eb5489b](https://github.com/hamster1963/Speed-Cron/commit/eb5489b477cd4f06647a2ee006e5c96377d10285))
- 增加缺失的后端地址 ([a4704c1](https://github.com/hamster1963/Speed-Cron/commit/a4704c1c1db2ba804e1712258158e3cc6f0d26f3))
- 修复了需要传入变量的注释错误 ([1ca07ed](https://github.com/hamster1963/Speed-Cron/commit/1ca07ed9e5aaf91049ca94e063b70c917761f6a1))

### <!-- 2 -->🚜 Refactor

- 重构测速功能 ([a8f4680](https://github.com/hamster1963/Speed-Cron/commit/a8f4680664c79b8dc8196344a9ed02cd12e95dae))
- 重构了启动任务的执行逻辑 ([badf60e](https://github.com/hamster1963/Speed-Cron/commit/badf60e2c9ca2a8b62edd5dd612b7f222346adcc))
- 重构了获取最新Release信息模块 ([a6d6e35](https://github.com/hamster1963/Speed-Cron/commit/a6d6e356925a35e0d68a6be868e42a12c11602f8))
- 现在版本与后端地址的默认值为unknown ([1ca07ed](https://github.com/hamster1963/Speed-Cron/commit/1ca07ed9e5aaf91049ca94e063b70c917761f6a1))

## Release 0.3.1 - 2023-07-17

### <!-- 0 -->⛰️  Features

- Changelog增加Updates字段 ([e01e6e3](https://github.com/hamster1963/Speed-Cron/commit/e01e6e38fc97114dc157f1b9553d2306973327bf))

### <!-- 1 -->🐛 Bug Fixes

- 在命令行处理前再次检查更新核心状态，避免出现冲突 ([5965d65](https://github.com/hamster1963/Speed-Cron/commit/5965d65afc1f2e52dde41bfd576fa8eaab4be231))
- 直接使用go build进行编译测试 ([b900615](https://github.com/hamster1963/Speed-Cron/commit/b900615cfca8c362207777393f388845d22f3d29))
- 修复启动测试服务与自动更新服务冲突的bug ([124b718](https://github.com/hamster1963/Speed-Cron/commit/124b718eb6d6fdbe023672a0cb8100543c145d42))

### <!-- 2 -->🚜 Refactor

- 重构后端接口地址定义格式 ([4f92f7a](https://github.com/hamster1963/Speed-Cron/commit/4f92f7abb4f7ab849dc11af3a6a2fde23bdc9d06))
- 使用单独g_structs文件夹存放全局结构体 ([a1069da](https://github.com/hamster1963/Speed-Cron/commit/a1069da9097731639a1dc320a30e952f7bf84ef4))
- 使用统一的缓存键管理 ([a1069da](https://github.com/hamster1963/Speed-Cron/commit/a1069da9097731639a1dc320a30e952f7bf84ef4))

### <!-- 9 -->🔼 Updates

- 测速命令行处理间隔变为500毫秒 ([fe3c620](https://github.com/hamster1963/Speed-Cron/commit/fe3c620f9f0f1711e09da95a478584bee19e3a79))
- 测速状态缓存时间改为不过期 ([515380d](https://github.com/hamster1963/Speed-Cron/commit/515380dab55a14abe88d805bcaa916ec5a408669))
- 自动更新服务时间间隔改为20s ([78cf61a](https://github.com/hamster1963/Speed-Cron/commit/78cf61a51fab83f8d7b1d501dd4127830833cada))
- 更新完成后等待重启时间改为1秒 ([25f5155](https://github.com/hamster1963/Speed-Cron/commit/25f515575d130197dcb1f97da82eddff5af7881a))

## Release 0.3.0 - 2023-07-15

### <!-- 1 -->🐛 Bug Fixes

- 精简了没用的注释 ([ee72dab](https://github.com/hamster1963/Speed-Cron/commit/ee72dab73e09437251b080d384058645df79392f))

### <!-- 3 -->📚 Documentation

- 新增了一些逻辑注释 ([b83ce97](https://github.com/hamster1963/Speed-Cron/commit/b83ce97592ab47a7950078d444bf7916cb5294ae))

### <!-- 4 -->⚡ Performance

- InitData不再进行类型转换 ([b83ce97](https://github.com/hamster1963/Speed-Cron/commit/b83ce97592ab47a7950078d444bf7916cb5294ae))
- 精简了版本比较的分支逻辑 ([b83ce97](https://github.com/hamster1963/Speed-Cron/commit/b83ce97592ab47a7950078d444bf7916cb5294ae))

## Release 0.2.7 - 2023-07-15

### <!-- 0 -->⛰️  Features

- 初次启动增加HTTP延迟测试服务 ([08022fd](https://github.com/hamster1963/Speed-Cron/commit/08022fd087aa5efeabe23cc6a86574ab20e20e67))

## Release 0.2.6 - 2023-07-15

### <!-- 1 -->🐛 Bug Fixes

- 新增测速前检测更新核心状态判断，避免更新过程中开始测速造成下次启动失败。 ([8660965](https://github.com/hamster1963/Speed-Cron/commit/86609653b8877f90f26ff0ec2d5bea2388421661))

## Release 0.2.5 - 2023-07-15

### <!-- 1 -->🐛 Bug Fixes

- 修复404报错，获取版本前进行判断是否存在二进制包以供下载 ([0228354](https://github.com/hamster1963/Speed-Cron/commit/022835418defacb9a5de13ec2ae0b9fd211253fb))

## Release 0.2.4 - 2023-07-15

### <!-- 1 -->🐛 Bug Fixes

- 自动更新模块下载后增加文件大小判断 ([7890531](https://github.com/hamster1963/Speed-Cron/commit/789053117cab5450b4753d7f91488b28a6aa71e7))
- 修复入口描述文字错误的问题 ([3767dc9](https://github.com/hamster1963/Speed-Cron/commit/3767dc97594e898147aad46c7021a314514a5f9c))
- 增加下载完成后的等待时间 ([6897480](https://github.com/hamster1963/Speed-Cron/commit/6897480e390e34bcc154f049d636440984335535))

### <!-- 3 -->📚 Documentation

- 增加一些注释 ([9c33857](https://github.com/hamster1963/Speed-Cron/commit/9c338571cc839a33d89f98992d6d6914cee379c8))

### <!-- 4 -->⚡ Performance

- 精简一些变量初始化 ([636fe4f](https://github.com/hamster1963/Speed-Cron/commit/636fe4fe78d33331ae92a5b89441b2a387920601))
- 精简了未用到的单次测速 ([9a30820](https://github.com/hamster1963/Speed-Cron/commit/9a3082063d0c1c87adf1e0746629c709b0c8e017))

## Release 0.2.3 - 2023-07-14

### <!-- 7 -->⚙️ Miscellaneous Tasks

- 将Go版本降级为1.18，尽量降低杀毒软件报毒风险 ([4d4c5a3](https://github.com/hamster1963/Speed-Cron/commit/4d4c5a315aee1b5910b64147eb07c842343fd90e))

## Release 0.2.2 - 2023-07-14

### <!-- 1 -->🐛 Bug Fixes

- 修复下载错误未正确捕获的问题 ([a1830aa](https://github.com/hamster1963/Speed-Cron/commit/a1830aae226ebd430f59bb1ee3ee0cff9ac4d5e5))

## Release 0.2.1 - 2023-07-14

### <!-- 0 -->⛰️  Features

- 现在设备会自动上报版本号了 ([4405593](https://github.com/hamster1963/Speed-Cron/commit/440559377665962118e975db2ec96301d8c44c1d))

### <!-- 6 -->🧪 Testing

- 将获取更新间隔变为20s ([4405593](https://github.com/hamster1963/Speed-Cron/commit/440559377665962118e975db2ec96301d8c44c1d))

## Release 0.2.0 - 2023-07-14

### <!-- 0 -->⛰️  Features

- 增加定时检测更新定时任务 ([27c2e8e](https://github.com/hamster1963/Speed-Cron/commit/27c2e8e790ced8cbe7e4b28236a781836bc38586))
- 增加自动更新模块 ([27c2e8e](https://github.com/hamster1963/Speed-Cron/commit/27c2e8e790ced8cbe7e4b28236a781836bc38586))

### <!-- 1 -->🐛 Bug Fixes

- 精简了无用的函数注释 ([469df67](https://github.com/hamster1963/Speed-Cron/commit/469df67a50a7485d771ff4a7d43aa087305d16c0))

### <!-- 2 -->🚜 Refactor

- 新增测速状态缓存标识 ([27c2e8e](https://github.com/hamster1963/Speed-Cron/commit/27c2e8e790ced8cbe7e4b28236a781836bc38586))
- 重构了初始化数据的结构，采用结构体，以指针形式传递。 ([61b03c1](https://github.com/hamster1963/Speed-Cron/commit/61b03c1c3cde8d2db79343bfc41448673e4aa008))

## Release 0.1.7 - 2023-07-13

### <!-- 7 -->⚙️ Miscellaneous Tasks

- Bump golang.org/x/net ([2490488](https://github.com/hamster1963/Speed-Cron/commit/2490488b27f875ab35cf559fb9339c64edcf662a))
- Bump golang.org/x/net from 0.0.0-20211123202848-9e5a29745d54 to 0.7.0 ([16ace9e](https://github.com/hamster1963/Speed-Cron/commit/16ace9e8adcce9d0000e57d25902c77460e3859f))

### Signed-off-by

- Dependabot[bot] <support@github.com> ([2490488](https://github.com/hamster1963/Speed-Cron/commit/2490488b27f875ab35cf559fb9339c64edcf662a))

## Release 0.1.6 - 2023-07-13

### <!-- 0 -->⛰️  Features

- 测速流程结束后退出speedtest CLI ([3dcd2de](https://github.com/hamster1963/Speed-Cron/commit/3dcd2debd3a315c26b7d24e5b642759ee2d5348a))

### <!-- 1 -->🐛 Bug Fixes

- 精简了未用到的配置文件 ([97eb700](https://github.com/hamster1963/Speed-Cron/commit/97eb700f969bc3cab05623900fb0351c5743ac46))
- 不再对配置进行打包，避免gf cli打包出错 ([785cb6d](https://github.com/hamster1963/Speed-Cron/commit/785cb6d74207e9e5a140e420386f886b9b223959))

### <!-- 3 -->📚 Documentation

- README.md中增加了一张演示截图 ([f60fa62](https://github.com/hamster1963/Speed-Cron/commit/f60fa62f1a7304798f91382f826bc246d7ab8593))
- 增加了注释，删去了不必要的参数注释 ([ed559e8](https://github.com/hamster1963/Speed-Cron/commit/ed559e84519a80f76928fdfa7e6fb6c7dda6f160))

## Release 0.1.5 - 2023-07-13

### <!-- 1 -->🐛 Bug Fixes

- 修复发布正式版时错误获取最近标签为beta的错误，会获取最近的正式版进行比对 ([c372d7f](https://github.com/hamster1963/Speed-Cron/commit/c372d7fd1f7bb330bfc0cbd292e92194921634e9))

### <!-- 7 -->⚙️ Miscellaneous Tasks

- 修改changelog中commit的排序方式 ([c372d7f](https://github.com/hamster1963/Speed-Cron/commit/c372d7fd1f7bb330bfc0cbd292e92194921634e9))

## Release 0.1.4 - 2023-07-13

### <!-- 0 -->⛰️  Features

- 增加解析多行commit内容 ([bc1f1c3](https://github.com/hamster1963/Speed-Cron/commit/bc1f1c349e42529593a6f98d292309ac32f2d1c5))

### <!-- 1 -->🐛 Bug Fixes

- 修复多行commit会造成无法正常生成changelog的问题 ([bc1f1c3](https://github.com/hamster1963/Speed-Cron/commit/bc1f1c349e42529593a6f98d292309ac32f2d1c5))

### <!-- 6 -->🧪 Testing

- 测试跳过beta标签 ([d72ad0c](https://github.com/hamster1963/Speed-Cron/commit/d72ad0c8a634d719478ac102feb4c1669c974bfe))

## Release 0.1.3-rc3 - 2023-07-13

### <!-- 1 -->🐛 Bug Fixes

- 精简了go-release Actions配置文件中未用到的binary_name选项 ([4ed489c](https://github.com/hamster1963/Speed-Cron/commit/4ed489c50bd7c07e5d45d63deba78150ca065f07))

## Release 0.1.3-rc2 - 2023-07-13

### <!-- 1 -->🐛 Bug Fixes

- 修改了cmd包获取版本的变量名 ([8f0d1ba](https://github.com/hamster1963/Speed-Cron/commit/8f0d1ba320d4b14d39af2087d3c9137748cc5758))

## Release 0.1.3-rc1 - 2023-07-13

### <!-- 6 -->🧪 Testing

- 通过GitHub Actions注入版本号 ([4539659](https://github.com/hamster1963/Speed-Cron/commit/453965961680e9effc1528277a503b776ca6946e))

## Release 0.1.2 - 2023-07-12

### <!-- 0 -->⛰️  Features

- 增加快速获取最新tag的脚本 ([237fe43](https://github.com/hamster1963/Speed-Cron/commit/237fe430cb9a1205c6467b4bda3add7eb2b03815))
- 增加自动生成changelog GitHub Actions ([cd2af63](https://github.com/hamster1963/Speed-Cron/commit/cd2af63c8d524a39f9999b75cf109b1e2c7f1c8c))

## Release 0.1.1 - 2023-07-11

### <!-- 1 -->🐛 Bug Fixes

- 精简了go-release生成架构的设置 ([822203a](https://github.com/hamster1963/Speed-Cron/commit/822203a3905d136ef06ef91f9b31083e878fac8f))
- 精简了客户端代码文件 ([5d571c2](https://github.com/hamster1963/Speed-Cron/commit/5d571c27562346ca015849db20e485b4bcde3ff3))

### <!-- 3 -->📚 Documentation

- README.md增加本地测试参数提示 ([3ca7fe5](https://github.com/hamster1963/Speed-Cron/commit/3ca7fe536fd9f3495c988ac44ef1a9da7c775a53))

### <!-- 9 -->🔼 Updates

- 升级版本到v0.1.0 ([cbc81d4](https://github.com/hamster1963/Speed-Cron/commit/cbc81d430df985e6bfb9fe1a0c046a6990f36269))

## Release 0.0.5 - 2023-07-11

### <!-- 1 -->🐛 Bug Fixes

- 修正版本号输出不正确的问题 ([f4486d3](https://github.com/hamster1963/Speed-Cron/commit/f4486d3b75ebafb1ce349d622bd34dd838eb4787))
- 修正版本号为v0.0.5 ([f4486d3](https://github.com/hamster1963/Speed-Cron/commit/f4486d3b75ebafb1ce349d622bd34dd838eb4787))

## Release 0.0.4 - 2023-07-11

### <!-- 0 -->⛰️  Features

- 增加version指令获取当前版本 ([40274ee](https://github.com/hamster1963/Speed-Cron/commit/40274eecc3de73e03dd7d224373f8369bdd4a72e))

### <!-- 1 -->🐛 Bug Fixes

- 修正版本号为v0.0.4 ([fa73f40](https://github.com/hamster1963/Speed-Cron/commit/fa73f408b985e26f271373f11cb8176b4eb2cd70))

## Release 0.0.3 - 2023-07-10

### <!-- 1 -->🐛 Bug Fixes

- 目前只生成Windows Amd64架构二进制文件 ([8cb3501](https://github.com/hamster1963/Speed-Cron/commit/8cb35013f21c55afa50b8ce0b7f7adf9c7500b95))
- 修改生成的文件名与格式 ([8cb3501](https://github.com/hamster1963/Speed-Cron/commit/8cb35013f21c55afa50b8ce0b7f7adf9c7500b95))

## Release 0.0.2 - 2023-07-10

### <!-- 0 -->⛰️  Features

- 增加Go编译测试与release二进制Actions ([19d16d9](https://github.com/hamster1963/Speed-Cron/commit/19d16d9e5d7874587fe0b4b7e6300f8e390e5461))

### <!-- 1 -->🐛 Bug Fixes

- 修正模板为定时任务后台服务 ([caaf521](https://github.com/hamster1963/Speed-Cron/commit/caaf521978f753a60c474c7ae1433ad3a05e2018))
- 清理初始化进度条冗余代码 ([74f5d6e](https://github.com/hamster1963/Speed-Cron/commit/74f5d6e359a188c6f2b14e62de0bd0f9c580e8fe))
- 在闭包中进行http客户端错误处理 ([e1ea5e6](https://github.com/hamster1963/Speed-Cron/commit/e1ea5e61acf4d603ecd8107940acdcef736bd281))
- 采用glog替换原本的标准输出 ([bdac5f1](https://github.com/hamster1963/Speed-Cron/commit/bdac5f180bba6b88db551b6a877193f8dc0c7f18))
- 去除ip网段限制，可能有获取不到真正内网ip的风险 ([e0dfe0a](https://github.com/hamster1963/Speed-Cron/commit/e0dfe0aac4958aa409a9f628e29a1b8f5a6e0326))

### <!-- 3 -->📚 Documentation

- 新增了readme文档 ([1b2d5d4](https://github.com/hamster1963/Speed-Cron/commit/1b2d5d4e9b838edc3a2e1db821cb389afe355d1c))

### <!-- 7 -->⚙️ Miscellaneous Tasks

- 精简了go模块依赖 ([9390a12](https://github.com/hamster1963/Speed-Cron/commit/9390a12e4718f3a1d38f983aac39977fc86e0d79))

### <!-- 9 -->🔼 Updates

- 完成测速定时任务 ([6fd441b](https://github.com/hamster1963/Speed-Cron/commit/6fd441bf8267b904d055e4f3050cb14be8a99cbb))
- 完成延迟检测定时任务 ([6fd441b](https://github.com/hamster1963/Speed-Cron/commit/6fd441bf8267b904d055e4f3050cb14be8a99cbb))
- 更新环境为正式环境 ([9d2d8ff](https://github.com/hamster1963/Speed-Cron/commit/9d2d8ffb58de86260a293166d44c6425dfaa52d5))
- 改为单应用架构 ([01eb86c](https://github.com/hamster1963/Speed-Cron/commit/01eb86cd94412a7d4d6fd116713b99ed17ad334f))
- 新增设备认证 ([01eb86c](https://github.com/hamster1963/Speed-Cron/commit/01eb86cd94412a7d4d6fd116713b99ed17ad334f))
- 新增配置获取 ([01eb86c](https://github.com/hamster1963/Speed-Cron/commit/01eb86cd94412a7d4d6fd116713b99ed17ad334f))
- 新增定时任务运控 ([01eb86c](https://github.com/hamster1963/Speed-Cron/commit/01eb86cd94412a7d4d6fd116713b99ed17ad334f))
- 精简配置 ([3891dfd](https://github.com/hamster1963/Speed-Cron/commit/3891dfd9afd2ba509caf1f97f789c3f88d65e303))

### Delete

- 移除全局函数与中间件 ([8d51260](https://github.com/hamster1963/Speed-Cron/commit/8d51260d93d7636cba35291c0b8fa1ffda040458))
- 删除了开发产生的编译文件 ([0bc734d](https://github.com/hamster1963/Speed-Cron/commit/0bc734dd8330dc03aac1a07a3f95bd3058ac3c1a))

