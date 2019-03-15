from php:7.2-fpm-alpine

RUN apk add --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/edge/community/ gnu-libiconv \
	libpng libpng-devel \
	&& docker-php-ext-install pdo_mysql gd
	
ENV LD_PRELOAD /usr/lib/preloadable_libiconv.so php