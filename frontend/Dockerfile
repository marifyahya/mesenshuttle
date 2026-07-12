# Stage 1: Build
FROM node:20-alpine AS builder

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy the rest of the application
COPY . .

# Build the Nuxt application
RUN npm run build

# Stage 2: Production
FROM node:20-alpine AS runner

WORKDIR /app

# Set environment variables for production
ENV NODE_ENV=production
ENV HOST=0.0.0.0
ENV PORT=3000

# Copy only the necessary files from builder stage
COPY --from=builder /app/.output /app/.output

# Expose port 3000
EXPOSE 3000

# Use a non-root user for security
RUN addgroup -S nuxtgroup && adduser -S nuxtuser -G nuxtgroup
USER nuxtuser

# Run the built server
CMD ["node", ".output/server/index.mjs"]
