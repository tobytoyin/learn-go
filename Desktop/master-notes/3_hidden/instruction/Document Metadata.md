
## Document Naming

Document filename should be named using `d-yyyy-mm-dd-HH-MM-SS`. Documents are stored under `0_pages/<randomPartition>/`. Document can have different states, however as to keep it simple and static, these states are maintain using [[#Special Tags]]. 

---

## Metadata YAML Keys

The document follows several predefined Keys for metadata management: 

- `creation-date` - the creation timestamp of this document
- `type` - types of this document (see [[#Document Type]])
- `aliases` - keywords that are useful to surface this document
- `tags` - tag/ groups that are useful to group similar documents (see [[#Special Tags]] aswell)
- `catalogs` - resources ID used to construct this document (see [[#Catalogs Management]])
- `title` - title of this document

---
## Document Type

- `type`: 
	- `concept` - concepts & knowledge that should be correct
	- `ideas` - ideas that can be incorrect and not derived from concepts
	- `summary` - summary/ interpretation of concepts that can be incorrect
	- `walkthrough` - examples/ implementation of concepts/ summary/ ideas
	- `chat` - conversation that initialised with some LLM chatbots

---
## Special Tags

- `__wip` - indicate that the note is still work-in-progress
- `__deprecated` - indicate that the note is no longer relevant (see [[3_hidden/instruction/Note Deprecation|Note Deprecation]])
- `__revision` - indicate that the note is under major revision & updates

---
## Catalogs Management

Each document might contains resources that relevant when constructing it but these resources might not exist in the current file system. To allow linking these relationships, these resources can be appended into `catalogs` based on their ID. 

For example, a document relies on a DrawIO file with ID `00001`, we can add this ID: 

```yaml
catalogs: 
	- 00001
```
