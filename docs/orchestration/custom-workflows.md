# Custom Workflows

Create your own workflow templates to define custom orchestration patterns. Templates are automatically loaded into the workflow picker.

## Template Location

Place custom templates in:

```
~/.perles/workflows/{your_workflow_name}/
```

Each workflow directory contains a `template.yaml` file and markdown templates for each task.

---

## Template Structure

A workflow template consists of:

```
~/.perles/workflows/joke-contest/
├── template.yaml                    # DAG definition and metadata
├── v1-joke-contest-epic.md          # Epic instructions (coordinator)
├── v1-joke-contest-joke-1.md        # Task 1 template
├── v1-joke-contest-joke-2.md        # Task 2 template
├── v1-joke-contest-judge.md         # Task 3 template
└── v1-human-review.md               # Human review gate
```

---

## Template YAML Reference

### Top-Level Registration Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `namespace` | string | Yes | Always `"workflow"` |
| `key` | string | Yes | Unique identifier (e.g., `"joke-contest"`) |
| `version` | string | Yes | Version identifier (e.g., `"v1"`) |
| `name` | string | Yes | Human-readable name for the workflow picker |
| `description` | string | Yes | Description of what the workflow does |
| `epic_template` | string | Yes | Filename of the markdown template for epic content |
| `system_prompt` | string | No | Override system prompt (usually leave empty) |
| `path` | string | No | Path prefix for artifact inputs/outputs (e.g., `".spec"`) |
| `labels` | list | No | Tags for filtering (e.g., `["category:meta", "lang:go"]`) |
| `arguments` | list | No | User-configurable parameters |
| `nodes` | list | No | DAG of workflow tasks |

### Argument Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `key` | string | Yes | Unique identifier, accessed as `{{.Args.key}}` |
| `label` | string | Yes | Human-readable label for form field |
| `description` | string | No | Help text / placeholder |
| `type` | string | Yes | Input type: `text`, `number`, `textarea`, `select`, `multi-select` |
| `required` | bool | No | Whether the field must be filled (default: `false`) |
| `default` | string | No | Default value |
| `options` | list | Conditional | Required for `select` and `multi-select` types |

### Node Fields (DAG Tasks)

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `key` | string | Yes | Unique identifier within the workflow |
| `name` | string | Yes | Display name for the task |
| `template` | string | Yes | Filename of the markdown template |
| `assignee` | string | No | Worker role (e.g., `"worker-1"`, `"human"` for review gates) |
| `after` | list | No | Node keys this node depends on (runs after these complete) |
| `inputs` | list | No | Artifacts consumed by this node |
| `outputs` | list | No | Artifacts produced by this node |

### Artifact Fields (Inputs/Outputs)

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `key` | string | Yes | Stable identifier for template access (e.g., `{{.Inputs.plan}}`) |
| `file` | string | Yes | Filename, may contain template syntax |

---

## Complete Example

### template.yaml

```yaml
registry:
  - namespace: "workflow"
    key: "joke-contest"
    version: "v1"
    name: "Joke Contest"
    description: "Two workers write jokes in parallel, then a third judges"
    epic_template: "v1-joke-contest-epic.md"
    labels:
      - "category:meta"

    arguments:
      - key: "theme"
        label: "Joke Theme"
        description: "Optional theme or topic for the jokes"
        type: "text"
        required: false

    nodes:
      # Phase 1: Parallel joke writing
      - key: "joke-1"
        name: "Joker 1 - Write Joke"
        template: "v1-joke-contest-joke-1.md"
        assignee: "worker-1"

      - key: "joke-2"
        name: "Joker 2 - Write Joke"
        template: "v1-joke-contest-joke-2.md"
        assignee: "worker-2"

      # Phase 2: Human review gate
      - key: "review"
        name: "Human Review"
        template: "v1-human-review.md"
        assignee: "human"
        after:
          - "joke-1"
          - "joke-2"

      # Phase 3: Judge picks winner
      - key: "judge"
        name: "Judge - Pick Winner"
        template: "v1-joke-contest-judge.md"
        assignee: "worker-3"
        after:
          - "review"
```

### Epic Template (v1-joke-contest-epic.md)

Epic templates use Go template syntax and have access to workflow arguments:

```markdown
# Joke Contest: {{.Name}}

You are the **Coordinator** for a joke contest workflow.

## Context

{{- if .Args.theme}}
- **Theme:** {{.Args.theme}}
{{- else}}
- **Theme:** Any topic (no theme specified)
{{- end}}

## Your Workers

| Worker | Role | Phase |
|--------|------|-------|
| worker-1 | Joker 1 | 1 |
| worker-2 | Joker 2 | 1 |
| human | Reviewer | 2 |
| worker-3 | Judge | 3 |

**NOTE:** You (the Coordinator) are NOT a worker. Start immediately.
```

---

## Template Variables

Templates have access to these variables:

| Variable | Description |
|----------|-------------|
| `{{.Name}}` | Workflow name |
| `{{.Args.<key>}}` | User-provided argument values |
| `{{.Date}}` | Current date |
| `{{.Inputs.<key>}}` | Input artifact references |
| `{{.Outputs.<key>}}` | Output artifact references |

---

## DAG Execution

Nodes define a directed acyclic graph (DAG):

- Nodes with no `after` field start immediately (can run in parallel)
- Nodes with `after` wait for all listed dependencies to complete
- `assignee: "human"` creates a human review gate that pauses for manual approval

```
joke-1 ──┐
          ├──> review (human) ──> judge
joke-2 ──┘
```
