# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AiPromptAiPromptApiFormat',
    'AiPromptAiPromptTemplateType',
    'AiPromptAiPromptType',
    'AssistantAssociationAssociationType',
    'AssistantType',
    'KnowledgeBaseType',
]


class AiPromptAiPromptApiFormat(str, Enum):
    ANTHROPIC_CLAUDE_MESSAGES = "ANTHROPIC_CLAUDE_MESSAGES"
    ANTHROPIC_CLAUDE_TEXT_COMPLETIONS = "ANTHROPIC_CLAUDE_TEXT_COMPLETIONS"


class AiPromptAiPromptTemplateType(str, Enum):
    TEXT = "TEXT"


class AiPromptAiPromptType(str, Enum):
    ANSWER_GENERATION = "ANSWER_GENERATION"
    INTENT_LABELING_GENERATION = "INTENT_LABELING_GENERATION"
    QUERY_REFORMULATION = "QUERY_REFORMULATION"


class AssistantAssociationAssociationType(str, Enum):
    KNOWLEDGE_BASE = "KNOWLEDGE_BASE"


class AssistantType(str, Enum):
    AGENT = "AGENT"


class KnowledgeBaseType(str, Enum):
    EXTERNAL = "EXTERNAL"
    CUSTOM = "CUSTOM"
    MESSAGE_TEMPLATES = "MESSAGE_TEMPLATES"
