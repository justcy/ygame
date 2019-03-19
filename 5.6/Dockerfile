from php:5.6.37-fpm-alpine

RUN apk add --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/edge/community/ gnu-libiconv \
    libpng libpng-dev tzdata \
	&& docker-php-ext-install pdo_mysql gd \
	&& ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
	
ENV LD_PRELOAD /usr/lib/preloadable_libiconv.so php
