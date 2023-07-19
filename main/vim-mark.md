# vim mark

## set marks

| cmd         | description                           |
| ----------- | ------------------------------------- |
| `m{a-zA-Z}` | set mark a at current cursor location |

## use marks

| cmd  | description                                                   |
| ---- | ------------------------------------------------------------- |
| `ma` | set mark a at current cursor location                         |
| `'a` | jump to line of mark a (first non-blank character in line)    |
| `a   | jump to position (line and column) of mark a                  |
| d'a  | delete from current line to line of mark a                    |
| d`a  | delete from current cursor position to position of mark a     |
| c'a  | change text from current line to line of mark a               |
| y`a  | yank text to unnamed buffer from cursor to position of mark a |

## jump marks

| Command    | Description                                                   |
| ---------- | ------------------------------------------------------------- |
| `.         | jump to position where last change occurred in current buffer |
| `"         | jump to position where last exited current buffer             |
| `0         | jump to position in last file edited (when exited Vim)        |
| `1         | like \`0 but the previous file (also \`2 etc)                 |
| `''`       | jump back (to line in current buffer where jumped from)       |
| ``         | jump back (to position in current buffer where jumped from)   |
| \`[ or \`] | jump to beginning/end of previously changed or yanked text    |
| \`< or \`> | jump to beginning/end of last visual selection                |

## list marks

| cmd       | description                |
| --------- | -------------------------- |
| :marks    | list all the current marks |
| :marks aB | list marks a, B            |

## delete marks

| cmd            | description                                             |
| -------------- | ------------------------------------------------------- |
| :delmarks a    | delete mark a                                           |
| :delmarks a-d  | delete marks a, b, c, d                                 |
| :delmarks abxy | delete marks a, b, x, y                                 |
| :delmarks aA   | delete marks a, A                                       |
| :delmarks!     | delete all lowercase marks for the current buffer (a-z) |