# Mini OpenRTB Bidder (Pair Programming Exercise)

This project is a simplified OpenRTB bidder written in Go.\
It simulates a real production request flow:

HTTP Bid Request → Parsing → Filtering → Decision Server → Response →
Event Streaming 

The goal of this project is **not** to build a perfect bidder, but to
work with an intentionally simplified system that contains realistic
architectural and performance problems.

It is used as a pair-programming exercise to discuss production
readiness, concurrency, and system design.

------------------------------------------------------------------------

## Architecture Overview

    Client
      │
      ▼
    HTTP Bidder (/openrtb2/bid)
      │
      ├── Parser (OpenRTB JSON → structs)
      ├── Filter (eligibility & frequency cap)
      ├── Decision Client (external service call)
      ├── Response Builder
      └── Event Producer 

A mock **Decision Server** runs locally and simulates an external
dependency.

------------------------------------------------------------------------

## Project Structure

    cmd/bidder/            application entrypoint (composition root)
    internal/app/          HTTP wiring and handlers
    internal/openrtb/      request types
    internal/parser/       request parsing
    internal/filter/       filtering logic
    internal/decision/     decision server + client
    internal/response/     bid response builder
    internal/producer/        async event producer 


