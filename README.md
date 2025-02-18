
# Dailbot

Dailbot is an AI-powered caller assistant designed to handle inbound and outbound calls for businesses. The goal is to automate customer interactions, reducing the manual effort required for businesses to follow up with customers.

## Project Overview

This repository contains the web server implementation of Dailbot, written in Golang using the Fiber web framework. The web server integrates with Airtable and provides an API for managing webhooks, bases, and tables.

### Inspiration

The idea for Dailbot came from observing a business owner who had to manually call customers after they booked an appointment. This system simplifies that process by automatically triggering calls based on form submissions, ensuring timely and efficient customer engagement.

## Features Implemented

- **Airtable OAuth Integration**: Establishes authentication and authorization for interacting with Airtable.
- **Webhook Management**: Creates and manages webhooks for Airtable bases and tables.
- **In-Memory Caching**: Implements an in-memory cache for performance optimization.
- **Redis Caching**: Integrates Redis to manage persistent cache data.
- **Fiber-based Web Server**: Uses the Fiber framework for handling HTTP requests.
- **AI-Driven Call Manager**: Triggers conversation flow using AI text-to-speech (TTS), transcription, and large language models (LLMs).

## Features To Be Added

- **Event Queue System**: Implement an event queue using RabbitMQ to handle message processing.
- **Caller Server**: A dedicated service to manage call routing.
- **Voice Server**: Handles voice interactions and synthesis.
- **Analytics Service**: Provides insights and reports on call interactions.

## Getting Started

### Prerequisites
Ensure you have the following installed:

- [Go](https://go.dev/) (latest version recommended)
- [Redis](https://redis.io/)

### Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/emperorsixpacks/Dailbot-server.git
   cd Dailbot-server
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Run the server:
   ```sh
   go run src/main.go
   ```

## Contributing
Contributions are welcome! Feel free to submit issues or pull requests to improve the project.

## License
This project is licensed under the MIT License.

## Contact
For any inquiries or support, feel free to reach out via the repository's issue tracker.

