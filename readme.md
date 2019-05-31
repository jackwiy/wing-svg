下载go1.12.5.linux-amd64.tar.gz放到项目根目录，运行：
````
docker build -t svg
docker run svg
````
最终会看到png输出的图片内容
````
FROM centos

RUN yum install -y make gcc gcc-c++ pkgconfig cmake librsvg2 librsvg2-devel git
ENV PKG_CONFIG_PATH=/usr/local/lib64/pkgconfig
ENV LD_LIBRARY_PATH="LD_LIBRARY_PATH:/usr/local/lib64/:/usr/local/lib/"

ARG SERVICE_NAME="svg"
WORKDIR /usr/local/${SERVICE_NAME}
RUN mkdir -p /usr/local/${SERVICE_NAME}
COPY . /usr/local/${SERVICE_NAME}
COPY lib/libresvg.dylib /usr/local/lib
COPY lib/libresvg.so /usr/local/lib

## 设置下go基础环境环境
COPY go1.12.5.linux-amd64.tar.gz /usr/local
RUN cd /usr/local && tar -zxvf go1.12.5.linux-amd64.tar.gz
ENV GOROOT=/usr/local/go
ENV PATH=$PATH:/usr/local/go/bin
RUN go get github.com/ungerik/go-cairo
RUN go get github.com/chikamim/go-resvg

## 编译glibc-2.18
COPY glibc-2.18.tar.gz /usr/local
RUN cd /usr/local && tar zxf glibc-2.18.tar.gz
RUN cd /usr/local/glibc-2.18/ && mkdir build && cd build/ && ../configure --prefix=/usr && make -j2 && make install

RUN echo "cd /usr/local/${SERVICE_NAME}/ && go run main.go" >/tmp/start.sh

CMD ["sh", "/tmp/start.sh"]
````