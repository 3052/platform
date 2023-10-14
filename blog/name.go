package stream

import "time"

type episode interface {
   series() string
   season() (int64, error)
   episode() (int64, error)
   title() string
}

type film interface {
   title() string
   date() (time.Time, error)
}

func format_episode(episode) (string, error) {
   return "", nil
}

func format_film(film) (string, error) {
   return "", nil
}

type episode_film interface {
   episode
   film
}

func Format_Episode_Film(e episode_film) (string, error) {
   if true {
      return format_film(e)
   }
   return format_episode(e)
}
