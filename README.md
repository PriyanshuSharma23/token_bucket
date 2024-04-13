# Token Bucket Implementation

**NOT MEANT FOR PRODUCTION USE**

This repository contains an implementation of a token bucket in Go. A token bucket is a rate limiting algorithm used in networking and distributed systems to control the rate of traffic or requests.

## Implementation Details

The Bucket struct implements the token bucket algorithm. It tracks the current size of the bucket, its capacity, and the rate at which tokens are added to the bucket. Tokens are added to the bucket at a constant rate (r tokens per second) until the bucket is full (c tokens). When requests are made, tokens are consumed from the bucket if available. If the bucket is empty, requests are denied until tokens are refilled.

## Usage

To use the token bucket implementation, follow these steps:

1. Create a new bucket instance using the NewBucket function:

```go
// Create a bucket with a fill rate of 2 tokens per second and a capacity of 10 tokens
b := bucket.NewBucket(2, 10)
```

2. Use the Check method to check if a request can be served:

```go
// Check if a single token can be consumed from the bucket
if b.Check(1) {
    // Token available, serve the request
} else {
    // Bucket empty, deny the request
}
```

3. Optionally, use the Size method to get the current size of the bucket:

```go
size := b.Size()
```

Tests

The repository includes unit tests to ensure the correctness and robustness of the implementation. To run the tests, use the following command:

```bash
$ go test -v
```

Additional Tests

Several additional test cases have been added to ensure the robustness of the implementation. These test cases cover scenarios such as exhausting the bucket's capacity, negative fill rates, negative capacity inputs, and ensure thread safety by testing concurrent access.
