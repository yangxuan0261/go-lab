#!/usr/bin/env python
# -*- coding: utf-8 -*-
import os
import sys
import stat
import shutil
import glob
import urllib
import zipfile
from plistlib import *
from hashlib import md5
import json
import gzip
import tarfile
import traceback

try:
    import xml.etree.cElementTree as ET
except ImportError:
    import xml.etree.ElementTree as ET

FILE_MODE = stat.S_IRUSR | stat.S_IWUSR | stat.S_IRGRP | stat.S_IROTH

SelfPath = os.path.abspath(os.path.dirname(__file__))
ProjPath = os.path.join(SelfPath, "..", "..")

def getProjPath():
    return ProjPath

def execute(cmd):
    status = os.system(cmd)
    status >>= 8
    if status != 0:
        print("\n")
        print("system execute '%s' return %d" % (cmd, status))
        traceback.print_stack()
        raise Exception("os.system fail, cmd:%s" % cmd)
    return status


def svnUpdata(url, destDir):
    if os.path.exists(destDir):
        curPath = os.getcwd()
        os.chdir(destDir)
        execute("svn up")
        os.chdir(curPath)
    else:
        cmd = "svn checkout %s %s" % (url, destDir)
        execute(cmd)


def getFiles(path, ignoreDirs=[".svn", ".git"], ignoreFiles=[".meta"]):
    fileList = []
    for root, dirs, files in os.walk(path):
        root = root.replace("\\", "/")
        index = root.rfind("/")
        rootName = root[index+1:]
        if rootName in ignoreDirs:
            break

        for name in files:
            extName = os.path.splitext(name)[-1]
            if not (extName in ignoreFiles):
                fileList.append(os.path.join(root, name).replace("\\", "/"))
    return fileList


def chmod(fileName, mode=FILE_MODE):
    if os.path.isfile(fileName):
        os.chmod(fileName, mode)
        return

    for root, dirs, files in os.walk(fileName):
        for f in files:
            f = os.path.join(root, f)
            os.chmod(f, mode)


def removeFile(fileName):
    fileName = os.path.abspath(fileName)
    if not os.path.exists(fileName):
        print("File: %s is not existed" % fileName)
        return

    chmod(fileName)
    if os.path.isfile(fileName):
        os.remove(fileName)
    elif os.path.isdir(fileName):
        print(fileName)
        shutil.rmtree(fileName)
    else:
        print("removeFile error: %s" % filePath)


def copyFile(srcFile, destFile, ignores=[".svn", ".git", ".get_date.dat", "Thumbs.dao"]):
    srcFile = os.path.abspath(srcFile)
    destFile = os.path.abspath(destFile)
    if not os.path.exists(srcFile):
        return

    if not os.path.exists(os.path.dirname(destFile)):
        os.makedirs(os.path.dirname(destFile))

    if os.path.exists(destFile):
        removeFile(destFile)

    if os.path.isfile(srcFile):
        shutil.copyfile(srcFile, destFile)
    elif os.path.isdir(srcFile):
        ignores = shutil.ignore_patterns(*ignores)
        shutil.copytree(srcFile, destFile, ignore=ignores)

    chmod(destFile)


def copyFiles(srcDir, dstDir, flags):
    for flag in flags:
        for file in glob.glob(os.path.join(srcDir, flag)):
            name = os.path.basename(file)
            dstFile = os.path.join(dstDir, name)
            copyFile(file, dstFile)


def moveFiles(srcDir, dstDir, flags):
    for flag in flags:
        for file in glob.glob(os.path.join(srcDir, flag)):
            name = os.path.basename(file)
            dstFile = os.path.join(dstDir, name)
            shutil.move(file, dstFile)


def removeFiles(srcDir, flags):
    for flag in flags:
        for file in glob.glob(os.path.join(srcDir, flag)):
            removeFile(file)


# isContainSrcDir true:src文件夹整个拷过去, false: 只拷里面的文件
def copyDir(srcDir, destDir, isContainSrcDir=True, ignoreDirs=[".svn", ".git"], ignoreFiles=[".meta"]):
    if isContainSrcDir:
        basename = os.path.basename(srcDir)
        destDir = os.path.join(destDir, basename)

    for f in getFiles(srcDir, ignoreDirs, ignoreFiles):
        # print("--- file:%s" % f)
        copyFile(f, f.replace(srcDir, destDir))

# 压缩 srcDir 目录 成 dstFile


def zipDir(dstFile, srcDir, mode=zipfile.ZIP_DEFLATED, isIncludeParent=True):
    with zipfile.ZipFile(dstFile, 'w', mode) as zip:
        pDir = isIncludeParent and os.path.dirname(srcDir) or srcDir
        # print("--- pDir:"+ pDir)
        for root, dirs, files in os.walk(srcDir):
            # print("--- root:" + root)
            for file in files:
                absDir = os.path.join(root, file)
                relativeDir = absDir.replace(pDir, "")
                zip.write(absDir, relativeDir)

# 解压缩 filename 到指定目录


def unzip(filename, destDir=None):
    print("start to unzip %s" % filename)
    destDir = destDir or os.path.dirname(filename)
    zfile = zipfile.ZipFile(filename, "r")

    def handle(f):
        dest = os.path.join(destDir, f)
        dPath = os.path.dirname(dest)
        if not os.path.exists(dPath):
            os.makedirs(dPath)
        with open(dest, "wb") as fd:
            fd.write(zfile.read(f))
    map(handle, zfile.namelist())


def md5File(filename):
    m = md5()
    with open(filename, "rb") as fd:
        m.update(fd.read())
    return m.hexdigest()


def get_tmp_path(src_path):
    tmp = 1000
    while True:
        tmp_path = os.path.join(src_path, "%s" % tmp)
        if not os.path.exists(tmp_path):
            break
        tmp = tmp + 1
    return tmp_path


class SVNCommands(object):

    @classmethod
    def export(cls, url, dest_dir):
        execute("svn export %s %s" % (url, dest_dir))

    @classmethod
    def checkout(cls, url, dest_dir):
        execute("svn co %s %s" % (url, dest_dir))


def get_android_version_code(xmlPath):
    tree = ET.ElementTree(file=xmlPath)
    xmlns = "http://schemas.android.com/apk/res/android"
    attName = "{%s}versionCode" % (xmlns)
    versionCode = tree.getroot().attrib[attName]
    return versionCode


def get_android_version_name(xmlPath):
    tree = ET.ElementTree(file=xmlPath)
    xmlns = "http://schemas.android.com/apk/res/android"
    attName = "{%s}versionName" % (xmlns)
    versionName = tree.getroot().attrib[attName]
    return versionName


def get_ios_version_name(xmlPath):
    plist = Plist.fromFile(xmlPath)
    return plist.CFBundleShortVersionString


def showAndGetMenus(menus, prompt=None):
    prompt = prompt or "please choose one item:"
    while True:
        print(prompt)
        for k, v in enumerate(menus):
            print("%s. %s" % (k + 1, v))

        index = raw_input(">>")
        try:
            index = int(index)
        except ValueError:
            pass

        if index in range(1, len(menus) + 1):
            return menus[index - 1]
    return None


def download(url, filename):
    def callback(blocknum, blocksize, totalsize):
        percent = 100.0 * blocknum * blocksize / totalsize
        if percent > 100:
            percent = 100
        # sys.stdout.write("downloading:  %.2f%% %s\r" % (percent, "#" * int(percent)))
        sys.stdout.write("downloading:  %.2f%%\r" % percent)
        sys.stdout.flush()
    print("start to download: %s" % url)
    urllib.urlretrieve(url, filename, callback)
    print("finished...")


def loadJson(path):
    with open(path, "r") as fd:
        content = fd.read()
    return json.loads(content)


# 一次性打包整个根目录。空子目录会被打包。
# 如果只打包不压缩，将"w:gz"参数改为"w:"或"w"即可。
def make_targz(output_filename, source_dir, mod="w:gz"):
    with tarfile.open(output_filename, mod) as tar:
        tar.add(source_dir, arcname=os.path.basename(source_dir))

# 逐个添加文件打包，未打包空子目录。可过滤文件。
# 如果只打包不压缩，将"w:gz"参数改为"w:"或"w"即可。


def make_targz_one_by_one(output_filename, source_dir, mod="w:gz"):
    tar = tarfile.open(output_filename, mod)
    for root, dir, files in os.walk(source_dir):
        for file in files:
            pathfile = os.path.join(root, file)
            tar.add(pathfile)
    tar.close()


def del_empty_dir(path, isDel=False, isDelMeta=True):
    if not os.path.isdir(path):
        return

    foundList = []
    for root, dirs, files in os.walk(path):
        if len(dirs) == 0 and len(files) == 0:
            foundList.append(root)

    for dir in foundList:
        print(("--- need del:" + dir))
        if isDel:
            os.rmdir(dir)

        if isDelMeta:
            metaPath = dir + ".meta"
            if os.path.isfile(metaPath):
                # print("--- exit meta:" + metaPath)
                os.remove(metaPath)

# 递归修改文件名，去空格


def del_space_of_file(path):
    class CItem():
        def __init__(self, dir, src, dst):
            self.dir = dir
            self.src = src
            self.dst = dst

    srcFlag = " "
    dstFlag = "_"

    def hasSpace(str):
        if srcFlag in str:
            retStr = str.replace(srcFlag, dstFlag)
            return retStr

    def rename(srcFile, destFile):
        shutil.copyfile(srcFile, destFile)
        os.remove(srcFile)
        pass

    fndList = []
    for root, dirs, files in os.walk(path):
        for name in files:
            dstName = hasSpace(name)
            if dstName != None:
                item = CItem(root, name, dstName)
                fndList.append(item)

        # for name in dirs: # 暂时不做目录变更
        #     hasSpace(name)
        #     pass

    for item in fndList:
        print("--- found: dir (%s), from (%s) to (%s)" %
              (item.dir, item.src, item.dst))
        rename(os.path.join(item.dir, item.src),
               os.path.join(item.dir, item.dst))

# 读取文件


def readFile(filePath):
    contentStr = ""
    with open(filePath, "rb") as fd:
        contentStr = fd.read()
    return contentStr.decode('utf-8')

# 读取文件, 并返回每一行


def readFileLines(filePath):
    contentStr = readFile(filePath)
    return contentStr.split("\n")


def writeFile(filePath, str, mod="wb"):
    # a+ 为追加
    with open(filePath, mod) as fd:
        fd.write(str.encode('utf-8'))
    pass


if __name__ == "__main__":
    pass
