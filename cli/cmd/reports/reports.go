package reports

import (
	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	apireports "github.com/matthiasbruns/ecwid-go/ecwid/reports"
)

// Cmd is the top-level reports command.
var Cmd = &cobra.Command{
	Use:   "reports",
	Short: "View store reports and stats",
}

var getCmd = &cobra.Command{
	Use:   "get <reportType>",
	Short: "Get a report by type",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := &apireports.ReportOptions{}

		if v, _ := cmd.Flags().GetInt64("started-from"); v > 0 {
			opts.StartedFrom = v
		}
		if v, _ := cmd.Flags().GetInt64("ended-at"); v > 0 {
			opts.EndedAt = v
		}
		if v, _ := cmd.Flags().GetString("time-scale"); v != "" {
			opts.TimeScaleValue = v
		}
		if v, _ := cmd.Flags().GetString("compare-period"); v != "" {
			opts.ComparePeriod = v
		}

		result, err := cmdutil.AppClient.Reports.GetReport(cmd.Context(), args[0], opts)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get latest store stats",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		result, err := cmdutil.AppClient.Reports.LatestStats(cmd.Context(), nil)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

func init() {
	Cmd.AddCommand(getCmd, statsCmd)

	// Get report flags.
	getCmd.Flags().Int64("started-from", 0, "Start UNIX timestamp")
	getCmd.Flags().Int64("ended-at", 0, "End UNIX timestamp")
	getCmd.Flags().String("time-scale", "", "Time scale: hour, day, week, month, year")
	getCmd.Flags().String("compare-period", "", "Comparison period")
}
