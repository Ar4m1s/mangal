package downloader

import (
	"fmt"
	"github.com/metafates/mangal/constant"
	"github.com/metafates/mangal/converter"
	"github.com/metafates/mangal/history"
	"github.com/metafates/mangal/log"
	"github.com/metafates/mangal/open"
	"github.com/metafates/mangal/source"
	"github.com/metafates/mangal/style"
	"github.com/spf13/viper"
)

// Read the chapter by downloading it with the given source
// and opening it with the configured reader.
func Read(chapter *source.Chapter, progress func(string)) error {
	if viper.GetBool(constant.ReaderReadInBrowser) {
		return open.StartWith(
			chapter.URL,
			viper.GetString(constant.ReaderBrowser),
		)
	}

	if viper.GetBool(constant.DownloaderReadDownloaded) && chapter.IsDownloaded() {
		path, err := chapter.Path(false)
		if err == nil {
			return openRead(path, progress)
		}
	}

	log.Info(fmt.Sprintf("downloading %s for reading. Provider is %s", chapter.Name, chapter.Source().ID()))
	log.Info("getting pages of " + chapter.Name)
	progress("Getting pages")
	pages, err := chapter.Source().PagesOf(chapter)
	if err != nil {
		log.Error(err)
		return err
	}

	err = chapter.DownloadPages(true, progress)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("getting " + viper.GetString(constant.FormatsUse) + " converter")
	conv, err := converter.Get(viper.GetString(constant.FormatsUse))
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("converting " + viper.GetString(constant.FormatsUse))
	progress(fmt.Sprintf(
		"Converting %d pages to %s %s",
		len(pages),
		style.Yellow(viper.GetString(constant.FormatsUse)),
		style.Faint(chapter.SizeHuman())),
	)
	path, err := conv.SaveTemp(chapter)
	if err != nil {
		log.Error(err)
		return err
	}

	err = openRead(path, progress)
	if err != nil {
		log.Error(err)
		return err
	}

	if viper.GetBool(constant.HistorySaveOnRead) {
		go func() {
			err := history.Save(chapter)
			if err != nil {
				log.Warn(err)
			} else {
				log.Info("history saved")
			}
		}()
	}

	progress("Done")
	return nil
}

func openRead(path string, progress func(string)) error {
	var (
		reader string
		err    error
	)

	switch viper.GetString(constant.FormatsUse) {
	case constant.PDF:
		reader = viper.GetString(constant.ReaderPDF)
	case constant.CBZ:
		reader = viper.GetString(constant.ReaderCBZ)
	case constant.ZIP:
		reader = viper.GetString(constant.ReaderZIP)
	case constant.Plain:
		reader = viper.GetString(constant.RaderPlain)
	}

	if reader != "" {
		log.Info("opening with " + reader)
		progress(fmt.Sprintf("Opening %s", reader))
	} else {
		log.Info("no reader specified. opening with default")
		progress("Opening")
	}

	err = open.RunWith(path, reader)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("could not open %s with %s: %s", path, reader, err.Error())
	}

	log.Info("opened without errors")

	return nil
}
