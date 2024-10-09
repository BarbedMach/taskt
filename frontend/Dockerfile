FROM nginx:alpine

COPY ./html /usr/share/nginx/html
COPY ./css /usr/share/nginx/css
COPY ./script /usr/share/nginx/script
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80