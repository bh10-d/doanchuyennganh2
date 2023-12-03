FROM node:21

WORKDIR /node-microservice

COPY /nodejs/package*.json .

RUN npm install

COPY /nodejs .

EXPOSE 9090

CMD [ "node", "index.js" ]