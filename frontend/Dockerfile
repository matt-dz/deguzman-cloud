# Stage 1: Build the SvelteKit application
FROM --platform=$BUILDPLATFORM node:latest AS build
ARG VITE_AUTH_BACKEND_URL
ENV ADAPTER=node
ENV VITE_BASE_URL=https://deguzman.cloud
ENV VITE_AUTH_URL=https://auth.deguzman.cloud
ENV VITE_AUTH_BACKEND_URL=${VITE_AUTH_BACKEND_URL}
WORKDIR /app
COPY package*.json ./
COPY . .

RUN npm ci
RUN npm run build
RUN npm prune --production

# Stage 2: Serve the application
FROM node:current-alpine
WORKDIR /app
COPY --from=build /app/build ./build/
COPY --from=build /app/package.json ./package.json
COPY --from=build /app/node_modules ./node_modules
EXPOSE 3000

# Start the SvelteKit application
CMD ["node", "build"]
