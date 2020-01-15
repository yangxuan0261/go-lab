#!/usr/bin/env python
# -*- coding: utf-8 -*-

import os
import sys
import traceback
import glob
import utils as utils

SelfPath = sys.path[0]
XlsToolDir = os.path.join(SelfPath, "xlsx")
XlsToolPath = os.path.join(XlsToolDir, "xls_deploy_tool.py")

TempSheet = "skin"
TempXls = os.path.join(SelfPath, "../角色表.xlsx")
TempFlag = "s"  # c: 客户端使用, s: 服务器使用

class Excel2Pb:
    "excel 转成 protobuf 和 二进制数据文件"

    def __init__(self):
        self.name = ""

    def genPbBin(self):
        os.chdir(XlsToolDir)

        xlsPath = TempXls
        sheetName = TempSheet
        flag = TempFlag

        if not os.path.exists(xlsPath):
            raise Exception("--- xls path dont exist, path:%s" % xlsPath)

        cmd = "python %s %s %s %s" % (XlsToolPath, sheetName, xlsPath, flag)
        print("--- cmd:", cmd)
        utils.execute(cmd)

        pass

    def genGo(self):
        self.genPbBin()

        os.chdir(XlsToolDir)

        # 生成 xxx.pb.go
        os.system("protoc -I . --go_out=. ./*.proto")

        # 移动文件到目的文件夹
        DstGoPbDir = os.path.join(SelfPath, "../gen")

        utils.moveFiles(XlsToolDir, DstGoPbDir, ["*.pb.go"])
        utils.moveFiles(XlsToolDir, DstGoPbDir, ["*.bytes"])

        # 清除 多余文件
        utils.removeFiles(
            XlsToolDir, ["*_pb2.py", "*.pyc", "*.log", "*.txt", "*.proto"])

    def genCSharp(self):
        pass

    def genLua(self):
        pass


if __name__ == "__main__":
    try:
        ins = Excel2Pb()
        ins.genGo()
    except Exception as e:
        print("--- fail, err:", e)
    else:
        print("--- success")
