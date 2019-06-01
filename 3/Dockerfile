FROM centos

RUN yum install -y make gcc gcc-c++ pkgconfig cmake librsvg2 librsvg2-devel git wget
RUN find / -name *svg*
#
#/var/lib/yum/yumdb/l/5d0e419b48b01cd87baf99b7f25cfea01402128e-librsvg2-devel-2.40.20-1.el7-x86_64
#/var/lib/yum/yumdb/l/b7458b2a8ec7f16f79a8fc400f8976a7fa663291-librsvg2-2.40.20-1.el7-x86_64
#/usr/lib64/girepository-1.0/Rsvg-2.0.typelib
#/usr/lib64/pkgconfig/cairo-svg.pc
#/usr/lib64/pkgconfig/librsvg-2.0.pc
#/usr/lib64/librsvg-2.so
#/usr/lib64/gdk-pixbuf-2.0/2.10.0/loaders/libpixbufloader-svg.so
#/usr/lib64/librsvg-2.so.2.40.20
#/usr/lib64/librsvg-2.so.2rsvg-2
#/usr/include/cairo/cairo-svg.h
#/usr/include/librsvg-2.0
#/usr/include/librsvg-2.0/librsvg
#/usr/include/librsvg-2.0/librsvg/rsvg-cairo.h
#/usr/include/librsvg-2.0/librsvg/librsvg-features.h
#/usr/include/librsvg-2.0/librsvg/rsvg.h
#/usr/include/librsvg-2.0/librsvg/librsvg-enum-types.h
#/usr/share/licenses/librsvg2-2.40.20
#/usr/share/mime/image/svg+xml.xml
#/usr/share/mime/image/svg+xml-compressed.xml
#/usr/share/vala/vapi/librsvg-2.0.vapi
#/usr/share/gir-1.0/Rsvg-2.0.gir
#/usr/share/thumbnailers/librsvg.thumbnailer


ENV PKG_CONFIG_PATH=/usr/local/lib64/pkgconfig
ENV LD_LIBRARY_PATH="LD_LIBRARY_PATH:/usr/local/lib64/:/usr/local/lib/"

ARG SERVICE_NAME="svg"
WORKDIR /usr/local/${SERVICE_NAME}
RUN mkdir -p /usr/local/${SERVICE_NAME}
COPY . /usr/local/${SERVICE_NAME}
#RUN cd /usr/local/lib && wget https://github.com/chikamim/go-resvg/raw/master/lib/libresvg.so
#RUN cd /usr/local/lib && wget https://github.com/chikamim/go-resvg/raw/master/lib/libresvg.dylib

## 设置下go基础环境环境
##RUN cd /usr/local && wget https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz
COPY go1.12.5.linux-amd64.tar.gz /usr/local
RUN cd /usr/local && tar -zxvf go1.12.5.linux-amd64.tar.gz
ENV GOROOT=/usr/local/go
ENV PATH=$PATH:/usr/local/go/bin

## 编译glibc-2.18
#RUN cd /usr/local && wget http://ftp.gnu.org/gnu/glibc/glibc-2.18.tar.gz
#RUN cd /usr/local && tar zxf glibc-2.18.tar.gz
#RUN cd /usr/local/glibc-2.18/ && mkdir build && cd build/ && ../configure --prefix=/usr && make -j2 && make install

RUN echo "cd /usr/local/${SERVICE_NAME}/ && go run main.go" >/tmp/start.sh
CMD ["sh", "/tmp/start.sh"]
