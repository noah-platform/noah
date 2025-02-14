FROM node:20-alpine AS builder

WORKDIR /app

RUN npm install -g pnpm

COPY . .

RUN pnpm install --frozen-lockfile --ignore-scripts

RUN pnpm --filter frontend build

####################################################################################################

FROM node:20-alpine AS runner

WORKDIR /app

COPY --from=builder /app/node_modules /app/node_modules
COPY --from=builder /app/frontend/node_modules /app/frontend/node_modules
COPY --from=builder /app/frontend/build /app/frontend/build

RUN npm install -g @react-router/serve

WORKDIR /app/frontend

EXPOSE 3000

CMD ["npx", "react-router-serve", "./build/server/index.mjs"]
