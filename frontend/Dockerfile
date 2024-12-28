# Stage 01 - Build
FROM node:22.5.0-alpine AS builder

WORKDIR /app

COPY package.json pnpm-lock.yaml ./

RUN npm install -g pnpm

RUN pnpm i

COPY . /app

RUN pnpm run build

# Stage 02 - Serve with Nginx
FROM nginx:alpine

# Copy built files from the builder stage to Nginx's default directory
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose port 80 for serving the app
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]
