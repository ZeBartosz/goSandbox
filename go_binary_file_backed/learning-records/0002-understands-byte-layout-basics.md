# Understands byte layout basics

The user now understands that 1 byte is 8 bits, fixed-size integer choices like `uint32` consume fixed space in the file format, strings are variable-size byte sequences, and `copy` places already-created string bytes into an allocated row buffer. This unlocks implementing the first row encoder/decoder without re-teaching the difference between appending and writing into explicit offsets.
