# image
FROM node:12-alpine as build

# update repositories
RUN apk update && apk upgrade && \
    apk add --no-cache bash make npm && \ 
    apk --no-cache add curl

WORKDIR /app
COPY package.json /app/package.json

# install libraries from package.json (only prod)
RUN npm install --only=prod

# copy app
COPY public/ /app/public/
COPY src /app/src/
COPY src/styles /app/styles
COPY .env /app

# run app
EXPOSE 3000
RUN npm run build
CMD ["npm", "start"]
