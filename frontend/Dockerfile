# Stage 1: Build the React app
FROM node:22-alpine AS builder

WORKDIR /app

# Copy the package first.json for caching dependencies
COPY package.json package-lock.json ./
RUN npm install

COPY . .

RUN npm run build

# Stage 2: Serve the static files with Nginx
FROM nginx:stable-alpine

# Copying the compiled application
COPY --from=builder /app/dist /usr/share/nginx/html

# Replacing the default Nginx config with our own
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]