# mbuf

mbuf is the Elephant Message Relay, a package for sending messages asynchronosly. 

Features:
- Simple interface with two methods: Add(message) and Diagnostics()
- Simple, one-line configuration with only one required field (the service URL we want to send to)
- Add(message) method forks a new thread and returns immediately for fast "fire and forget" performance
- All data is thread-local and synch-locked where necessary, preventing race conditions
- Buffered messages are sent either when the buffer is full or when a polling timer fires
- Complete control over message formatting and addressing via pluggable interface helpers
- Sensible default helpers format messages and HTTP requests using simple JSON protocol
- Pluggable Azure helper handles Azure signing and header complexities
- Handles network outages gracefully by automatically caching messages, waiting, and retrying
- Provides methods for "flushing" messages that cannot be sent
- The default Flusher dumps unsendable messages to stdout
- Additional Flushers can call callback methods or perform other special handling
- Provides extensive feedback and diagnostic output
- Easily configured for single-message (no-wait) sending
- Easily configured for full-buffer-only (no polling) or polling-only (unlimited buffer size) behavior
- Completely self-contained, so that multiple relays can be active at once, each configured differently
