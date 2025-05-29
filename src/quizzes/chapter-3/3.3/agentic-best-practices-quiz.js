const rawQuizdown = `

# What is the primary reason for documenting prompts when working with LLMs?

1. [ ] To comply with company security policies
   > While security compliance might benefit from documentation, it's not the primary reason.
1. [x] To provide a record of what prompts were used and what results were obtained for future reference
1. [ ] To prevent LLM hallucinations
   > Documentation doesn't prevent hallucinations, though it may help identify them.
1. [ ] To reduce token usage and costs
   > Documentation doesn't directly affect token usage or costs.

# In Harper Reed's LLM Codegen Workflow, what is the correct sequence of stages?

1. [ ] Planning, Execution, Idea Honing
   > This order is incorrect. The workflow starts with developing the specification.
1. [ ] Execution, Idea Honing, Planning
   > This order is incorrect. Planning precedes execution.
1. [x] Idea Honing, Planning, Execution
1. [ ] Planning, Idea Honing, Execution
   > This order is incorrect. Idea honing (developing the specification) comes first.

# What is the "Second Opinion" technique primarily used for?

1. [ ] Generating initial code for a new feature
   > This technique is applied to existing work, not for initial generation.
1. [ ] Finding bugs in existing code
   > While it might identify bugs, this isn't its primary purpose.
1. [x] Reviewing existing technical work to find simpler or more idiomatic solutions
1. [ ] Automating unit test creation
   > This isn't the focus of the Second Opinion technique.

# Why are LLMs described as "dumber than they look" in the context of professional software development?

1. [ ] They can't write code as well as human developers
   > This isn't the core reason for the description.
1. [ ] They take longer to complete tasks than humans
   > Speed isn't the issue being addressed.
1. [ ] They require constant supervision
   > While supervision is important, this doesn't capture the core limitation.
1. [x] They are statistical text predictors without true understanding, despite appearing intelligent

# What's the recommended approach when using LLMs for a complex development task?

1. [ ] Ask the LLM to complete the entire task at once
   > This approach doesn't maintain sufficient control over the process.
1. [x] Break down the task into small, well-defined steps and use the LLM for each step with human oversight
1. [ ] Have multiple LLMs work on different parts simultaneously
   > This doesn't address the core need for human oversight and integration.
1. [ ] Generate as many variations as possible and pick the best one
   > This is inefficient and doesn't leverage structured workflows.

# What is the "Throwaway Debugging Scripts" technique best used for?

1. [ ] Creating production-ready utilities
   > These scripts are explicitly not intended for production use.
1. [ ] Documenting code
   > Documentation isn't the primary purpose of these scripts.
1. [x] Automating specific debugging steps without requiring codebase integration
1. [ ] Generating test data
   > While it could be used for this, it's not the primary purpose.

# When using LLMs to "plug technical gaps" in unfamiliar domains, what is a critical best practice?

1. [ ] Only use the most advanced LLM available
   > The specific LLM is less important than proper review.
1. [ ] Generate at least three alternatives for every solution
   > Multiple alternatives aren't necessarily required if proper review is conducted.
1. [ ] Immediately implement the solution in production
   > This would be dangerous without proper review.
1. [x] Have the generated code reviewed by a domain expert

`;

export { rawQuizdown }
