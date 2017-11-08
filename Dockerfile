FROM golang:1.9.2-alpine

ENV PATH /api:$PATH

ARG API_PORT
ENV PORT $API_PORT

WORKDIR /api

# Copy binary demo
ADD ./api-linux-amd64 /api/demo

# Modified files for titan
RUN chmod 555 /api/demo

# Expose ports
EXPOSE $PORT

# Run Titan
CMD demo -port $PORT
