# log
The Elephant logging package, which allows applications to create and send structured log messages.


## Logs
log assumes that every microservice will have one default log, which it shares with many other microservices. The MicroserviceLog type encapsulates the default application log for a microservice. In addition to generating certain automatic log messages, the MicroserviceLog handles a wide range of basic logging needs including local debugging, performance monitoring, and session tracking.

Most of the code and functionality of the log package is devoted to the MicroserviceLog. But the package also offers the RealTimeBidLog, a specialized log for realtime bidding appplications. This log offers a much smaller feature set than the MicroserviceLog, but shares the same basic design.

It's also quite easy to create an app-specific log sharing the same design, using the RealTimeBidLog as a pattern.


### The Log design
All logs have three sets of handlers for dealing with log messages: immediates, relays, and flushers. By default a log has one of each, but it is easy to add or remove from each set, as needed. Communications to message handlers only goes in one direction; none of them returns errors or other values to the caller. Every log should implement a Diagnostics method to expose diagnostic data to the consuming app.

#### Immediate Handlers 
Immediate handlers do something with a log message immediately, i.e., on the calling thread. By default, this is where we implement console log handlers, which log directly to stdout and stderr. Immediate handlers should be very lightweight, as the calling thread will have to wait while they execute. Any process that might take significant time, such as writing to a file or sending a web request, should be performed elsewhere.

#### Relay Handlers
A Relay Handler is a handler that is designed for longer-running processes. Every Relay Handler must implement an Add(msg) function, which should fire off a new thread to perform its work and return immediately. 

The default Relay Handler is mbuf.MessageRelay, which sends messages to remote message handlers. The mbuf package offers pre-configured MessageRelays for both Loggly and Azure, and it's quite simple to write an implementation for an arbitrary service. MessageRelay also provides its own Diagnostics.

Other valid purposes for a Relay Handler would be to write log messages to a file, save them to a database, or add them to an Azure message bus.

#### Flushers
Flushers are called by a Relay Handler when the handler is unable to complete its work successfully. For example, mbuf.MessageRelay will call its flushers after it has tried the max number of times to send its messages to its configured service. The default flusher writes its contents to stdout using JSON.Marshal. From stdout, the messages may be further processed by the host, such as by writing them to a file in var/log, or relaying them a backup log service.


## The MicroserviceLog
MicroserviceLog implements all of the standard, familiar logging features, such as standard syslog fields, levels, fatal and panic, etc. The log package normally creates a MicroserivceLog at startup and sets this as its default MicroserviceLog. Package functions such as ForFunc() use the default MicroserviceLog. While it is technically possible that a microservice could have more than one MicroserviceLog, it's doubtful that this would ever be a sound design. So for the purposes of this documentation, it is safe to assume that there is always one MicroserviceLog, the default log for the microservice.

### Log Levels
When the MicroserivceLog is created, it is assigned a level from the usual set of seven log levels (trace, debug, info, warn, error, fatal, panic). This level is what we think of as the log level for the application itself. Additionally, each of its handlers (usually one of each, an immediate for the console and a MessageRelay) can have its own logging level. Finally, every MicroserivceLog message is assigned a level when it is created, as indicated by the function used (Trace(msg), Debug(msg), etc.).

The level rules are extremely simple:
* If the level of a message would be LESS URGENT than the level of its MicroserivceLog, it will never be created, and
* If the level of a message is LESS URGENT than the level of a message handler, the handler will ignore it.

For clarity, this means that if the level of the MicroserivceLog is Error, then a message logged using Debug(msg) will not be created. In fact, the function will short-circuit very quickly after evaluating msg. This means that it is safe to use as many Debug(msg) statements as you like without worrying about performance, so long as the "msg" argument is not itself a function that takes a long time to run. 



## Best Practices

Start each function with a call to log.ForFunc(ccontext.Context). You can either keep the logger that ForFunc returns and use it to write other messages, or you can let the return value go. Either way, you will see the function fire when you turn on trace-level logging.

Use const strings for your msg values. Remember that anything you pass as an argument to a function in Go WILL be evaluated, even if the function itself does no work. So if you build a big string inside your Debug(), like 
```go
    logger.Debug(fmt.Sprintf(`processed %v messages in %v, error: %v`, msgCount, elapsed, err))
```
that string is ALWAYS going to be built and copied at least once, even if the application log level is all the way up at Panic. 

Instead of 