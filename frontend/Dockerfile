# Decidi usar alpine poqeue s liviana y simple para este ejemplo. 
FROM nginx:alpine

# copiar en la imagen el html de frontend
COPY index.html /usr/share/nginx/html/index.html

# HARDENING & CREACIÓN DE USUARIO 'ARADIA'
RUN adduser -D -u 1000 aradia

# Ajustar Nginx para que no use Root
RUN chown -R aradia:aradia /var/cache/nginx \
    && chown -R aradia:aradia /var/log/nginx \
    && chown -R aradia:aradia /etc/nginx/conf.d

# Cambiamos el puerto a 8080 (los puertos < 1024 son solo para root)
RUN sed -i '/user  nginx;/d' /etc/nginx/nginx.conf \
    && sed -i 's,/run/nginx.pid,/tmp/nginx.pid,g' /etc/nginx/nginx.conf \
    && sed -i 's,/var/run/nginx.pid,/tmp/nginx.pid,g' /etc/nginx/nginx.conf \
    && sed -i 's/listen       80;/listen       8080;/g' /etc/nginx/conf.d/default.conf

# aradia puede administrar y escribir logs
RUN chown -R aradia:aradia /var/cache/nginx /var/log/nginx /etc/nginx/conf.d /usr/share/nginx/html

# levantar el usuario administrador
USER aradia

#puerto
EXPOSE 8080

# levantar nginx
CMD ["nginx", "-g", "daemon off;"]



