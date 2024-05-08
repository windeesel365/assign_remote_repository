# Git Commands

A collection of essential Git commands with explanations.

## Initialization and Cloning

- **`git init`**  
  Initializes a new Git repository in the selected directory.

- **`git clone [URL]`**  
  Copies a repository from the URL to the local system.

## Repository Status

- **`git status`**  
  Displays the repository's status, such as files that have been changed but not committed.

## Staging and Committing

- **`git add [filename]`**  
  Adds files to the staging area, preparing them for a commit.

- **`git commit -m "[message]"`**  
  Commits files in the staging area with the provided commit message.

## Working with Remotes

- **`git push [remote] [branch]`**  
  Sends the committed changes to a remote repository, such as GitHub or GitLab.

- **`git pull [remote] [branch]`**  
  Fetches changes from the remote repository to the local repository and merges them with the current branch.

- **`git fetch [remote]`**  
  Fetches the latest changes from a remote repository without automatically merging them into the current branch.

- **`git remote add [name] [URL]`**  
  Adds a new remote repository with the specified name and URL.

## Branch Management

- **`git branch [branch_name]`**  
  Creates a new branch from the current branch.

- **`git checkout [branch_name]`**  
  Switches to the specified branch or uses `-b` to create and switch to a new branch.

- **`git merge [branch_name]`**  
  Merges changes from the specified branch into the current branch.

- **`git rebase [branch]`**  
  Moves or integrates an existing branch into the specified branch without creating new commits.

## Other Useful Commands

- **`git log`**  
  Displays a history of commits.

- **`git reset [options]`**  
  Reverts changes, such as removing files from the staging area or rolling back a commit.

- **`git stash`**  
  Temporarily stores uncommitted changes to use them later.

- **`git tag [name]`**  
  Creates a tag for the current commit to mark a version or other significant notes.
