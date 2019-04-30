from node:alpine

run npm install -g docsify-cli@latest
run mkdir -p /usr/local/docsify

## Container dnvironment variables
env DEBUG 0
env PORT 3000
env DOCSIFY_VERSION latest
env NODE_VERSION alpine

## Container runtime configuration
expose 3000
workdir /usr/local/docsify

## Container entry point
entrypoint [ "docsify", "serve", "--port", "3000" ]

## Container entry point default arguments
cmd [ "." ]
