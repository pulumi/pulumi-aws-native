// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Comprehend
{
    [EnumType]
    public readonly struct DocumentClassifierAugmentedManifestsListItemSplit : IEquatable<DocumentClassifierAugmentedManifestsListItemSplit>
    {
        private readonly string _value;

        private DocumentClassifierAugmentedManifestsListItemSplit(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DocumentClassifierAugmentedManifestsListItemSplit Train { get; } = new DocumentClassifierAugmentedManifestsListItemSplit("TRAIN");
        public static DocumentClassifierAugmentedManifestsListItemSplit Test { get; } = new DocumentClassifierAugmentedManifestsListItemSplit("TEST");

        public static bool operator ==(DocumentClassifierAugmentedManifestsListItemSplit left, DocumentClassifierAugmentedManifestsListItemSplit right) => left.Equals(right);
        public static bool operator !=(DocumentClassifierAugmentedManifestsListItemSplit left, DocumentClassifierAugmentedManifestsListItemSplit right) => !left.Equals(right);

        public static explicit operator string(DocumentClassifierAugmentedManifestsListItemSplit value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DocumentClassifierAugmentedManifestsListItemSplit other && Equals(other);
        public bool Equals(DocumentClassifierAugmentedManifestsListItemSplit other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DocumentClassifierDocumentReaderConfigDocumentReadAction : IEquatable<DocumentClassifierDocumentReaderConfigDocumentReadAction>
    {
        private readonly string _value;

        private DocumentClassifierDocumentReaderConfigDocumentReadAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DocumentClassifierDocumentReaderConfigDocumentReadAction TextractDetectDocumentText { get; } = new DocumentClassifierDocumentReaderConfigDocumentReadAction("TEXTRACT_DETECT_DOCUMENT_TEXT");
        public static DocumentClassifierDocumentReaderConfigDocumentReadAction TextractAnalyzeDocument { get; } = new DocumentClassifierDocumentReaderConfigDocumentReadAction("TEXTRACT_ANALYZE_DOCUMENT");

        public static bool operator ==(DocumentClassifierDocumentReaderConfigDocumentReadAction left, DocumentClassifierDocumentReaderConfigDocumentReadAction right) => left.Equals(right);
        public static bool operator !=(DocumentClassifierDocumentReaderConfigDocumentReadAction left, DocumentClassifierDocumentReaderConfigDocumentReadAction right) => !left.Equals(right);

        public static explicit operator string(DocumentClassifierDocumentReaderConfigDocumentReadAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DocumentClassifierDocumentReaderConfigDocumentReadAction other && Equals(other);
        public bool Equals(DocumentClassifierDocumentReaderConfigDocumentReadAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DocumentClassifierDocumentReaderConfigDocumentReadMode : IEquatable<DocumentClassifierDocumentReaderConfigDocumentReadMode>
    {
        private readonly string _value;

        private DocumentClassifierDocumentReaderConfigDocumentReadMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DocumentClassifierDocumentReaderConfigDocumentReadMode ServiceDefault { get; } = new DocumentClassifierDocumentReaderConfigDocumentReadMode("SERVICE_DEFAULT");
        public static DocumentClassifierDocumentReaderConfigDocumentReadMode ForceDocumentReadAction { get; } = new DocumentClassifierDocumentReaderConfigDocumentReadMode("FORCE_DOCUMENT_READ_ACTION");

        public static bool operator ==(DocumentClassifierDocumentReaderConfigDocumentReadMode left, DocumentClassifierDocumentReaderConfigDocumentReadMode right) => left.Equals(right);
        public static bool operator !=(DocumentClassifierDocumentReaderConfigDocumentReadMode left, DocumentClassifierDocumentReaderConfigDocumentReadMode right) => !left.Equals(right);

        public static explicit operator string(DocumentClassifierDocumentReaderConfigDocumentReadMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DocumentClassifierDocumentReaderConfigDocumentReadMode other && Equals(other);
        public bool Equals(DocumentClassifierDocumentReaderConfigDocumentReadMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DocumentClassifierDocumentReaderConfigFeatureTypesItem : IEquatable<DocumentClassifierDocumentReaderConfigFeatureTypesItem>
    {
        private readonly string _value;

        private DocumentClassifierDocumentReaderConfigFeatureTypesItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DocumentClassifierDocumentReaderConfigFeatureTypesItem Tables { get; } = new DocumentClassifierDocumentReaderConfigFeatureTypesItem("TABLES");
        public static DocumentClassifierDocumentReaderConfigFeatureTypesItem Forms { get; } = new DocumentClassifierDocumentReaderConfigFeatureTypesItem("FORMS");

        public static bool operator ==(DocumentClassifierDocumentReaderConfigFeatureTypesItem left, DocumentClassifierDocumentReaderConfigFeatureTypesItem right) => left.Equals(right);
        public static bool operator !=(DocumentClassifierDocumentReaderConfigFeatureTypesItem left, DocumentClassifierDocumentReaderConfigFeatureTypesItem right) => !left.Equals(right);

        public static explicit operator string(DocumentClassifierDocumentReaderConfigFeatureTypesItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DocumentClassifierDocumentReaderConfigFeatureTypesItem other && Equals(other);
        public bool Equals(DocumentClassifierDocumentReaderConfigFeatureTypesItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DocumentClassifierInputDataConfigDataFormat : IEquatable<DocumentClassifierInputDataConfigDataFormat>
    {
        private readonly string _value;

        private DocumentClassifierInputDataConfigDataFormat(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DocumentClassifierInputDataConfigDataFormat ComprehendCsv { get; } = new DocumentClassifierInputDataConfigDataFormat("COMPREHEND_CSV");
        public static DocumentClassifierInputDataConfigDataFormat AugmentedManifest { get; } = new DocumentClassifierInputDataConfigDataFormat("AUGMENTED_MANIFEST");

        public static bool operator ==(DocumentClassifierInputDataConfigDataFormat left, DocumentClassifierInputDataConfigDataFormat right) => left.Equals(right);
        public static bool operator !=(DocumentClassifierInputDataConfigDataFormat left, DocumentClassifierInputDataConfigDataFormat right) => !left.Equals(right);

        public static explicit operator string(DocumentClassifierInputDataConfigDataFormat value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DocumentClassifierInputDataConfigDataFormat other && Equals(other);
        public bool Equals(DocumentClassifierInputDataConfigDataFormat other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DocumentClassifierInputDataConfigDocumentType : IEquatable<DocumentClassifierInputDataConfigDocumentType>
    {
        private readonly string _value;

        private DocumentClassifierInputDataConfigDocumentType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DocumentClassifierInputDataConfigDocumentType PlainTextDocument { get; } = new DocumentClassifierInputDataConfigDocumentType("PLAIN_TEXT_DOCUMENT");
        public static DocumentClassifierInputDataConfigDocumentType SemiStructuredDocument { get; } = new DocumentClassifierInputDataConfigDocumentType("SEMI_STRUCTURED_DOCUMENT");

        public static bool operator ==(DocumentClassifierInputDataConfigDocumentType left, DocumentClassifierInputDataConfigDocumentType right) => left.Equals(right);
        public static bool operator !=(DocumentClassifierInputDataConfigDocumentType left, DocumentClassifierInputDataConfigDocumentType right) => !left.Equals(right);

        public static explicit operator string(DocumentClassifierInputDataConfigDocumentType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DocumentClassifierInputDataConfigDocumentType other && Equals(other);
        public bool Equals(DocumentClassifierInputDataConfigDocumentType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DocumentClassifierLanguageCode : IEquatable<DocumentClassifierLanguageCode>
    {
        private readonly string _value;

        private DocumentClassifierLanguageCode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DocumentClassifierLanguageCode En { get; } = new DocumentClassifierLanguageCode("en");
        public static DocumentClassifierLanguageCode Es { get; } = new DocumentClassifierLanguageCode("es");
        public static DocumentClassifierLanguageCode Fr { get; } = new DocumentClassifierLanguageCode("fr");
        public static DocumentClassifierLanguageCode It { get; } = new DocumentClassifierLanguageCode("it");
        public static DocumentClassifierLanguageCode De { get; } = new DocumentClassifierLanguageCode("de");
        public static DocumentClassifierLanguageCode Pt { get; } = new DocumentClassifierLanguageCode("pt");

        public static bool operator ==(DocumentClassifierLanguageCode left, DocumentClassifierLanguageCode right) => left.Equals(right);
        public static bool operator !=(DocumentClassifierLanguageCode left, DocumentClassifierLanguageCode right) => !left.Equals(right);

        public static explicit operator string(DocumentClassifierLanguageCode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DocumentClassifierLanguageCode other && Equals(other);
        public bool Equals(DocumentClassifierLanguageCode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DocumentClassifierMode : IEquatable<DocumentClassifierMode>
    {
        private readonly string _value;

        private DocumentClassifierMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DocumentClassifierMode MultiClass { get; } = new DocumentClassifierMode("MULTI_CLASS");
        public static DocumentClassifierMode MultiLabel { get; } = new DocumentClassifierMode("MULTI_LABEL");

        public static bool operator ==(DocumentClassifierMode left, DocumentClassifierMode right) => left.Equals(right);
        public static bool operator !=(DocumentClassifierMode left, DocumentClassifierMode right) => !left.Equals(right);

        public static explicit operator string(DocumentClassifierMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DocumentClassifierMode other && Equals(other);
        public bool Equals(DocumentClassifierMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct FlywheelDocumentClassificationConfigMode : IEquatable<FlywheelDocumentClassificationConfigMode>
    {
        private readonly string _value;

        private FlywheelDocumentClassificationConfigMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static FlywheelDocumentClassificationConfigMode MultiClass { get; } = new FlywheelDocumentClassificationConfigMode("MULTI_CLASS");
        public static FlywheelDocumentClassificationConfigMode MultiLabel { get; } = new FlywheelDocumentClassificationConfigMode("MULTI_LABEL");

        public static bool operator ==(FlywheelDocumentClassificationConfigMode left, FlywheelDocumentClassificationConfigMode right) => left.Equals(right);
        public static bool operator !=(FlywheelDocumentClassificationConfigMode left, FlywheelDocumentClassificationConfigMode right) => !left.Equals(right);

        public static explicit operator string(FlywheelDocumentClassificationConfigMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FlywheelDocumentClassificationConfigMode other && Equals(other);
        public bool Equals(FlywheelDocumentClassificationConfigMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct FlywheelModelType : IEquatable<FlywheelModelType>
    {
        private readonly string _value;

        private FlywheelModelType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static FlywheelModelType DocumentClassifier { get; } = new FlywheelModelType("DOCUMENT_CLASSIFIER");
        public static FlywheelModelType EntityRecognizer { get; } = new FlywheelModelType("ENTITY_RECOGNIZER");

        public static bool operator ==(FlywheelModelType left, FlywheelModelType right) => left.Equals(right);
        public static bool operator !=(FlywheelModelType left, FlywheelModelType right) => !left.Equals(right);

        public static explicit operator string(FlywheelModelType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FlywheelModelType other && Equals(other);
        public bool Equals(FlywheelModelType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct FlywheelTaskConfigLanguageCode : IEquatable<FlywheelTaskConfigLanguageCode>
    {
        private readonly string _value;

        private FlywheelTaskConfigLanguageCode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static FlywheelTaskConfigLanguageCode En { get; } = new FlywheelTaskConfigLanguageCode("en");
        public static FlywheelTaskConfigLanguageCode Es { get; } = new FlywheelTaskConfigLanguageCode("es");
        public static FlywheelTaskConfigLanguageCode Fr { get; } = new FlywheelTaskConfigLanguageCode("fr");
        public static FlywheelTaskConfigLanguageCode It { get; } = new FlywheelTaskConfigLanguageCode("it");
        public static FlywheelTaskConfigLanguageCode De { get; } = new FlywheelTaskConfigLanguageCode("de");
        public static FlywheelTaskConfigLanguageCode Pt { get; } = new FlywheelTaskConfigLanguageCode("pt");

        public static bool operator ==(FlywheelTaskConfigLanguageCode left, FlywheelTaskConfigLanguageCode right) => left.Equals(right);
        public static bool operator !=(FlywheelTaskConfigLanguageCode left, FlywheelTaskConfigLanguageCode right) => !left.Equals(right);

        public static explicit operator string(FlywheelTaskConfigLanguageCode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FlywheelTaskConfigLanguageCode other && Equals(other);
        public bool Equals(FlywheelTaskConfigLanguageCode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}