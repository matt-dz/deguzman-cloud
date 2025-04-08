FROM nginx:alpine
RUN mkdir -p /usr/share/nginx/redirect
COPY index.html /usr/share/nginx/html/index.html
COPY redirect.html /usr/share/nginx/html/redirect.html
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
