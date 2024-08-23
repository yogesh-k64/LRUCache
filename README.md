# Cache-Sys and Cache-App Project

## Overview

This repository contains two projects:

- **Cache-Sys**: A Go-based LRU Cache backend that provides API endpoints to set, get, and delete key-value pairs with expiration support.
- **Cache-App**: A React-based frontend that interacts with the Cache-Sys backend to allow users to manage cache entries through a web interface.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Setup](#setup)

## Features

- **Cache-Sys (Go Backend)**:
  - Implements a Least Recently Used (LRU) cache.
  - Provides RESTful API endpoints for cache operations.
  - Supports setting expiration times for cache entries.
  - Concurrent processing to handle multiple requests efficiently.

- **Cache-App (React Frontend)**:
  - User interface to interact with the cache system.
  - Forms to set, get, and delete cache entries.
  - Smooth animations and transitions for a better user experience.

## Prerequisites

Before running the projects, ensure you have the following installed:

- **Go** (version 1.18 or higher)
- **Node.js** (version 16 or higher)
- **npm** or **yarn**

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/cache-project.git
cd cache-project

cd cache-sys
go mod tidy

cd ../cache-app
npm install

- Run go app

cd cache-sys
go run .

- Run React app

cd ../cache-app
npm start


API Endpoints
The Go backend provides the following API endpoints:

1. GET /get?key=<key>
Description: Retrieve the value associated with the given key.
Parameters: key (query parameter)
Response:
200 OK with JSON containing the key-value pair.
404 Not Found if the key does not exist or has expired.

2. POST /set
Description: Set a key-value pair in the cache with an optional expiration time.
Request Body:
json
Copy code
{
  "key": "<key>",
  "value": "<value>",
  "expiration": <seconds>
}
Response: 200 OK on success.

3. DELETE /delete?key=<key>
Description: Delete a key-value pair from the cache.
Parameters: key (query parameter)
Response: 200 OK on success.

Frontend UI
The React frontend provides a simple interface to interact with the cache:

Set Cache: Enter a key, value, and expiration time to store in the cache.
Get Cache: Retrieve the value associated with a given key.
Delete Cache: Delete a key-value pair from the cache.
