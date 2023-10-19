// Code generated by cli_gen.go DO NOT EDIT.
// source: github.com/rancher/opni/pkg/apis/capability/v1/capability.proto

package v1

import (
	context "context"
	errors "errors"
	cli "github.com/rancher/opni/internal/codegen/cli"
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
	cliutil "github.com/rancher/opni/pkg/opni/cliutil"
	cobra "github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	strings "strings"
)

type contextKey_Backend_type struct{}

var contextKey_Backend contextKey_Backend_type

func ContextWithBackendClient(ctx context.Context, client BackendClient) context.Context {
	return context.WithValue(ctx, contextKey_Backend, client)
}

func BackendClientFromContext(ctx context.Context) (BackendClient, bool) {
	client, ok := ctx.Value(contextKey_Backend).(BackendClient)
	return client, ok
}

type contextKey_Node_type struct{}

var contextKey_Node contextKey_Node_type

func ContextWithNodeClient(ctx context.Context, client NodeClient) context.Context {
	return context.WithValue(ctx, contextKey_Node, client)
}

func NodeClientFromContext(ctx context.Context) (NodeClient, bool) {
	client, ok := ctx.Value(contextKey_Node).(NodeClient)
	return client, ok
}

type contextKey_RBACManager_type struct{}

var contextKey_RBACManager contextKey_RBACManager_type

func ContextWithRBACManagerClient(ctx context.Context, client RBACManagerClient) context.Context {
	return context.WithValue(ctx, contextKey_RBACManager, client)
}

func RBACManagerClientFromContext(ctx context.Context) (RBACManagerClient, bool) {
	client, ok := ctx.Value(contextKey_RBACManager).(RBACManagerClient)
	return client, ok
}

var extraCmds_Backend []*cobra.Command

func addExtraBackendCmd(custom *cobra.Command) {
	extraCmds_Backend = append(extraCmds_Backend, custom)
}

func BuildBackendCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "backend",
		Short:             ``,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
	}

	cliutil.AddSubcommands(cmd, append([]*cobra.Command{
		BuildBackendInfoCmd(),
		BuildBackendListCmd(),
		BuildBackendInstallCmd(),
		BuildBackendStatusCmd(),
		BuildBackendUninstallCmd(),
		BuildBackendUninstallStatusCmd(),
		BuildBackendCancelUninstallCmd(),
	}, extraCmds_Backend...)...)
	cli.AddOutputFlag(cmd)
	return cmd
}

func BuildBackendInfoCmd() *cobra.Command {
	in := &v1.Reference{}
	cmd := &cobra.Command{
		Use:               "info",
		Short:             "Returns info about the backend, including capability name.",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := BackendClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.Info(cmd.Context(), in)
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildBackendListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "list",
		Short:             "Returns a list of capabilities available in the backend.",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := BackendClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.List(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	return cmd
}

func BuildBackendInstallCmd() *cobra.Command {
	in := &InstallRequest{}
	cmd := &cobra.Command{
		Use:               "install",
		Short:             "Installs the capability on an agent.",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := BackendClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.Install(cmd.Context(), in)
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildBackendStatusCmd() *cobra.Command {
	in := &StatusRequest{}
	cmd := &cobra.Command{
		Use:               "status",
		Short:             "Returns common runtime config info for this capability from a specific agent.",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := BackendClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.Status(cmd.Context(), in)
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildBackendUninstallCmd() *cobra.Command {
	in := &UninstallRequest{}
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Requests the backend to clean up any resources it owns and prepare",
		Long: `
for uninstallation. This process is asynchronous. The status of the
operation can be queried using the UninstallStatus method, or canceled
using the CancelUninstall method.
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := BackendClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.Uninstall(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildBackendUninstallStatusCmd() *cobra.Command {
	in := &UninstallStatusRequest{}
	cmd := &cobra.Command{
		Use:               "uninstall-status",
		Short:             "Gets the status of the uninstall task for the given cluster.",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := BackendClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.UninstallStatus(cmd.Context(), in)
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildBackendCancelUninstallCmd() *cobra.Command {
	in := &CancelUninstallRequest{}
	cmd := &cobra.Command{
		Use:               "cancel-uninstall",
		Short:             "Cancels an uninstall task for the given cluster, if it is still pending.",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := BackendClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.CancelUninstall(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

var extraCmds_Node []*cobra.Command

func addExtraNodeCmd(custom *cobra.Command) {
	extraCmds_Node = append(extraCmds_Node, custom)
}

func BuildNodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "node",
		Short:             ``,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
	}

	cmd.AddCommand(BuildNodeSyncNowCmd())
	cli.AddOutputFlag(cmd)
	return cmd
}

func BuildNodeSyncNowCmd() *cobra.Command {
	in := &Filter{}
	cmd := &cobra.Command{
		Use:               "sync-now",
		Short:             "",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := NodeClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.SyncNow(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

var extraCmds_RBACManager []*cobra.Command

func addExtraRBACManagerCmd(custom *cobra.Command) {
	extraCmds_RBACManager = append(extraCmds_RBACManager, custom)
}

func BuildRBACManagerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "rbac-manager",
		Short:             ``,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
	}

	cliutil.AddSubcommands(cmd, append([]*cobra.Command{
		BuildRBACManagerInfoCmd(),
		BuildRBACManagerGetAvailablePermissionsCmd(),
		BuildRBACManagerGetRoleCmd(),
		BuildRBACManagerCreateRoleCmd(),
		BuildRBACManagerUpdateRoleCmd(),
		BuildRBACManagerDeleteRoleCmd(),
		BuildRBACManagerListRolesCmd(),
	}, extraCmds_RBACManager...)...)
	cli.AddOutputFlag(cmd)
	return cmd
}

func BuildRBACManagerInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "info",
		Short:             "Returns info about the manager, including capability name",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := RBACManagerClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.Info(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	return cmd
}

func BuildRBACManagerGetAvailablePermissionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "get-available-permissions",
		Short:             "",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := RBACManagerClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.GetAvailablePermissions(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	return cmd
}

func BuildRBACManagerGetRoleCmd() *cobra.Command {
	in := &v1.Reference{}
	cmd := &cobra.Command{
		Use:               "get-role",
		Short:             "",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := RBACManagerClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.GetRole(cmd.Context(), in)
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildRBACManagerCreateRoleCmd() *cobra.Command {
	in := &v1.Role{}
	cmd := &cobra.Command{
		Use:               "create-role",
		Short:             "",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := RBACManagerClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.CreateRole(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildRBACManagerUpdateRoleCmd() *cobra.Command {
	in := &v1.Role{}
	cmd := &cobra.Command{
		Use:               "update-role",
		Short:             "",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := RBACManagerClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.UpdateRole(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildRBACManagerDeleteRoleCmd() *cobra.Command {
	in := &v1.Reference{}
	cmd := &cobra.Command{
		Use:               "delete-role",
		Short:             "",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := RBACManagerClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.DeleteRole(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildRBACManagerListRolesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "list-roles",
		Short:             "",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := RBACManagerClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.ListRoles(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	return cmd
}

func (in *InstallRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("InstallRequest", pflag.ExitOnError)
	fs.SortFlags = true
	if in.Capability == nil {
		in.Capability = &v1.Reference{}
	}
	fs.AddFlagSet(in.Capability.FlagSet(append(prefix, "capability")...))
	if in.Agent == nil {
		in.Agent = &v1.Reference{}
	}
	fs.AddFlagSet(in.Agent.FlagSet(append(prefix, "agent")...))
	fs.BoolVar(&in.IgnoreWarnings, strings.Join(append(prefix, "ignore-warnings"), "."), false, "")
	return fs
}

func (in *StatusRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("StatusRequest", pflag.ExitOnError)
	fs.SortFlags = true
	if in.Capability == nil {
		in.Capability = &v1.Reference{}
	}
	fs.AddFlagSet(in.Capability.FlagSet(append(prefix, "capability")...))
	if in.Agent == nil {
		in.Agent = &v1.Reference{}
	}
	fs.AddFlagSet(in.Agent.FlagSet(append(prefix, "agent")...))
	return fs
}

func (in *UninstallRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("UninstallRequest", pflag.ExitOnError)
	fs.SortFlags = true
	if in.Capability == nil {
		in.Capability = &v1.Reference{}
	}
	fs.AddFlagSet(in.Capability.FlagSet(append(prefix, "capability")...))
	if in.Agent == nil {
		in.Agent = &v1.Reference{}
	}
	fs.AddFlagSet(in.Agent.FlagSet(append(prefix, "agent")...))
	if in.Options == nil {
		in.Options = &structpb.Struct{}
	}
	fs.AddFlagSet(in.Options.FlagSet(append(prefix, "options")...))
	return fs
}

func (in *UninstallStatusRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("UninstallStatusRequest", pflag.ExitOnError)
	fs.SortFlags = true
	if in.Capability == nil {
		in.Capability = &v1.Reference{}
	}
	fs.AddFlagSet(in.Capability.FlagSet(append(prefix, "capability")...))
	if in.Agent == nil {
		in.Agent = &v1.Reference{}
	}
	fs.AddFlagSet(in.Agent.FlagSet(append(prefix, "agent")...))
	return fs
}

func (in *CancelUninstallRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("CancelUninstallRequest", pflag.ExitOnError)
	fs.SortFlags = true
	if in.Capability == nil {
		in.Capability = &v1.Reference{}
	}
	fs.AddFlagSet(in.Capability.FlagSet(append(prefix, "capability")...))
	if in.Agent == nil {
		in.Agent = &v1.Reference{}
	}
	fs.AddFlagSet(in.Agent.FlagSet(append(prefix, "agent")...))
	return fs
}

func (in *Filter) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("Filter", pflag.ExitOnError)
	fs.SortFlags = true
	fs.StringSliceVar(&in.CapabilityNames, strings.Join(append(prefix, "capability-names"), "."), nil, "")
	return fs
}
