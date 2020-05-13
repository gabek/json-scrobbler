# json-scrobbler
Monitor a JSON file for track information and when it changes scrobble the track to Last.FM.

# Purpose

This was built originally as a companion for [Supbox](https://github.com/gabek/supbox), but it will work with any JSON file as long as it has an `artist` and `track` key.

For example [Supbox](https://github.com/gabek/supbox) renders the following `nowplaying.json` file:

```
{"ID":"6d655f64-f2e5-4ec9-bd51-3a2c86984c8e","artist":"VNV Nation","track":"Retaliate","imagePath":"/Users/gabek/Library/Pioneer/rekordbox/share/PIONEER/Artwork/bd3/82718-334f-482d-ad0a-82a1f8ba2507/artwork.jpg"}
```

So point `json-scrobbler` at that file in the `config/config.yaml` and run.

## Authentication

### Keys

You'll have to create your own set of Last.FM authentication keys, but that's easy enough.  Just go to https://www.last.fm/api/account/create, fill out the form (Put some dummy value as `Callback URL`) and then put the resulting `key` and `secret` into the `config.yaml` file.

### User

Put your Last.FM username and password in the `config.yaml` file to enable authenticating yourself for scrobbling.