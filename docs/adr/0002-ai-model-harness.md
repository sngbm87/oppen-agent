# ADR 0002: AI Model Harness as First-Class Citizen

**Date**: 2026-05-28
**Status**: Accepted
**Context**: Most agents treat models as simple API clients. For Tier 1 quality and real optimization we need routing, quantization, cost governance, guardrails, and telemetry as a dedicated, observable, and extensible harness.

**Decision**: Create `internal/harness/models/` with clear sub-components: Router/Selector, Quantizer, Guardrails, Telemetry, ContextManager. All model calls go through this harness.

**Consequences**:
- Positive: Cost control, performance optimization, safety, A/B testing, easy addition of new models (local or cloud)
- Positive: Full observability of model behavior (latency, tokens, success rate, cost)
- Negative: One extra layer of abstraction (worth it for the gains)

---