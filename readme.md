## 使用docker直接编译运行
## 主要用于svg转换等操作
## 运行：
````
docker build -t svg
docker run -v `pwd`:/usr/local/svg/ svg
````
最终会看到png输出的图片内容，当前执行目录会看到test.png文件