FROM node:18-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install --omit=dev

COPY server.js ./
COPY src/ ./src/

RUN mkdir -p /data && chown -R node:node /data

USER node

EXPOSE 3000

CMD ["npm", "start"]
