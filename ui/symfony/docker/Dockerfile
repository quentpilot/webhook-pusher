FROM php:8.3-cli

RUN apt-get update && apt-get install -y \
    git unzip libpq-dev zlib1g-dev libzip-dev \
    && docker-php-ext-install pdo pdo_pgsql zip

COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

WORKDIR /var/www/symfony

COPY . .

ENV COMPOSER_ALLOW_SUPERUSER=1

RUN composer install --no-interaction --prefer-dist

EXPOSE 8001