FROM node:alpine

WORKDIR /app/
RUN apk update
RUN apk add --virtual builds-deps build-base python

copy . .
RUN npm config set python /usr/bin/python
WORKDIR /app/sec_backend
RUN npm install
WORKDIR /app


CMD ["node", "/app/index.js"]