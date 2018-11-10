package config

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	gcfg "gopkg.in/gcfg.v1"

	kexec "k8s.io/utils/exec"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/cert"
)

// The following are global config parameters that other modules may access directly
var (
	// ovn-kubernetes version, to be changed with every release
	Version = "0.3.0"

	// Default holds parsed config file parameters and command-line overrides
	Default = DefaultConfig{
		MTU: 1400,
	}

	// Logging holds logging-related parsed config file parameters and command-line overrides
	Logging = LoggingConfig{
		File:  "", // do not log to a file by default
		Level: 4,
	}

	// CNI holds CNI-related parsed config file parameters and command-line overrides
	CNI = CNIConfig{
		ConfDir: "/etc/cni/net.d",
		Plugin:  "ovn4nfvk8s-cni",
	}

	// Kubernetes holds Kubernetes-related parsed config file parameters and command-line overrides
	Kubernetes = KubernetesConfig{
		APIServer: "http://localhost:8080",
	}
)

// DefaultConfig holds parsed config file parameters and command-line overrides
type DefaultConfig struct {
	// MTU value used for the overlay networks.
	MTU int `gcfg:"mtu"`
}

// LoggingConfig holds logging-related parsed config file parameters and command-line overrides
type LoggingConfig struct {
	// File is the path of the file to log to
	File string `gcfg:"logfile"`
	// Level is the logging verbosity level
	Level int `gcfg:"loglevel"`
}

// CNIConfig holds CNI-related parsed config file parameters and command-line overrides
type CNIConfig struct {
	// ConfDir specifies the CNI config directory in which to write the overlay CNI config file
	ConfDir string `gcfg:"conf-dir"`
	// Plugin specifies the name of the CNI plugin
	Plugin string `gcfg:"plugin"`
}

// KubernetesConfig holds Kubernetes-related parsed config file parameters and command-line overrides
type KubernetesConfig struct {
	Kubeconfig string `gcfg:"kubeconfig"`
	CACert     string `gcfg:"cacert"`
	APIServer  string `gcfg:"apiserver"`
	Token      string `gcfg:"token"`
}

// Config is used to read the structured config file and to cache config in testcases
type config struct {
	Default    DefaultConfig
	Logging    LoggingConfig
	CNI        CNIConfig
	Kubernetes KubernetesConfig
}

// copy members of struct 'src' into the corresponding field in struct 'dst'
// if the field in 'src' is a non-zero int or a non-zero-length string. This
// function should be called with pointers to structs.
func overrideFields(dst, src interface{}) {
	dstStruct := reflect.ValueOf(dst).Elem()
	srcStruct := reflect.ValueOf(src).Elem()
	if dstStruct.Kind() != srcStruct.Kind() || dstStruct.Kind() != reflect.Struct {
		panic("mismatched value types")
	}
	if dstStruct.NumField() != srcStruct.NumField() {
		panic("mismatched struct types")
	}

	for i := 0; i < dstStruct.NumField(); i++ {
		dstField := dstStruct.Field(i)
		srcField := srcStruct.Field(i)
		if dstField.Kind() != srcField.Kind() {
			panic("mismatched struct fields")
		}
		switch srcField.Kind() {
		case reflect.String:
			if srcField.String() != "" {
				dstField.Set(srcField)
			}
		case reflect.Int:
			if srcField.Int() != 0 {
				dstField.Set(srcField)
			}
		default:
			panic(fmt.Sprintf("unhandled struct field type: %v", srcField.Kind()))
		}
	}
}

var cliConfig config

// Flags are general command-line flags. Apps should add these flags to their
// own urfave/cli flags and call InitConfig() early in the application.
var Flags = []cli.Flag{
	cli.StringFlag{
		Name:  "config-file",
		Usage: "configuration file path (default: /etc/openvswitch/ovn4nfv_k8s.conf)",
	},

	// Generic options
	cli.IntFlag{
		Name:        "mtu",
		Usage:       "MTU value used for the overlay networks (default: 1400)",
		Destination: &cliConfig.Default.MTU,
	},

	// Logging options
	cli.IntFlag{
		Name:        "loglevel",
		Usage:       "log verbosity and level: 5=debug, 4=info, 3=warn, 2=error, 1=fatal (default: 4)",
		Destination: &cliConfig.Logging.Level,
	},
	cli.StringFlag{
		Name:        "logfile",
		Usage:       "path of a file to direct log output to",
		Destination: &cliConfig.Logging.File,
	},

	// CNI options
	cli.StringFlag{
		Name:        "cni-conf-dir",
		Usage:       "the CNI config directory in which to write the overlay CNI config file (default: /etc/cni/net.d)",
		Destination: &cliConfig.CNI.ConfDir,
	},
	cli.StringFlag{
		Name:        "cni-plugin",
		Usage:       "the name of the CNI plugin (default: ovn4nfvk8s-cni)",
		Destination: &cliConfig.CNI.Plugin,
	},

	// Kubernetes-related options
	cli.StringFlag{
		Name:        "k8s-kubeconfig",
		Usage:       "absolute path to the Kubernetes kubeconfig file (not required if the --k8s-apiserver, --k8s-ca-cert, and --k8s-token are given)",
		Destination: &cliConfig.Kubernetes.Kubeconfig,
	},
	cli.StringFlag{
		Name:        "k8s-apiserver",
		Usage:       "URL of the Kubernetes API server (not required if --k8s-kubeconfig is given) (default: http://localhost:8443)",
		Destination: &cliConfig.Kubernetes.APIServer,
	},
	cli.StringFlag{
		Name:        "k8s-cacert",
		Usage:       "the absolute path to the Kubernetes API CA certificate (not required if --k8s-kubeconfig is given)",
		Destination: &cliConfig.Kubernetes.CACert,
	},
	cli.StringFlag{
		Name:        "k8s-token",
		Usage:       "the Kubernetes API authentication token (not required if --k8s-kubeconfig is given)",
		Destination: &cliConfig.Kubernetes.Token,
	},
}

type Defaults struct {
	K8sAPIServer bool
	K8sToken     bool
	K8sCert      bool
}

const (
	ovsVsctlCommand = "ovs-vsctl"
)

func buildKubernetesConfig(exec kexec.Interface, cli, file *config, defaults *Defaults) error {

	// Copy config file values over default values
	overrideFields(&Kubernetes, &file.Kubernetes)
	// And CLI overrides over config file and default values
	overrideFields(&Kubernetes, &cli.Kubernetes)

	if Kubernetes.Kubeconfig != "" && !pathExists(Kubernetes.Kubeconfig) {
		return fmt.Errorf("kubernetes kubeconfig file %q not found", Kubernetes.Kubeconfig)
	}
	if Kubernetes.CACert != "" && !pathExists(Kubernetes.CACert) {
		return fmt.Errorf("kubernetes CA certificate file %q not found", Kubernetes.CACert)
	}

	url, err := url.Parse(Kubernetes.APIServer)
	if err != nil {
		return fmt.Errorf("kubernetes API server address %q invalid: %v", Kubernetes.APIServer, err)
	} else if url.Scheme != "https" && url.Scheme != "http" {
		return fmt.Errorf("kubernetes API server URL scheme %q invalid", url.Scheme)
	}

	return nil
}

// getConfigFilePath returns config file path and 'true' if the config file is
// the fallback path (eg not given by the user), 'false' if given explicitly
// by the user
func getConfigFilePath(ctx *cli.Context) (string, bool) {
	configFile := ctx.String("config-file")
	if configFile != "" {
		return configFile, false
	}

	// default
	return filepath.Join("/etc", "openvswitch", "ovn4nfv_k8s.conf"), true

}

// InitConfig reads the config file and common command-line options and
// constructs the global config object from them. It returns the config file
// path (if explicitly specified) or an error
func InitConfig(ctx *cli.Context, exec kexec.Interface, defaults *Defaults) (string, error) {
	return InitConfigWithPath(ctx, exec, "", defaults)
}

// InitConfigWithPath reads the given config file (or if empty, reads the config file
// specified by command-line arguments, or empty, the default config file) and
// common command-line options and constructs the global config object from
// them. It returns the config file path (if explicitly specified) or an error
func InitConfigWithPath(ctx *cli.Context, exec kexec.Interface, configFile string, defaults *Defaults) (string, error) {
	var cfg config
	var retConfigFile string
	var configFileIsDefault bool

	// If no specific config file was given, try to find one from command-line
	// arguments, or the platform-specific default config file path
	if configFile == "" {
		configFile, configFileIsDefault = getConfigFilePath(ctx)
	}

	logrus.SetOutput(os.Stderr)

	if !configFileIsDefault {
		// Only return explicitly specified config file
		retConfigFile = configFile
	}

	f, err := os.Open(configFile)
	// Failure to find a default config file is not a hard error
	if err != nil && !configFileIsDefault {
		return "", fmt.Errorf("failed to open config file %s: %v", configFile, err)
	}
	if f != nil {
		defer f.Close()

		// Parse ovn4nfvk8s config file.
		if err = gcfg.ReadInto(&cfg, f); err != nil {
			return "", fmt.Errorf("failed to parse config file %s: %v", f.Name(), err)
		}
		logrus.Infof("Parsed config file %s", f.Name())
		logrus.Infof("Parsed config: %+v", cfg)
	}

	if defaults == nil {
		defaults = &Defaults{}
	}

	// Build config that needs no special processing
	overrideFields(&Default, &cfg.Default)
	overrideFields(&Default, &cliConfig.Default)
	overrideFields(&CNI, &cfg.CNI)
	overrideFields(&CNI, &cliConfig.CNI)

	// Logging setup
	overrideFields(&Logging, &cfg.Logging)
	overrideFields(&Logging, &cliConfig.Logging)
	logrus.SetLevel(logrus.Level(Logging.Level))
	if Logging.File != "" {
		var file *os.File
		file, err = os.OpenFile(Logging.File, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
		if err != nil {
			logrus.Errorf("failed to open logfile %s (%v). Ignoring..", Logging.File, err)
		} else {
			logrus.SetOutput(file)
		}
	}

	if err = buildKubernetesConfig(exec, &cliConfig, &cfg, defaults); err != nil {
		return "", err
	}
	logrus.Debugf("Default config: %+v", Default)
	logrus.Debugf("Logging config: %+v", Logging)
	logrus.Debugf("CNI config: %+v", CNI)
	logrus.Debugf("Kubernetes config: %+v", Kubernetes)

	return retConfigFile, nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// NewClientset creates a Kubernetes clientset from either a kubeconfig,
// TLS properties, or an apiserver URL
func NewClientset(conf *KubernetesConfig) (*kubernetes.Clientset, error) {
	var kconfig *rest.Config
	var err error

	if conf.Kubeconfig != "" {
		// uses the current context in kubeconfig
		kconfig, err = clientcmd.BuildConfigFromFlags("", conf.Kubeconfig)
	} else if strings.HasPrefix(conf.APIServer, "https") {
		if conf.APIServer == "" || conf.Token == "" {
			return nil, fmt.Errorf("TLS-secured apiservers require token and CA certificate")
		}
		kconfig = &rest.Config{
			Host:        conf.APIServer,
			BearerToken: conf.Token,
		}
		if conf.CACert != "" {
			if _, err := cert.NewPool(conf.CACert); err != nil {
				return nil, err
			}
			kconfig.TLSClientConfig = rest.TLSClientConfig{CAFile: conf.CACert}
		}
	} else if strings.HasPrefix(conf.APIServer, "http") {
		kconfig, err = clientcmd.BuildConfigFromFlags(conf.APIServer, "")
	} else {
		// Assume we are running from a container managed by kubernetes
		// and read the apiserver address and tokens from the
		// container's environment.
		kconfig, err = rest.InClusterConfig()
	}
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(kconfig)
}
