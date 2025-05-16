# A FOSS Story of Golangci-lint v2: From the Idea to the IDE

## Status

### ❓ On Evaluation

## Description

Golangci-lint, the most popular Go static analysis tool, recently introduced a major upgrade.
This talk explores the motivation behind the changes, the designs that shaped the transition, and how the ecosystem adapts.

Creating the first major version with 7 years of history is complex and requires much time.
As golangci-lint is widely used, we needed to support the change as much as possible to ease the migration's user experience and avoid the "endless wishlist" trap.
At the same time, we needed to add value with new features.
All those elements were challenges to create this new release and exposed some interesting thoughts about handling those from the perspective of a FOSS maintainer.

Golangci-lint v2 also involves adaptations in the ecosystem tools.
We'll share the story of migrating the VS Code integration, highlighting concrete issues, why they mattered, and how the community resolved them.
Through this story of open-source collaboration, you'll gain practical insight into migrating your projects to golangci-lint v2.

Join us for a deep dive into this milestone upgrade through the collaborative story in the open-source world!

---

Outline:

The journey to Golangci-lint v2 (Ludovic's part)

1. A quick history of golangci-lint
2. To be breaking or not breaking?
   - 6 years is too long
   - Breaking but smoothly
   - The `migrate` command
3. A formatter is a linter?
   - Define the definitions
   - Formatters
   - The `fmt` command
4. Exclusion of the default exclusions
   - Follow the standard or common practice?
   - new configuration (presets)
5. The fight between options to enable/disable/filter linters
   - The "fast" problem
   - Multiple options on the same topic that cannot be used together constitute a single option.
6. The hell of the file path management inside the configuration
   - Matching vs user-facing
   - Exclusions (linters vs formatters)
   - Configuration path matching
   - Linters' configuration path matching

How to adapt to Golangci-lint v2 (Takuto's part)

1. Breaking changes in Golangci-lint v2 (Recap of Ludovic's section)
   - Config file structure overhaul
   - Partial modifications to the default config
   - Deprecated flags removed
2. Inevitably broken tools, but why?
   - Many tools relied on output flags to control formatting for result parsers.
   - With those flags gone, using them now triggers errors.
   - Examples: GoLand, reviewdog/action-golangci-lint, VS Code
3. The migration of the VS Code extension (<https://go-review.googlesource.com/c/vscode-go/+/660415>)
   - Minimal fix: Replace deleted flags
   - Advanced requirement: Support toggling between v1 and v2 per workspace
   - Planned solution: Addition of the new installable tool "golangci-lint-v2" / Addition of the new lintTool "golangci-lint-v2" / Make "golangci-lint" option work with v2
   - New installable tool: The extension installs the binary in a temp dir, then moves it to GOBIN under a versioned name.
   - Change in the lintTool: When using the "golangci-lint" option, the extension detects the version and applies the right flags.
4. The "path-mode" maze (<https://go-review.googlesource.com/c/vscode-go/+/664877> / <https://go-review.googlesource.com/c/vscode-go/+/667735>)
   - Golangci-lint changed the default "relative-path-mode" from "wd" to "cfg", breaking error display in file/package linting.
   - Paths in the output no longer match files, breaking diagnostic errors.
   - Solution: Use the "path-mode" option to force absolute paths
5. What users should do
   - Check if your tools support v2. If not, staying on v1 is fine.
   - To migrate, upgrade to v2, and run "golangci-lint migrate" to auto-convert the config file.
   - If you unnecessarily use flags in your original scripts or tools, replace them with options in the config file.

## What type of submission is this?

60 minute Talk

## Speaker Biography

Takuto is an undergraduate student at Chiba Institute of Technology in Japan. He pursues a degree in computer science to become a cloud platform developer. His expertise lies in web development (backend) and cloud infrastructure.  
He organizes Gophers EX, a project to support the overseas expansion of the Japan Go community. With his experience as a speaker at an international conference, he provides practical methods to challenge conferences overseas.
