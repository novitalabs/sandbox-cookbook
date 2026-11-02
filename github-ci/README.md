# GitHub Actions CI on Novita Sandbox

Run GitHub Actions jobs inside Novita Sandbox via a self-hosted runner scale set
(`runs-on: novita-scale-set`). Three tiny calculator projects (Go / Python /
JavaScript) each have their own workflow in `.github/workflows/`.

> This recipe is for reference only. The `.github/` directory lives inside
> `github-ci/` so GitHub will not pick it up from this repo. To try it, copy
> `.github/` together with the demo directories to the root of your own repo.

| Dir            | Lint                            | Test                   | Workflow                       |
|----------------|---------------------------------|------------------------|--------------------------------|
| `go-demo/`     | gofmt, go vet, staticcheck      | `go test -cover`       | `.github/workflows/go.yml`     |
| `python-demo/` | ruff check, ruff format --check | `pytest` (3.10–3.12)   | `.github/workflows/python.yml` |
| `js-demo/`     | eslint                          | `node --test` (20, 22) | `.github/workflows/js.yml`     |

Each workflow only triggers when its own directory (or its workflow file)
changes.

## How it works

1. A GitHub Actions runner scale set named `novita-scale-set` is registered to
   this repository / organization and backed by Novita Sandbox instances.
2. Workflows request it with `runs-on: novita-scale-set`.
3. Each job starts in a fresh sandbox, checks out the code, installs its
   toolchain with the standard `actions/setup-*` actions, then runs lint and
   tests like on any other runner.

## Run locally

```bash
# Go
cd go-demo && go vet ./... && go test -cover ./...

# Python
cd python-demo && uv sync --extra dev && uv run pytest

# JavaScript
cd js-demo && npm ci && npm run lint && npm test
```
