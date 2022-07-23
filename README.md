# k4Const: prevent `kVarName`-like variables from modifying it.

## Example

TBD

## Rationale

1. When the codebase is big, I really miss the `const` type qualifier to prevent unexpected modification.
2. Golang uses identifier names to represent public/private, so it's similar to use it for other purpose.
3. It's a general notation to use `k` in naming, to represent it's a constant.