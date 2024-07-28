FROM node:22.5-alpine as builder
RUN apk update && apk add git
WORKDIR /app
COPY package-lock.json package.json ./
RUN npm install
COPY . .
RUN ls -a
RUN ls -a
RUN npm run docs:build

FROM nginx:alpine3.19 as host
COPY --from=builder /app/docs/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
