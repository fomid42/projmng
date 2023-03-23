# project

## the structure of the manage

```
manage main
|- settings
	|- settings.json
	|- info.json
|- archive
|- work
	|- (project block dir)
```

## the structure of the project block

```
project block main
|- settings
	|- settings.json
	|- info.json
|- README.md
|- projects
	|- (project dir)
```

## the structure of the project

```
project main
|- settings
	|- settings.json
	|- info.json
|- notes
	|- (idea contents)
|- main
	|- README.md
	|- (project contents)
```

## the structure of the archive

```
archive main
|- settings
	|- settings.json
	|- info.json
|- notes
	|- (notes_dirs)
|- projects
	|- (project_dirs)
```

## project flow

```sequence
Note over project flow: creating(creating the directory of the structure of the project)
Note over project flow: start(start to edit)
Note over project flow: edit(edit to finish the project)
Note over project flow: end (finish the project)
Note over project flow: expire (the project is not finished, but there are no need to edit the project any further)
Note over project flow: move to archive(When a project block expires, it is moved to the archive and managed as a project, not as a project block.)
```

## project manager UI

- project block
	- title : project list : start : expire
- project
	- title : start : edit : expire : last access

## project manage

- project name
	- Project names should be in the format (title)+"_"+(id), with a unique number assigned to each project as its id to avoid potential conflicts if not protected.
- project expire
	- The user managing the project must decide whether to set an expiration date for the project and when the expiration date will expire.
- project note
	- The user must record the note of idea for next time. So, the note of idea should be separeted from the project in which the note is included, because idea may be useful in other project unrelated to this project.

## archive manage

- note
	- note names should be in the format (title)+"_note_"+(id), with a unique number assigned to each project as its id to avoid potential conflicts if not protected.