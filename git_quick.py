#!/usr/bin/env python
# -*- coding: utf-8 -*-

import os
import sys

SelfPath = sys.path[0]

if __name__ == "__main__":
    envName = "My_Python"
    pyDir = os.environ.get(envName)  # 从环境变量中获取
    if pyDir == None:
        raise Exception("Error: 配置 python 工具目录 环境变量 {}".format(envName))

    syncTool = os.path.join(pyDir, "tool/git_sync.py")

    if not os.path.exists(syncTool):
        raise Exception("Error: 找不到工具脚本: {}".format(syncTool))

    os.system("python3 {} {}".format(syncTool, SelfPath))
    os.system("pause")
    sys.exit(0)
