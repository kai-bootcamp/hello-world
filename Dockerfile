FROM golang:1.17.4-alpine3.15 AS builder

ENV GIN_MODE=release

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY ./backend/go.mod .
COPY ./backend/go.sum .

RUN go mod download

COPY ./backend/src ./src
RUN go build -o ./app ./src/main.go

FROM alpine:latest as backend

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir -p /api

WORKDIR /api

COPY --from=builder /api/app .

EXPOSE 4000

ENTRYPOINT ["./app"]

FROM node:14-alpine AS deps
# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.

RUN apk add --no-cache libc6-compat
WORKDIR /app

COPY frontend/package.json frontend/yarn.lock ./
RUN yarn install --frozen-lockfile

FROM node:14-alpine AS node_builder

WORKDIR /app
COPY ./frontend .
COPY --from=deps /app/node_modules ./node_modules
RUN yarn build

FROM node:14-alpine AS frontend
WORKDIR /app

ENV NODE_ENV production

RUN addgroup -g 1001 -S nodejs
RUN adduser -S nextjs -u 1001

COPY --from=node_builder /app/public ./public
COPY --from=node_builder --chown=nextjs:nodejs /app/.next ./.next
COPY --from=node_builder /app/node_modules ./node_modules
COPY --from=node_builder /app/package.json ./package.json

USER nextjs

EXPOSE 3000

ENV PORT 3000

CMD ["node_modules/.bin/next", "start"]

FROM nginx
COPY ./nginx/nginx.conf /etc/nginx/conf.d/configfile.template

ENV PORT 8080
ENV HOST 0.0.0.0

EXPOSE 8080

CMD sh -c "envsubst '\$PORT' < /etc/nginx/conf.d/configfile.template > /etc/nginx/conf.d/default.conf && nginx -g 'daemon off;'"
