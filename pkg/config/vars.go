package config

import (
	"github.com/urfave/cli"
)

type (
	configVar struct {
		Type    configVarType
		Default interface{}
		CLIFlag cli.Flag
	}

	configVarType string
)

// Will be populated in init() below
var CLIFlags []cli.Flag

var (
	stringType configVarType = "string"
	intType    configVarType = "int"
	boolType   configVarType = "bool"
)

var configVars = map[string]configVar{
	"genindex": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "gen-index",
			Usage:  "generate index.yaml, print to stdout and exit",
			EnvVar: "GEN_INDEX",
		},
	},
	"debug": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "debug",
			Usage:  "show debug messages",
			EnvVar: "DEBUG",
		},
	},
	"logjson": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "log-json",
			Usage:  "output structured logs as json",
			EnvVar: "LOG_JSON",
		},
	},
	"disablemetrics": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "disable-metrics",
			Usage:  "disable Prometheus metrics",
			EnvVar: "DISABLE_METRICS",
		},
	},
	"disableapi": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "disable-api",
			Usage:  "disable all routes prefixed with /api",
			EnvVar: "DISABLE_API",
		},
	},
	"allowoverwrite": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "allow-overwrite",
			Usage:  "allow chart versions to be re-uploaded",
			EnvVar: "ALLOW_OVERWRITE",
		},
	},
	"port": {
		Type:    intType,
		Default: 8080,
		CLIFlag: cli.IntFlag{
			Name:   "port",
			Usage:  "port to listen on",
			EnvVar: "PORT",
		},
	},
	"charturl": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "chart-url",
			Usage:  "absolute url for .tgzs in index.yaml",
			EnvVar: "CHART_URL",
		},
	},
	"basicauth.user": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "basic-auth-user",
			Usage:  "username for basic http authentication",
			EnvVar: "BASIC_AUTH_USER",
		},
	},
	"basicauth.pass": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "basic-auth-pass",
			Usage:  "password for basic http authentication",
			EnvVar: "BASIC_AUTH_PASS",
		},
	},
	"authanonymousget": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "auth-anonymous-get",
			Usage:  "allow anonymous GET operations when auth is used",
			EnvVar: "AUTH_ANONYMOUS_GET",
		},
	},
	"tls.cert": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "tls-cert",
			Usage:  "path to tls certificate chain file",
			EnvVar: "TLS_CERT",
		},
	},
	"tls.key": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "tls-key",
			Usage:  "path to tls key file",
			EnvVar: "TLS_KEY",
		},
	},
	"storage.backend": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage",
			Usage:  "storage backend, can be one of: local, amazon, google",
			EnvVar: "STORAGE",
		},
	},
	"storage.local.rootdir": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-local-rootdir",
			Usage:  "directory to store charts for local storage backend",
			EnvVar: "STORAGE_LOCAL_ROOTDIR",
		},
	},
	"storage.amazon.bucket": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-amazon-bucket",
			Usage:  "s3 bucket to store charts for amazon storage backend",
			EnvVar: "STORAGE_AMAZON_BUCKET",
		},
	},
	"storage.amazon.prefix": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-amazon-prefix",
			Usage:  "prefix to store charts for --storage-amazon-bucket",
			EnvVar: "STORAGE_AMAZON_PREFIX",
		},
	},
	"storage.amazon.region": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-amazon-region",
			Usage:  "region of --storage-amazon-bucket",
			EnvVar: "STORAGE_AMAZON_REGION",
		},
	},
	"storage.amazon.endpoint": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-amazon-endpoint",
			Usage:  "alternative s3 endpoint",
			EnvVar: "STORAGE_AMAZON_ENDPOINT",
		},
	},
	"storage.amazon.sse": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-amazon-sse",
			Usage:  "server side encryption algorithm",
			EnvVar: "STORAGE_AMAZON_SSE",
		},
	},
	"storage.google.bucket": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-google-bucket",
			Usage:  "gcs bucket to store charts for google storage backend",
			EnvVar: "STORAGE_GOOGLE_BUCKET",
		},
	},
	"storage.google.prefix": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-google-prefix",
			Usage:  "prefix to store charts for --storage-google-bucket",
			EnvVar: "STORAGE_GOOGLE_PREFIX",
		},
	},
	"storage.microsoft.container": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-microsoft-container",
			Usage:  "container to store charts for microsoft storage backend",
			EnvVar: "STORAGE_MICROSOFT_CONTAINER",
		},
	},
	"storage.microsoft.prefix": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-microsoft-prefix",
			Usage:  "prefix to store charts for --storage-microsoft-prefix",
			EnvVar: "STORAGE_MICROSOFT_PREFIX",
		},
	},
	"storage.alibaba.bucket": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-alibaba-bucket",
			Usage:  "OSS bucket to store charts for Alibaba Cloud storage backend",
			EnvVar: "STORAGE_ALIBABA_BUCKET",
		},
	},
	"storage.alibaba.prefix": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-alibaba-prefix",
			Usage:  "prefix to store charts for --storage-alibaba-cloud-bucket",
			EnvVar: "STORAGE_ALIBABA_PREFIX",
		},
	},
	"storage.alibaba.endpoint": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-alibaba-endpoint",
			Usage:  "OSS endpoint",
			EnvVar: "STORAGE_ALIBABA_ENDPOINT",
		},
	},
	"storage.alibaba.sse": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-alibaba-sse",
			Usage:  "server side encryption algorithm for Alibaba Cloud storage backend, AES256 or KMS",
			EnvVar: "STORAGE_ALIBABA_SSE",
		},
	},
	"chartpostformfieldname": {
		Type:    stringType,
		Default: "chart",
		CLIFlag: cli.StringFlag{
			Name:   "chart-post-form-field-name",
			Usage:  "form field which will be queried for the chart file content",
			EnvVar: "CHART_POST_FORM_FIELD_NAME",
		},
	},
	"provpostformfieldname": {
		Type:    stringType,
		Default: "prov",
		CLIFlag: cli.StringFlag{
			Name:   "prov-post-form-field-name",
			Usage:  "form field which will be queried for the provenance file content",
			EnvVar: "PROV_POST_FORM_FIELD_NAME",
		},
	},
	"indexlimit": {
		Type:    intType,
		Default: 0,
		CLIFlag: cli.IntFlag{
			Name:   "index-limit",
			Usage:  "parallel scan limit for the repo indexer",
			EnvVar: "INDEX_LIMIT",
		},
	},
	"contextpath": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "context-path",
			Usage:  "base context path",
			EnvVar: "CONTEXT_PATH",
		},
	},
	"depth": {
		Type:    intType,
		Default: 0,
		CLIFlag: cli.IntFlag{
			Name:   "depth",
			Usage:  "levels of nested repos for multitenancy",
			EnvVar: "DEPTH",
		},
	},
}

func populateCLIFlags() {
	CLIFlags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Usage:  "chartmuseum configuration file",
			EnvVar: "CONFIG",
		},
	}
	for _, configVar := range configVars {
		if flag := configVar.CLIFlag; flag != nil {
			CLIFlags = append(CLIFlags, flag)
		}
	}
}

func init() {
	populateCLIFlags()
}
