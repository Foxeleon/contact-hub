# Stage 1: Build the React app
FROM node:18-alpine AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

COPY . .

# Set the API URL during the build
ARG VITE_API_URL_PERSONS=http://localhost:8080/persons
ENV VITE_API_URL_PERSONS=${VITE_API_URL_PERSONS}

RUN npm run build

# Stage 2: Serve the static files with Nginx
FROM nginx:stable-alpine

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]