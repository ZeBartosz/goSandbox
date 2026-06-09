# Binary file-backed row store Glossary

This glossary covers the language used while moving LunaSQL from JSON-backed table files to explicit binary row storage.

## Terms

**Byte**:
A unit of storage made of 8 bits. In this workspace, row layouts are described by byte positions, not by Go source-level values.
_Avoid_: Bite

**Fixed-size field**:
A field that always uses the same number of bytes in the file format, such as a `uint32` stored as 4 bytes.
_Avoid_: Normal field

**Variable-size field**:
A field whose byte length depends on the value being stored, such as a string encoded as UTF-8 bytes.
_Avoid_: Dynamic thing

**Offset**:
A byte position measured from a known starting point, such as the beginning of a row or file.
_Avoid_: Index, place

