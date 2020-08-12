FROM node:14.4-alpine

RUN apk update && \
    npm install -g npm && \
    npm install -g @vue/cli

RUN mkdir -p /app
WORKDIR /app

COPY dist/assets /app

# RUN cd assets && yarn install && yarn serve
RUN yarn install && yarn build

# Exposing port 8080
EXPOSE 80
CMD ["yarn", "serve"]
