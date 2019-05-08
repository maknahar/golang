# [Quest Ingester](https://github.com/maknahar/go-web-skeleton)
Service to ingest searchable data into database.

# Used Packages
1. [chi](https://github.com/go-chi/chi) : For Routing
2. [flag](https://github.com/namsral/flag): For configurations (environment variables and command line arguments)
3. [logrus](https://github.com/sirupsen/logrus): For Logging 
4. [pgx](https://github.com/jackc/pgx): For PostgreSQL Database

# Usage
1. [Fork](https://help.github.com/articles/fork-a-repo/) this repo
2. [Rename](https://help.github.com/articles/renaming-a-repository/) repo name
3. Replace `/maknahar/go-web-skeleton/` with `/<your-username>/<your-repo-name>/` in forked project.

OR [duplicate](https://help.github.com/articles/duplicating-a-repository) the repo instead of step 1 and 2.


##### Pre-commit Hook

Pre-commit hooks for golang are used from https://github.com/dnephin/pre-commit-golang

Install pre-commit from https://pre-commit.com/#install.
For Mac User `brew install pre-commit`

Run `pre-commit install --install-hooks`

Install validate-toml from https://github.com/BurntSushi/toml/tree/master/cmd/tomlv

Install golangci-lint from https://github.com/golangci/golangci-lint#install
golangci-lint contains go-lint and go-critic and hence they are commented in yaml file.
They might not be enabled by default. To enable them, run below command.
`golangci-lint linters -E gocritic`

Before committing any files, the hooks mentioned in yaml will be executed.

# Dependency Management
Service uses Go Module.

# Contribution
If you know any open source package that you believe in must have in a Go Web Service, Please let us know.
