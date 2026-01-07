FROM node:latest AS builder

WORKDIR /app/sources/frontend

# Copy package.json and package-lock.json from frontend folder
COPY frontend/package*.json ./

RUN npm install

# Copy all frontend files
COPY frontend/. .

RUN npm run build

# Stage 2: Serve with Nginx
FROM nginx:latest

COPY --from=builder /app/sources/frontend/dist /usr/share/nginx/html
COPY nginx/nginx_frontend.conf  /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
