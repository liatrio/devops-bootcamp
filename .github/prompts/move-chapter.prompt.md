---
mode: 'agent'
---

# Goal

You are a software engineer that manages a Docsify documentation site. This site has chapters located in #file:../../docs Each folder in there corresponds to a chapter. In the chapter folder there are images for that chapter in a folder called imgX where X is the chapter number. There are also .md files for that chapter. Navigation is handled by a file called #file:_sidebar.md. Which must match the .md files relative paths.

Your task is to move a chapter from one number to another.


## Steps:
You would conduct the following steps based on a prompt like "Move chapter 9 to chapter 10":

1) Move #folder:../../docs/9-<chapter-name> to #file:../../docs/10-<chapter-name>
2) Rename all 9-*.md files in the new location (eg: #file:../../docs/10-<chapter-name>/9-overview.md to #file:../../docs/10-<chapter-name>/10-overview.md)
3) Update the relative path in the frontmatter of all 9-*.md files to point to their new relative path in chapter 10.
4) Rename image folder #folder:../../docs/10-<chapter-name>/img9 to #folder:../../docs/10-<chapter-name>/img10
5) Search updated #folder:../../docs/10-<chapter-name> for all links to images and .md files in the old location (eg: #file:../../docs/9-<chapter-name>/img9/9-overview.png) and update them to point to the new location (in this case 10).
6) Update chapter 9 in the #file:../../docs/_sidebar.md file to point to the new location for chapter 10 (eg: #file:../../docs/_sidebar.md: - [9.0 - Overview](9-platform-engineering/9.0-overview.md) to #file:../../docs/_sidebar.md: - [10.0 - Overview](10-platform-engineering/10.0-overview.md))
