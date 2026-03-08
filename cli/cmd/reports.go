package cmd

import (
	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/ecwid/reports"
)

var reportsCmd = &cobra.Command{
	Use:   "reports",
	Short: "View store reports and stats",
}

var reportsGetCmd = &cobra.Command{
	Use:   "get <reportType>",
	Short: "Get a report by type",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := &reports.ReportOptions{}

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

		result, err := AppClient.Reports.GetReport(cmd.Context(), args[0], opts)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var reportsStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get latest store stats",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		result, err := AppClient.Reports.LatestStats(cmd.Context(), nil)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

func init() {
	rootCmd.AddCommand(reportsCmd)
	reportsCmd.AddCommand(reportsGetCmd)
	reportsCmd.AddCommand(reportsStatsCmd)

	// Get report flags.
	reportsGetCmd.Flags().Int64("started-from", 0, "Start UNIX timestamp")
	reportsGetCmd.Flags().Int64("ended-at", 0, "End UNIX timestamp")
	reportsGetCmd.Flags().String("time-scale", "", "Time scale: hour, day, week, month, year")
	reportsGetCmd.Flags().String("compare-period", "", "Comparison period")
}
