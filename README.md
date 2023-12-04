# Crypto Price Alert Terminal App

## Project Goal

This project aims to create a terminal application that allows users to track real-time price changes of selected cryptocurrencies. Users can specify a set of cryptocurrencies and define threshold percentages for price fluctuations. When the price change of a selected cryptocurrency crosses the specified threshold, the user is notified through the terminal.

## Technologies

- **Language**: Go (Golang)
- **Messaging**: RabbitMQ
- **API**: CoinMarketCap API
- **Architecture**: Event-driven architecture

## Folder Structure

- **/api/repository**: Contains logic for interacting with the CoinMarketCap API.
    - `coinmarketcap.go`: Handles API requests.
    - `types.go`: Defines structures for API response data.

- **/cmd**: Contains the main entry points.
    - `/consumer`:
        - `main.go`: Entry point for the consumer application.
    - `main.go`: Primary entry point for the terminal application.

- **/pkg/events**: Encapsulates the core event handling logic.
    - `event_dispatcher.go`: Manages event dispatching.
    - `event_handlers.go`: Processes different event types.
    - `events.go`: Defines event structures.
    - `interfaces.go`: Specifies interfaces for events and handlers.

- **/rabbitmq**: Dedicated to RabbitMQ interaction.
    - `rabbitmq.go`: Handles channel opening and message consumption.

## Usage

1. **Setting Up**: Instructions for setting up the Go environment, RabbitMQ, and any dependencies.
2. **Running the Application**: Step-by-step guide to running the application and configuring the environment.
3. **User Interaction**: How to choose cryptocurrencies, set thresholds, and receive alerts.

## Architecture

- **Event-Driven Design**: Description of the flow of events from receiving price updates to user notifications.
- **RabbitMQ**: Explanation of RabbitMQ's role in decoupling data fetching and event processing.
- **CoinMarketCap API**: Overview of interactions with the CoinMarketCap API for real-time price data.
- **Event Handling**: Discussion of event handlers' roles in processing price updates and threshold crossings.

## To-Do List

- [x] Implement API calls to fetch real-time price data from CoinMarketCap.
- [ ] Set up RabbitMQ for messaging.
- [ ] Design `PriceUpdateEvent` and `ThresholdCrossedEvent` structures.
- [ ] Develop the event dispatcher to manage event processing.
- [ ] Implement event handlers for `PriceUpdateEvent` and `ThresholdCrossedEvent`.
- [ ] Create a user interface in the terminal for selecting coins and setting thresholds.
- [ ] Implement logic for users to manage and store their preferences.
- [ ] Schedule regular price data fetching from the API.
- [ ] Develop logic to compare fetched prices with user-defined thresholds.
- [ ] Integrate event handling with RabbitMQ for efficient message processing.
- [ ] Conduct thorough testing to ensure functional accuracy and reliability.
- [ ] Address any issues or enhancements identified during testing.

