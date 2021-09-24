FROM nginx

ENV mysql_host='localhost'
ENV mysql_port=3306
ENV mysql_user='root'
ENV mysql_passwd=''

COPY lucky-draw-ui/dist /usr/share/nginx/html/luckydraw
COPY conf/nginx.conf /etc/nginx/conf.d/default.conf

RUN mkdir conf
COPY conf/app.conf conf/app.conf
ADD lucky-draw lucky-draw
RUN chmod +x lucky-draw

ENTRYPOINT service nginx start && ./lucky-draw
