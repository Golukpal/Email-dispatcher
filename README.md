# 🚀 Concurrent Email Campaign Dispatcher

A high-performance, Mailchimp-like bulk email dispatcher built with **Golang**. This project demonstrates the power of Go's native concurrency model by utilizing **Goroutines**, **Channels**, and the **Producer-Consumer architecture** to process and send thousands of templated emails efficiently.

---

## ✨ Features

- **Blazing Fast Concurrency:** Distributes email dispatching tasks across multiple worker threads (Goroutines) running in parallel.
- **Producer-Consumer Pattern:** Uses unbuffered channels for safe, non-blocking data hand-offs between the CSV reader and email senders.
- **Memory Efficient:** Streams CSV records directly into channels instead of loading the entire massive dataset into memory at once.
- **Safe Synchronization:** Utilizes `sync.WaitGroup` to ensure the main program waits gracefully for all background workers to finish their tasks.

---

Tech Stack
Language: Go (Golang)

Concurrency: Goroutines, Channels

Testing/Environment: Docker,

