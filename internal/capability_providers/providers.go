// Package providers holds example capability providers — an ADVANCED, opt-in
// feature most workers do not need.
//
// A capability provider replaces the platform's BUILT-IN implementation behind
// one agent capability seat (today: "tool_search" — the relevance scorer, and
// "memory" — history selection). Agents get built-in tool search and memory
// without any provider; a workflow only uses yours when it explicitly binds it
// on an agent node:
//
//	capability_providers:
//	  memory: starter-recency-memory     # seat -> provider name
//	  # NOTE: no history_policy line — binding a memory provider implies "custom"
//
// Requires sdk-go >= v0.0.20. To try these, uncomment the registration lines
// in cmd/worker/main.go. Both examples are deliberately "marker-style": their
// effect is unmistakable in a run trace, so you can prove YOUR code ran
// instead of the built-in before writing a real implementation.
//
// Data boundary: the memory seat receives the thread's FULL conversation
// history on every bound agent turn. Treat a memory provider's worker as
// holding user conversation data.
package capabilityproviders

import (
	"log"
	"sort"
	"strings"

	sdk "github.com/dibbla-agents/sdk-go"
)

// MemoryMarker makes the memory provider's effect visible in a run trace:
// if this text shows up in the injected history, your provider ran.
const MemoryMarker = "[starter-memory-provider] custom history injected"

// RecencyMemoryProvider keeps only the most recent stored turn and prepends a
// marker turn. Not a useful strategy — a visibility harness. A real provider
// would summarize, embed-rank, or budget-pack the turns it returns.
func RecencyMemoryProvider() sdk.MemoryProvider {
	return sdk.MemoryProvider{
		Name:        "starter-recency-memory",
		Description: "Example memory provider: injects a marker turn + the last stored turn.",
		Version:     "1.0.0",
		// Token ceiling as a fraction of the model context (0 = platform default).
		MaxHistoryFraction: 0.5,
		Transform: func(currentMessage string, turns []sdk.Turn, tokenBudget int, meta sdk.ThreadMeta) ([]sdk.Turn, error) {
			// Log what the engine sends so you can see the incoming shape
			// before trusting your transform. This call also appears in the
			// run-log dock as a capability_provider_call entry.
			log.Printf("memory provider: thread=%q turns_in=%d budget=%d", meta.ThreadID, len(turns), tokenBudget)

			marker := sdk.Turn{
				Role: "assistant",
				Parts: []sdk.Part{{
					Type: sdk.PartTypeText,
					Text: &sdk.TextPart{Text: MemoryMarker},
				}},
			}
			out := []sdk.Turn{marker}
			if len(turns) > 0 {
				out = append(out, turns[len(turns)-1]) // keep the last stored turn verbatim
			}
			return out, nil
		},
	}
}

// PrefixToolSearchProvider ranks the offered tools with a deterministic,
// obviously-not-the-built-in ordering: query-prefix matches first, then
// everything else, each group reverse-alphabetical, truncated to topN.
func PrefixToolSearchProvider() sdk.ToolSearchProvider {
	return sdk.ToolSearchProvider{
		Name:        "starter-prefix-search",
		Description: "Example tool_search provider: prefix matches first, then reverse-alphabetical.",
		Version:     "1.0.0",
		Select: func(query string, stubs []sdk.ProviderStub, topN int) ([]string, error) {
			log.Printf("tool_search provider: query=%q offered=%d topN=%d", query, len(stubs), topN)

			q := strings.ToLower(strings.TrimSpace(query))
			var prefix, rest []string
			for _, s := range stubs {
				if q != "" && strings.HasPrefix(strings.ToLower(s.Name), q) {
					prefix = append(prefix, s.Name)
				} else {
					rest = append(rest, s.Name)
				}
			}
			sort.Sort(sort.Reverse(sort.StringSlice(prefix)))
			sort.Sort(sort.Reverse(sort.StringSlice(rest)))
			selected := append(prefix, rest...)
			if topN > 0 && len(selected) > topN {
				selected = selected[:topN]
			}
			return selected, nil
		},
	}
}

// Register registers both example providers.
func Register(server *sdk.Server) error {
	if err := server.RegisterCapabilityProvider(RecencyMemoryProvider()); err != nil {
		return err
	}
	return server.RegisterCapabilityProvider(PrefixToolSearchProvider())
}

// MustRegister is the one-line enable used from cmd/worker/main.go — leave the
// call commented unless you are actually building a custom capability provider.
func MustRegister(server *sdk.Server) {
	if err := Register(server); err != nil {
		log.Fatalf("failed to register capability providers: %v", err)
	}
}
