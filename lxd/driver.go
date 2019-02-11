package lxd

import (
	"context"
	"time"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad/plugins/base"
	"github.com/hashicorp/nomad/plugins/drivers"
	"github.com/hashicorp/nomad/plugins/shared/hclspec"
	cstructs "github.com/hashicorp/nomad/client/structs"
)

const (
	// pluginName is the name of the plugin
	pluginName = "lxc"

	// fingerprintPeriod is the interval at which the driver will send fingerprint responses
	fingerprintPeriod = 30 * time.Second

	// taskHandleVersion is the version of task handle which this driver sets
	// and understands how to decode driver state
	taskHandleVersion = 1
)

var (
	// pluginInfo is the response returned for the PluginInfo RPC
	pluginInfo = &base.PluginInfoResponse{
		Type:              base.PluginTypeDriver,
		PluginApiVersions: []string{drivers.ApiVersion010},
		PluginVersion:     "0.1.0-dev",
		Name:              pluginName,
	}

	// configSpec is the hcl specification returned by the ConfigSchema RPC
	configSpec = hclspec.NewObject(map[string]*hclspec.Spec{
		"enabled": hclspec.NewDefault(
			hclspec.NewAttr("enabled", "bool", false),
			hclspec.NewLiteral("true"),
		),
		"volumes_enabled": hclspec.NewDefault(
			hclspec.NewAttr("volumes_enabled", "bool", false),
			hclspec.NewLiteral("true"),
		),
		"lxc_path": hclspec.NewAttr("lxc_path", "string", false),
	})

	// taskConfigSpec is the hcl specification for the driver config section of
	// a task within a job. It is returned in the TaskConfigSchema RPC
	taskConfigSpec = hclspec.NewObject(map[string]*hclspec.Spec{
		"template":       hclspec.NewAttr("template", "string", true),
		"distro":         hclspec.NewAttr("distro", "string", false),
		"release":        hclspec.NewAttr("release", "string", false),
		"arch":           hclspec.NewAttr("arch", "string", false),
		"image_variant":  hclspec.NewAttr("image_variant", "string", false),
		"image_server":   hclspec.NewAttr("image_server", "string", false),
		"gpg_key_id":     hclspec.NewAttr("gpg_key_id", "string", false),
		"gpg_key_server": hclspec.NewAttr("gpg_key_server", "string", false),
		"disable_gpg":    hclspec.NewAttr("disable_gpg", "string", false),
		"flush_cache":    hclspec.NewAttr("flush_cache", "string", false),
		"force_cache":    hclspec.NewAttr("force_cache", "string", false),
		"template_args":  hclspec.NewAttr("template_args", "list(string)", false),
		"log_level":      hclspec.NewAttr("log_level", "string", false),
		"verbosity":      hclspec.NewAttr("verbosity", "string", false),
		"volumes":        hclspec.NewAttr("volumes", "list(string)", false),
	})

	// capabilities is returned by the Capabilities RPC and indicates what
	// optional features this driver supports
	capabilities = &drivers.Capabilities{
		SendSignals: false,
		Exec:        false,
		FSIsolation: drivers.FSIsolationImage,
	}
)

// Driver is a driver for running LXC containers
type Driver struct {

}

// Config is the driver configuration set by the SetConfig RPC call
type Config struct {

}

// NewLXDDriver returns a new DriverPlugin implementation
func NewLXDDriver(logger hclog.Logger) drivers.DriverPlugin {
	ctx, cancel := context.WithCancel(context.Background())
	print(ctx)
	print(cancel)
	logger = logger.Named(pluginName)
	return &Driver{}
}

func (d *Driver) PluginInfo() (*base.PluginInfoResponse, error) {
	return nil, nil
}
func (d *Driver) SetConfig(config *base.Config) error {
	return nil
}

func (d *Driver) ConfigSchema() (*hclspec.Spec, error) {
	return nil, nil
}

func (d *Driver) TaskConfigSchema() (*hclspec.Spec, error) {
	return nil, nil
}

func (d *Driver) Capabilities() (*drivers.Capabilities, error) {
	return nil, nil
}

func (d *Driver) Fingerprint(context.Context) (<-chan *drivers.Fingerprint, error) {
	return nil, nil
}

func (d *Driver) RecoverTask(*drivers.TaskHandle) error {
	return nil
}
func (d *Driver) StartTask(*drivers.TaskConfig) (*drivers.TaskHandle, *drivers.DriverNetwork, error) {
	return nil, nil, nil
}
func (d *Driver) WaitTask(ctx context.Context, taskID string) (<-chan *drivers.ExitResult, error) {
	return nil, nil
}
func (d *Driver) StopTask(taskID string, timeout time.Duration, signal string) error {
	return nil
}
func (d *Driver) DestroyTask(taskID string, force bool) error {
	return nil
}
func (d *Driver) InspectTask(taskID string) (*drivers.TaskStatus, error) {
	return nil, nil
}
func (d *Driver) TaskStats(context.Context, string, time.Duration) (<-chan *cstructs.TaskResourceUsage, error) {
	return nil, nil
}
func (d *Driver) TaskEvents(context.Context) (<-chan *drivers.TaskEvent, error) {
	return nil, nil
}

func (d *Driver) SignalTask(taskID string, signal string) error {
	return nil
}
func (d *Driver) ExecTask(taskID string, cmd []string, timeout time.Duration) (*drivers.ExecTaskResult, error) {
	return nil, nil
}