# KES定时测速客户端
[![build-go-binary](https://github.com/KES-IT/Speed-Cron/actions/workflows/go-release.yml/badge.svg?branch=main)](https://github.com/KES-IT/Speed-Cron/actions/workflows/go-release.yml)
## 1. 项目介绍
本项目是一个基于Speed CLI的KES定时测速客户端，可以在指定时间段内定时测速并将测速结果发送到后端进行数据汇总。
![](https://s2.loli.net/2023/07/13/gnMZJQStCkoEXNK.png)
## 2. 项目开发
安装GF CLI工具
```bash
go mod tidy
```
```bash
gf run main.go --department=IT --name=KhalilFong
```
## 3. 项目部署
```bash
gf build main.go
```
