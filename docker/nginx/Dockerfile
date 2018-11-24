FROM nginx
RUN mkdir -p /usr/share/nginx/app
COPY app/index.json /usr/share/nginx/app/index.json
COPY app/404.json /usr/share/nginx/app/404.json
COPY app/update.json /usr/share/nginx/app/update.json
COPY nginx.conf /etc/nginx/nginx.conf
COPY audit_integral.conf /etc/nginx/conf.d/audit_integral.conf