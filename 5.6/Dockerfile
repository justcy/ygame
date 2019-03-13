from php:5.6.37-fpm-alpine

RUN apk add -U --no-cache \
    && docker-php-ext-install pdo_mysql

RUN apk add --no-cache --repository http://dl-3.alpinelinux.org/alpine/edge/testing gnu-libiconv
ENV LD_PRELOAD /usr/lib/preloadable_libiconv.so php