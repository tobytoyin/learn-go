A Note can be deprecated for several reasons: 
- the contents is outdated, incorrect, or there's a better to substitute it
- the contents are migrated to somewhere else

To indicate the Node is deprecated, follow the below format: 
- duplicate the note into a new one, so the links can maintain at the exist one
- add a `__deprecated` tag in the new note
- add a first header section name "Deprecated Update" at the beginning of the new note

For example

```md
---
...rest of the frontmatter
tags:
	- __deprecated
---

# Deprecated Update

...reason for this note to be deprecated and links to the updated note...

---

Rest of the note contents

```