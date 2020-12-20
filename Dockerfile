# STEP 1 build the executable
FROM golang:1.15 AS builder

# Create Working Directory and cd into it
WORKDIR /server

# Copy over the files
COPY . /server

# Download the packages
RUN ["go", "mod", "download"]

# Build the application
RUN ["go", "build", "."]

# STEP 2 build a small image
FROM gcr.io/distroless/base

# Create the working dir and cd into it
WORKDIR /server

COPY --from=builder /server /server

# Expose port 3000
EXPOSE 3000

# Run the application
ENTRYPOINT [ "./swapiql" ]

# Useful links
# https://github.com/GoogleContainerTools/distroless
# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324