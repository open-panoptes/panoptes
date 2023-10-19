// Code generated by cli_gen.go DO NOT EDIT.
// source: github.com/rancher/opni/pkg/apis/management/v1/management.proto

package v1

import (
	context "context"
	errors "errors"
	cli "github.com/rancher/opni/internal/codegen/cli"
	v11 "github.com/rancher/opni/pkg/apis/capability/v1"
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
	cliutil "github.com/rancher/opni/pkg/opni/cliutil"
	flagutil "github.com/rancher/opni/pkg/util/flagutil"
	cobra "github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	strings "strings"
)

type contextKey_Management_type struct{}

var contextKey_Management contextKey_Management_type

func ContextWithManagementClient(ctx context.Context, client ManagementClient) context.Context {
	return context.WithValue(ctx, contextKey_Management, client)
}

func ManagementClientFromContext(ctx context.Context) (ManagementClient, bool) {
	client, ok := ctx.Value(contextKey_Management).(ManagementClient)
	return client, ok
}

var extraCmds_Management []*cobra.Command

func addExtraManagementCmd(custom *cobra.Command) {
	extraCmds_Management = append(extraCmds_Management, custom)
}

func BuildManagementCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "management",
		Short:             ``,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
	}

	cliutil.AddSubcommands(cmd, append([]*cobra.Command{
		BuildManagementCreateBootstrapTokenCmd(),
		BuildManagementRevokeBootstrapTokenCmd(),
		BuildManagementListBootstrapTokensCmd(),
		BuildManagementGetBootstrapTokenCmd(),
		BuildManagementListClustersCmd(),
		BuildManagementDeleteClusterCmd(),
		BuildManagementCertsInfoCmd(),
		BuildManagementGetClusterCmd(),
		BuildManagementGetClusterHealthStatusCmd(),
		BuildManagementEditClusterCmd(),
		BuildManagementListRBACBackendsCmd(),
		BuildManagementGetAvailableBackendPermissionsCmd(),
		BuildManagementCreateBackendRoleCmd(),
		BuildManagementUpdateBackendRoleCmd(),
		BuildManagementDeleteBackendRoleCmd(),
		BuildManagementGetBackendRoleCmd(),
		BuildManagementListBackendRolesCmd(),
		BuildManagementCreateRoleBindingCmd(),
		BuildManagementUpdateRoleBindingCmd(),
		BuildManagementDeleteRoleBindingCmd(),
		BuildManagementGetRoleBindingCmd(),
		BuildManagementListRoleBindingsCmd(),
		BuildManagementAPIExtensionsCmd(),
		BuildManagementGetConfigCmd(),
		BuildManagementUpdateConfigCmd(),
		BuildManagementListCapabilitiesCmd(),
		BuildManagementCapabilityInstallerCmd(),
		BuildManagementInstallCapabilityCmd(),
		BuildManagementUninstallCapabilityCmd(),
		BuildManagementCapabilityStatusCmd(),
		BuildManagementCapabilityUninstallStatusCmd(),
		BuildManagementCancelCapabilityUninstallCmd(),
		BuildManagementGetDashboardSettingsCmd(),
		BuildManagementUpdateDashboardSettingsCmd(),
	}, extraCmds_Management...)...)
	cli.AddOutputFlag(cmd)
	return cmd
}

func BuildManagementCreateBootstrapTokenCmd() *cobra.Command {
	in := &CreateBootstrapTokenRequest{}
	cmd := &cobra.Command{
		Use:   "create-bootstrap-token",
		Short: "",
		Long: `
HTTP handlers for this method:
- POST /tokens
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.CreateBootstrapToken(cmd.Context(), in)
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

func BuildManagementRevokeBootstrapTokenCmd() *cobra.Command {
	in := &v1.Reference{}
	cmd := &cobra.Command{
		Use:   "revoke-bootstrap-token",
		Short: "",
		Long: `
HTTP handlers for this method:
- DELETE /tokens/{id}
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.RevokeBootstrapToken(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildManagementListBootstrapTokensCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-bootstrap-tokens",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /tokens
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.ListBootstrapTokens(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	return cmd
}

func BuildManagementGetBootstrapTokenCmd() *cobra.Command {
	in := &v1.Reference{}
	cmd := &cobra.Command{
		Use:   "get-bootstrap-token",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /tokens/{id}
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.GetBootstrapToken(cmd.Context(), in)
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

func BuildManagementListClustersCmd() *cobra.Command {
	in := &ListClustersRequest{}
	cmd := &cobra.Command{
		Use:   "list-clusters",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /clusters
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.ListClusters(cmd.Context(), in)
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	cmd.RegisterFlagCompletionFunc("match-options", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"Default", "EmptySelectorMatchesAll", "EmptySelectorMatchesNone"}, cobra.ShellCompDirectiveDefault
	})
	return cmd
}

func BuildManagementDeleteClusterCmd() *cobra.Command {
	in := &v1.Reference{}
	cmd := &cobra.Command{
		Use:   "delete-cluster",
		Short: "",
		Long: `
HTTP handlers for this method:
- DELETE /clusters/{id}
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.DeleteCluster(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildManagementCertsInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "certs-info",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /certs
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.CertsInfo(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	return cmd
}

func BuildManagementGetClusterCmd() *cobra.Command {
	in := &v1.Reference{}
	cmd := &cobra.Command{
		Use:   "get-cluster",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /clusters/{id}
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.GetCluster(cmd.Context(), in)
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

func BuildManagementGetClusterHealthStatusCmd() *cobra.Command {
	in := &v1.Reference{}
	cmd := &cobra.Command{
		Use:   "get-cluster-health-status",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /clusters/{id}/health
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.GetClusterHealthStatus(cmd.Context(), in)
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

func BuildManagementEditClusterCmd() *cobra.Command {
	in := &EditClusterRequest{}
	cmd := &cobra.Command{
		Use:   "edit-cluster",
		Short: "",
		Long: `
HTTP handlers for this method:
- PUT /clusters/{cluster.id}
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.EditCluster(cmd.Context(), in)
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

func BuildManagementListRBACBackendsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-rbac-backends",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /rbac/backend
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.ListRBACBackends(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	return cmd
}

func BuildManagementGetAvailableBackendPermissionsCmd() *cobra.Command {
	in := &v1.CapabilityType{}
	cmd := &cobra.Command{
		Use:   "get-available-backend-permissions",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /rbac/backend/{name}/permissions
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.GetAvailableBackendPermissions(cmd.Context(), in)
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

func BuildManagementCreateBackendRoleCmd() *cobra.Command {
	in := &v1.BackendRole{}
	cmd := &cobra.Command{
		Use:   "backend-role create",
		Short: "",
		Long: `
HTTP handlers for this method:
- POST /rbac/backend/{capability.name}/roles
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if cmd.Flags().Lookup("interactive").Value.String() == "true" {
				if edited, err := cliutil.EditInteractive(in); err != nil {
					return err
				} else {
					in = edited
				}
			} else if fileName := cmd.Flags().Lookup("file").Value.String(); fileName != "" {
				if err := cliutil.LoadFromFile(in, fileName); err != nil {
					return err
				}
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.CreateBackendRole(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP("file", "f", "", "path to a file containing the config, or - to read from stdin")
	cmd.Flags().BoolP("interactive", "i", false, "edit the config interactively in an editor")
	cmd.MarkFlagsMutuallyExclusive("file", "interactive")
	cmd.MarkFlagFilename("file")
	return cmd
}

func BuildManagementUpdateBackendRoleCmd() *cobra.Command {
	in := &v1.BackendRole{}
	cmd := &cobra.Command{
		Use:   "backend-role update",
		Short: "",
		Long: `
HTTP handlers for this method:
- PUT /rbac/backend/{capability.name}/roles
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if cmd.Flags().Lookup("interactive").Value.String() == "true" {
				if edited, err := cliutil.EditInteractive(in); err != nil {
					return err
				} else {
					in = edited
				}
			} else if fileName := cmd.Flags().Lookup("file").Value.String(); fileName != "" {
				if err := cliutil.LoadFromFile(in, fileName); err != nil {
					return err
				}
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.UpdateBackendRole(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP("file", "f", "", "path to a file containing the config, or - to read from stdin")
	cmd.Flags().BoolP("interactive", "i", false, "edit the config interactively in an editor")
	cmd.MarkFlagsMutuallyExclusive("file", "interactive")
	cmd.MarkFlagFilename("file")
	return cmd
}

func BuildManagementDeleteBackendRoleCmd() *cobra.Command {
	in := &v1.BackendRoleRequest{}
	cmd := &cobra.Command{
		Use:   "backend-role delete",
		Short: "",
		Long: `
HTTP handlers for this method:
- DELETE /rbac/backend/{capability.name}/roles/{roleRef.id}
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.DeleteBackendRole(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildManagementGetBackendRoleCmd() *cobra.Command {
	in := &v1.BackendRoleRequest{}
	cmd := &cobra.Command{
		Use:   "backend-role get",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /rbac/backend/{capability.name}/roles/{roleRef.id}
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.GetBackendRole(cmd.Context(), in)
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

func BuildManagementListBackendRolesCmd() *cobra.Command {
	in := &v1.CapabilityType{}
	cmd := &cobra.Command{
		Use:   "backend-role list",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /rbac/backend/{name}/roles
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.ListBackendRoles(cmd.Context(), in)
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

func BuildManagementCreateRoleBindingCmd() *cobra.Command {
	in := &v1.RoleBinding{}
	cmd := &cobra.Command{
		Use:   "create-role-binding",
		Short: "",
		Long: `
HTTP handlers for this method:
- POST /rolebindings
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.CreateRoleBinding(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildManagementUpdateRoleBindingCmd() *cobra.Command {
	in := &v1.RoleBinding{}
	cmd := &cobra.Command{
		Use:   "update-role-binding",
		Short: "",
		Long: `
HTTP handlers for this method:
- PUT /rolebindings
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.UpdateRoleBinding(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildManagementDeleteRoleBindingCmd() *cobra.Command {
	in := &v1.Reference{}
	cmd := &cobra.Command{
		Use:   "delete-role-binding",
		Short: "",
		Long: `
HTTP handlers for this method:
- DELETE /rolebindings/{id}
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.DeleteRoleBinding(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildManagementGetRoleBindingCmd() *cobra.Command {
	in := &v1.Reference{}
	cmd := &cobra.Command{
		Use:   "get-role-binding",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /rolebindings/{id}
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.GetRoleBinding(cmd.Context(), in)
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

func BuildManagementListRoleBindingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-role-bindings",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /rolebindings
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.ListRoleBindings(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	return cmd
}

func BuildManagementAPIExtensionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api-extensions",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /apiextensions
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.APIExtensions(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	return cmd
}

func BuildManagementGetConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-config",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /config
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.GetConfig(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	return cmd
}

func BuildManagementUpdateConfigCmd() *cobra.Command {
	in := &UpdateConfigRequest{}
	cmd := &cobra.Command{
		Use:   "update-config",
		Short: "",
		Long: `
HTTP handlers for this method:
- PUT /config
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.UpdateConfig(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildManagementListCapabilitiesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-capabilities",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /capabilities
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.ListCapabilities(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	return cmd
}

func BuildManagementCapabilityInstallerCmd() *cobra.Command {
	in := &CapabilityInstallerRequest{}
	cmd := &cobra.Command{
		Use:   "capability-installer",
		Short: "Deprecated: For agent v2, use InstallCapability instead.",
		Long: `
HTTP handlers for this method:
- POST /capabilities/{name}/installer
`[1:],
		Deprecated:        "CapabilityInstaller is deprecated.",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.CapabilityInstaller(cmd.Context(), in)
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

func BuildManagementInstallCapabilityCmd() *cobra.Command {
	in := &CapabilityInstallRequest{}
	cmd := &cobra.Command{
		Use:   "install-capability",
		Short: "",
		Long: `
HTTP handlers for this method:
- POST /clusters/{target.cluster.id}/capabilities/{name}/install
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.InstallCapability(cmd.Context(), in)
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

func BuildManagementUninstallCapabilityCmd() *cobra.Command {
	in := &CapabilityUninstallRequest{}
	cmd := &cobra.Command{
		Use:   "uninstall-capability",
		Short: "",
		Long: `
HTTP handlers for this method:
- POST /clusters/{target.cluster.id}/capabilities/{name}/uninstall
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.UninstallCapability(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildManagementCapabilityStatusCmd() *cobra.Command {
	in := &CapabilityStatusRequest{}
	cmd := &cobra.Command{
		Use:   "capability-status",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /clusters/{cluster.id}/capabilities/{name}/status
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.CapabilityStatus(cmd.Context(), in)
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

func BuildManagementCapabilityUninstallStatusCmd() *cobra.Command {
	in := &CapabilityStatusRequest{}
	cmd := &cobra.Command{
		Use:   "capability-uninstall-status",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /clusters/{cluster.id}/capabilities/{name}/uninstall/status
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.CapabilityUninstallStatus(cmd.Context(), in)
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

func BuildManagementCancelCapabilityUninstallCmd() *cobra.Command {
	in := &CapabilityUninstallCancelRequest{}
	cmd := &cobra.Command{
		Use:   "cancel-capability-uninstall",
		Short: "",
		Long: `
HTTP handlers for this method:
- POST /clusters/{cluster.id}/capabilities/{name}/uninstall/cancel
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.CancelCapabilityUninstall(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildManagementGetDashboardSettingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-dashboard-settings",
		Short: "",
		Long: `
HTTP handlers for this method:
- GET /dashboard/settings
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.GetDashboardSettings(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	return cmd
}

func BuildManagementUpdateDashboardSettingsCmd() *cobra.Command {
	in := &DashboardSettings{}
	cmd := &cobra.Command{
		Use:   "update-dashboard-settings",
		Short: "",
		Long: `
HTTP handlers for this method:
- PUT /dashboard/settings
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ManagementClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.UpdateDashboardSettings(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func (in *CreateBootstrapTokenRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("CreateBootstrapTokenRequest", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.DurationpbValue(nil, &in.Ttl), strings.Join(append(prefix, "ttl"), "."), "")
	fs.StringToStringVar(&in.Labels, strings.Join(append(prefix, "labels"), "."), nil, "")
	fs.Int64Var(&in.MaxUsages, strings.Join(append(prefix, "max-usages"), "."), 0, "")
	return fs
}

func (in *ListClustersRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("ListClustersRequest", pflag.ExitOnError)
	fs.SortFlags = true
	if in.MatchLabels == nil {
		in.MatchLabels = &v1.LabelSelector{}
	}
	fs.AddFlagSet(in.MatchLabels.FlagSet(append(prefix, "match-labels")...))
	fs.Var(flagutil.EnumValue(&in.MatchOptions), strings.Join(append(prefix, "match-options"), "."), "")
	return fs
}

func (in *EditClusterRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("EditClusterRequest", pflag.ExitOnError)
	fs.SortFlags = true
	if in.Cluster == nil {
		in.Cluster = &v1.Reference{}
	}
	fs.AddFlagSet(in.Cluster.FlagSet(append(prefix, "cluster")...))
	fs.StringToStringVar(&in.Labels, strings.Join(append(prefix, "labels"), "."), nil, "")
	return fs
}

func (in *UpdateConfigRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("UpdateConfigRequest", pflag.ExitOnError)
	fs.SortFlags = true
	return fs
}

func (in *CapabilityInstallerRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("CapabilityInstallerRequest", pflag.ExitOnError)
	fs.SortFlags = true
	fs.StringVar(&in.Name, strings.Join(append(prefix, "name"), "."), "", "")
	fs.StringVar(&in.Token, strings.Join(append(prefix, "token"), "."), "", "")
	fs.StringVar(&in.Pin, strings.Join(append(prefix, "pin"), "."), "", "")
	return fs
}

func (in *CapabilityInstallRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("CapabilityInstallRequest", pflag.ExitOnError)
	fs.SortFlags = true
	fs.StringVar(&in.Name, strings.Join(append(prefix, "name"), "."), "", "")
	if in.Target == nil {
		in.Target = &v11.InstallRequest{}
	}
	fs.AddFlagSet(in.Target.FlagSet(append(prefix, "target")...))
	return fs
}

func (in *CapabilityUninstallRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("CapabilityUninstallRequest", pflag.ExitOnError)
	fs.SortFlags = true
	fs.StringVar(&in.Name, strings.Join(append(prefix, "name"), "."), "", "")
	if in.Target == nil {
		in.Target = &v11.UninstallRequest{}
	}
	fs.AddFlagSet(in.Target.FlagSet(append(prefix, "target")...))
	return fs
}

func (in *CapabilityStatusRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("CapabilityStatusRequest", pflag.ExitOnError)
	fs.SortFlags = true
	fs.StringVar(&in.Name, strings.Join(append(prefix, "name"), "."), "", "")
	if in.Cluster == nil {
		in.Cluster = &v1.Reference{}
	}
	fs.AddFlagSet(in.Cluster.FlagSet(append(prefix, "cluster")...))
	return fs
}

func (in *CapabilityUninstallCancelRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("CapabilityUninstallCancelRequest", pflag.ExitOnError)
	fs.SortFlags = true
	fs.StringVar(&in.Name, strings.Join(append(prefix, "name"), "."), "", "")
	if in.Cluster == nil {
		in.Cluster = &v1.Reference{}
	}
	fs.AddFlagSet(in.Cluster.FlagSet(append(prefix, "cluster")...))
	return fs
}

func (in *DashboardSettings) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("DashboardSettings", pflag.ExitOnError)
	fs.SortFlags = true
	if in.Global == nil {
		in.Global = &DashboardGlobalSettings{}
	}
	fs.AddFlagSet(in.Global.FlagSet(append(prefix, "global")...))
	fs.StringToStringVar(&in.User, strings.Join(append(prefix, "user"), "."), nil, "")
	return fs
}

func (in *DashboardGlobalSettings) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("DashboardGlobalSettings", pflag.ExitOnError)
	fs.SortFlags = true
	fs.StringVar(&in.DefaultImageRepository, strings.Join(append(prefix, "default-image-repository"), "."), "", "")
	fs.Var(flagutil.DurationpbValue(nil, &in.DefaultTokenTtl), strings.Join(append(prefix, "default-token-ttl"), "."), "")
	fs.StringToStringVar(&in.DefaultTokenLabels, strings.Join(append(prefix, "default-token-labels"), "."), nil, "")
	return fs
}
