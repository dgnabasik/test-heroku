# Dockerfile.react
FROM node:12-alpine as dev
WORKDIR /root/
CMD echo "getting latest npm modules" \
    && yarn install \
    && echo "installation complete, starting" \
    && yarn start

FROM node:12-alpine as build
WORKDIR /root
COPY ./package.json ./package.json
RUN yarn install
COPY . ./ 
RUN yarn run build -o app

# multistage build of container that is deployed to Heroku.
FROM nginx:alpine as prod
EXPOSE 80 443
COPY --from=build /root/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
