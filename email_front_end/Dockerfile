# build stage
FROM node:lts-alpine as build-stage
WORKDIR /app
COPY /app/package*.json .
RUN ls -la && pwd && cat package.json && sleep 10s
COPY /app/ .
RUN npm install
#RUN ls -la && pwd && sleep 30s
RUN npm run build

# production stage
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]