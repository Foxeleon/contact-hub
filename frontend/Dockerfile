# Stage 1: Build the React app
FROM node:22-alpine AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

COPY . .

ARG VITE_API_URL_PERSONS=/api/persons
ENV VITE_API_URL_PERSONS=${VITE_API_URL_PERSONS}

RUN npm run build

# Stage 2: Serve the static files with Nginx
FROM nginx:stable-alpine

COPY --from=builder /app/dist /usr/share/nginx/html

RUN rm /etc/nginx/conf.d/default.conf

COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]