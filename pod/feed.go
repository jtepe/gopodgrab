package pod

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"
)

type podTime time.Time

func (p *podTime) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	tok, err := dec.Token()
	if err != nil {
		return err
	}
	data := string(tok.(xml.CharData))
	t, err := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", data)
	if err != nil {
		return err
	}
	*p = podTime(t)
	err = dec.Skip()
	return err
}

type podFile struct {
	URL  string
	Size int64
	Enc  string
}

func (f *podFile) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		switch a.Name.Local {
		case "url":
			f.URL = a.Value
		case "type":
			f.Enc = a.Value
		case "length":
			size, err := strconv.ParseInt(a.Value, 10, 64)
			if err != nil {
				return err
			}
			f.Size = size
		}
	}
	err := dec.Skip()
	return err
}

func (f *podFile) String() string {
	return fmt.Sprintf("URL: %s\nSize (bytes): %d\nEncoding: %s", f.URL, f.Size, f.Enc)
}

type Episode struct {
	XMLName  xml.Name `xml:"item"`
	Title    string   `xml:"title"`
	PubDate  *podTime `xml:"pubDate"`
	File     *podFile `xml:"enclosure"`
	Duration int      `xml:"duration"`
	Bytes    int64    `xml:"-"`
}

func (e *Episode) String() string {
	return fmt.Sprintf("Title: %s\nPubDate: %s\nURL: %v\nDuration: %v",
		e.Title, (time.Time(*e.PubDate)).Format("Mon, 2 Jan 2006 15:04:05 -0700"),
		e.File,
		time.Duration(e.Duration)*time.Second)
}

func ParseFeed(r io.Reader) ([]*Episode, error) {
	dec := xml.NewDecoder(r)
	var episodes []*Episode

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		switch el := tok.(type) {
		case xml.StartElement:
			if el.Name.Local != "item" {
				continue
			}

			epi := new(Episode)
			err := dec.DecodeElement(epi, &el)
			if err != nil {
				log.Printf("failed to parse episode from feed: %v", err)
			}
			episodes = append(episodes, epi)
		}
	}

	return episodes, nil
}

var data = `
<root>
<item>
            <title>Deus Ex! (Stay Forever, Folge 2)</title>
            <itunes:title>Deus Ex! (Stay Forever, Folge 2)</itunes:title>
            <description>Das Gespräch geht um Warren Spector, die Deus Ex-Serie und die Firma Ion Storm. Jedenfalls am
                Anfang und am Ende, zwischendurch schweift es ein bisschen ab. Ahem.
            </description>
            <pubDate>Fri, 09 Sep 2011 22:33:27 +0000</pubDate>
            <link>https://podcastd45a61.podigee.io/2-deus-ex-stay-forever-folge-2</link>
            <guid isPermaLink="false">139893a4ae83fa923a28aa0466b40a1b</guid>
            <content:encoded>
                <![CDATA[
<p>Wir erwähnen, ausführlich oder nebenbei, folgende Spiele: die Deus Ex-Serie, Martian Dreams, Daikatana, Half-Life, Bioshock, Mass Effect, Dominion: Storm over Gift 3, Ultima Underworld, System Shock, Anachronox, Baphomets Fluch, Savage Empire, Thief: Deadly Shadows, Unreal Tournament.</p>]]>
            </content:encoded>
            <image>
                <url>
                    https://images.podigee.com/0x,sZP1NxnUSqBeeBEJGQ81kzeZgzQYxQsKMF3vWUk7yaW4=/https://cdn.podigee.com/uploads/u4307/986bdc78-0431-4b07-80f0-1c6c27039a67.jpg
                </url>
                <title>Deus Ex! (Stay Forever, Folge 2)</title>
                <link>https://podcastd45a61.podigee.io/2-deus-ex-stay-forever-folge-2</link>
            </image>
            <itunes:image
                    href="https://images.podigee.com/0x,sZP1NxnUSqBeeBEJGQ81kzeZgzQYxQsKMF3vWUk7yaW4=/https://cdn.podigee.com/uploads/u4307/986bdc78-0431-4b07-80f0-1c6c27039a67.jpg"/>
            <itunes:episode>2</itunes:episode>
            <itunes:episodeType>full</itunes:episodeType>
            <itunes:subtitle></itunes:subtitle>
            <itunes:summary>Das Gespräch geht um Warren Spector, die Deus Ex-Serie und die Firma Ion Storm. Jedenfalls
                am Anfang und am Ende, zwischendurch schweift es ein bisschen ab. Ahem.
            </itunes:summary>
            <itunes:explicit>no</itunes:explicit>
            <itunes:keywords>DEUS EX ION STORM JOHN ROMERO WARREN SPECTOR</itunes:keywords>
            <itunes:author>Gunnar Lott, Christian Schmidt, Fabian Käufer</itunes:author>
            <enclosure
                    url="https://cdn.podigee.com/uploads/u4307/0915efe6-c75e-4de4-8604-1ac4874b81b1.m4a?v=1541672487&amp;source=feed"
                    type="audio/aac" length="27191433"/>
            <itunes:duration>2602</itunes:duration>
        </item>
        <item>
            <title>Revolution! (Stay Forever, Folge 1 Classic)</title>
            <itunes:title>Revolution! (Stay Forever, Folge 1 Classic)</itunes:title>
            <description>Willkommen bei Stay Forever, dem brandneuen Podcast von Gunnar Lott und Christian Schmidt.

                Erstaunlicherweise ist uns, nach all den Jahren, auch aufgefallen, dass so ein Podcast möglicherweise
                ein sicherer Weg zu Ruhm und Reichtum ist. Also dachten wir, wir versuchen’s mal. Hier ist unser
                Erstling. Das Gespräch dreht sich um Revolution Software, Charles Cecil und die Baphomets Fluch-Spiele,
                mäandert aber zwischendurch auch mal in andere Bereiche.
            </description>
            <pubDate>Mon, 29 Aug 2011 19:51:48 +0000</pubDate>
            <link>https://www.stayforever.de/2011/08/folge-1-revolution/</link>
            <guid isPermaLink="false">e9c9efed20ed4580f88e54f955082dbb</guid>
            <content:encoded>
                <![CDATA[Wir sprechen über Baphomets Fluch und so
<p>Willkommen bei Stay Forever, dem brandneuen Podcast von Gunnar Lott und Christian Schmidt.</p>
<p>Erstaunlicherweise ist uns, nach all den Jahren, auch aufgefallen, dass so ein Podcast möglicherweise ein sicherer Weg zu Ruhm und Reichtum ist. Also dachten wir, wir versuchen’s mal. Hier ist unser Erstling. Das Gespräch dreht sich um Revolution Software, Charles Cecil und die Baphomets Fluch-Spiele, mäandert aber zwischendurch auch mal in andere Bereiche.</p>
<p>[Danksagungen] Wir bedanken uns bei Toni Schwaiger für die Überarbeitung und Nino Kerl für die Ansage. Das Logo des Podcasts und der Header dieser Seite stammen von Bastian Stock; die Musik am Anfang verwenden wir mit freundlicher Genehmigung von Chris Hülsbeck. Das Sprachsample „Stay awhile, stay forever“ kommt aus dem uralten Spiel Impossible Mission.</p>
<p>[Nachschlager-Service] Folgende Spiele werden en passant oder ausdrücklich erwähnt: Baphomets Fluch-Serie, Lure of the Temptress, Beneath a Steel Sky, The Colonel’s Bequest, Vollgas, Indiana Jones und der letzte Kreuzzug, Sword and Sworcery, Deadline, Livingstone, I presume?, One Single Life, Beyond Good and Evil, Monkey Island 2.</p>]]>
            </content:encoded>
            <image>
                <url>
                    https://images.podigee.com/0x,sEc01Y9mYLCSQcFSofL4vAm5W1z_a6-3beOsZH5TWEUg=/https://cdn.podigee.com/uploads/u4307/57e1e434-3afb-4747-83dc-b2f18abdb69e.jpg
                </url>
                <title>Revolution! (Stay Forever, Folge 1 Classic)</title>
                <link>https://www.stayforever.de/2011/08/folge-1-revolution/</link>
            </image>
            <itunes:image
                    href="https://images.podigee.com/0x,sEc01Y9mYLCSQcFSofL4vAm5W1z_a6-3beOsZH5TWEUg=/https://cdn.podigee.com/uploads/u4307/57e1e434-3afb-4747-83dc-b2f18abdb69e.jpg"/>
            <itunes:episode>1</itunes:episode>
            <itunes:episodeType>full</itunes:episodeType>
            <itunes:subtitle>Wir sprechen über Baphomets Fluch und so</itunes:subtitle>
            <itunes:summary>Willkommen bei Stay Forever, dem brandneuen Podcast von Gunnar Lott und Christian Schmidt.

                Erstaunlicherweise ist uns, nach all den Jahren, auch aufgefallen, dass so ein Podcast möglicherweise
                ein sicherer Weg zu Ruhm und Reichtum ist. Also dachten wir, wir versuchen’s mal. Hier ist unser
                Erstling. Das Gespräch dreht sich um Revolution Software, Charles Cecil und die Baphomets Fluch-Spiele,
                mäandert aber zwischendurch auch mal in andere Bereiche.
            </itunes:summary>
            <itunes:explicit>no</itunes:explicit>
            <itunes:keywords>Charles Cecil,Revolution Software</itunes:keywords>
            <itunes:author>Gunnar Lott, Christian Schmidt, Fabian Käufer</itunes:author>
            <enclosure
                    url="https://cdn.podigee.com/uploads/u4307/c016793b-ccdc-46bc-9808-461818a9a568.m4a?v=1541672244&amp;source=feed"
                    type="audio/aac" length="26537979"/>
            <itunes:duration>2561</itunes:duration>
        </item>
<item>
            <title>Baphomets Fluch (Stay Forever, Folge 1 REMAKE)</title>
            <itunes:title>Baphomets Fluch (Stay Forever, Folge 1 REMAKE) Itunes-Name</itunes:title>
            <description>1996 kam, schon relativ spät in der Hochzeit der Adventures, das aufwändig gestaltete Spiel
                Baphomets Fluch (Broken Sword) auf den Markt. Entwickler war das kleine englische Studio Revolution
                Software unter Charles Cecil. Das Spiel fand eine enorm große Fangemeinde und legte den Grundstein für
                eine der langlebigsten Adventure-Serien. Christian und Gunnar haben die allererste Folge von Stay
                Forever nochmal ganz neu aufgenommen.

                Infos zum Spiel:

                Thema: Baphomets Fluch
                Erscheinungsjahr: 1996
                Genre: Adventure
                Plattform: MS-DOS, Windows (später jede Plattform von GBA bis iOS)
                Entwickler: Revolution Software
                Publisher: Virgin Interactive Ltd.
                Designer: Charles Cecil, Steve Ince

                Podcast-Credits:

                Sprecher: Christian Schmidt, Gunnar Lott
                Audioproduktion: Fabian Langer, Christian Schmidt
                Titelgrafik: Paul Schmidt
                Intro, Outro: Nino Kerl (Ansage); Chris Hülsbeck (Musik)
                Chronist: Herr Anym
                Community Management: Christian Beuster
            </description>
            <pubDate>Sun, 28 Aug 2011 10:20:00 +0000</pubDate>
            <link>https://www.stayforever.de/2019/02/baphomets-fluch-folge-1-remake/</link>
            <guid isPermaLink="false">b68f0199b5092974e8f663cbaedab0f5</guid>
            <content:encoded>
                <![CDATA[Neue Aufnahme der (legendären?) Folge 1
<p>1996 kam, schon relativ spät in der Hochzeit der Adventures, das aufwändig gestaltete Spiel Baphomets Fluch (Broken Sword) auf den Markt. Entwickler war das kleine englische Studio Revolution Software unter Charles Cecil. Das Spiel fand eine enorm große Fangemeinde und legte den Grundstein für eine der langlebigsten Adventure-Serien.</p>
<p><img src="https://www.stayforever.de/wp-content/uploads/2019/02/32751-circle-of-blood-dos-front-cover-233x300.jpg" alt="" width="117" height="150" class="alignright size-medium wp-image-4215" /></p>
<p>Thema: Baphomets Fluch<br>
Erscheinungsjahr: 1996<br>
Genre: Adventure<br>
Plattform: MS-DOS, Windows (später jede Plattform von GBA bis iOS)<br>
Entwickler: Revolution Software<br>
Publisher: Virgin Interactive Ltd.<br>
Designer: Charles Cecil, Steve Ince<br></p>
<p><strong>Podcast-Credits:</strong></p>
<p>Sprecher: Christian Schmidt, Gunnar Lott<br>
Audioproduktion: Fabian Langer, Christian Schmidt<br>
Titelgrafik: Paul Schmidt<br>
Intro, Outro: Nino Kerl (Ansage); Chris Hülsbeck (Musik)<br>
Chronist: Herr Anym<br>
Community Management: Christian Beuster<br></p>
<hr>
<p><strong>Hinweise:</strong> Bitte hier kommentieren oder auf <a href="https://www.reddit.com/r/stayforever/">Reddit</a>. Diese Folge gibt es auch auf <a href="https://www.youtube.com/playlist?list=PLk81FXjsqZvcFtNnm0lzwJ_9EoD1q1td5">Youtube</a>, <a href="https://open.spotify.com/show/0HrgGvhjzvg1Qd9yF0cJ2a">Spotify</a>, <a href="https://podcastd45a61.podigee.io/feed/mp3">im Feed</a> und natürlich auf <a href="https://itunes.apple.com/de/podcast/stay-forever/id461077931?mt=2">iTunes</a>. Wir freuen uns über Reaktionen und Empfehlungen auf <a href="https://twitter.com/stayforeverDE">Twitter</a> oder <a href="https://www.facebook.com/StayForeverPodcast/">Facebook</a>. Wer uns unterstützen möchte, kann das auf <a href="https://steadyhq.com/de/stayforever">Steady</a>, <a href="https://www.patreon.com/stayforever">Patreon</a> oder per Kauf von <a href="https://amzn.to/2Q6D9Fb">irgendwas auf Amazon</a> tun (Affiliate-Link).</p>]]>
            </content:encoded>
            <image>
                <url>
                    https://images.podigee.com/0x,sofpG_4qf7QOLKyCkgmqPbRxaPo3wd_nvfr8G-cuyiXo=/https://cdn.podigee.com/uploads/u4307/fc4d81c4-b8bf-4a9c-956d-5d1a71e63506.jpg
                </url>
                <title>Baphomets Fluch (Stay Forever, Folge 1 REMAKE)</title>
                <link>https://www.stayforever.de/2019/02/baphomets-fluch-folge-1-remake/</link>
            </image>
            <itunes:image
                    href="https://images.podigee.com/0x,sofpG_4qf7QOLKyCkgmqPbRxaPo3wd_nvfr8G-cuyiXo=/https://cdn.podigee.com/uploads/u4307/fc4d81c4-b8bf-4a9c-956d-5d1a71e63506.jpg"/>
            <itunes:episode>150</itunes:episode>
            <itunes:episodeType>full</itunes:episodeType>
            <itunes:subtitle>Neue Aufnahme der (legendären?) Folge 1</itunes:subtitle>
            <itunes:summary>1996 kam, schon relativ spät in der Hochzeit der Adventures, das aufwändig gestaltete Spiel
                Baphomets Fluch (Broken Sword) auf den Markt. Entwickler war das kleine englische Studio Revolution
                Software unter Charles Cecil. Das Spiel fand eine enorm große Fangemeinde und legte den Grundstein für
                eine der langlebigsten Adventure-Serien. Christian und Gunnar haben die allererste Folge von Stay
                Forever nochmal ganz neu aufgenommen.

                Infos zum Spiel:

                Thema: Baphomets Fluch
                Erscheinungsjahr: 1996
                Genre: Adventure
                Plattform: MS-DOS, Windows (später jede Plattform von GBA bis iOS)
                Entwickler: Revolution Software
                Publisher: Virgin Interactive Ltd.
                Designer: Charles Cecil, Steve Ince

                Podcast-Credits:

                Sprecher: Christian Schmidt, Gunnar Lott
                Audioproduktion: Fabian Langer, Christian Schmidt
                Titelgrafik: Paul Schmidt
                Intro, Outro: Nino Kerl (Ansage); Chris Hülsbeck (Musik)
                Chronist: Herr Anym
                Community Management: Christian Beuster
            </itunes:summary>
            <itunes:explicit>no</itunes:explicit>
            <itunes:keywords>Charles Cecil,Steve Ince,Broken Sword,Revolution Software,1996,Dosgames,Stay
                Forever,retrogames,podcast,Baphomets Fluch
            </itunes:keywords>
            <itunes:author>Gunnar Lott, Christian Schmidt, Fabian Käufer</itunes:author>
            <enclosure
                    url="https://cdn.podigee.com/media/podcast_6853_stay_forever_episode_150_baphomets_fluch_stay_forever_folge_1_remake.m4a?v=1549019798&amp;source=feed"
                    type="audio/aac" length="73110523"/>
            <itunes:duration>7110</itunes:duration>
        </item>
</root>
`
