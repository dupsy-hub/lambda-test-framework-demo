# Go Lambda Testing Framework

A standardized, reusable framework for writing **unit tests, benchmarks, and integration tests** for all AWS Lambda functions written in Go.

It ensures **code quality**, **reliability**, and **automation** across all Lambda projects.

---

## Framework Overview — How It Works

```mermaid
flowchart LR
    A[Lambda Handler Code] --> B[Dependency Interfaces (awsiface)]
    B --> C[Generated Mocks (mockgen)]
    C --> D[Unit Tests (table-driven)]
    D --> E[Makefile + Coverage Script]
    E --> F[CI Pipeline (GitHub Actions)]
    F --> G[Coverage Gate ≥90% ✅]
    G --> H[Confident, Tested, Deployable Lambda]
```

## Repository Structure

```python
unit-test/
├── Makefile                   # Test, coverage, mock, and benchmark commands
├── go.mod / go.sum            # Go module + dependency management
├── internal/
│   ├── awsiface/              # AWS SDK interfaces for mocking
│   │   └── s3.go
│   ├── handlers/              # Lambda function logic (with tests)
│   │   └── sample/
│       ├── handler.go
│       ├── handler_test.go
│       ├── handler_bench_test.go
│       └── handler_integration_test.go
├── testkit/
│   ├── assert.go
│   └── localstack.go
├── mocks/                     # Auto-generated mock files
│   └── mock_s3.go
├── scripts/                   # Automation scripts
│   └── coverage_check.sh
├── testdata/                  # JSON fixtures for integration tests
└── .github/workflows/ci.yml   # CI pipeline (optional)

```

---

## Getting Started

## 1️. Install dependencies

```bash
go mod tidy

```

## 2. Generate AWS mocks

```bash
mingw32-make mocks

```

## 3. Run all unit tests

```bash
mingw32-make unit

```

Runs all tests
Generates a coverage report (coverage.out)
Enforces ≥90% coverage via coverage_check.sh

## 4. Run performance benchmarks (optional)

```bash

mingw32-make bench

```

## 5. View coverage report

```bash
mingw32-make cover
```

Then open coverage.html in your browser.

## Acceptance Criteria Mapping

Requirement How It’s Achieved

> 90% coverage Enforced automatically in make unit and CI pipeline
> Mock AWS services Implemented via internal/awsiface/ + mocks/
> Success/failure test cases See handler_test.go (valid, missing ID, S3 failure)
> Performance benchmarks handler_bench_test.go
> Integration test utilities internal/testkit/ (add LocalStack helpers)
> Automated CI/CD testing .github/workflows/ci.yml ready

## Adding a New Lambda

Create a new directory:

```bash
mkdir -p internal/handlers/<lambda-name>

```

Add your handler logic in handler.go.

Write tests in handler_test.go (table-driven style).

Add benchmarks in handler_bench_test.go.

Run and verify:

```bash
mingw32-make mocks
mingw32-make unit

```

## Adding New AWS Service Interfaces

Example (internal/awsiface/dynamodb.go):

```bash

package awsiface

import (
  "context"
  "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoPutItem interface {
  PutItem(ctx context.Context, in *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

var _ DynamoPutItem = (*dynamodb.Client)(nil)

```

Generate the mock:

```bash
mingw32-make mocks

```

## Optional Integration Tests (LocalStack)

To test AWS calls locally:

Add internal/testkit/localstack.go:

```bash
//go:build integration
package testkit
// Add constructors for S3, DynamoDB, etc., using LOCALSTACK_ENDPOINT.

```

Run integration tests:

```bash

LOCALSTACK_ENDPOINT=http://localhost:4566 go test -tags=integration ./...

```

## Guardrails (Future Enhancements)

| Guardrail                  | Purpose                                                                             |
| -------------------------- | ----------------------------------------------------------------------------------- |
| **Branch protection rule** | Require CI to pass and ≥90% coverage before merge                                   |
| **Lint rule**              | Fail build if a new `internal/handlers/<lambda>` has `.go` files but no `*_test.go` |
| **Pre-commit hook**        | Auto-run `make mocks` and `make unit` before commits                                |

## Handy Commands (Windows Git Bash)

If make isn’t recognized:

```bash
echo 'alias make=mingw32-make' >> ~/.bashrc
source ~/.bashrc
```

Then you can just use:

```bash
make mocks
make unit
```

## Example Reference

| Component              | File                                             | Purpose                            |
| ---------------------- | ------------------------------------------------ | ---------------------------------- |
| **Sample Lambda**      | `internal/handlers/sample/handler.go`            | Core Lambda logic                  |
| **Unit tests**         | `internal/handlers/sample/handler_test.go`       | Success + failure + mock scenarios |
| **Benchmark**          | `internal/handlers/sample/handler_bench_test.go` | Performance testing example        |
| **Mock AWS interface** | `internal/awsiface/s3.go`                        | Mockable S3 client definition      |
| **Generated mock**     | `mocks/mock_s3.go`                               | Auto-generated via `mockgen`       |
| **Coverage script**    | `scripts/coverage_check.sh`                      | Enforces ≥90% coverage             |
| **Test utilities**     | `internal/testkit/assert.go`                     | Reusable assertions                |

## Outcome

This framework ensures:

Consistent Lambda testing patterns across the team

Isolated, reproducible unit tests (no real AWS calls)

Benchmarks for performance awareness

CI/CD enforcement of coverage and quality gates

## Quickstart

## 1. Clone the repo

```bash
git clone https://github.com/your-org/unit-test.git
cd unit-test

```

## 2. Generate mocks

```bash
mingw32-make mocks

```

## 3. Run tests with coverage gate

```bash
mingw32-make unit

```
