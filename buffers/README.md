# Readers, Writers, Buffers

This is one of those concepts in Go that for some reason I can't _really_ wrap my head around, so I'm going to write this whole thing out in an attempt to force myself to understand it. Go has the concept of Readers, Writers, and in general they use Buffers. The long and the short of it is that Reading and Writing, when doing a lot in a row, are computationally expensive. They are syscalls which do a lot under the hood when they get executed for each read/write. Buffers allow a read/write syscall to work with a bunch of data at a time and store it in memory so that it can be worked on. When the buffer is full, the data is flushed and a new syscall is generated to read or write more of the data instead of one syscall for each byte. It's faster on a long-enough time scale and can also allow large amounts of data to be worked on rather than load the whole thing into memory.

Readers and Writers on the other hand, are basically common functions that many modules use within an interface so they contextually do the same thing. A Reader object can be passed around to as many modules as needed as long as they implement the Reader interface.

## bytes.Buffer

This is kind of the swiss army knife of buffers. It's a buffer that you can fill with pretty much anything. So let's play with it.