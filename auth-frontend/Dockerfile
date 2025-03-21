# Stage 1: Build the SvelteKit application
FROM --platform=$BUILDPLATFORM node:latest AS build
WORKDIR /app
COPY package*.json ./
COPY . .
ENV ADAPTER=node
RUN npm ci
RUN npm run build
RUN npm prune --production

# Stage 2: Serve the application
FROM node:current-alpine
WORKDIR /app
COPY --from=build /app/build ./build/
COPY --from=build /app/package.json ./package.json
COPY --from=build /app/node_modules ./node_modules
ENV NODE_ENV=production
EXPOSE 3000

# Start the SvelteKit application
CMD ["node", "build"]
