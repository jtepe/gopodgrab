# gopodgrab
A command-line tool to download and manage your favourite podcasts from an XML feed.

This is developed as a small side-project for personal use. Features will be added as seen fit, needed, appropriate,
nice to have, etc.

## Usage
### Add a new podcast
`$ gopodgrab add --feed-url https://path/to/podcast/feed --storage /path/to/store/episodes --name foocast`

Adds a new podcast "foocast" to be managed by `gopodgrab` specifying where to store the episodes and the location of the
cast's feed file.

### Update podcast
`$ gopodgrab update foocast`

Update foocast, downloading all new epsiodes since the last update to the local storage directory. Here "foocast" is
the name of a managed podcast. The special name "all" updates all managed podcasts.

`$ gopodgrab update all`

### Show all currently managed podcasts
`$ gopodgrab list`

Lists all currently managed podcasts with some additional information for each.

`$ gopodgrab show foocast`

Show a more detailed summary for podcast "foocast".
