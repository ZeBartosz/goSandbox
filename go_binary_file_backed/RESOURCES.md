# Binary file-backed row store Resources

## Knowledge

- [Go package docs: `encoding/binary`](https://pkg.go.dev/encoding/binary)
  Canonical Go API for converting fixed-size numbers to and from byte slices. Use for: integer encoding, byte order, `PutUint32`, `Uint32`.
- [Go package docs: `os`](https://pkg.go.dev/os)
  Canonical Go API for opening, reading, writing, seeking, and syncing files. Use for: `OpenFile`, file modes, append/read workflows.
- [Go blog: Go Slices: usage and internals](https://go.dev/blog/slices-intro)
  Official explanation of slices as descriptors over arrays. Use for: understanding `[]byte`, slicing, length vs capacity, copying bytes.
- [SQLite Database File Format](https://www.sqlite.org/fileformat.html)
  High-trust real-world database file format documentation. Use for: pages, varints, payloads, and how serious databases describe bytes on disk.
- [PostgreSQL docs: Database Page Layout](https://www.postgresql.org/docs/current/storage-page-layout.html)
  Production database page layout reference. Use later for: page headers, item pointers, tuple storage, and why row stores rarely stay as one giant row list.
- [CMU 15-445/645 Database Systems: Storage notes](https://15445.courses.cs.cmu.edu/)
  University database systems course. Use for: storage manager, pages, tuples, buffer pool, and B+ tree concepts as LunaSQL grows.

## Wisdom (Communities)

- [Database Internals Discord / community around Alex Petrov's _Database Internals_](https://databass.dev/)
  Useful for asking storage-engine design questions after you have a concrete design or bug.
- [r/databasedevelopment](https://www.reddit.com/r/databasedevelopment/)
  Niche subreddit for people building databases. Use for: design sanity checks, but verify answers against docs/source.

## Gaps

- Need a specific small row-format reference implementation in Go that matches LunaSQL's learning scope. We will create one in this workspace.
