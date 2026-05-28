# ADR 0001: Initial Architecture - Hexagonal + DDD + AI-Model Optimized Harnesses

**Date**: 2026-05-28
**Status**: Accepted
**Context**: We need a production-grade, testable, extensible foundation for a multi-model, self-evolving autonomous coding agent that can scale to swarm level while remaining observable and safe.

**Decision**: Use Hexagonal Architecture (Ports & Adapters) + Domain-Driven Design with explicit harness boundaries. Core domain (Task, Agent, Memory, Tool, Model, Verification) is isolated. All AI Model concerns (routing, context, safety, observability) are first-class harnesses that can be swapped or extended.

**Consequences**:
- Positive: High testability, clear boundaries, easy to add new models/harnesses/benchmarks
- Positive: AI Model optimization (quantization, routing, guardrails) is explicit and measurable
- Negative: Slightly more boilerplate initially (mitigated by code generation and strong conventions)

**Alternatives Considered**:
- Pure Clean Architecture: Too heavy for our needs
- Simple layered: Insufficient boundaries for complex harness interactions
- Actor model (Erlang/Elixir): Interesting but higher operational complexity for polyglot FFI

---