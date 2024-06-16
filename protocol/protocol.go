package protocol

import(
	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

func CheckSimilarity(s string, c string) float64 {
	similarity := strutil.Similarity(s, c, metrics.NewJaroWinkler())
	return similarity
}