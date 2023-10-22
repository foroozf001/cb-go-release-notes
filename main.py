from git import Repo

repo = Repo("/Users/fforoozan/Repositories/vw/cb-go-release-notes")
assert not repo.bare

hc = repo.head.commit
hct = hc.tree

assert hc != hct
assert hc == repo.head.reference.commit

print(hc.message)

fifty_first_commits = list(repo.iter_commits("main", max_count=50))
assert len(fifty_first_commits) <= 50