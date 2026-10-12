# KubeCon + CloudNativeCon Brief

Conference-specific knowledge for the generic proposal agents (`Orchestrator`, `Writer`, `Reviewer`). Generic definitions stay reusable; everything KubeCon-specific lives here. Pass this file as the conference brief per task.

## Event

- KubeCon + CloudNativeCon Europe 2027, Barcelona, 15-18 March 2027.
- CFP: <https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/program/cfp/>
- Event page: <https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/>
- Caution (2026-10-12): the CFP detail page returned 404 once; re-verify field specs against the live CFP before submitting.

## Sessionize Fields

- Session Title: Title Case, shown on the schedule as-is.
- Description: max 1000 characters. Voice: third person (origin: reviewer-cited CFP; verify against the live CFP before submitting).
- Submitting for: KubeCon + CloudNativeCon vs CNCF Project Opportunity.
- Poster Session: Yes/No selector, no extra text.
- Value to the community: the heaviest-weighted field.
- Case study: Yes/No selector, no extra text. CFP defines a case study as "a report on an organization's real-world implementation of a practice, product, system, service, or combination thereof, and how well it worked" — an individual contribution story does not qualify (verified 2026-10-12).
- Talk presented before: Yes/No selector, no extra text.
- Relevant CNCF/Open Source Projects: list featured projects.
- Bio: human-owned. Agents never touch it.

## Speaker Constraints on File (2026-10-12)

- Eligible format: Session Presentation (30 minutes) only; no Project LT eligibility. Include a concrete 30-minute structure in Description.

## Strictness

High. KubeCon is very expensive: Value must prove worth-the-ticket with Monday-morning actions or reusable artifacts (procedures, templates, checklists, flowcharts). Title + Description + Value must hook a busy reviewer as a set.

## Standing Exceptions

- Rule of Three: not flagged when the items are real content.
- textlint `slopless/sentence-case` vs CFP Title Case: known config conflict, excluded from content failure.

## Maintenance

- Edition-specific facts (city, dates, CFP URLs) live under Event and must be refreshed per edition; the rest applies KubeCon-wide.
