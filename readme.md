# 验证码登录测试
1. 入口：cmd/main.go
1. 前端目录：frontend
1. 后端处理逻辑：internal/controllers
1. 配置文件示例：configs/app.ini.example 实际使用过程中，自己修改对应配置

#部署说明
部署文件已经打包到cmd/captcha.zip,下载解压后，

1. configs/app.ini.example为app.ini并添加对应配置
1. 安装nginx（yum -y install nginx），拷贝配置文件configs/captcha.conf /etc/nginx/conf.d/captcha.conf，启动nginx
1. 拷贝静态文件frontend到nginx配置的根目录的captcha目录下
mkdir /usr/share/nginx/html/captcha;<br/>
cp -r frontend/* /usr/share/nginx/html/captcha/*<br/>
并修改index.html中data-appid为自己的captcha appid
