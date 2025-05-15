const { getUserInfo, getRepoInfo } = require("./github");

//// USER CHECKS FOR chrisjblackburn ////

async function checkUserLogin(username, expectedLogin = "chrisjblackburn") {
  const userInfo = await getUserInfo(username);
  if (userInfo.login === expectedLogin) {
    console.log(`User login matches expected: ${userInfo.login}`);
  } else {
    console.log(
      `User login does NOT match expected. Got: ${userInfo.login}, Expected: ${expectedLogin}`,
    );
  }
}

async function checkUserName(username, expectedName = "Chris Blackburn") {
  const userInfo = await getUserInfo(username);
  if (userInfo.name === expectedName) {
    console.log(`User name matches expected: ${userInfo.name}`);
  } else {
    console.log(
      `User name does NOT match expected. Got: ${userInfo.name}, Expected: ${expectedName}`,
    );
  }
}

async function checkUserBio(username, keyword = "DevOps") {
  const userInfo = await getUserInfo(username);
  if (userInfo.bio && userInfo.bio.includes(keyword)) {
    console.log(`User bio contains keyword "${keyword}": ${userInfo.bio}`);
  } else {
    console.log(`User bio does NOT contain keyword "${keyword}".`);
  }
}

async function checkUserFollowers(username, minFollowers = 100) {
  const userInfo = await getUserInfo(username);
  if (userInfo.followers >= minFollowers) {
    console.log(`User has a good number of followers: ${userInfo.followers}`);
  } else {
    console.log(
      `User has less than ${minFollowers} followers: ${userInfo.followers}`,
    );
  }
}

async function checkUserPublicRepos(username, minRepos = 10) {
  const userInfo = await getUserInfo(username);
  if (userInfo.public_repos >= minRepos) {
    console.log(
      `User has a healthy number of public repos: ${userInfo.public_repos}`,
    );
  } else {
    console.log(
      `User has less than ${minRepos} public repos: ${userInfo.public_repos}`,
    );
  }
}

async function checkUserUrl(username) {
  const userInfo = await getUserInfo(username);
  if (userInfo.html_url && userInfo.html_url.startsWith("https://")) {
    console.log(`User URL is secure: ${userInfo.html_url}`);
  } else {
    console.log(`User URL is not secure or missing: ${userInfo.html_url}`);
  }
}

//// REPO CHECKS FOR DEVOPS-BOOTCAMP ////

async function checkRepoName(owner, repo, expectedName = "devops-bootcamp") {
  const repoInfo = await getRepoInfo(owner, repo);
  if (repoInfo.name === expectedName) {
    console.log(`Repo name matches expected: ${repoInfo.name}`);
  } else {
    console.log(
      `Repo name does NOT match expected. Got: ${repoInfo.name}, Expected: ${expectedName}`,
    );
  }
}

async function checkRepoDescription(owner, repo, keyword = "DevOps") {
  const repoInfo = await getRepoInfo(owner, repo);
  if (repoInfo.description && repoInfo.description.includes(keyword)) {
    console.log(
      `Description contains keyword "${keyword}": ${repoInfo.description}`,
    );
  } else {
    console.log(`Description does NOT contain keyword "${keyword}".`);
  }
}

async function checkRepoStars(owner, repo, minStars = 100) {
  const repoInfo = await getRepoInfo(owner, repo);
  if (repoInfo.stargazers_count >= minStars) {
    console.log(
      `Repo has a good number of stars: ${repoInfo.stargazers_count}`,
    );
  } else {
    console.log(
      `Repo has less than ${minStars} stars: ${repoInfo.stargazers_count}`,
    );
  }
}

async function checkRepoForks(owner, repo, minForks = 20) {
  const repoInfo = await getRepoInfo(owner, repo);
  if (repoInfo.forks_count >= minForks) {
    console.log(`Repo has a healthy number of forks: ${repoInfo.forks_count}`);
  } else {
    console.log(
      `Repo has less than ${minForks} forks: ${repoInfo.forks_count}`,
    );
  }
}

async function checkRepoOpenIssues(owner, repo, warningLevel = 10) {
  const repoInfo = await getRepoInfo(owner, repo);
  if (repoInfo.open_issues_count > warningLevel) {
    console.log(
      `Warning: Repo has a high number of open issues: ${repoInfo.open_issues_count}`,
    );
  } else {
    console.log(`Open issues are under control: ${repoInfo.open_issues_count}`);
  }
}

async function checkRepoUrl(owner, repo) {
  const repoInfo = await getRepoInfo(owner, repo);
  if (repoInfo.html_url && repoInfo.html_url.startsWith("https://")) {
    console.log(`Repo URL is secure: ${repoInfo.html_url}`);
  } else {
    console.log(`Repo URL is not secure or missing: ${repoInfo.html_url}`);
  }
}

module.exports = {
  checkUserLogin,
  checkUserName,
  checkUserBio,
  checkUserFollowers,
  checkUserPublicRepos,
  checkUserUrl,
  checkRepoName,
  checkRepoDescription,
  checkRepoStars,
  checkRepoForks,
  checkRepoOpenIssues,
  checkRepoUrl,
};
