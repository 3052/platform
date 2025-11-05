package justwatch

import (
   "cmp"
   "maps"
   "slices"
)

// keep order
type Locale struct {
   FullLocale  string
   Country     string
   CountryName string
}

type StandardWebUrl string

// `presentationType` data seems to be incorrect in some cases. For example,
// JustWatch reports this as SD: fetchtv.com.au/movie/details/19285
// when the site itself reports as HD
type Offer struct {
   ElementCount     int
   MonetizationType string
   StandardWebUrl   StandardWebUrl
}

type EnrichedOffer struct {
   Offer  Offer
   Locale *Locale
}

func UniqueEnrichedOffers(rawOffersByLocale map[*Locale][]Offer) []EnrichedOffer {
   // Use a map to deduplicate OFFERS, keeping the first LOCALE seen for each.
   uniqueOffers := make(map[Offer]*Locale)
   // In a real app, you would iterate through your API calls.
   // Here, we iterate through the simulated map of results.
   for locale, offers := range rawOffersByLocale {
      for _, offer := range offers {
         // If the offer is NOT in the map, add it with its locale.
         // This performs the deduplication on the Offer struct.
         if _, found := uniqueOffers[offer]; !found {
            uniqueOffers[offer] = locale
         }
      }
   }
   // Now, build the final enriched slice from the unique offers.
   enrichedOffers := make([]EnrichedOffer, 0, len(uniqueOffers))
   for offer, locale := range uniqueOffers {
      enrichedOffers = append(enrichedOffers, EnrichedOffer{
         Offer:  offer,
         Locale: locale,
      })
   }
   return enrichedOffers
}

// FilterOffers function
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

// GroupAndSortByURL function
func GroupAndSortByURL(offers []EnrichedOffer) ([]StandardWebUrl, map[StandardWebUrl][]EnrichedOffer) {
   groupedOffers := make(map[StandardWebUrl][]EnrichedOffer)
   for _, offer := range offers {
      key := offer.Offer.StandardWebUrl
      groupedOffers[key] = append(groupedOffers[key], offer)
   }
   // Secondary sort within each group by Country is necessary.
   for _, offerGroup := range groupedOffers {
      slices.SortFunc(offerGroup, func(a, b EnrichedOffer) int {
         return cmp.Compare(a.Locale.Country, b.Locale.Country)
      })
   }
   return slices.Sorted(maps.Keys(groupedOffers)), groupedOffers
}
