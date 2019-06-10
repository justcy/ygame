from php:5.6.37-fpm-alpine

RUN apk add --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/edge/community/ gnu-libiconv \
    libpng tzdata libpng-dev libmcrypt-dev libxml2 libxml2-dev libmemcached-dev zlib-dev \
	&& docker-php-ext-install pdo_mysql gd bcmath  mysqli  shmop soap socket \
	&& pecl install redis-4.0.1 memcached-2.2.0 \
	&& docker-php-ext-enable redis memcached \
	&& ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ENV LD_PRELOAD /usr/lib/preloadable_libiconv.so php
