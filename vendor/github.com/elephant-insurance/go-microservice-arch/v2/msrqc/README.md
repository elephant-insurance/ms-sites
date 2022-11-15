# msrqc

msrqc is the Elephant Microservice Request Context. It is an interface that isolates certain features of the gin.Context. msrqc.Context also implements context.Context, so a *gin.Context is a valid msrqc.Context, and an msrqc.Context is a valid context.Context.

A concrete implementation and a generic contructor is also provided. 