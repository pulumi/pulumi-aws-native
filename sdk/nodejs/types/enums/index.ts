// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// Export sub-modules:
import * as amplify from "./amplify";
import * as amplifyuibuilder from "./amplifyuibuilder";
import * as apigateway from "./apigateway";
import * as apigatewayv2 from "./apigatewayv2";
import * as appconfig from "./appconfig";
import * as appflow from "./appflow";
import * as applicationinsights from "./applicationinsights";
import * as applicationsignals from "./applicationsignals";
import * as apprunner from "./apprunner";
import * as appsync from "./appsync";
import * as arczonalshift from "./arczonalshift";
import * as athena from "./athena";
import * as auditmanager from "./auditmanager";
import * as autoscaling from "./autoscaling";
import * as b2bi from "./b2bi";
import * as backup from "./backup";
import * as batch from "./batch";
import * as bedrock from "./bedrock";
import * as billing from "./billing";
import * as budgets from "./budgets";
import * as cassandra from "./cassandra";
import * as ce from "./ce";
import * as chatbot from "./chatbot";
import * as cleanrooms from "./cleanrooms";
import * as cleanroomsml from "./cleanroomsml";
import * as cloudformation from "./cloudformation";
import * as cloudfront from "./cloudfront";
import * as cloudtrail from "./cloudtrail";
import * as codeartifact from "./codeartifact";
import * as codebuild from "./codebuild";
import * as codeguruprofiler from "./codeguruprofiler";
import * as codegurureviewer from "./codegurureviewer";
import * as codepipeline from "./codepipeline";
import * as codestarconnections from "./codestarconnections";
import * as codestarnotifications from "./codestarnotifications";
import * as cognito from "./cognito";
import * as comprehend from "./comprehend";
import * as connect from "./connect";
import * as connectcampaignsv2 from "./connectcampaignsv2";
import * as controltower from "./controltower";
import * as customerprofiles from "./customerprofiles";
import * as databrew from "./databrew";
import * as datasync from "./datasync";
import * as datazone from "./datazone";
import * as deadline from "./deadline";
import * as devicefarm from "./devicefarm";
import * as devopsguru from "./devopsguru";
import * as dms from "./dms";
import * as dynamodb from "./dynamodb";
import * as ec2 from "./ec2";
import * as ecr from "./ecr";
import * as ecs from "./ecs";
import * as efs from "./efs";
import * as eks from "./eks";
import * as elasticache from "./elasticache";
import * as emr from "./emr";
import * as emrserverless from "./emrserverless";
import * as entityresolution from "./entityresolution";
import * as events from "./events";
import * as evidently from "./evidently";
import * as evs from "./evs";
import * as finspace from "./finspace";
import * as fis from "./fis";
import * as fms from "./fms";
import * as forecast from "./forecast";
import * as frauddetector from "./frauddetector";
import * as fsx from "./fsx";
import * as gamelift from "./gamelift";
import * as globalaccelerator from "./globalaccelerator";
import * as glue from "./glue";
import * as grafana from "./grafana";
import * as greengrassv2 from "./greengrassv2";
import * as groundstation from "./groundstation";
import * as guardduty from "./guardduty";
import * as healthimaging from "./healthimaging";
import * as healthlake from "./healthlake";
import * as iam from "./iam";
import * as imagebuilder from "./imagebuilder";
import * as inspectorv2 from "./inspectorv2";
import * as internetmonitor from "./internetmonitor";
import * as iot from "./iot";
import * as iotanalytics from "./iotanalytics";
import * as iotevents from "./iotevents";
import * as iotsitewise from "./iotsitewise";
import * as iottwinmaker from "./iottwinmaker";
import * as iotwireless from "./iotwireless";
import * as ivs from "./ivs";
import * as ivschat from "./ivschat";
import * as kafkaconnect from "./kafkaconnect";
import * as kendra from "./kendra";
import * as kinesis from "./kinesis";
import * as kinesisanalyticsv2 from "./kinesisanalyticsv2";
import * as kinesisfirehose from "./kinesisfirehose";
import * as kinesisvideo from "./kinesisvideo";
import * as kms from "./kms";
import * as lakeformation from "./lakeformation";
import * as lambda from "./lambda";
import * as launchwizard from "./launchwizard";
import * as lex from "./lex";
import * as lightsail from "./lightsail";
import * as location from "./location";
import * as logs from "./logs";
import * as lookoutmetrics from "./lookoutmetrics";
import * as m2 from "./m2";
import * as macie from "./macie";
import * as mediaconnect from "./mediaconnect";
import * as medialive from "./medialive";
import * as mediapackage from "./mediapackage";
import * as mediapackagev2 from "./mediapackagev2";
import * as mediatailor from "./mediatailor";
import * as memorydb from "./memorydb";
import * as msk from "./msk";
import * as mwaa from "./mwaa";
import * as networkfirewall from "./networkfirewall";
import * as oam from "./oam";
import * as odb from "./odb";
import * as omics from "./omics";
import * as opensearchserverless from "./opensearchserverless";
import * as opensearchservice from "./opensearchservice";
import * as organizations from "./organizations";
import * as osis from "./osis";
import * as panorama from "./panorama";
import * as paymentcryptography from "./paymentcryptography";
import * as pcaconnectorad from "./pcaconnectorad";
import * as pcaconnectorscep from "./pcaconnectorscep";
import * as pcs from "./pcs";
import * as personalize from "./personalize";
import * as pinpoint from "./pinpoint";
import * as pipes from "./pipes";
import * as proton from "./proton";
import * as qbusiness from "./qbusiness";
import * as quicksight from "./quicksight";
import * as rbin from "./rbin";
import * as rds from "./rds";
import * as redshift from "./redshift";
import * as redshiftserverless from "./redshiftserverless";
import * as refactorspaces from "./refactorspaces";
import * as resiliencehub from "./resiliencehub";
import * as resourceexplorer2 from "./resourceexplorer2";
import * as resourcegroups from "./resourcegroups";
import * as robomaker from "./robomaker";
import * as rolesanywhere from "./rolesanywhere";
import * as route53 from "./route53";
import * as route53recoverycontrol from "./route53recoverycontrol";
import * as route53resolver from "./route53resolver";
import * as rum from "./rum";
import * as s3 from "./s3";
import * as s3express from "./s3express";
import * as s3outposts from "./s3outposts";
import * as s3tables from "./s3tables";
import * as sagemaker from "./sagemaker";
import * as scheduler from "./scheduler";
import * as securityhub from "./securityhub";
import * as securitylake from "./securitylake";
import * as servicecatalog from "./servicecatalog";
import * as servicecatalogappregistry from "./servicecatalogappregistry";
import * as ses from "./ses";
import * as shield from "./shield";
import * as signer from "./signer";
import * as sns from "./sns";
import * as ssm from "./ssm";
import * as ssmcontacts from "./ssmcontacts";
import * as ssmincidents from "./ssmincidents";
import * as ssmquicksetup from "./ssmquicksetup";
import * as sso from "./sso";
import * as stepfunctions from "./stepfunctions";
import * as supportapp from "./supportapp";
import * as synthetics from "./synthetics";
import * as systemsmanagersap from "./systemsmanagersap";
import * as timestream from "./timestream";
import * as transfer from "./transfer";
import * as verifiedpermissions from "./verifiedpermissions";
import * as vpclattice from "./vpclattice";
import * as wafv2 from "./wafv2";
import * as wisdom from "./wisdom";
import * as workspaces from "./workspaces";
import * as workspacesinstances from "./workspacesinstances";
import * as workspacesthinclient from "./workspacesthinclient";
import * as workspacesweb from "./workspacesweb";

export {
    amplify,
    amplifyuibuilder,
    apigateway,
    apigatewayv2,
    appconfig,
    appflow,
    applicationinsights,
    applicationsignals,
    apprunner,
    appsync,
    arczonalshift,
    athena,
    auditmanager,
    autoscaling,
    b2bi,
    backup,
    batch,
    bedrock,
    billing,
    budgets,
    cassandra,
    ce,
    chatbot,
    cleanrooms,
    cleanroomsml,
    cloudformation,
    cloudfront,
    cloudtrail,
    codeartifact,
    codebuild,
    codeguruprofiler,
    codegurureviewer,
    codepipeline,
    codestarconnections,
    codestarnotifications,
    cognito,
    comprehend,
    connect,
    connectcampaignsv2,
    controltower,
    customerprofiles,
    databrew,
    datasync,
    datazone,
    deadline,
    devicefarm,
    devopsguru,
    dms,
    dynamodb,
    ec2,
    ecr,
    ecs,
    efs,
    eks,
    elasticache,
    emr,
    emrserverless,
    entityresolution,
    events,
    evidently,
    evs,
    finspace,
    fis,
    fms,
    forecast,
    frauddetector,
    fsx,
    gamelift,
    globalaccelerator,
    glue,
    grafana,
    greengrassv2,
    groundstation,
    guardduty,
    healthimaging,
    healthlake,
    iam,
    imagebuilder,
    inspectorv2,
    internetmonitor,
    iot,
    iotanalytics,
    iotevents,
    iotsitewise,
    iottwinmaker,
    iotwireless,
    ivs,
    ivschat,
    kafkaconnect,
    kendra,
    kinesis,
    kinesisanalyticsv2,
    kinesisfirehose,
    kinesisvideo,
    kms,
    lakeformation,
    lambda,
    launchwizard,
    lex,
    lightsail,
    location,
    logs,
    lookoutmetrics,
    m2,
    macie,
    mediaconnect,
    medialive,
    mediapackage,
    mediapackagev2,
    mediatailor,
    memorydb,
    msk,
    mwaa,
    networkfirewall,
    oam,
    odb,
    omics,
    opensearchserverless,
    opensearchservice,
    organizations,
    osis,
    panorama,
    paymentcryptography,
    pcaconnectorad,
    pcaconnectorscep,
    pcs,
    personalize,
    pinpoint,
    pipes,
    proton,
    qbusiness,
    quicksight,
    rbin,
    rds,
    redshift,
    redshiftserverless,
    refactorspaces,
    resiliencehub,
    resourceexplorer2,
    resourcegroups,
    robomaker,
    rolesanywhere,
    route53,
    route53recoverycontrol,
    route53resolver,
    rum,
    s3,
    s3express,
    s3outposts,
    s3tables,
    sagemaker,
    scheduler,
    securityhub,
    securitylake,
    servicecatalog,
    servicecatalogappregistry,
    ses,
    shield,
    signer,
    sns,
    ssm,
    ssmcontacts,
    ssmincidents,
    ssmquicksetup,
    sso,
    stepfunctions,
    supportapp,
    synthetics,
    systemsmanagersap,
    timestream,
    transfer,
    verifiedpermissions,
    vpclattice,
    wafv2,
    wisdom,
    workspaces,
    workspacesinstances,
    workspacesthinclient,
    workspacesweb,
};

export const Region = {
    /**
     * Africa (Cape Town)
     */
    AfSouth1: "af-south-1",
    /**
     * Asia Pacific (Hong Kong)
     */
    ApEast1: "ap-east-1",
    /**
     * Asia Pacific (Taipei)
     */
    ApEast2: "ap-east-2",
    /**
     * Asia Pacific (Tokyo)
     */
    ApNortheast1: "ap-northeast-1",
    /**
     * Asia Pacific (Seoul)
     */
    ApNortheast2: "ap-northeast-2",
    /**
     * Asia Pacific (Osaka)
     */
    ApNortheast3: "ap-northeast-3",
    /**
     * Asia Pacific (Mumbai)
     */
    ApSouth1: "ap-south-1",
    /**
     * Asia Pacific (Hyderabad)
     */
    ApSouth2: "ap-south-2",
    /**
     * Asia Pacific (Singapore)
     */
    ApSoutheast1: "ap-southeast-1",
    /**
     * Asia Pacific (Sydney)
     */
    ApSoutheast2: "ap-southeast-2",
    /**
     * Asia Pacific (Jakarta)
     */
    ApSoutheast3: "ap-southeast-3",
    /**
     * Asia Pacific (Melbourne)
     */
    ApSoutheast4: "ap-southeast-4",
    /**
     * Asia Pacific (Malaysia)
     */
    ApSoutheast5: "ap-southeast-5",
    /**
     * Asia Pacific (Thailand)
     */
    ApSoutheast7: "ap-southeast-7",
    /**
     * Canada (Central)
     */
    CaCentral1: "ca-central-1",
    /**
     * Canada West (Calgary)
     */
    CaWest1: "ca-west-1",
    /**
     * China (Beijing)
     */
    CnNorth1: "cn-north-1",
    /**
     * China (Ningxia)
     */
    CnNorthwest1: "cn-northwest-1",
    /**
     * Europe (Frankfurt)
     */
    EuCentral1: "eu-central-1",
    /**
     * Europe (Zurich)
     */
    EuCentral2: "eu-central-2",
    /**
     * EU ISOE West
     */
    EuIsoeWest1: "eu-isoe-west-1",
    /**
     * Europe (Stockholm)
     */
    EuNorth1: "eu-north-1",
    /**
     * Europe (Milan)
     */
    EuSouth1: "eu-south-1",
    /**
     * Europe (Spain)
     */
    EuSouth2: "eu-south-2",
    /**
     * Europe (Ireland)
     */
    EuWest1: "eu-west-1",
    /**
     * Europe (London)
     */
    EuWest2: "eu-west-2",
    /**
     * Europe (Paris)
     */
    EuWest3: "eu-west-3",
    /**
     * EU (Germany)
     */
    EuscDeEast1: "eusc-de-east-1",
    /**
     * Israel (Tel Aviv)
     */
    IlCentral1: "il-central-1",
    /**
     * Middle East (UAE)
     */
    MeCentral1: "me-central-1",
    /**
     * Middle East (Bahrain)
     */
    MeSouth1: "me-south-1",
    /**
     * Mexico (Central)
     */
    MxCentral1: "mx-central-1",
    /**
     * South America (Sao Paulo)
     */
    SaEast1: "sa-east-1",
    /**
     * US East (N. Virginia)
     */
    UsEast1: "us-east-1",
    /**
     * US East (Ohio)
     */
    UsEast2: "us-east-2",
    /**
     * AWS GovCloud (US-East)
     */
    UsGovEast1: "us-gov-east-1",
    /**
     * AWS GovCloud (US-West)
     */
    UsGovWest1: "us-gov-west-1",
    /**
     * US ISO East
     */
    UsIsoEast1: "us-iso-east-1",
    /**
     * US ISO WEST
     */
    UsIsoWest1: "us-iso-west-1",
    /**
     * US ISOB East (Ohio)
     */
    UsIsobEast1: "us-isob-east-1",
    /**
     * US ISOF EAST
     */
    UsIsofEast1: "us-isof-east-1",
    /**
     * US ISOF SOUTH
     */
    UsIsofSouth1: "us-isof-south-1",
    /**
     * US West (N. California)
     */
    UsWest1: "us-west-1",
    /**
     * US West (Oregon)
     */
    UsWest2: "us-west-2",
} as const;

/**
 * A Region represents any valid Amazon region that may be targeted with deployments.
 */
export type Region = (typeof Region)[keyof typeof Region];
