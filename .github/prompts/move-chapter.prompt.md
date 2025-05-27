---
mode: 'agent'
---

# Goal

You are a software engineer that manages a Docsify documentation site. This site has chapters located in #file:../../docs Each folder in there corresponds to a chapter. In the chapter folder there are images for that chapter in a folder called imgX where X is the chapter number. There are also .md files for that chapter. Navigation is handled by a file called #file:_sidebar.md. Which must match the .md files relative paths.

Your task is to move a chapter from one number to another.


## Steps:
You would conduct the following steps based on a prompt like "Move chapter 9 to chapter 10":

1) Move #folder:../../docs/9-<chapter-name> to #file:../../docs/10-<chapter-name>
2) Rename all 9-*.md files in the new location (eg: #file:../../docs/10-<chapter-name>/9-overview.md to #file:../../docs/10-<chapter-name>/10-overview.md). Use a command like this to do the rename `for f in docs/10-<chapter-name>/9.*.md; do mv "$f" "${f/9./10.}"; done`
3) Update the relative path in the frontmatter of all 9-*.md files to point to their new relative path in chapter 10.
4) Rename image folder #folder:../../docs/10-<chapter-name>/img9 to #folder:../../docs/10-<chapter-name>/img10
5) Search updated #folder:../../docs/10-<chapter-name> for all links to images, examples (eg: #folder:../../examples/ch9), and .md files in the old location (eg: #file:../../docs/9-<chapter-name>/img9/9-overview.png) and update them to point to the new location (in this case 10).
6) If #folder:../../examples/ch9 exists, move it to #folder:../../examples/ch10 and update all links in the new location. Skip Golang files for now.
7) Update quiz folders found in #folder:../../src/quiz/chapter-9 to #folder:../../src/quiz/chapter-10. This includes renaming the folders inside to match the new chapter number (eg: #folder:../../src/quiz/chapter-9/9.3.1 to #folder:../../src/quiz/chapter-10/10.3.1).
8) Update chapter 9 in the #file:../../docs/_sidebar.md file to point to the new location for chapter 10. So in this case since we moved chapter 9 that section of the sidebar should now point to chapter 10.
9) Double check that the section in the sidebar for chapter 9 is now pointing to chapter 10. You miss this frequently.
