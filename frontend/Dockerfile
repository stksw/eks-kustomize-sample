FROM node:16-alpine

WORKDIR /app
COPY ./ /app
RUN yarn install
RUN yarn build

EXPOSE 3000

ENTRYPOINT ["yarn", "start"]