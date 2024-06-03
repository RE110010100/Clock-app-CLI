#!/bin/bash

# Name of the Docker image
IMAGE_NAME="clock-app"

# Function to check if the Docker image exists
image_exists() {
  docker images --format '{{.Repository}}:{{.Tag}}' | grep -q "^${IMAGE_NAME}:latest$"
}

# Check if the Docker image already exists
if image_exists; then
  echo "Docker image $IMAGE_NAME already exists. Skipping build."
else
  # Build the Docker image
  echo "Building the Docker image..."
  docker build -t $IMAGE_NAME .

  # Check if the build was successful
  if [ $? -ne 0 ]; then
    echo "Docker build failed. Exiting."
    exit 1
  fi
fi

# Run the Docker container interactively
echo "Running the Docker container..."
docker run -it $IMAGE_NAME:latest

# Check if the container run was successful
if [ $? -ne 0 ]; then
  echo "Docker run failed. Exiting."
  exit 1
fi

echo "Docker container stopped."
