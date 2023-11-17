---
id: C68E3F1E-A8DC-41CF-9567-DAD2077B0113
creation-date: 2023-01-27T22:49:03
type: summary
aliases:
  - vim cheatsheet
tags:
  - cli
  - setup
  - vim
---

## Common Keys

> [!NOTE]
> `[n]` - indicates a variable number 

- `[n]W` - forward N *words*
- `[n]B` - *backward* N words
- `D` - *delete*
	- `DW` - at start of word, *delete* the word
	- `D[n]W` - at start of word, *delete* N words
	- `DD` - delete line
	- `D$` - delete to end of line
	- `D0` - delete to start of line
- `C` - *change*, (i.e., delete & insert at the same time)
	- `CW` - change word 
	- `C[n]W` - change N words
- `Y` - *yank* , i.e., copy
	- `YY` - copy line
- `P` - *paste* what you yanked 
- `J` - move downward (down arrow)
	- `[n]J` - move downward N lines
- `K` - move upward (up arrow) 
	- `[n]K` - move upward N lines
- `U` - undo
- `V` - visual mode, i.e., highlight & select multiple lines
- `>`/ `<` - indent and back-indent this and next lines
- `>>` / `<<` - indent and back-ident this line
- `$` - move to end of line
- `0` - move to start of line
- `%` - move to start of matching bracket

---
## Useful Commands

- search text - `/<text>` and use `N` to next result
- global text find & replace - `:%s/<find>/<replace>/g`  - this is similar to [[0_pages/01/2023-03-25-16-32-12-31800|sed]]

---
## Useful Setups 

Some useful setups when using non-local Vim (e.g., cloud server): 
- relative line numbers - `:set relativenumber`
- set 2 spaces tab - `:set tabstop=2 softtabstop=2 shiftwidth=2`

---
## ℹ️  Resources
- [How to Find and Replace in Vim - YouTube](https://www.youtube.com/watch?v=PzmLJy0o6qo)