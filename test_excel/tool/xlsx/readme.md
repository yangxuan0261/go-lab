### 配置表 Excel

网上使用到比较多的工具是 *xls_deploy_tool.py* 这个脚本

执行这个脚本需要的几个前置条件

1. 使用的 python 版本是 2.x (2.7最好)

2. protoc 生成 pb 的可执行文件, 版本是 2.5. (传送门:  https://github.com/protocolbuffers/protobuf/releases/tag/v2.5.0 )

3. xlrd, python 读取 Excel 的库, 直接使用 pip 安装:

    ```json
    $ pip install setuptools
    $ pip install xlrd
    ```

4. 生成

    ![](http://yxbl.itengshe.com/20191201205527-1.png)

    ```json
    python xls_deploy_tool.py skin f:/z_mywiki/test_script/python/excel2pb/角色表.xlsx s
    ```

    1. skin : 表格页
    2. xlsx : Excel 表
    3. s : 读取含有 S 的字段 ( c 则为 含有 C 的字段, 用来区分 服务器/客户端 )



附: 貌似有 python3 及 protoc3.x 的使用: UE4表格工具三部曲之一【环境配置】 - https://blog.csdn.net/yuxikuo_1/article/details/102693663 