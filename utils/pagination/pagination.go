package pagination

import (
	"encoding/base64"
	"encoding/gob"
	"strings"
)

type PageToken struct {
	Offset int64
}

func ParsePageToken(pageSize int32, pageToken string) (PageToken, error) {
	if pageToken == "" {
		return PageToken{
			Offset: 0,
		}, nil
	}
	var result PageToken
	if err := decodePageToken(pageToken, &result); err != nil {
		return PageToken{}, err
	}
	
	return result, nil
}

func (p PageToken) String() string {
	return encodePageRank(p)
}

func (p PageToken) Next(pageSize int32) PageToken {
	p.Offset += int64(pageSize)
	return p
}

func encodePageRank(pageToken PageToken) string {
	var b strings.Builder
	base64Encoder := base64.NewEncoder(base64.RawStdEncoding, &b)
	gobEncoder := gob.NewEncoder(base64Encoder)
	_ = gobEncoder.Encode(pageToken)
	_ = base64Encoder.Close()
	return b.String()
}

func decodePageToken(pageToken string, result *PageToken) error {
	dec := gob.NewDecoder(base64.NewDecoder(base64.RawStdEncoding, strings.NewReader(pageToken)))
	if err := dec.Decode(result); err != nil {
		return err
	}
	return nil
}
