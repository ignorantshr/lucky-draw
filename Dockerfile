FROM node:16 AS front
WORKDIR /front
COPY lucky-draw-ui .
RUN npm install
RUN npm run build

FROM golang:1.16-alpine AS backend
WORKDIR /backend
COPY . .
ENV GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 go build -o lucky-draw

FROM nginx

ENV mysql_host='localhost'
ENV mysql_port=3306
ENV mysql_db='lucky_draw'
ENV mysql_user='root'
ENV mysql_passwd=''

COPY conf/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=front /front/dist /usr/share/nginx/html/luckydraw

RUN mkdir conf
COPY conf/app.conf conf/app.conf
COPY --from=backend /backend/lucky-draw lucky-draw
RUN chmod +x lucky-draw

ENTRYPOINT service nginx start && ./lucky-draw
