#!/usr/bin/env python
# -*- coding: utf-8 -*-

import os
import sys
import platform

# protoc 可执行程序所在目录都必须在环境变量中
# 生成 grpc

class CGen():
    def __init__(self, *args):
        self.path = sys.path[0]

    def do(self):
        print("--- path", self.path)
        os.chdir(self.path)
        os.system("protoc -I .\protos\ --go_out=plugins=grpc:./aaa .\protos/*.proto")
        print("--- gen success")
        pass

if __name__ == "__main__":
    ins = CGen()
    ins.do()
    sys.exit(0)
