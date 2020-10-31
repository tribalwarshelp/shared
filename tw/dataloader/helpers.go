package dataloader

import (
	"compress/gzip"
	"encoding/csv"
	"encoding/xml"
	"io"
	"io/ioutil"
)

func newCsvReader(r io.Reader) *csv.Reader {
	csvReader := csv.NewReader(r)
	csvReader.LazyQuotes = true
	csvReader.FieldsPerRecord = -1
	return csvReader
}

func uncompressAndReadCsvLines(r io.Reader) ([][]string, error) {
	uncompressedStream, err := gzip.NewReader(r)
	if err != nil {
		return [][]string{}, err
	}
	defer uncompressedStream.Close()
	return newCsvReader(uncompressedStream).ReadAll()
}

func (d *dataLoader) getCSVData(url string, compressed bool) ([][]string, error) {
	resp, err := d.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if !compressed {
		return newCsvReader(resp.Body).ReadAll()
	}
	return uncompressAndReadCsvLines(resp.Body)
}

func (d *dataLoader) getXML(url string, decode interface{}) error {
	resp, err := d.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return xml.Unmarshal(bytes, decode)
}
