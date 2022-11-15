# rbuf

rbuf is the Elephant Ring Buffer, an in-memory data structure that does not grow, can't be overrun, and can be configured to act as a buffer, queue, or stack for arbitrary data types.

## Features

+ Flexible, in-memory storage buffer that can grow and shrink without (usually) allocating more RAM
+ Buffer may be optionally pre-filled with dummy data for RAM usage stability
+ Buffer may hold any kind of data, including pointers to structs (not just bytes) without type constraints
+ Buffer will never be overrun or panic, even when it is full and nothing is reading from it
+ Can be used as a persistent cache (non-destructive reads) or a read-once buffer (destructive reads)
+ Can be read in FIFO or LIFO order
+ Single-threaded (for best performance in single-threaded code) and thread-safe versions
+ Buffer supports blocking reads for dedicated reader loops
+ Fully unit-tested with completely implemented example code