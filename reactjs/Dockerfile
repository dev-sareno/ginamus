FROM node:18-alpine AS build
WORKDIR /app
COPY package.json .
COPY package-lock.json .
RUN npm ci
COPY . .
COPY .env .
RUN npm run build

FROM nginx:alpine
WORKDIR /usr/share/nginx/html
COPY --from=build /app/build .
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
