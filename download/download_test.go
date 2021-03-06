package download

import (
	"log"
	"os"
	"testing"
)

func BenchmarkDownload(b *testing.B) {
	b.Run("thinkbroadband 1GB", func(b *testing.B) {
		url := "http://ipv4.download.thinkbroadband.com/1GB.zip"

		b.Run("1 chunk", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				downloadChunked(b, url, 1)
			}
		})

		b.Run("10 chunks", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				downloadChunked(b, url, 10)
			}
		})

		b.Run("20 chunks", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				downloadChunked(b, url, 20)
			}
		})

		b.Run("40 chunks", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				downloadChunked(b, url, 40)
			}
		})

		b.Run("60 chunks", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				downloadChunked(b, url, 60)
			}
		})
	})

	b.Run("hetzner 1GB", func(b *testing.B) {
		url := "https://speed.hetzner.de/1GB.bin"

		b.Run("1 chunk", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				downloadChunked(b, url, 1)
			}
		})

		b.Run("10 chunks", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				downloadChunked(b, url, 10)
			}
		})

		b.Run("20 chunks", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				downloadChunked(b, url, 20)
			}
		})

		b.Run("40 chunks", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				downloadChunked(b, url, 40)
			}
		})

		b.Run("80 chunks", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				downloadChunked(b, url, 80)
			}
		})
	})
}

func downloadChunked(b *testing.B, url string, numChunks int) {
	err := New(url, numChunks, false).Run()
	if err != nil {
		b.Fatalf("error downloading file: %s", err)
	}

	b.StopTimer()

	err = os.Remove(filename(url))
	if err != nil {
		log.Printf("error removing file: %s\n", err)
	}

	b.StartTimer()
}
