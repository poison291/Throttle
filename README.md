# Go Token Bucket Rate Limiter (Learning Project)

This project is a simple **in-memory rate limiter** implemented in Go using the **Token Bucket algorithm**.  
It is built as a learning exercise to understand rate limiting, basic security controls, and time-based logic in backend systems.

No frameworks, no external services — just core Go.

---

## Features

- Token Bucket rate limiting
- Time-based token refill
- Request blocking after exceeding limits
- Temporary user lock using `time.Time`
- Middleware-style request control
- Real API request simulation

---

## How It Works

### Token Bucket
- Each request consumes one token
- Tokens refill over time based on a refill rate
- Requests are rejected when no tokens are available

### Abuse Protection
- Repeated blocked requests increase a counter
- After a threshold, the user is temporarily locked
- Lock duration is enforced using time comparison

---

## Project Structure

- `tokenBucket`  
  Holds rate limiting and lock state

- `refil()`  
  Refills tokens based on elapsed time

- `AllowRequest()`  
  Determines whether a request should pass

- `middleware()`  
  Acts as a gatekeeper before executing a request

- `apiFetch()`  
  Makes a real HTTP request for testing purposes

---

## Usage

```bash
go run bucket.go
