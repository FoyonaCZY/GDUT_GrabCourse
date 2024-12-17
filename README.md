# 使用教程

## 运行程序

在 Releases 中下载打包好的 exe 文件并运行，或运行 main.go 源代码。

## 获取cookies

![image-20231227163820149](image-20231227163820149.png)

浏览器登录教务系统，按下f12，进入“网络”模块，勾选“保留日志”，随后刷新网页。在左侧找到login!welcome.action，点击，在“标头”处找到cookie:JSESSIONID=xxxxxxxx，将JSESSIONID=xxxxxxxx复制下来，粘贴到程序中运行。
