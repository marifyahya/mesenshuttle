FROM node:20-slim AS base

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

FROM base AS builder
RUN npm run build

FROM node:20-slim AS runner

WORKDIR /app

ENV NODE_ENV=production
ENV HOST=0.0.0.0
ENV PORT=3000

COPY --from=builder /app/.output /app/.output

EXPOSE 3000

RUN groupadd -r nuxtgroup && useradd -r -g nuxtgroup nuxtuser
USER nuxtuser

CMD ["node", ".output/server/index.mjs"]
