from __future__ import annotations

import re
import sys
from pathlib import Path


ROOT = Path(__file__).resolve().parents[3]


REQUIRED_ARTIFACTS = [
    "AGENTS.md",
    "docs/ai/manifest.yaml",
    "docs/ai/runtime-rule-index.md",
    "docs/ai/status/current.yaml",
    "docs/ai/requirements/ledger.yaml",
    "docs/ai/tasks/bootstrap-tree.yaml",
    "docs/ai/tasks/forest.yaml",
    "docs/ai/tasks/main-tree.yaml",
    "docs/ai/tasks/branches/vibe-coding-infra/tree.yaml",
    "docs/ai/tasks/current-slice.yaml",
    "docs/ai/capabilities/index.md",
    "docs/ai/capabilities/dependency-candidates.yaml",
    "docs/ai/evidence/index.md",
    "docs/ai/knowledge/index.md",
    "docs/ai/skills/index.md",
    "docs/ai/skills/full-project-lifecycle/SKILL.md",
    "docs/ai/templates/full-project-lifecycle-workflow-sop.md",
    "docs/ai/analysis/mainline-full-project-lifecycle-gap-analysis.md",
    "docs/ai/architecture/full-project-lifecycle-workflow.md",
    "docs/ai/knowledge/entries/kb_007_full_project_lifecycle.md",
    "docs/ai/skills/project-requirement-discovery/SKILL.md",
    "docs/ai/templates/project-requirement-discovery-sop.md",
    "docs/ai/templates/project-requirement-discovery-record.yaml",
    "docs/ai/requirements/intake/README.md",
    "docs/ai/requirements/workflow_engine.yaml",
    "docs/ai/requirements/template_discovery.yaml",
    "docs/ai/requirements/state_machine.yaml",
    "docs/ai/skills/requirement-workflow-engine/SKILL.md",
    "docs/ai/templates/requirement-workflow-engine-sop.md",
    "docs/ai/analysis/requirement-workflow-engine-gap-analysis.md",
    "docs/ai/architecture/requirement-workflow-engine.md",
    "docs/ai/knowledge/entries/kb_011_requirement_workflow_engine.md",
    "docs/ai/analysis/project-requirement-discovery-gap-analysis.md",
    "docs/ai/architecture/project-requirement-discovery-workflow.md",
    "docs/ai/knowledge/entries/kb_009_project_requirement_discovery.md",
    "docs/ai/skills/project-lifecycle-downstream-detailing/SKILL.md",
    "docs/ai/templates/project-lifecycle-downstream-gates-sop.md",
    "docs/ai/templates/project-lifecycle-downstream-record.yaml",
    "docs/ai/lifecycle/README.md",
    "docs/ai/analysis/project-lifecycle-downstream-gap-analysis.md",
    "docs/ai/architecture/project-lifecycle-downstream-workflow.md",
    "docs/ai/knowledge/entries/kb_010_project_lifecycle_downstream.md",
    "docs/ai/skills/compiler-runtime-assimilation/SKILL.md",
    "docs/ai/templates/compiler-runtime-assimilation-sop.md",
    "docs/ai/analysis/compiler-runtime-assimilation.md",
    "docs/ai/architecture/compiler-runtime-assimilation.md",
    "docs/ai/knowledge/entries/kb_008_compiler_runtime_assimilation.md",
    "docs/ai/handoff/current.md",
]

CATALOG_ONLY_FORBIDDEN_ARTIFACTS = [
    "package.json",
    "package-lock.json",
    "pnpm-lock.yaml",
    "yarn.lock",
    "pyproject.toml",
    "requirements.txt",
    "src",
    "app",
    "components",
    "tools",
]


def read_text(path: str) -> str:
    return (ROOT / path).read_text(encoding="utf-8")


def fail(message: str) -> None:
    print(f"Runtime validation failed: {message}", file=sys.stderr)
    raise SystemExit(1)


def extract_yaml_scalar(text: str, field: str, source: str) -> str:
    match = re.search(
        rf"(?m)^{re.escape(field)}:\s*([A-Za-z0-9_./\-\[\]]+)\s*$", text
    )
    if not match:
        fail(f"{source} must define {field}.")
    return match.group(1)


def main() -> None:
    missing = [path for path in REQUIRED_ARTIFACTS if not (ROOT / path).exists()]
    if missing:
        fail(f"Missing required runtime artifacts: {', '.join(missing)}")

    agents = read_text("AGENTS.md")
    if "No Prompt Runtime Dependency" not in agents:
        fail("AGENTS.md must include the no prompt runtime dependency section.")

    status = read_text("docs/ai/status/current.yaml")
    current_slice = read_text("docs/ai/tasks/current-slice.yaml")

    status_slice = extract_yaml_scalar(
        status, "current_slice", "docs/ai/status/current.yaml"
    )
    status_branch = extract_yaml_scalar(
        status, "current_branch", "docs/ai/status/current.yaml"
    )
    status_tree = extract_yaml_scalar(
        status, "current_tree", "docs/ai/status/current.yaml"
    )
    slice_id = extract_yaml_scalar(
        current_slice, "id", "docs/ai/tasks/current-slice.yaml"
    )
    slice_branch = extract_yaml_scalar(
        current_slice, "branch_id", "docs/ai/tasks/current-slice.yaml"
    )
    if status_slice != slice_id:
        fail(
            "docs/ai/status/current.yaml current_slice does not match "
            "docs/ai/tasks/current-slice.yaml id."
        )
    if status_branch != slice_branch:
        fail(
            "docs/ai/status/current.yaml current_branch does not match "
            "docs/ai/tasks/current-slice.yaml branch_id."
        )
    if not (ROOT / status_tree).exists():
        fail(f"current_tree does not exist: {status_tree}")

    candidates = read_text("docs/ai/capabilities/dependency-candidates.yaml")
    if "install_allowed: false" not in candidates:
        fail("Capability candidates must keep install_allowed: false.")
    if "scaffold_allowed: false" not in candidates:
        fail("Capability candidates must keep scaffold_allowed: false.")
    forbidden_existing = [
        path for path in CATALOG_ONLY_FORBIDDEN_ARTIFACTS if (ROOT / path).exists()
    ]
    if forbidden_existing:
        fail(
            "Catalog-only strategy forbids these root artifacts: "
            + ", ".join(forbidden_existing)
        )

    lifecycle_text = "\n".join(
        [
            read_text("docs/ai/runtime-rule-index.md"),
            read_text("docs/ai/skills/full-project-lifecycle/SKILL.md"),
            read_text("docs/ai/templates/full-project-lifecycle-workflow-sop.md"),
            read_text("docs/ai/knowledge/entries/kb_007_full_project_lifecycle.md"),
        ]
    )
    required_lifecycle_terms = [
        "task_line_routing_gate",
        "idea_intake_gate",
        "requirement_analysis_gate",
        "requirement_baseline_confirmation_gate",
        "current_source_research_gate",
        "research_confirmation_gate",
        "task_analysis_gate",
        "task_analysis_confirmation_gate",
        "architecture_and_stack_design_gate",
        "architecture_confirmation_gate",
        "infra_mode_recommendation_gate",
        "infra_mode_confirmation_gate",
        "agent_driving_infra_gate",
        "task_tree_and_slice_gate",
        "standard_light_risk_escalated",
        "enterprise_high_assurance",
        "implementation_test_docs_state_loop",
        "next_step_protocol_gate",
        "round_closure_gate",
        "next_round_intake[n+1]",
    ]
    missing_lifecycle_terms = [
        term for term in required_lifecycle_terms if term not in lifecycle_text
    ]
    if missing_lifecycle_terms:
        fail(
            "Full-project lifecycle artifacts are missing required terms: "
            + ", ".join(missing_lifecycle_terms)
        )

    discovery_text = "\n".join(
        [
            read_text("docs/ai/runtime-rule-index.md"),
            read_text("docs/ai/skills/project-requirement-discovery/SKILL.md"),
            read_text("docs/ai/templates/project-requirement-discovery-sop.md"),
            read_text("docs/ai/templates/project-requirement-discovery-record.yaml"),
            read_text("docs/ai/requirements/workflow_engine.yaml"),
            read_text("docs/ai/requirements/template_discovery.yaml"),
            read_text("docs/ai/requirements/state_machine.yaml"),
            read_text(
                "docs/ai/knowledge/entries/kb_009_project_requirement_discovery.md"
            ),
            read_text("docs/ai/architecture/project-requirement-discovery-workflow.md"),
        ]
    )
    required_discovery_terms = [
        "project_requirement_discovery",
        "idea_seed_intake_gate",
        "idea_interpretation_gate",
        "domain_surface_mapping_gate",
        "requirement_inventory_gate",
        "question_backlog_gate",
        "requirement_collection_round_gate",
        "requirement_plan_gate",
        "requirement_persistence_sync_gate",
        "baseline_readiness_gate",
        "requirement_discovery_record",
        "skill_project_requirement_discovery",
        "kb_009",
    ]
    missing_discovery_terms = [
        term for term in required_discovery_terms if term not in discovery_text
    ]
    if missing_discovery_terms:
        fail(
            "Project requirement discovery artifacts are missing required terms: "
            + ", ".join(missing_discovery_terms)
        )

    requirement_engine_text = "\n".join(
        [
            read_text("docs/ai/runtime-rule-index.md"),
            read_text("docs/ai/requirements/workflow_engine.yaml"),
            read_text("docs/ai/requirements/template_discovery.yaml"),
            read_text("docs/ai/requirements/state_machine.yaml"),
            read_text("docs/ai/skills/requirement-workflow-engine/SKILL.md"),
            read_text("docs/ai/templates/requirement-workflow-engine-sop.md"),
            read_text("docs/ai/analysis/requirement-workflow-engine-gap-analysis.md"),
            read_text("docs/ai/architecture/requirement-workflow-engine.md"),
            read_text("docs/ai/knowledge/entries/kb_011_requirement_workflow_engine.md"),
        ]
    )
    required_requirement_engine_terms = [
        "requirement_workflow_engine",
        "RequirementDiscoveryWorkflowEngine",
        "RequirementDiscoveryTemplateCatalog",
        "RequirementDiscoveryStateMachine",
        "REQUIREMENTS_GATHERING",
        "code_generation_allowed: false",
        "template_discovery.yaml",
        "workflow_engine.yaml",
        "state_machine.yaml",
        "commerce",
        "saas",
        "blog",
        "generic_app",
        "skill_requirement_workflow_engine",
        "kb_011",
    ]
    missing_requirement_engine_terms = [
        term
        for term in required_requirement_engine_terms
        if term not in requirement_engine_text
    ]
    if missing_requirement_engine_terms:
        fail(
            "Requirement workflow engine artifacts are missing required terms: "
            + ", ".join(missing_requirement_engine_terms)
        )

    downstream_text = "\n".join(
        [
            read_text("docs/ai/runtime-rule-index.md"),
            read_text("docs/ai/skills/project-lifecycle-downstream-detailing/SKILL.md"),
            read_text("docs/ai/templates/project-lifecycle-downstream-gates-sop.md"),
            read_text("docs/ai/templates/project-lifecycle-downstream-record.yaml"),
            read_text("docs/ai/lifecycle/README.md"),
            read_text(
                "docs/ai/knowledge/entries/kb_010_project_lifecycle_downstream.md"
            ),
            read_text("docs/ai/architecture/project-lifecycle-downstream-workflow.md"),
        ]
    )
    required_downstream_terms = [
        "project_lifecycle_downstream",
        "research_plan_gate",
        "current_source_research_execution_gate",
        "research_synthesis_confirmation_gate",
        "task_graph_analysis_gate",
        "task_analysis_confirmation_packet_gate",
        "architecture_dossier_gate",
        "architecture_confirmation_packet_gate",
        "infra_mode_risk_matrix_gate",
        "infra_mode_confirmation_packet_gate",
        "agent_driving_infra_plan_gate",
        "task_tree_slice_contract_gate",
        "implementation_iteration_ledger_gate",
        "verification_evidence_packet_gate",
        "acceptance_closure_packet_gate",
        "next_round_reentry_gate",
        "project_lifecycle_downstream_record",
        "skill_project_lifecycle_downstream_detailing",
        "kb_010",
    ]
    missing_downstream_terms = [
        term for term in required_downstream_terms if term not in downstream_text
    ]
    if missing_downstream_terms:
        fail(
            "Project lifecycle downstream artifacts are missing required terms: "
            + ", ".join(missing_downstream_terms)
        )

    assimilation_text = "\n".join(
        [
            read_text("docs/ai/runtime-rule-index.md"),
            read_text("docs/ai/skills/compiler-runtime-assimilation/SKILL.md"),
            read_text("docs/ai/templates/compiler-runtime-assimilation-sop.md"),
            read_text(
                "docs/ai/knowledge/entries/kb_008_compiler_runtime_assimilation.md"
            ),
            read_text("docs/ai/architecture/compiler-runtime-assimilation.md"),
        ]
    )
    required_assimilation_terms = [
        "compiler_runtime_assimilation",
        "INV-18",
        "compiler_context_intake_gate",
        "repository_fit_gate",
        "semantic_layer_extraction_gate",
        "runtime_artifact_translation_gate",
        "no_prompt_runtime_dependency_recheck_gate",
        "validation_and_evidence_gate",
        "handoff_and_next_condition_gate",
        "skill_compiler_runtime_assimilation",
        "kb_008",
    ]
    missing_assimilation_terms = [
        term for term in required_assimilation_terms if term not in assimilation_text
    ]
    if missing_assimilation_terms:
        fail(
            "Compiler-runtime assimilation artifacts are missing required terms: "
            + ", ".join(missing_assimilation_terms)
        )

    skills_index = read_text("docs/ai/skills/index.md")
    knowledge_index = read_text("docs/ai/knowledge/index.md")
    if "skill_full_project_lifecycle" not in skills_index:
        fail("Skill index must reference skill_full_project_lifecycle.")
    if "kb_007" not in knowledge_index:
        fail("Knowledge index must reference kb_007.")
    if "skill_project_requirement_discovery" not in skills_index:
        fail("Skill index must reference skill_project_requirement_discovery.")
    if "kb_009" not in knowledge_index:
        fail("Knowledge index must reference kb_009.")
    if "skill_requirement_workflow_engine" not in skills_index:
        fail("Skill index must reference skill_requirement_workflow_engine.")
    if "kb_011" not in knowledge_index:
        fail("Knowledge index must reference kb_011.")
    if "skill_project_lifecycle_downstream_detailing" not in skills_index:
        fail(
            "Skill index must reference "
            "skill_project_lifecycle_downstream_detailing."
        )
    if "kb_010" not in knowledge_index:
        fail("Knowledge index must reference kb_010.")
    if "skill_compiler_runtime_assimilation" not in skills_index:
        fail("Skill index must reference skill_compiler_runtime_assimilation.")
    if "kb_008" not in knowledge_index:
        fail("Knowledge index must reference kb_008.")

    forest = read_text("docs/ai/tasks/forest.yaml")
    selected_tree = read_text(status_tree)
    if status_branch not in forest:
        fail("Task forest does not contain the current branch.")
    if status_tree not in forest:
        fail("Task forest does not reference the current tree path.")
    if f"branch_id: {status_branch}" not in selected_tree:
        fail("Current tree does not declare the current branch_id.")

    extract_yaml_scalar(
        status, "next_allowed_phase", "docs/ai/status/current.yaml"
    )

    runtime_text = "\n".join(
        [
            read_text("AGENTS.md"),
            read_text("docs/ai/runtime-rule-index.md"),
            read_text("docs/ai/handoff/current.md"),
        ]
    )
    forbidden_prompt_dependency = re.compile(
        r"(must|should|required to|requires)\s+read(ing)?\s+the\s+original\s+compiler\s+prompt",
        re.IGNORECASE,
    )
    if forbidden_prompt_dependency.search(runtime_text):
        fail("Runtime artifacts appear to require reading the original compiler prompt.")

    print("Runtime validation passed.")


if __name__ == "__main__":
    main()
