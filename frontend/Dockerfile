FROM node:21.2 AS builder
WORKDIR /app
COPY package* ./
RUN npm install
COPY public ./public
COPY src ./src
COPY .env.production ./
RUN npm run build


FROM nginx:alpine
COPY --from=builder /app/build /usr/share/nginx/html
