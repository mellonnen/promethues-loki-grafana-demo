# Build the frontend app.
FROM node:18.0-alpine as build-stage
# Copy source code into container.
RUN mkdir /usr/app
COPY . /usr/app
WORKDIR /usr/app
# Install dependencies.
RUN yarn
ENV PATH /usr/src/app/node_modules/.bin:$PATH
# Build app.
RUN yarn build

# Package app with nginx.
FROM nginx:alpine

# Copy contents from build container.
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=build-stage /usr/app/dist .

ENTRYPOINT ["nginx", "-g", "daemon off;"]
