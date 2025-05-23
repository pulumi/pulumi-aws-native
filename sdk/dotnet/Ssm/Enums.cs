// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Ssm
{
    /// <summary>
    /// The severity level that is assigned to the association.
    /// </summary>
    [EnumType]
    public readonly struct AssociationComplianceSeverity : IEquatable<AssociationComplianceSeverity>
    {
        private readonly string _value;

        private AssociationComplianceSeverity(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AssociationComplianceSeverity Critical { get; } = new AssociationComplianceSeverity("CRITICAL");
        public static AssociationComplianceSeverity High { get; } = new AssociationComplianceSeverity("HIGH");
        public static AssociationComplianceSeverity Medium { get; } = new AssociationComplianceSeverity("MEDIUM");
        public static AssociationComplianceSeverity Low { get; } = new AssociationComplianceSeverity("LOW");
        public static AssociationComplianceSeverity Unspecified { get; } = new AssociationComplianceSeverity("UNSPECIFIED");

        public static bool operator ==(AssociationComplianceSeverity left, AssociationComplianceSeverity right) => left.Equals(right);
        public static bool operator !=(AssociationComplianceSeverity left, AssociationComplianceSeverity right) => !left.Equals(right);

        public static explicit operator string(AssociationComplianceSeverity value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AssociationComplianceSeverity other && Equals(other);
        public bool Equals(AssociationComplianceSeverity other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The mode for generating association compliance. You can specify `AUTO` or `MANUAL` . In `AUTO` mode, the system uses the status of the association execution to determine the compliance status. If the association execution runs successfully, then the association is `COMPLIANT` . If the association execution doesn't run successfully, the association is `NON-COMPLIANT` .
    /// 
    /// In `MANUAL` mode, you must specify the `AssociationId` as a parameter for the `PutComplianceItems` API action. In this case, compliance data is not managed by State Manager. It is managed by your direct call to the `PutComplianceItems` API action.
    /// 
    /// By default, all associations use `AUTO` mode.
    /// </summary>
    [EnumType]
    public readonly struct AssociationSyncCompliance : IEquatable<AssociationSyncCompliance>
    {
        private readonly string _value;

        private AssociationSyncCompliance(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AssociationSyncCompliance Auto { get; } = new AssociationSyncCompliance("AUTO");
        public static AssociationSyncCompliance Manual { get; } = new AssociationSyncCompliance("MANUAL");

        public static bool operator ==(AssociationSyncCompliance left, AssociationSyncCompliance right) => left.Equals(right);
        public static bool operator !=(AssociationSyncCompliance left, AssociationSyncCompliance right) => !left.Equals(right);

        public static explicit operator string(AssociationSyncCompliance value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AssociationSyncCompliance other && Equals(other);
        public bool Equals(AssociationSyncCompliance other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The key of a key-value pair that identifies the location of an attachment to a document.
    /// </summary>
    [EnumType]
    public readonly struct DocumentAttachmentsSourceKey : IEquatable<DocumentAttachmentsSourceKey>
    {
        private readonly string _value;

        private DocumentAttachmentsSourceKey(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DocumentAttachmentsSourceKey SourceUrl { get; } = new DocumentAttachmentsSourceKey("SourceUrl");
        public static DocumentAttachmentsSourceKey S3FileUrl { get; } = new DocumentAttachmentsSourceKey("S3FileUrl");
        public static DocumentAttachmentsSourceKey AttachmentReference { get; } = new DocumentAttachmentsSourceKey("AttachmentReference");

        public static bool operator ==(DocumentAttachmentsSourceKey left, DocumentAttachmentsSourceKey right) => left.Equals(right);
        public static bool operator !=(DocumentAttachmentsSourceKey left, DocumentAttachmentsSourceKey right) => !left.Equals(right);

        public static explicit operator string(DocumentAttachmentsSourceKey value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DocumentAttachmentsSourceKey other && Equals(other);
        public bool Equals(DocumentAttachmentsSourceKey other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specify the document format for the request. The document format can be either JSON or YAML. JSON is the default format.
    /// </summary>
    [EnumType]
    public readonly struct DocumentFormat : IEquatable<DocumentFormat>
    {
        private readonly string _value;

        private DocumentFormat(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DocumentFormat Yaml { get; } = new DocumentFormat("YAML");
        public static DocumentFormat Json { get; } = new DocumentFormat("JSON");
        public static DocumentFormat Text { get; } = new DocumentFormat("TEXT");

        public static bool operator ==(DocumentFormat left, DocumentFormat right) => left.Equals(right);
        public static bool operator !=(DocumentFormat left, DocumentFormat right) => !left.Equals(right);

        public static explicit operator string(DocumentFormat value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DocumentFormat other && Equals(other);
        public bool Equals(DocumentFormat other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of document to create.
    /// </summary>
    [EnumType]
    public readonly struct DocumentType : IEquatable<DocumentType>
    {
        private readonly string _value;

        private DocumentType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DocumentType ApplicationConfiguration { get; } = new DocumentType("ApplicationConfiguration");
        public static DocumentType ApplicationConfigurationSchema { get; } = new DocumentType("ApplicationConfigurationSchema");
        public static DocumentType Automation { get; } = new DocumentType("Automation");
        public static DocumentType AutomationChangeTemplate { get; } = new DocumentType("Automation.ChangeTemplate");
        public static DocumentType AutoApprovalPolicy { get; } = new DocumentType("AutoApprovalPolicy");
        public static DocumentType ChangeCalendar { get; } = new DocumentType("ChangeCalendar");
        public static DocumentType CloudFormation { get; } = new DocumentType("CloudFormation");
        public static DocumentType Command { get; } = new DocumentType("Command");
        public static DocumentType DeploymentStrategy { get; } = new DocumentType("DeploymentStrategy");
        public static DocumentType ManualApprovalPolicy { get; } = new DocumentType("ManualApprovalPolicy");
        public static DocumentType Package { get; } = new DocumentType("Package");
        public static DocumentType Policy { get; } = new DocumentType("Policy");
        public static DocumentType ProblemAnalysis { get; } = new DocumentType("ProblemAnalysis");
        public static DocumentType ProblemAnalysisTemplate { get; } = new DocumentType("ProblemAnalysisTemplate");
        public static DocumentType Session { get; } = new DocumentType("Session");

        public static bool operator ==(DocumentType left, DocumentType right) => left.Equals(right);
        public static bool operator !=(DocumentType left, DocumentType right) => !left.Equals(right);

        public static explicit operator string(DocumentType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DocumentType other && Equals(other);
        public bool Equals(DocumentType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Update method - when set to 'Replace', the update will replace the existing document; when set to 'NewVersion', the update will create a new version.
    /// </summary>
    [EnumType]
    public readonly struct DocumentUpdateMethod : IEquatable<DocumentUpdateMethod>
    {
        private readonly string _value;

        private DocumentUpdateMethod(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DocumentUpdateMethod Replace { get; } = new DocumentUpdateMethod("Replace");
        public static DocumentUpdateMethod NewVersion { get; } = new DocumentUpdateMethod("NewVersion");

        public static bool operator ==(DocumentUpdateMethod left, DocumentUpdateMethod right) => left.Equals(right);
        public static bool operator !=(DocumentUpdateMethod left, DocumentUpdateMethod right) => !left.Equals(right);

        public static explicit operator string(DocumentUpdateMethod value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DocumentUpdateMethod other && Equals(other);
        public bool Equals(DocumentUpdateMethod other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The data type of the parameter, such as ``text`` or ``aws:ec2:image``. The default is ``text``.
    /// </summary>
    [EnumType]
    public readonly struct ParameterDataType : IEquatable<ParameterDataType>
    {
        private readonly string _value;

        private ParameterDataType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ParameterDataType Text { get; } = new ParameterDataType("text");
        public static ParameterDataType Awsec2image { get; } = new ParameterDataType("aws:ec2:image");

        public static bool operator ==(ParameterDataType left, ParameterDataType right) => left.Equals(right);
        public static bool operator !=(ParameterDataType left, ParameterDataType right) => !left.Equals(right);

        public static explicit operator string(ParameterDataType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ParameterDataType other && Equals(other);
        public bool Equals(ParameterDataType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The parameter tier.
    /// </summary>
    [EnumType]
    public readonly struct ParameterTier : IEquatable<ParameterTier>
    {
        private readonly string _value;

        private ParameterTier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ParameterTier Standard { get; } = new ParameterTier("Standard");
        public static ParameterTier Advanced { get; } = new ParameterTier("Advanced");
        public static ParameterTier IntelligentTiering { get; } = new ParameterTier("Intelligent-Tiering");

        public static bool operator ==(ParameterTier left, ParameterTier right) => left.Equals(right);
        public static bool operator !=(ParameterTier left, ParameterTier right) => !left.Equals(right);

        public static explicit operator string(ParameterTier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ParameterTier other && Equals(other);
        public bool Equals(ParameterTier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of parameter.
    /// </summary>
    [EnumType]
    public readonly struct ParameterType : IEquatable<ParameterType>
    {
        private readonly string _value;

        private ParameterType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ParameterType String { get; } = new ParameterType("String");
        public static ParameterType StringList { get; } = new ParameterType("StringList");

        public static bool operator ==(ParameterType left, ParameterType right) => left.Equals(right);
        public static bool operator !=(ParameterType left, ParameterType right) => !left.Equals(right);

        public static explicit operator string(ParameterType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ParameterType other && Equals(other);
        public bool Equals(ParameterType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. The default value is UNSPECIFIED.
    /// </summary>
    [EnumType]
    public readonly struct PatchBaselineApprovedPatchesComplianceLevel : IEquatable<PatchBaselineApprovedPatchesComplianceLevel>
    {
        private readonly string _value;

        private PatchBaselineApprovedPatchesComplianceLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PatchBaselineApprovedPatchesComplianceLevel Critical { get; } = new PatchBaselineApprovedPatchesComplianceLevel("CRITICAL");
        public static PatchBaselineApprovedPatchesComplianceLevel High { get; } = new PatchBaselineApprovedPatchesComplianceLevel("HIGH");
        public static PatchBaselineApprovedPatchesComplianceLevel Medium { get; } = new PatchBaselineApprovedPatchesComplianceLevel("MEDIUM");
        public static PatchBaselineApprovedPatchesComplianceLevel Low { get; } = new PatchBaselineApprovedPatchesComplianceLevel("LOW");
        public static PatchBaselineApprovedPatchesComplianceLevel Informational { get; } = new PatchBaselineApprovedPatchesComplianceLevel("INFORMATIONAL");
        public static PatchBaselineApprovedPatchesComplianceLevel Unspecified { get; } = new PatchBaselineApprovedPatchesComplianceLevel("UNSPECIFIED");

        public static bool operator ==(PatchBaselineApprovedPatchesComplianceLevel left, PatchBaselineApprovedPatchesComplianceLevel right) => left.Equals(right);
        public static bool operator !=(PatchBaselineApprovedPatchesComplianceLevel left, PatchBaselineApprovedPatchesComplianceLevel right) => !left.Equals(right);

        public static explicit operator string(PatchBaselineApprovedPatchesComplianceLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PatchBaselineApprovedPatchesComplianceLevel other && Equals(other);
        public bool Equals(PatchBaselineApprovedPatchesComplianceLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Defines the operating system the patch baseline applies to. The Default value is WINDOWS.
    /// </summary>
    [EnumType]
    public readonly struct PatchBaselineOperatingSystem : IEquatable<PatchBaselineOperatingSystem>
    {
        private readonly string _value;

        private PatchBaselineOperatingSystem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PatchBaselineOperatingSystem Windows { get; } = new PatchBaselineOperatingSystem("WINDOWS");
        public static PatchBaselineOperatingSystem AmazonLinux { get; } = new PatchBaselineOperatingSystem("AMAZON_LINUX");
        public static PatchBaselineOperatingSystem AmazonLinux2 { get; } = new PatchBaselineOperatingSystem("AMAZON_LINUX_2");
        public static PatchBaselineOperatingSystem AmazonLinux2022 { get; } = new PatchBaselineOperatingSystem("AMAZON_LINUX_2022");
        public static PatchBaselineOperatingSystem AmazonLinux2023 { get; } = new PatchBaselineOperatingSystem("AMAZON_LINUX_2023");
        public static PatchBaselineOperatingSystem Ubuntu { get; } = new PatchBaselineOperatingSystem("UBUNTU");
        public static PatchBaselineOperatingSystem RedhatEnterpriseLinux { get; } = new PatchBaselineOperatingSystem("REDHAT_ENTERPRISE_LINUX");
        public static PatchBaselineOperatingSystem Suse { get; } = new PatchBaselineOperatingSystem("SUSE");
        public static PatchBaselineOperatingSystem Centos { get; } = new PatchBaselineOperatingSystem("CENTOS");
        public static PatchBaselineOperatingSystem OracleLinux { get; } = new PatchBaselineOperatingSystem("ORACLE_LINUX");
        public static PatchBaselineOperatingSystem Debian { get; } = new PatchBaselineOperatingSystem("DEBIAN");
        public static PatchBaselineOperatingSystem Macos { get; } = new PatchBaselineOperatingSystem("MACOS");
        public static PatchBaselineOperatingSystem Raspbian { get; } = new PatchBaselineOperatingSystem("RASPBIAN");
        public static PatchBaselineOperatingSystem RockyLinux { get; } = new PatchBaselineOperatingSystem("ROCKY_LINUX");
        public static PatchBaselineOperatingSystem AlmaLinux { get; } = new PatchBaselineOperatingSystem("ALMA_LINUX");

        public static bool operator ==(PatchBaselineOperatingSystem left, PatchBaselineOperatingSystem right) => left.Equals(right);
        public static bool operator !=(PatchBaselineOperatingSystem left, PatchBaselineOperatingSystem right) => !left.Equals(right);

        public static explicit operator string(PatchBaselineOperatingSystem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PatchBaselineOperatingSystem other && Equals(other);
        public bool Equals(PatchBaselineOperatingSystem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The key for the filter.
    /// 
    /// For information about valid keys, see [PatchFilter](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchFilter.html) in the *AWS Systems Manager API Reference* .
    /// </summary>
    [EnumType]
    public readonly struct PatchBaselinePatchFilterKey : IEquatable<PatchBaselinePatchFilterKey>
    {
        private readonly string _value;

        private PatchBaselinePatchFilterKey(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PatchBaselinePatchFilterKey AdvisoryId { get; } = new PatchBaselinePatchFilterKey("ADVISORY_ID");
        public static PatchBaselinePatchFilterKey Arch { get; } = new PatchBaselinePatchFilterKey("ARCH");
        public static PatchBaselinePatchFilterKey BugzillaId { get; } = new PatchBaselinePatchFilterKey("BUGZILLA_ID");
        public static PatchBaselinePatchFilterKey Classification { get; } = new PatchBaselinePatchFilterKey("CLASSIFICATION");
        public static PatchBaselinePatchFilterKey CveId { get; } = new PatchBaselinePatchFilterKey("CVE_ID");
        public static PatchBaselinePatchFilterKey Epoch { get; } = new PatchBaselinePatchFilterKey("EPOCH");
        public static PatchBaselinePatchFilterKey MsrcSeverity { get; } = new PatchBaselinePatchFilterKey("MSRC_SEVERITY");
        public static PatchBaselinePatchFilterKey Name { get; } = new PatchBaselinePatchFilterKey("NAME");
        public static PatchBaselinePatchFilterKey PatchId { get; } = new PatchBaselinePatchFilterKey("PATCH_ID");
        public static PatchBaselinePatchFilterKey PatchSet { get; } = new PatchBaselinePatchFilterKey("PATCH_SET");
        public static PatchBaselinePatchFilterKey Priority { get; } = new PatchBaselinePatchFilterKey("PRIORITY");
        public static PatchBaselinePatchFilterKey Product { get; } = new PatchBaselinePatchFilterKey("PRODUCT");
        public static PatchBaselinePatchFilterKey ProductFamily { get; } = new PatchBaselinePatchFilterKey("PRODUCT_FAMILY");
        public static PatchBaselinePatchFilterKey Release { get; } = new PatchBaselinePatchFilterKey("RELEASE");
        public static PatchBaselinePatchFilterKey Repository { get; } = new PatchBaselinePatchFilterKey("REPOSITORY");
        public static PatchBaselinePatchFilterKey Section { get; } = new PatchBaselinePatchFilterKey("SECTION");
        public static PatchBaselinePatchFilterKey Security { get; } = new PatchBaselinePatchFilterKey("SECURITY");
        public static PatchBaselinePatchFilterKey Severity { get; } = new PatchBaselinePatchFilterKey("SEVERITY");
        public static PatchBaselinePatchFilterKey Version { get; } = new PatchBaselinePatchFilterKey("VERSION");

        public static bool operator ==(PatchBaselinePatchFilterKey left, PatchBaselinePatchFilterKey right) => left.Equals(right);
        public static bool operator !=(PatchBaselinePatchFilterKey left, PatchBaselinePatchFilterKey right) => !left.Equals(right);

        public static explicit operator string(PatchBaselinePatchFilterKey value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PatchBaselinePatchFilterKey other && Equals(other);
        public bool Equals(PatchBaselinePatchFilterKey other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The action for Patch Manager to take on patches included in the RejectedPackages list.
    /// </summary>
    [EnumType]
    public readonly struct PatchBaselineRejectedPatchesAction : IEquatable<PatchBaselineRejectedPatchesAction>
    {
        private readonly string _value;

        private PatchBaselineRejectedPatchesAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PatchBaselineRejectedPatchesAction AllowAsDependency { get; } = new PatchBaselineRejectedPatchesAction("ALLOW_AS_DEPENDENCY");
        public static PatchBaselineRejectedPatchesAction Block { get; } = new PatchBaselineRejectedPatchesAction("BLOCK");

        public static bool operator ==(PatchBaselineRejectedPatchesAction left, PatchBaselineRejectedPatchesAction right) => left.Equals(right);
        public static bool operator !=(PatchBaselineRejectedPatchesAction left, PatchBaselineRejectedPatchesAction right) => !left.Equals(right);

        public static explicit operator string(PatchBaselineRejectedPatchesAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PatchBaselineRejectedPatchesAction other && Equals(other);
        public bool Equals(PatchBaselineRejectedPatchesAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A compliance severity level for all approved patches in a patch baseline. Valid compliance severity levels include the following: `UNSPECIFIED` , `CRITICAL` , `HIGH` , `MEDIUM` , `LOW` , and `INFORMATIONAL` .
    /// </summary>
    [EnumType]
    public readonly struct PatchBaselineRuleComplianceLevel : IEquatable<PatchBaselineRuleComplianceLevel>
    {
        private readonly string _value;

        private PatchBaselineRuleComplianceLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PatchBaselineRuleComplianceLevel Critical { get; } = new PatchBaselineRuleComplianceLevel("CRITICAL");
        public static PatchBaselineRuleComplianceLevel High { get; } = new PatchBaselineRuleComplianceLevel("HIGH");
        public static PatchBaselineRuleComplianceLevel Informational { get; } = new PatchBaselineRuleComplianceLevel("INFORMATIONAL");
        public static PatchBaselineRuleComplianceLevel Low { get; } = new PatchBaselineRuleComplianceLevel("LOW");
        public static PatchBaselineRuleComplianceLevel Medium { get; } = new PatchBaselineRuleComplianceLevel("MEDIUM");
        public static PatchBaselineRuleComplianceLevel Unspecified { get; } = new PatchBaselineRuleComplianceLevel("UNSPECIFIED");

        public static bool operator ==(PatchBaselineRuleComplianceLevel left, PatchBaselineRuleComplianceLevel right) => left.Equals(right);
        public static bool operator !=(PatchBaselineRuleComplianceLevel left, PatchBaselineRuleComplianceLevel right) => !left.Equals(right);

        public static explicit operator string(PatchBaselineRuleComplianceLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PatchBaselineRuleComplianceLevel other && Equals(other);
        public bool Equals(PatchBaselineRuleComplianceLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
