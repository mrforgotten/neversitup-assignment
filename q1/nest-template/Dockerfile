# 1. Base
FROM node:20-alpine AS base

WORKDIR /app
RUN corepack enable
RUN corepack prepare yarn@stable --activate

COPY package.json yarn.lock .yarnrc.yml ./

# Install project dependencies using Yarn Berry
RUN yarn install --immutable

# 2. Builder
FROM base AS build

# Copy the rest of the application code
COPY . .

# Build the application
RUN yarn build

# 3. runner
FROM base AS run

WORKDIR /app

# Copy only the necessary files from the build stage
COPY --from=build /app/package.json /app/yarn.lock /app/.yarn /app/.yarnrc.yml ./
COPY --from=build /app/dist ./dist

# Install production dependencies using Yarn Berry
RUN yarn workspaces focus --production

# Expose the application port
EXPOSE 3000

# Set the default command to run the application
CMD ["node", "dist/main"]
