# Web Applications 📚

A collection of small–to–medium-sized web apps I’ve built to practice modern stacks and to showcase in my résumé.  
Every project lives in its own sub-directory so you can browse or run them independently.

| Folder | Stack | What it is |
|--------|-------|------------|
| **NotesManager** | PostgreSQL · Go (Gin) · React (Vite) | A full-stack note-taking service with lists, JWT auth and Dockerised dev/prod workflow |

&nbsp;

---

## Running a project

```bash
# clone repo
git clone github.com/berezovskiydeval/web-applications.git
cd web-applications

# each project has its own README
cd NotesManager
docker compose up           # example – spins up db, backend & frontend
