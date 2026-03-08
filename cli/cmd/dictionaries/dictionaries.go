package dictionaries

import (
	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	apidict "github.com/matthiasbruns/ecwid-go/ecwid/dictionaries"
)

// Cmd is the top-level dictionaries command.
var Cmd = &cobra.Command{
	Use:   "dictionaries",
	Short: "Query reference dictionaries",
}

var countriesCmd = &cobra.Command{
	Use:   "countries",
	Short: "List all countries",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		opts := &apidict.CountriesOptions{}

		if v, _ := cmd.Flags().GetString("lang"); v != "" {
			opts.Lang = v
		}
		if v, _ := cmd.Flags().GetBool("with-states"); v {
			opts.WithStates = true
		}

		result, err := cmdutil.AppClient.Dictionaries.Countries(cmd.Context(), opts)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var currenciesCmd = &cobra.Command{
	Use:   "currencies",
	Short: "List all currencies",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		lang, _ := cmd.Flags().GetString("lang")

		result, err := cmdutil.AppClient.Dictionaries.Currencies(cmd.Context(), lang)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var statesCmd = &cobra.Command{
	Use:   "states",
	Short: "List states for a country",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		countryCode, _ := cmd.Flags().GetString("country")
		lang, _ := cmd.Flags().GetString("lang")

		result, err := cmdutil.AppClient.Dictionaries.States(cmd.Context(), countryCode, lang)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var taxClassesCmd = &cobra.Command{
	Use:   "tax-classes",
	Short: "List tax classes for a country",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		countryCode, _ := cmd.Flags().GetString("country")
		lang, _ := cmd.Flags().GetString("lang")

		result, err := cmdutil.AppClient.Dictionaries.TaxClasses(cmd.Context(), countryCode, lang)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

func init() {
	Cmd.AddCommand(countriesCmd, currenciesCmd, statesCmd, taxClassesCmd)

	// Countries flags.
	countriesCmd.Flags().String("lang", "", "Language ISO code (e.g. en, de)")
	countriesCmd.Flags().Bool("with-states", false, "Include states for each country")

	// Currencies flags.
	currenciesCmd.Flags().String("lang", "", "Language ISO code (e.g. en, de)")

	// States flags.
	statesCmd.Flags().String("country", "", "Country code (required)")
	statesCmd.Flags().String("lang", "", "Language ISO code (e.g. en, de)")
	_ = statesCmd.MarkFlagRequired("country")

	// Tax classes flags.
	taxClassesCmd.Flags().String("country", "", "Country code (required)")
	taxClassesCmd.Flags().String("lang", "", "Language ISO code (e.g. en, de)")
	_ = taxClassesCmd.MarkFlagRequired("country")
}
