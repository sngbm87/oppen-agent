# Architecture - Oppen Agent

**Version**: 0.1.0
**Status**: Draft

## 1. High-Level Architecture
Hexagonal / Ports & Adapters for maximum testability and swapability of harnesses.

Core: Go daemon (cmd/oppen-agent)
Brain: Eshkol (cognitive, AD, HoTT, consciousness)
Quantum: QGTL/Moonlab hybrid layer
Harnesses: Pluggable (Model, Context, Tool, Safety, Observability, Self-Evolution)
Polyglot: FFI to Rust/Python/TS/Eshkol/C/C++

## 2. Key Harnesses (AI-Model Optimized)
- **Model Harness**: Router, quantizer, batcher, cost tracker, fallback. Supports all major providers + local (llama.cpp, vLLM, Ollama).
- **Context Harness**: Hierarchical memory, RAG, compression (geometric), token budget, eviction.
- **Tool Harness**: Registry, validator, executor (sandboxed), parallel, retry, trace.
- **Safety Harness**: Guardrails, hallucination detector (Witness), PII, policy engine.
- **Observability Harness**: OpenTelemetry tracing, metrics, replay.
- **Self-Evolution Harness**: Trace analysis, prompt/skill optimization, benchmark-driven improvement.

## 3. Data Flow
User/Messaging -> Daemon -> Agent Loop (Plan-Execute-Reflect) -> Model Harness -> Tool Harness -> Memory -> Verification -> Response

## 4. Technology Choices
- Runtime: Go (daemon, orchestration) + Rust (performance critical)
- Brain: Eshkol
- Polyglot: Python (tools/ecosystem), TypeScript (UI/messaging), Swift/Kotlin (edge)
- Memory: Geometric tensors + λ-Memory + vector DB
- Deployment: Docker, systemd, K8s ready

## 5. Quality Attributes
- Reliability: Circuit breakers, retries, checkpointing
- Performance: Sub-second for most operations, efficient token use
- Security: OAuth, sandbox, encryption at rest, audit
- Maintainability: Clean architecture, ADRs, full tests
- Extensibility: Plugin system for new models/tools/benchmarks

## 6. Decision Records
See docs/adr/

## 7. Diagrams
(PlantUML or Mermaid to be added in later iterations)

---
*Living document. Updated per stage.*