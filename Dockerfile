# Base Image
FROM frolvlad/alpine-glibc

# Copy log folders
COPY log log

# Copy environment
COPY environment environment

# Add bash shell in image
RUN apk add --no-cache bash

# Copy main app to image
ADD main /

# Expose port apps
EXPOSE 8080

# Run app when run image to container
CMD ["/main"]
