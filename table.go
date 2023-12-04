package dtable

import "github.com/wenooij/dtable/dtwire"

type Table interface {
	Scan(dtwire.Scanner) error
}
