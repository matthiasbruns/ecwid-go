package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/ecwid/categories"
)

var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Manage store categories",
}

// categories list

var (
	catListKeyword string
	catListParent  int64
	catListLimit   int
	catListOffset  int
)

var categoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List categories",
	RunE: func(cmd *cobra.Command, _ []string) error {
		opts := &categories.SearchOptions{
			Keyword: catListKeyword,
			Parent:  catListParent,
			Limit:   catListLimit,
			Offset:  catListOffset,
		}
		result, err := AppClient.Categories.Search(cmd.Context(), opts)
		if err != nil {
			return err
		}
		return outputResult(cmd, result.Items)
	},
}

// categories get

var categoriesGetCmd = &cobra.Command{
	Use:   "get <categoryId>",
	Short: "Get a category by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid category ID %q: %w", args[0], err)
		}
		result, err := AppClient.Categories.Get(cmd.Context(), id)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

// categories create

var (
	catCreateName        string
	catCreateParentID    int64
	catCreateDescription string
	catCreateEnabled     bool
)

var categoriesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new category",
	RunE: func(cmd *cobra.Command, _ []string) error {
		if catCreateName == "" {
			return fmt.Errorf("--name is required")
		}
		cat := &categories.Category{
			Name:        catCreateName,
			ParentID:    catCreateParentID,
			Description: catCreateDescription,
			Enabled:     &catCreateEnabled,
		}
		result, err := AppClient.Categories.Create(cmd.Context(), cat)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

// categories update

var (
	catUpdateName        string
	catUpdateParentID    int64
	catUpdateDescription string
	catUpdateEnabled     bool
)

var categoriesUpdateCmd = &cobra.Command{
	Use:   "update <categoryId>",
	Short: "Update an existing category",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid category ID %q: %w", args[0], err)
		}
		cat := &categories.Category{}
		if cmd.Flags().Changed("name") {
			cat.Name = catUpdateName
		}
		if cmd.Flags().Changed("parent-id") {
			cat.ParentID = catUpdateParentID
		}
		if cmd.Flags().Changed("description") {
			cat.Description = catUpdateDescription
		}
		if cmd.Flags().Changed("enabled") {
			cat.Enabled = &catUpdateEnabled
		}
		result, err := AppClient.Categories.Update(cmd.Context(), id, cat)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

// categories delete

var categoriesDeleteCmd = &cobra.Command{
	Use:   "delete <categoryId>",
	Short: "Delete a category by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid category ID %q: %w", args[0], err)
		}
		result, err := AppClient.Categories.Delete(cmd.Context(), id)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

func init() {
	// list flags
	categoriesListCmd.Flags().StringVar(&catListKeyword, "keyword", "", "filter by keyword")
	categoriesListCmd.Flags().Int64Var(&catListParent, "parent", 0, "filter by parent category ID")
	categoriesListCmd.Flags().IntVar(&catListLimit, "limit", 0, "maximum number of results")
	categoriesListCmd.Flags().IntVar(&catListOffset, "offset", 0, "result offset for pagination")

	// create flags
	categoriesCreateCmd.Flags().StringVar(&catCreateName, "name", "", "category name (required)")
	categoriesCreateCmd.Flags().Int64Var(&catCreateParentID, "parent-id", 0, "parent category ID")
	categoriesCreateCmd.Flags().StringVar(&catCreateDescription, "description", "", "category description")
	// Default true: new categories are enabled unless the caller opts out.
	categoriesCreateCmd.Flags().BoolVar(&catCreateEnabled, "enabled", true, "whether the category is enabled")

	// update flags
	categoriesUpdateCmd.Flags().StringVar(&catUpdateName, "name", "", "category name")
	categoriesUpdateCmd.Flags().Int64Var(&catUpdateParentID, "parent-id", 0, "parent category ID")
	categoriesUpdateCmd.Flags().StringVar(&catUpdateDescription, "description", "", "category description")
	// Default false here is irrelevant: the field is only sent when the flag is explicitly set.
	categoriesUpdateCmd.Flags().BoolVar(&catUpdateEnabled, "enabled", false, "whether the category is enabled")

	categoriesCmd.AddCommand(
		categoriesListCmd,
		categoriesGetCmd,
		categoriesCreateCmd,
		categoriesUpdateCmd,
		categoriesDeleteCmd,
	)
	rootCmd.AddCommand(categoriesCmd)
}
