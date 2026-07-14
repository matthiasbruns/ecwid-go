// Package instantsite implements the "instantsite" CLI command group for the
// Ecwid Instant Site API.
package instantsite

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	api "github.com/matthiasbruns/ecwid-go/ecwid/instantsite"
)

// Cmd is the top-level instantsite command.
var Cmd = &cobra.Command{
	Use:   "instantsite",
	Short: "Manage the store's Instant Site (pages, tiles, themes, redirects)",
}

// decodeInput reads JSON from --file or stdin and unmarshals it into v.
func decodeInput(cmd *cobra.Command, v any, what string) error {
	data, err := cmdutil.ReadInput(cmd)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("invalid %s JSON: %w", what, err)
	}
	return nil
}

// fileFlag adds the standard --file flag to a body-reading command.
func fileFlag(c *cobra.Command) *cobra.Command {
	c.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
	return c
}

func init() {
	Cmd.AddCommand(
		tokenCmd(),
		profileCmd(),
		lifecycleCmds(),
		pagesCmd(),
		tilesCmd(),
		imagesCmd(),
		themesCmd(),
		textLabelsCmd(),
		redirectsCmd(),
	)
}

// ── Token ─────────────────────────────────────────────────────────────────

func tokenCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "token",
		Short: "Exchange an app secret token for an Instant Site access token",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			siteID, _ := cmd.Flags().GetString("site-id")
			code, _ := cmd.Flags().GetString("code")
			grantType, _ := cmd.Flags().GetString("grant-type")
			result, err := cmdutil.AppClient.InstantSite.Token(cmd.Context(), &api.TokenRequest{
				GrantType: grantType,
				SiteID:    siteID,
				Code:      code,
			})
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}
	c.Flags().String("site-id", "", "store ID (site_id) to exchange the token for")
	c.Flags().String("code", "", "app secret token with the manage_instant_site scope")
	c.Flags().String("grant-type", api.GrantTypeAuthorizationCode, "grant type: authorization_code or draft_code")
	_ = c.MarkFlagRequired("site-id")
	_ = c.MarkFlagRequired("code")
	return c
}

// ── Profile ───────────────────────────────────────────────────────────────

func profileCmd() *cobra.Command {
	c := &cobra.Command{Use: "profile", Short: "Manage the Instant Site profile"}

	get := &cobra.Command{
		Use:   "get",
		Short: "Get the Instant Site profile",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			result, err := cmdutil.AppClient.InstantSite.GetProfile(cmd.Context())
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}

	create := fileFlag(&cobra.Command{
		Use:   "create",
		Short: "Create the Instant Site profile (reads JSON from stdin or --file)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			var profile api.Profile
			if err := decodeInput(cmd, &profile, "profile"); err != nil {
				return err
			}
			result, err := cmdutil.AppClient.InstantSite.CreateProfile(cmd.Context(), &profile)
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	})

	update := fileFlag(&cobra.Command{
		Use:   "update",
		Short: "Update the Instant Site profile (reads JSON from stdin or --file)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			var profile api.Profile
			if err := decodeInput(cmd, &profile, "profile"); err != nil {
				return err
			}
			result, err := cmdutil.AppClient.InstantSite.UpdateProfile(cmd.Context(), &profile)
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	})

	c.AddCommand(get, create, update)
	return c
}

// ── Lifecycle ─────────────────────────────────────────────────────────────

func lifecycleCmds() *cobra.Command {
	c := &cobra.Command{Use: "draft", Short: "Publish, discard, or clone Instant Site draft changes"}

	publish := &cobra.Command{
		Use:   "publish",
		Short: "Publish the current Instant Site draft",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := cmdutil.AppClient.InstantSite.Publish(cmd.Context()); err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, map[string]any{"published": true})
		},
	}

	discard := &cobra.Command{
		Use:   "discard",
		Short: "Discard unpublished Instant Site draft changes",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := cmdutil.AppClient.InstantSite.Discard(cmd.Context()); err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, map[string]any{"discarded": true})
		},
	}

	clone := fileFlag(&cobra.Command{
		Use:   "clone",
		Short: "Clone an Instant Site from another store (reads JSON from stdin or --file)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			var req api.CloneRequest
			if err := decodeInput(cmd, &req, "clone request"); err != nil {
				return err
			}
			if err := cmdutil.AppClient.InstantSite.Clone(cmd.Context(), &req); err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, map[string]any{"cloned": true})
		},
	})

	c.AddCommand(publish, discard, clone)
	return c
}

// ── Pages ─────────────────────────────────────────────────────────────────

func pagesCmd() *cobra.Command {
	c := &cobra.Command{Use: "pages", Short: "Manage Instant Site pages"}

	list := &cobra.Command{
		Use:   "list",
		Short: "List Instant Site pages",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			published, _ := cmd.Flags().GetBool("published")
			result, err := cmdutil.AppClient.InstantSite.ListPages(cmd.Context(), published)
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}
	list.Flags().Bool("published", false, "list published pages (default: draft)")

	create := fileFlag(&cobra.Command{
		Use:   "create",
		Short: "Create a page (reads JSON from stdin or --file)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			var page api.Page
			if err := decodeInput(cmd, &page, "page"); err != nil {
				return err
			}
			result, err := cmdutil.AppClient.InstantSite.CreatePage(cmd.Context(), &page)
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	})

	update := fileFlag(&cobra.Command{
		Use:   "update <pageId>",
		Short: "Update a page (reads JSON from stdin or --file)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var page api.Page
			if err := decodeInput(cmd, &page, "page"); err != nil {
				return err
			}
			result, err := cmdutil.AppClient.InstantSite.UpdatePage(cmd.Context(), args[0], &page)
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	})

	del := &cobra.Command{
		Use:   "delete <pageId>",
		Short: "Delete a page",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := cmdutil.AppClient.InstantSite.DeletePage(cmd.Context(), args[0])
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}

	c.AddCommand(list, create, update, del)
	return c
}

// ── Tiles ─────────────────────────────────────────────────────────────────

func tilesCmd() *cobra.Command {
	c := &cobra.Command{Use: "tiles", Short: "Manage Instant Site tiles"}

	list := &cobra.Command{
		Use:   "list",
		Short: "List tiles",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			published, _ := cmd.Flags().GetBool("published")
			pageID, _ := cmd.Flags().GetString("page-id")
			lang, _ := cmd.Flags().GetString("lang")
			result, err := cmdutil.AppClient.InstantSite.ListTiles(cmd.Context(), &api.TileListOptions{
				Published: published,
				PageID:    pageID,
				Lang:      lang,
			})
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}
	list.Flags().Bool("published", false, "list published tiles (default: draft)")
	list.Flags().String("page-id", "", "filter tiles to a single page")
	list.Flags().String("lang", "", "2-letter language code for translated text")

	get := &cobra.Command{
		Use:   "get <tileId>",
		Short: "Get a single tile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := cmdutil.AppClient.InstantSite.GetTile(cmd.Context(), args[0])
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}

	create := fileFlag(&cobra.Command{
		Use:   "create",
		Short: "Create a tile (reads JSON from stdin or --file)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			var req api.CreateTileRequest
			if err := decodeInput(cmd, &req, "tile"); err != nil {
				return err
			}
			result, err := cmdutil.AppClient.InstantSite.CreateTile(cmd.Context(), &req)
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	})

	update := fileFlag(&cobra.Command{
		Use:   "update <tileId>",
		Short: "Update a tile (reads JSON from stdin or --file)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var tile api.TileUpdate
			if err := decodeInput(cmd, &tile, "tile"); err != nil {
				return err
			}
			result, err := cmdutil.AppClient.InstantSite.UpdateTile(cmd.Context(), args[0], &tile)
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	})

	updateAll := fileFlag(&cobra.Command{
		Use:   "update-all",
		Short: "Update the whole tiles list (reads JSON from stdin or --file)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			var tiles api.TileBulkUpdate
			if err := decodeInput(cmd, &tiles, "tiles"); err != nil {
				return err
			}
			result, err := cmdutil.AppClient.InstantSite.UpdateTiles(cmd.Context(), &tiles)
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	})

	del := &cobra.Command{
		Use:   "delete <tileId>",
		Short: "Delete a tile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := cmdutil.AppClient.InstantSite.DeleteTile(cmd.Context(), args[0])
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}

	showcase := &cobra.Command{
		Use:   "showcase",
		Short: "List available tile templates",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			result, err := cmdutil.AppClient.InstantSite.TileShowcase(cmd.Context())
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}

	config := &cobra.Command{
		Use:   "config <configType>",
		Short: "Get the editor config for a tile category type",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := cmdutil.AppClient.InstantSite.TileConfig(cmd.Context(), args[0])
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}

	c.AddCommand(list, get, create, update, updateAll, del, showcase, config)
	return c
}

// ── Tile images ───────────────────────────────────────────────────────────

func imagesCmd() *cobra.Command {
	c := &cobra.Command{Use: "images", Short: "Manage Instant Site tile images"}

	reserve := &cobra.Command{
		Use:   "reserve <tileId>",
		Short: "Reserve an upload slot for a tile image",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := cmdutil.AppClient.InstantSite.ReserveTileImage(cmd.Context(), args[0])
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}

	get := &cobra.Command{
		Use:   "get <imageId>",
		Short: "Get an uploaded tile image",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := cmdutil.AppClient.InstantSite.GetImage(cmd.Context(), args[0])
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}

	buckets := &cobra.Command{
		Use:   "buckets",
		Short: "List image bucket CDN URLs",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			result, err := cmdutil.AppClient.InstantSite.ImageBuckets(cmd.Context())
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}

	c.AddCommand(reserve, get, buckets)
	return c
}

// ── Themes ────────────────────────────────────────────────────────────────

func themesCmd() *cobra.Command {
	c := &cobra.Command{Use: "themes", Short: "Manage Instant Site themes"}

	list := &cobra.Command{
		Use:   "list",
		Short: "List saved themes",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			result, err := cmdutil.AppClient.InstantSite.ListThemes(cmd.Context())
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}

	create := fileFlag(&cobra.Command{
		Use:   "create",
		Short: "Create a theme from a color palette (reads JSON from stdin or --file)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			var colors api.ThemeColors
			if err := decodeInput(cmd, &colors, "theme colors"); err != nil {
				return err
			}
			result, err := cmdutil.AppClient.InstantSite.CreateTheme(cmd.Context(), &colors)
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	})

	update := fileFlag(&cobra.Command{
		Use:   "update <themeId>",
		Short: "Update a theme (reads JSON from stdin or --file)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var colors api.ThemeColors
			if err := decodeInput(cmd, &colors, "theme colors"); err != nil {
				return err
			}
			result, err := cmdutil.AppClient.InstantSite.UpdateTheme(cmd.Context(), args[0], &colors)
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	})

	del := &cobra.Command{
		Use:   "delete <themeId>",
		Short: "Delete a theme",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := cmdutil.AppClient.InstantSite.DeleteTheme(cmd.Context(), args[0])
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}

	current := &cobra.Command{
		Use:   "current",
		Short: "Get the currently applied theme palette",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			result, err := cmdutil.AppClient.InstantSite.CurrentTheme(cmd.Context())
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}

	setCurrent := fileFlag(&cobra.Command{
		Use:   "set-current",
		Short: "Set the currently applied theme palette (reads JSON from stdin or --file)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			var colors api.ThemeColors
			if err := decodeInput(cmd, &colors, "theme colors"); err != nil {
				return err
			}
			result, err := cmdutil.AppClient.InstantSite.UpdateCurrentTheme(cmd.Context(), &colors)
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	})

	c.AddCommand(list, create, update, del, current, setCurrent)
	return c
}

// ── Text labels ───────────────────────────────────────────────────────────

func textLabelsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "text-labels",
		Short: "Get Instant Site text labels (i18n translations)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			result, err := cmdutil.AppClient.InstantSite.TextLabels(cmd.Context())
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}
}

// ── Redirects ─────────────────────────────────────────────────────────────

func redirectsCmd() *cobra.Command {
	c := &cobra.Command{Use: "redirects", Short: "Manage 301 URL redirects"}

	search := &cobra.Command{
		Use:   "search",
		Short: "Search redirects",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			keyword, _ := cmd.Flags().GetString("keyword")
			offset, err := cmdutil.GetNonNegativeInt(cmd, "offset")
			if err != nil {
				return err
			}
			limit, err := cmdutil.GetNonNegativeInt(cmd, "limit")
			if err != nil {
				return err
			}
			result, err := cmdutil.AppClient.InstantSite.SearchRedirects(cmd.Context(), &api.RedirectSearchOptions{
				Keyword: keyword,
				Offset:  offset,
				Limit:   limit,
			})
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}
	search.Flags().String("keyword", "", "filter by saved redirect URL")
	search.Flags().Int("offset", 0, "pagination offset")
	search.Flags().Int("limit", 0, "pagination limit")

	get := &cobra.Command{
		Use:   "get <redirectId>",
		Short: "Get a single redirect",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := cmdutil.AppClient.InstantSite.GetRedirect(cmd.Context(), args[0])
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	}

	create := fileFlag(&cobra.Command{
		Use:   "create",
		Short: "Create a redirect (reads JSON from stdin or --file)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			var redirect api.Redirect
			if err := decodeInput(cmd, &redirect, "redirect"); err != nil {
				return err
			}
			result, err := cmdutil.AppClient.InstantSite.CreateRedirect(cmd.Context(), &redirect)
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	})

	update := fileFlag(&cobra.Command{
		Use:   "update <redirectId>",
		Short: "Update a redirect (reads JSON from stdin or --file)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var redirect api.Redirect
			if err := decodeInput(cmd, &redirect, "redirect"); err != nil {
				return err
			}
			result, err := cmdutil.AppClient.InstantSite.UpdateRedirect(cmd.Context(), args[0], &redirect)
			if err != nil {
				return err
			}
			return cmdutil.OutputResult(cmd, result)
		},
	})

	c.AddCommand(search, get, create, update)
	return c
}
