FROM node:18.14-alpine
# specify tag

ENV PROJECT_NAME frontend

# define frontend workdir
RUN mkdir $PROJECT_NAME
WORKDIR /$PROJECT_NAME

RUN apk update && \
    apk upgrade && \
    apk add vim && \
    apk add git && \
    apk add --no-cache curl \
        && curl -sL https://unpkg.com/@pnpm/self-installer | node

# install nuxt
# CMD ["sh", "-c",  "cd ../ && pnpm dlx nuxi init --force ${PROJECT_NAME}"]
