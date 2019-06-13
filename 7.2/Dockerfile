from php:7.2-fpm-alpine

RUN apk add --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/edge/community/ gnu-libiconv \
    libpng libpng-dev tzdata libmcrypt-dev libxml2 libxml2-dev libmemcached-dev zlib-dev  $PHPIZE_DEPS \
	&& docker-php-ext-install pdo_mysql gd bcmath mysqli  shmop soap sockets \
	&& pecl install redis-4.0.1 memcached-2.2.0 memcache \
	&& docker-php-ext-enable redis \
	&& docker-php-ext-enable memcached \
	&& docker-php-ext-enable memcache  \
	&& ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
	
ENV LD_PRELOAD /usr/lib/preloadable_libiconv.so php
