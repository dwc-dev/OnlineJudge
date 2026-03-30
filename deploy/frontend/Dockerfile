FROM node:24-alpine AS builder

WORKDIR /app

RUN corepack enable

COPY frontend/package.json frontend/pnpm-lock.yaml ./
RUN --mount=type=cache,id=pnpm-store,target=/root/.local/share/pnpm/store pnpm install --frozen-lockfile

COPY frontend/ ./

ARG VITE_APP_API_BASE_URL=/api/v1
ENV VITE_APP_API_BASE_URL=${VITE_APP_API_BASE_URL}

RUN pnpm build

FROM nginx:1.28-alpine

COPY deploy/frontend/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
