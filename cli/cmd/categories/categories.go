package categories

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	apicategories "github.com/matthiasbruns/ecwid-go/ecwid/categories"
)

// Cmd is the top-level categories command.
var Cmd = &cobra.Command{
	Use:   "categories",
	Short: "Manage store categories",
}

var (
	catListKeyword string
	catListParent  int64
	catListLimit   int
	catListOffset  int
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List categories",
	RunE: func(cmd *cobra.Command, _ []string) error {
		opts := &apicategories.SearchOptions{
			Keyword: catListKeyword,
			Parent:  catListParent,
			Limit:   catListLimit,
			Offset:  catListOffset,
		}
		result, err := cmdutil.AppClient.Categories.Search(cmd.Context(), opts)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result.Items)
	},
}

var getCmd = &cobra.Command{
	Use:   "get <categoryId>",
	Short: "Get a category by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid category ID %q: %w", args[0], err)
		}
		if id <= 0 {
			return fmt.Errorf("category ID must be a positive integer, got %d", id)
		}
		result, err := cmdutil.AppClient.Categories.Get(cmd.Context(), id)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var (
	catCreateName        string
	catCreateParentID    int64
	catCreateDescription string
	catCreateEnabled     bool
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new category",
	RunE: func(cmd *cobra.Command, _ []string) error {
		if catCreateName == "" {
			return fmt.Errorf("--name is required")
		}
		cat := &apicategories.Category{
			Name:        catCreateName,
			ParentID:    catCreateParentID,
			Description: catCreateDescription,
			Enabled:     &catCreateEnabled,
		}
		result, err := cmdutil.AppClient.Categories.Create(cmd.Context(), cat)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var (
	catUpdateName        string
	catUpdateParentID    int64
	catUpdateDescription string
	catUpdateEnabled     bool
)

var updateCmd = &cobra.Command{
	Use:   "update <categoryId>",
	Short: "Update an existing category",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid category ID %q: %w", args[0], err)
		}
		if id <= 0 {
			return fmt.Errorf("category ID must be a positive integer, got %d", id)
		}
		cat := &apicategories.Category{}
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
		if cmd.Flags().Changed("parent-id") && catUpdateParentID == 0 {
			return fmt.Errorf("--parent-id 0 is not supported: parentId=0 cannot be expressed in the request payload")
		}
		if !cmd.Flags().Changed("name") && !cmd.Flags().Changed("parent-id") && !cmd.Flags().Changed("description") && !cmd.Flags().Changed("enabled") {
			return fmt.Errorf("no fields specified to update")
		}
		result, err := cmdutil.AppClient.Categories.Update(cmd.Context(), id, cat)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete <categoryId>",
	Short: "Delete a category by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid category ID %q: %w", args[0], err)
		}
		if id <= 0 {
			return fmt.Errorf("category ID must be a positive integer, got %d", id)
		}
		result, err := cmdutil.AppClient.Categories.Delete(cmd.Context(), id)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

func init() {
	// list flags
	listCmd.Flags().StringVar(&catListKeyword, "keyword", "", "filter by keyword")
	listCmd.Flags().Int64Var(&catListParent, "parent", 0, "filter by parent category ID")
	listCmd.Flags().IntVar(&catListLimit, "limit", 0, "maximum number of results")
	listCmd.Flags().IntVar(&catListOffset, "offset", 0, "result offset for pagination")

	// create flags
	createCmd.Flags().StringVar(&catCreateName, "name", "", "category name (required)")
	_ = createCmd.MarkFlagRequired("name")
	createCmd.Flags().Int64Var(&catCreateParentID, "parent-id", 0, "parent category ID")
	createCmd.Flags().StringVar(&catCreateDescription, "description", "", "category description")
	createCmd.Flags().BoolVar(&catCreateEnabled, "enabled", true, "whether the category is enabled")

	// update flags
	updateCmd.Flags().StringVar(&catUpdateName, "name", "", "category name")
	updateCmd.Flags().Int64Var(&catUpdateParentID, "parent-id", 0, "parent category ID")
	updateCmd.Flags().StringVar(&catUpdateDescription, "description", "", "category description")
	updateCmd.Flags().BoolVar(&catUpdateEnabled, "enabled", false, "whether the category is enabled")

	Cmd.AddCommand(listCmd, getCmd, createCmd, updateCmd, deleteCmd)
}
