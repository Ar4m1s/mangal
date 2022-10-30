package inline

import (
	"encoding/json"
	"github.com/metafates/mangal/anilist"
	"github.com/metafates/mangal/constant"
	"github.com/metafates/mangal/source"
	"github.com/spf13/viper"
)

type Manga struct {
	// Source that the manga belongs to.
	Source string `json:"source"`
	// Mangal variant of the manga
	Mangal *source.Manga `json:"mangal"`
	// Anilist is the closest anilist match to mangal manga
	Anilist *anilist.Manga `json:"anilist"`
}

type Output struct {
	Query  string   `json:"query"`
	Result []*Manga `json:"result"`
}

func asJson(manga []*source.Manga, options *Options) (marshalled []byte, err error) {
	var m = make([]*Manga, len(manga))
	for i, manga := range manga {
		al := manga.Anilist.OrElse(nil)
		if !options.IncludeAnilistManga {
			al = nil
		}

		m[i] = &Manga{
			Mangal:  manga,
			Anilist: al,
			Source:  manga.Source.Name(),
		}
	}

	return json.Marshal(&Output{
		Result: m,
		Query:  options.Query,
	})
}

func prepareManga(manga *source.Manga, options *Options) error {
	var err error

	if options.IncludeAnilistManga {
		err = manga.BindWithAnilist()
		if err != nil {
			return err
		}
	}

	if options.ChaptersFilter.IsPresent() {
		chapters, err := manga.Source.ChaptersOf(manga)
		if err != nil {
			return err
		}

		chapters, err = options.ChaptersFilter.MustGet()(chapters)
		if err != nil {
			return err
		}

		manga.Chapters = chapters

		if options.PopulatePages {
			for _, chapter := range chapters {
				_, err := chapter.Source().PagesOf(chapter)
				if err != nil {
					return err
				}
			}
		}
	} else {
		// clear chapters in case they were loaded from cache or something
		manga.Chapters = make([]*source.Chapter, 0)
	}

	if viper.GetBool(constant.MetadataFetchAnilist) {
		_ = manga.PopulateMetadata(func(string) {})
	}

	return nil
}
