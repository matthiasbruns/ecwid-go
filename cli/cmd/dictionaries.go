package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/ecwid/dictionaries"
)

var dictionariesCmd = &cobra.Command{
	Use:   "dictionaries",
	Short: "Query reference dictionaries",
}

var dictionariesCountriesCmd = &cobra.Command{
	Use:   "countries",
	Short: "List all countries",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		opts := &dictionaries.CountriesOptions{}

		if v, _ := cmd.Flags().GetString("lang"); v != "" {
			opts.Lang = v
		}
		if v, _ := cmd.Flags().GetBool("with-states"); v {
			opts.WithStates = true
		}

		result, err := AppClient.Dictionaries.Countries(cmd.Context(), opts)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var dictionariesCurrenciesCmd = &cobra.Command{
	Use:   "currencies",
	Short: "List all currencies",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		lang, _ := cmd.Flags().GetString("lang")

		result, err := AppClient.Dictionaries.Currencies(cmd.Context(), lang)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var dictionariesStatesCmd = &cobra.Command{
	Use:   "states",
	Short: "List states for a country",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		countryCode, _ := cmd.Flags().GetString("country")
		if countryCode == "" {
			return fmt.Errorf("--country flag is required")
		}
		lang, _ := cmd.Flags().GetString("lang")

		result, err := AppClient.Dictionaries.States(cmd.Context(), countryCode, lang)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var dictionariesTaxClassesCmd = &cobra.Command{
	Use:   "tax-classes",
	Short: "List tax classes for a country",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		countryCode, _ := cmd.Flags().GetString("country")
		if countryCode == "" {
			return fmt.Errorf("--country flag is required")
		}
		lang, _ := cmd.Flags().GetString("lang")

		result, err := AppClient.Dictionaries.TaxClasses(cmd.Context(), countryCode, lang)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

func init() {
	rootCmd.AddCommand(dictionariesCmd)
	dictionariesCmd.AddCommand(dictionariesCountriesCmd)
	dictionariesCmd.AddCommand(dictionariesCurrenciesCmd)
	dictionariesCmd.AddCommand(dictionariesStatesCmd)
	dictionariesCmd.AddCommand(dictionariesTaxClassesCmd)

	// Countries flags.
	dictionariesCountriesCmd.Flags().String("lang", "", "Language ISO code (e.g. en, de)")
	dictionariesCountriesCmd.Flags().Bool("with-states", false, "Include states for each country")

	// Currencies flags.
	dictionariesCurrenciesCmd.Flags().String("lang", "", "Language ISO code (e.g. en, de)")

	// States flags.
	dictionariesStatesCmd.Flags().String("country", "", "Country code (required)")
	dictionariesStatesCmd.Flags().String("lang", "", "Language ISO code (e.g. en, de)")

	// Tax classes flags.
	dictionariesTaxClassesCmd.Flags().String("country", "", "Country code (required)")
	dictionariesTaxClassesCmd.Flags().String("lang", "", "Language ISO code (e.g. en, de)")
}
