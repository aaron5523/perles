# Theming

Perles supports comprehensive theming with built-in presets and customizable color tokens.

## Quick Start

Use a built-in theme preset:

```yaml
theme:
  preset: catppuccin-mocha
```

---

## Available Presets

Run `perles themes` to see all available presets:

| Preset | Description |
|--------|-------------|
| `default` | Default perles theme |
| `catppuccin-mocha` | Warm, cozy dark theme |
| `catppuccin-latte` | Warm, cozy light theme |
| `dracula` | Dark theme with vibrant colors |
| `nord` | Arctic, north-bluish palette |
| `high-contrast` | High contrast for accessibility |

---

## Customizing Colors

Override specific colors while using a preset:

```yaml
theme:
  preset: dracula
  colors:
    status.error: "#FF0000"
    priority.critical: "#FF5555"
```

Or create a fully custom theme from scratch:

```yaml
theme:
  colors:
    text.primary: "#FFFFFF"
    text.muted: "#888888"
    status.success: "#00FF00"
    status.error: "#FF0000"
    border.default: "#444444"
    border.focus: "#FFFFFF"
```

---

## Color Tokens

Colors are organized by category:

### Text

| Token | Description |
|-------|-------------|
| `text.primary` | Primary text color |
| `text.secondary` | Secondary text color |
| `text.muted` | Muted/dimmed text |
| `text.description` | Description text |
| `text.placeholder` | Placeholder text |

### Border

| Token | Description |
|-------|-------------|
| `border.default` | Default border color |
| `border.focus` | Focused element border |
| `border.highlight` | Highlighted border |

### Status

| Token | Description |
|-------|-------------|
| `status.success` | Success state |
| `status.warning` | Warning state |
| `status.error` | Error state |

### Buttons

| Token | Description |
|-------|-------------|
| `button.text` | Button text color |
| `button.primary.bg` | Primary button background |
| `button.primary.focus` | Primary button focus state |
| `button.danger.bg` | Danger button background |

### Selection

| Token | Description |
|-------|-------------|
| `selection.indicator` | Selection indicator |
| `selection.background` | Selection background |

### Toasts

| Token | Description |
|-------|-------------|
| `toast.success` | Success toast |
| `toast.error` | Error toast |
| `toast.info` | Info toast |
| `toast.warn` | Warning toast |

### Issue Priority

| Token | Description |
|-------|-------------|
| `priority.critical` | P0 critical priority |
| `priority.high` | P1 high priority |
| `priority.medium` | P2 medium priority |
| `priority.low` | P3 low priority |
| `priority.backlog` | P4 backlog priority |

### Issue Status

| Token | Description |
|-------|-------------|
| `issue.status.open` | Open issue color |
| `issue.status.in_progress` | In-progress issue color |
| `issue.status.closed` | Closed issue color |

### Issue Type

| Token | Description |
|-------|-------------|
| `type.task` | Task type color |
| `type.bug` | Bug type color |
| `type.feature` | Feature type color |
| `type.epic` | Epic type color |
| `type.chore` | Chore type color |

### BQL Syntax Highlighting

| Token | Description |
|-------|-------------|
| `bql.keyword` | BQL keywords (and, or, not) |
| `bql.operator` | BQL operators (=, !=, ~) |
| `bql.field` | BQL field names |
| `bql.string` | BQL string values |
| `bql.literal` | BQL literal values |
