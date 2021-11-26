FROM golang:1.17

# Ignore APT warnings about not having a TTY
ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && \
    apt-get -q install -y wget build-essential pkg-config --no-install-recommends \
                               libjpeg-dev libpng-dev libtiff-dev \
                               libgif-dev libx11-dev --no-install-recommends

ENV IMAGEMAGICK_VERSION=6.9.9-35

RUN cd && \
	wget https://github.com/ImageMagick/ImageMagick6/archive/${IMAGEMAGICK_VERSION}.tar.gz && \
	tar xvzf ${IMAGEMAGICK_VERSION}.tar.gz && \
	cd ImageMagick* && \
	./configure \
	    --without-magick-plus-plus \
	    --without-perl \
	    --disable-openmp \
	    --with-gvc=no \
	    --disable-docs && \
	make -j$(nproc) && make install && \
	ldconfig /usr/local/lib

WORKDIR /go/lgtm-generator
COPY . .

RUN go install
CMD /go/bin/lgtm-generator
