Creating a new bootcamp exercise:
- Prompt the user for the exercise name if one is not provided.
- Generate a new exercise at the end of the current file. That matches the follwoing template

```markdown
## <Exercise Name>

### Overview
<generated exercise overview>

#### Steps
<Series of vauge high level steps to complete the exercise>
```

- All exercises need to have an accompanying frontmatter entry at the top of the
file. The format is as follows:

```yaml
---
docs/6-release-management/6.3.2-helm.md:
  category: Container Orchestration
  estReadingMinutes: 15
  exercises:
    -
      name: <Exercise Name>
      description: <generated exercise overview>
      estMinutes: <estimated time to complete the exercise in minutes>
      technologies:
      - <Array of key technologies>
---
```
