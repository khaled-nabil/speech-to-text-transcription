FROM node:22-alpine

RUN apk --no-cache add curl

WORKDIR /app

COPY package.json package-lock.json* ./

RUN npm ci

RUN npm i -g serve

COPY . .

RUN npm run build

EXPOSE 3000

CMD [ "serve", "-s", "dist" ]
