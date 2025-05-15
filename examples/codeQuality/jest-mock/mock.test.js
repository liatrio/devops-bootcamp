jest.mock("./github", () => ({
  getUserInfo: jest.fn(),
}));

const { getUserInfo } = require("./github");
const {
  checkUserLogin,
  checkUserName,
  checkUserBio,
  checkUserFollowers,
  checkUserPublicRepos,
  checkUserUrl,
} = require("./mock");

const mockUserInfo = {
  login: "chrisjblackburn",
  name: "Chris Blackburn",
  bio: "DevOps leader and enthusiast",
  followers: 250,
  public_repos: 42,
  html_url: "https://github.com/chrisjblackburn",
};

describe("User Info Checks (with mocked getUserInfo)", () => {
  let logSpy;

  beforeEach(() => {
    logSpy = jest.spyOn(console, "log").mockImplementation(() => {});
    getUserInfo.mockResolvedValue(mockUserInfo);
  });

  afterEach(() => {
    logSpy.mockRestore();
    getUserInfo.mockReset();
  });

  test("checkUserLogin matches expected", async () => {
    await checkUserLogin("chrisjblackburn", "chrisjblackburn");
    expect(getUserInfo).toHaveBeenCalledWith("chrisjblackburn");
    expect(logSpy).toHaveBeenCalledWith(
      "User login matches expected: chrisjblackburn",
    );
  });

  test("checkUserName matches expected", async () => {
    await checkUserName("chrisjblackburn", "Chris Blackburn");
    expect(logSpy).toHaveBeenCalledWith(
      "User name matches expected: Chris Blackburn",
    );
  });

  test("checkUserBio contains keyword", async () => {
    await checkUserBio("chrisjblackburn", "DevOps");
    expect(logSpy).toHaveBeenCalledWith(
      'User bio contains keyword "DevOps": DevOps leader and enthusiast',
    );
  });

  test("checkUserFollowers above threshold", async () => {
    await checkUserFollowers("chrisjblackburn", 100);
    expect(logSpy).toHaveBeenCalledWith(
      "User has a good number of followers: 250",
    );
  });

  test("checkUserPublicRepos above threshold", async () => {
    await checkUserPublicRepos("chrisjblackburn", 10);
    expect(logSpy).toHaveBeenCalledWith(
      "User has a healthy number of public repos: 42",
    );
  });

  test("checkUserUrl is secure", async () => {
    await checkUserUrl("chrisjblackburn");
    expect(logSpy).toHaveBeenCalledWith(
      "User URL is secure: https://github.com/chrisjblackburn",
    );
  });
});

describe("Repo Info Checks (with mocked getRepoInfo)", () => {
  test("TODO: WRITE REPO TESTS", async () => {
    fail("You haven't written repo tests yet!");
  });
});
