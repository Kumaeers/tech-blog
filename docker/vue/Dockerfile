FROM node:14.4-alpine

RUN apk update && \
    npm install -g npm && \
    npm install -g @vue/cli

WORKDIR /app

COPY ./assets .

RUN yarn install

CMD ["yarn","serve"]
