# EnvGuard

### This project is in process

---

EnvGuard is a CLI tool that helps you manage and validate environment variables across different environments and Docker Compose configurations. It ensures consistency between your .env files and helps prevent configuration-related issues.

## Features

- [x] Compare two .env files and identify unique variables in each
- [ ] Validate environment variables against Docker Compose files
- [ ] Type checking for environment variables
- [ ] Provide auto-fix

## Installation

Under construction... you need to compile from the source while:

```bash
git clone git@github.com:richardhapb/envguard
cd envguard
go build .
```

## Usage

### Compare .env Files

Compare two environment files to find missing variables:

```bash
# Compare two specific .env files
envguard check --env .env.development --env .env.production

# Check a single .env file (displays all variables)
envguard check --env .env

# Check default .env file
envguard check
```

Example output:
```
Unique values in .env.development

DB_DEBUG
REDIS_DEBUG

===============================================

Unique values in .env.production

SENTRY_DSN
NEW_RELIC_KEY
```

## Why EnvGuard?

Managing environment variables across different environments (development, staging, production) can be error-prone. EnvGuard helps you:

- Ensure all required variables are present across environments
- Prevent deployment issues due to missing environment variables
- Maintain consistency between different environment configurations
- Validate environment variable types against expected values

## Roadmap

- [ ] Docker Compose validation
  - Detect variables used in docker-compose.yml but missing in .env
  - Validate variable references and syntax
- [ ] Type checking
  - Ensure variables have consistent types across environments
  - Support for common types (string, number, boolean, URL)
- [ ] Variable templating support
- [ ] Auto-fix module
- [ ] Integration with CI/CD pipelines

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

