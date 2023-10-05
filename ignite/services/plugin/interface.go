package plugin

import (
	"context"

	v1 "github.com/ignite/cli/ignite/services/plugin/grpc/v1"
)

// Flag type aliases.
const (
	FlagTypeString      = v1.Flag_TYPE_FLAG_STRING_UNSPECIFIED
	FlagTypeInt         = v1.Flag_TYPE_FLAG_INT
	FlagTypeUint        = v1.Flag_TYPE_FLAG_UINT
	FlagTypeInt64       = v1.Flag_TYPE_FLAG_INT64
	FlagTypeUint64      = v1.Flag_TYPE_FLAG_UINT64
	FlagTypeBool        = v1.Flag_TYPE_FLAG_BOOL
	FlagTypeStringSlice = v1.Flag_TYPE_FLAG_STRING_SLICE
)

// Type aliases for the current plugin version.
type (
	Manifest        = v1.Manifest
	Command         = v1.Command
	Hook            = v1.Hook
	Flag            = v1.Flag
	FlagType        = v1.Flag_Type
	ExecutedHook    = v1.ExecutedHook
	ExecutedCommand = v1.ExecutedCommand
)

// Interface defines the interface that all Ignite App must implement.
//
//go:generate mockery --srcpkg . --name Interface --structname PluginInterface --filename interface.go --with-expecter
type Interface interface {
	// Manifest declares the app's Command(s) and Hook(s).
	Manifest(context.Context) (*Manifest, error)

	// Execute will be invoked by ignite when an app Command is executed.
	// It is global for all commands declared in Manifest, if you have declared
	// multiple commands, use cmd.Path to distinguish them.
	Execute(context.Context, *ExecutedCommand) error

	// ExecuteHookPre is invoked by ignite when a command specified by the Hook
	// path is invoked.
	// It is global for all hooks declared in Manifest, if you have declared
	// multiple hooks, use hook.Name to distinguish them.
	ExecuteHookPre(context.Context, *ExecutedHook) error

	// ExecuteHookPost is invoked by ignite when a command specified by the hook
	// path is invoked.
	// It is global for all hooks declared in Manifest, if you have declared
	// multiple hooks, use hook.Name to distinguish them.
	ExecuteHookPost(context.Context, *ExecutedHook) error

	// ExecuteHookCleanUp is invoked by ignite when a command specified by the
	// hook path is invoked. Unlike ExecuteHookPost, it is invoked regardless of
	// execution status of the command and hooks.
	// It is global for all hooks declared in Manifest, if you have declared
	// multiple hooks, use hook.Name to distinguish them.
	ExecuteHookCleanUp(context.Context, *ExecutedHook) error
}
