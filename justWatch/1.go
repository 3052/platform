package justWatch

import (
   "cmp"
   "maps"
   "slices"
)

// Struct definitions
type Locale struct {
   FullLocale  string
   Country     string
   CountryName string
}

type StandardWebUrl string

type EnrichedOffer struct {
   Offer  Offer
   Locale *Locale
}

// Deduplicate removes true duplicates where both the Offer and Locale are identical.
func Deduplicate(offers []EnrichedOffer) []EnrichedOffer {
   // 1. Sort the slice. This brings identical EnrichedOffers next to each other.
   slices.SortFunc(offers, func(a, b EnrichedOffer) int {
      if n := cmp.Compare(a.Offer.StandardWebUrl, b.Offer.StandardWebUrl); n != 0 {
         return n
      }
      if n := cmp.Compare(a.Offer.MonetizationType, b.Offer.MonetizationType); n != 0 {
         return n
      }
      if n := cmp.Compare(a.Offer.ElementCount, b.Offer.ElementCount); n != 0 {
         return n
      }
      return cmp.Compare(a.Locale.FullLocale, b.Locale.FullLocale)
   })

   // 2. Compact the sorted slice, removing consecutive duplicates.
   return slices.CompactFunc(offers, func(a, b EnrichedOffer) bool {
      return a.Offer == b.Offer && a.Locale == b.Locale
   })
}

// FilterOffers removes offers with unwanted monetization types.
func FilterOffers(offers []EnrichedOffer, unwantedTypes ...string) []EnrichedOffer {
   unwantedSet := make(map[string]struct{}, len(unwantedTypes))
   for _, t := range unwantedTypes {
      unwantedSet[t] = struct{}{}
   }
   var filteredOffers []EnrichedOffer
   for _, offer := range offers {
      if _, found := unwantedSet[offer.Offer.MonetizationType]; !found {
         filteredOffers = append(filteredOffers, offer)
      }
   }
   return filteredOffers
}

// GroupAndSortByURL groups offers by URL and sorts the groups.
// Within each group, offers are sorted by country.
func GroupAndSortByURL(offers []EnrichedOffer) ([]StandardWebUrl, map[StandardWebUrl][]EnrichedOffer) {
   groupedOffers := make(map[StandardWebUrl][]EnrichedOffer)
   for _, offer := range offers {
      key := offer.Offer.StandardWebUrl
      groupedOffers[key] = append(groupedOffers[key], offer)
   }

   for _, offerGroup := range groupedOffers {
      slices.SortFunc(offerGroup, func(a, b EnrichedOffer) int {
         return cmp.Compare(a.Locale.Country, b.Locale.Country)
      })
   }

   return slices.Sorted(maps.Keys(groupedOffers)), groupedOffers
}
