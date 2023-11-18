package dash

type ContentProtection struct {
   PSSH string `xml:"pssh"`
}

type Representation struct {
   s struct {
      ContentProtection []ContentProtection
   }
}

func (r Representation) ContentProtection() []ContentProtection {
   return r.s.ContentProtection
}
