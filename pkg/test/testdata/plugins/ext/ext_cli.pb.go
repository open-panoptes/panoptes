// Code generated by cli_gen.go DO NOT EDIT.
// source: github.com/open-panoptes/opni/pkg/test/testdata/plugins/ext/ext.proto

package ext

import (
	context "context"
	errors "errors"
	cli "github.com/open-panoptes/opni/internal/codegen/cli"
	cliutil "github.com/open-panoptes/opni/pkg/opni/cliutil"
	storage "github.com/open-panoptes/opni/pkg/storage"
	flagutil "github.com/open-panoptes/opni/pkg/util/flagutil"
	lo "github.com/samber/lo"
	cobra "github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
	errdetails "google.golang.org/genproto/googleapis/rpc/errdetails"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	strings "strings"
)

type contextKey_Ext_type struct{}

var contextKey_Ext contextKey_Ext_type

func ContextWithExtClient(ctx context.Context, client ExtClient) context.Context {
	return context.WithValue(ctx, contextKey_Ext, client)
}

func ExtClientFromContext(ctx context.Context) (ExtClient, bool) {
	client, ok := ctx.Value(contextKey_Ext).(ExtClient)
	return client, ok
}

type contextKey_Ext2_type struct{}

var contextKey_Ext2 contextKey_Ext2_type

func ContextWithExt2Client(ctx context.Context, client Ext2Client) context.Context {
	return context.WithValue(ctx, contextKey_Ext2, client)
}

func Ext2ClientFromContext(ctx context.Context) (Ext2Client, bool) {
	client, ok := ctx.Value(contextKey_Ext2).(Ext2Client)
	return client, ok
}

var extraCmds_Ext []*cobra.Command

func addExtraExtCmd(custom *cobra.Command) {
	extraCmds_Ext = append(extraCmds_Ext, custom)
}

func BuildExtCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "ext",
		Short:             ``,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
	}

	cliutil.AddSubcommands(cmd, append([]*cobra.Command{
		BuildExtFooCmd(),
		BuildExtBarCmd(),
		BuildExtBazCmd(),
		BuildExtSetCmd(),
	}, extraCmds_Ext...)...)
	cli.AddOutputFlag(cmd)
	return cmd
}

func BuildExtFooCmd() *cobra.Command {
	in := &FooRequest{}
	cmd := &cobra.Command{
		Use:   "foo",
		Short: "",
		Long: `
HTTP handlers for this method:
- POST /foo
- GET /foo
- PUT /foo
- DELETE /foo
- PATCH /foo
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ExtClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.Foo(cmd.Context(), in)
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

func BuildExtBarCmd() *cobra.Command {
	in := &BarRequest{}
	cmd := &cobra.Command{
		Use:   "bar",
		Short: "",
		Long: `
HTTP handlers for this method:
- POST /bar/{param1}/{param2}
- GET /bar/{param1}/{param2}/{param3}
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ExtClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.Bar(cmd.Context(), in)
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

func BuildExtBazCmd() *cobra.Command {
	in := &BazRequest{}
	cmd := &cobra.Command{
		Use:   "baz",
		Short: "",
		Long: `
HTTP handlers for this method:
- POST /baz
- POST /baz/{paramMsg.paramBool}/{paramMsg.paramString}/{paramMsg.paramEnum}
- POST /baz/{paramMsg.paramMsg.paramMsg.paramMsg.paramString}
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ExtClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.Baz(cmd.Context(), in)
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	cmd.RegisterFlagCompletionFunc("param-enum", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"UNKNOWN", "FOO", "BAR"}, cobra.ShellCompDirectiveDefault
	})
	return cmd
}

func BuildExtSetCmd() *cobra.Command {
	in := &SetRequest{}
	cmd := &cobra.Command{
		Use:   "set",
		Short: "",
		Long: `
HTTP handlers for this method:
- PUT /set/{node.id}
- PUT /set/example/{node.id}
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := ExtClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.Set(cmd.Context(), in)
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

var extraCmds_Ext2 []*cobra.Command

func addExtraExt2Cmd(custom *cobra.Command) {
	extraCmds_Ext2 = append(extraCmds_Ext2, custom)
}

func BuildExt2Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "ext-2",
		Short:             ``,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
	}

	cmd.AddCommand(BuildExt2FooCmd())
	cli.AddOutputFlag(cmd)
	return cmd
}

func BuildExt2FooCmd() *cobra.Command {
	in := &FooRequest{}
	cmd := &cobra.Command{
		Use:               "foo",
		Short:             "",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := Ext2ClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.Foo(cmd.Context(), in)
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

func (in *FooRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("FooRequest", pflag.ExitOnError)
	fs.SortFlags = true
	fs.StringVar(&in.Request, strings.Join(append(prefix, "request"), "."), "", "")
	return fs
}

func (in *BarRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("BarRequest", pflag.ExitOnError)
	fs.SortFlags = true
	fs.StringVar(&in.Param1, strings.Join(append(prefix, "param-1"), "."), "", "")
	fs.StringVar(&in.Param2, strings.Join(append(prefix, "param-2"), "."), "", "")
	fs.StringVar(&in.Param3, strings.Join(append(prefix, "param-3"), "."), "", "")
	return fs
}

func (in *BazRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("BazRequest", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Float64Var(&in.ParamFloat64, strings.Join(append(prefix, "param-float-64"), "."), 0.0, "")
	fs.Int64Var(&in.ParamInt64, strings.Join(append(prefix, "param-int-64"), "."), 0, "")
	fs.BoolVar(&in.ParamBool, strings.Join(append(prefix, "param-bool"), "."), false, "")
	fs.StringVar(&in.ParamString, strings.Join(append(prefix, "param-string"), "."), "", "")
	fs.BytesHexVar(&in.ParamBytes, strings.Join(append(prefix, "param-bytes"), "."), nil, "")
	fs.Var(flagutil.EnumValue(&in.ParamEnum), strings.Join(append(prefix, "param-enum"), "."), "")
	fs.Var(flagutil.DurationpbValue(nil, &in.ParamDuration), strings.Join(append(prefix, "param-duration"), "."), "")
	fs.StringSliceVar(&in.ParamRepeatedString, strings.Join(append(prefix, "param-repeated-string"), "."), nil, "")
	return fs
}

func (in *SetRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("SetRequest", pflag.ExitOnError)
	fs.SortFlags = true
	if in.Node == nil {
		in.Node = &Reference{}
	}
	fs.AddFlagSet(in.Node.FlagSet(append(prefix, "node")...))
	fs.StringVar(&in.Value, strings.Join(append(prefix, "value"), "."), "", "")
	if in.Example == nil {
		in.Example = &ExampleValue{}
	}
	fs.AddFlagSet(in.Example.FlagSet(append(prefix, "example")...))
	return fs
}

func (in *Reference) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("Reference", pflag.ExitOnError)
	fs.SortFlags = true
	fs.StringVar(&in.Id, strings.Join(append(prefix, "id"), "."), "", "")
	return fs
}

func (in *ExampleValue) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("ExampleValue", pflag.ExitOnError)
	fs.SortFlags = true
	fs.StringVar(&in.Value, strings.Join(append(prefix, "value"), "."), "", "")
	return fs
}

func (in *FooResponse) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("FooResponse", pflag.ExitOnError)
	fs.SortFlags = true
	fs.StringVar(&in.Response, strings.Join(append(prefix, "response"), "."), "", "")
	return fs
}

func (in *BarResponse) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("BarResponse", pflag.ExitOnError)
	fs.SortFlags = true
	fs.StringVar(&in.Param1, strings.Join(append(prefix, "param-1"), "."), "", "")
	fs.StringVar(&in.Param2, strings.Join(append(prefix, "param-2"), "."), "", "")
	fs.StringVar(&in.Param3, strings.Join(append(prefix, "param-3"), "."), "", "")
	return fs
}

func (in *SampleConfiguration) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("SampleConfiguration", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.BoolPtrValue(nil, &in.Enabled), strings.Join(append(prefix, "enabled"), "."), "")
	fs.Var(flagutil.StringPtrValue(nil, &in.StringField), strings.Join(append(prefix, "string-field"), "."), "")
	fs.Var(flagutil.StringPtrValue(nil, &in.SecretField), strings.Join(append(prefix, "secret-field"), "."), "\x1b[31m[secret]\x1b[0m ")
	fs.StringToStringVar(&in.MapField, strings.Join(append(prefix, "map-field"), "."), nil, "")
	fs.StringSliceVar(&in.RepeatedField, strings.Join(append(prefix, "repeated-field"), "."), nil, "")
	if in.MessageField == nil {
		in.MessageField = &SampleMessage{}
	}
	fs.AddFlagSet(in.MessageField.FlagSet(append(prefix, "message-field")...))
	return fs
}

func (in *SampleConfiguration) RedactSecrets() {
	if in == nil {
		return
	}
	if in.GetSecretField() != "" {
		in.SecretField = flagutil.Ptr("***")
	}
}

func (in *SampleConfiguration) UnredactSecrets(unredacted *SampleConfiguration) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if in.GetSecretField() == "***" {
		if unredacted.GetSecretField() == "" {
			details = append(details, &errdetails.ErrorInfo{
				Reason:   "DISCONTINUITY",
				Metadata: map[string]string{"field": "secretField"},
			})
		} else {
			*in.SecretField = *unredacted.SecretField
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *SampleMessage) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("SampleMessage", pflag.ExitOnError)
	fs.SortFlags = true
	if in.Field1 == nil {
		in.Field1 = &Sample1FieldMsg{}
	}
	fs.AddFlagSet(in.Field1.FlagSet(append(prefix, "field-1")...))
	if in.Field2 == nil {
		in.Field2 = &Sample2FieldMsg{}
	}
	fs.AddFlagSet(in.Field2.FlagSet(append(prefix, "field-2")...))
	if in.Field3 == nil {
		in.Field3 = &Sample3FieldMsg{}
	}
	fs.AddFlagSet(in.Field3.FlagSet(append(prefix, "field-3")...))
	if in.Field4 == nil {
		in.Field4 = &Sample4FieldMsg{}
	}
	fs.AddFlagSet(in.Field4.FlagSet(append(prefix, "field-4")...))
	if in.Field5 == nil {
		in.Field5 = &Sample5FieldMsg{}
	}
	fs.AddFlagSet(in.Field5.FlagSet(append(prefix, "field-5")...))
	if in.Field6 == nil {
		in.Field6 = &Sample6FieldMsg{}
	}
	fs.AddFlagSet(in.Field6.FlagSet(append(prefix, "field-6")...))
	if in.Msg == nil {
		in.Msg = &SampleMessage2{}
	}
	fs.AddFlagSet(in.Msg.FlagSet(append(prefix, "msg")...))
	return fs
}

func (in *Sample1FieldMsg) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("Sample1FieldMsg", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Int32Var(&in.Field1, strings.Join(append(prefix, "field-1"), "."), 0, "")
	return fs
}

func (in *Sample2FieldMsg) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("Sample2FieldMsg", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Int32Var(&in.Field1, strings.Join(append(prefix, "field-1"), "."), 0, "")
	fs.Int32Var(&in.Field2, strings.Join(append(prefix, "field-2"), "."), 0, "")
	return fs
}

func (in *Sample3FieldMsg) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("Sample3FieldMsg", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Int32Var(&in.Field1, strings.Join(append(prefix, "field-1"), "."), 0, "")
	fs.Int32Var(&in.Field2, strings.Join(append(prefix, "field-2"), "."), 0, "")
	fs.Int32Var(&in.Field3, strings.Join(append(prefix, "field-3"), "."), 0, "")
	return fs
}

func (in *Sample4FieldMsg) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("Sample4FieldMsg", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Int32Var(&in.Field1, strings.Join(append(prefix, "field-1"), "."), 0, "")
	fs.Int32Var(&in.Field2, strings.Join(append(prefix, "field-2"), "."), 0, "")
	fs.Int32Var(&in.Field3, strings.Join(append(prefix, "field-3"), "."), 0, "")
	fs.Int32Var(&in.Field4, strings.Join(append(prefix, "field-4"), "."), 0, "")
	return fs
}

func (in *Sample5FieldMsg) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("Sample5FieldMsg", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Int32Var(&in.Field1, strings.Join(append(prefix, "field-1"), "."), 0, "")
	fs.Int32Var(&in.Field2, strings.Join(append(prefix, "field-2"), "."), 0, "")
	fs.Int32Var(&in.Field3, strings.Join(append(prefix, "field-3"), "."), 0, "")
	fs.Int32Var(&in.Field4, strings.Join(append(prefix, "field-4"), "."), 0, "")
	fs.Int32Var(&in.Field5, strings.Join(append(prefix, "field-5"), "."), 0, "")
	return fs
}

func (in *Sample6FieldMsg) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("Sample6FieldMsg", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Int32Var(&in.Field1, strings.Join(append(prefix, "field-1"), "."), 0, "")
	fs.Int32Var(&in.Field2, strings.Join(append(prefix, "field-2"), "."), 0, "")
	fs.Int32Var(&in.Field3, strings.Join(append(prefix, "field-3"), "."), 0, "")
	fs.Int32Var(&in.Field4, strings.Join(append(prefix, "field-4"), "."), 0, "")
	fs.Int32Var(&in.Field5, strings.Join(append(prefix, "field-5"), "."), 0, "")
	fs.Int32Var(&in.Field6, strings.Join(append(prefix, "field-6"), "."), 0, "")
	return fs
}

func (in *SampleMessage2) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("SampleMessage2", pflag.ExitOnError)
	fs.SortFlags = true
	if in.Field1 == nil {
		in.Field1 = &Sample1FieldMsg{}
	}
	fs.AddFlagSet(in.Field1.FlagSet(append(prefix, "field-1")...))
	if in.Field2 == nil {
		in.Field2 = &Sample2FieldMsg{}
	}
	fs.AddFlagSet(in.Field2.FlagSet(append(prefix, "field-2")...))
	if in.Field3 == nil {
		in.Field3 = &Sample3FieldMsg{}
	}
	fs.AddFlagSet(in.Field3.FlagSet(append(prefix, "field-3")...))
	if in.Field4 == nil {
		in.Field4 = &Sample4FieldMsg{}
	}
	fs.AddFlagSet(in.Field4.FlagSet(append(prefix, "field-4")...))
	if in.Field5 == nil {
		in.Field5 = &Sample5FieldMsg{}
	}
	fs.AddFlagSet(in.Field5.FlagSet(append(prefix, "field-5")...))
	if in.Field6 == nil {
		in.Field6 = &Sample6FieldMsg{}
	}
	fs.AddFlagSet(in.Field6.FlagSet(append(prefix, "field-6")...))
	return fs
}

func (in *SampleDryRunRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("SampleDryRunRequest", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.EnumValue(&in.Target), strings.Join(append(prefix, "target"), "."), "")
	fs.Var(flagutil.EnumValue(&in.Action), strings.Join(append(prefix, "action"), "."), "")
	if in.Spec == nil {
		in.Spec = &SampleConfiguration{}
	}
	fs.AddFlagSet(in.Spec.FlagSet(append(prefix, "spec")...))
	return fs
}

func (in *SampleDryRunRequest) RedactSecrets() {
	if in == nil {
		return
	}
	in.Spec.RedactSecrets()
}

func (in *SampleDryRunRequest) UnredactSecrets(unredacted *SampleDryRunRequest) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if err := in.Spec.UnredactSecrets(unredacted.GetSpec()); storage.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "spec." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *SampleDryRunResponse) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("SampleDryRunResponse", pflag.ExitOnError)
	fs.SortFlags = true
	if in.Current == nil {
		in.Current = &SampleConfiguration{}
	}
	fs.AddFlagSet(in.Current.FlagSet(append(prefix, "current")...))
	if in.Modified == nil {
		in.Modified = &SampleConfiguration{}
	}
	fs.AddFlagSet(in.Modified.FlagSet(append(prefix, "modified")...))
	return fs
}

func (in *SampleDryRunResponse) RedactSecrets() {
	if in == nil {
		return
	}
	in.Current.RedactSecrets()
	in.Modified.RedactSecrets()
}

func (in *SampleDryRunResponse) UnredactSecrets(unredacted *SampleDryRunResponse) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if err := in.Current.UnredactSecrets(unredacted.GetCurrent()); storage.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "current." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if err := in.Modified.UnredactSecrets(unredacted.GetModified()); storage.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "modified." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *SampleConfigurationHistoryResponse) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("SampleConfigurationHistoryResponse", pflag.ExitOnError)
	fs.SortFlags = true
	return fs
}
