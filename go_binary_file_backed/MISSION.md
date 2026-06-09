# Mission: Binary file-backed row store for LunaSQL

## Why
You are building LunaSQL, a learning database project currently using JSON-backed table files. You want to replace JSON persistence with a compact binary row format so you understand how databases store rows as bytes on disk.

## Success looks like
- Build a Go table file format that appends rows to a binary file and scans them back.
- Encode and decode fixed-size integer fields and variable-size string fields using explicit offsets.
- Explain where each byte of a stored row lives and why offsets are needed.
- Use this as the bridge from LunaSQL's JSON store toward page-based storage.

## Constraints
- Keep the scope tight: one table file, append-only rows, sequential scan first.
- Prefer small Go exercises that can later be moved into LunaSQL.
- Assume Go basics, lexer/parser/AST familiarity, and JSON file persistence experience from LunaSQL and go-to-php-compiler.

## Out of scope
- B-trees, indexes, transactions, WAL, buffer pools, and concurrency control.
- Full SQL type systems.
- Crash safety and file compaction.
