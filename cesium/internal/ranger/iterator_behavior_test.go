package ranger_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/synnaxlabs/cesium/internal/ranger"
	"github.com/synnaxlabs/x/io/fs"
	"github.com/synnaxlabs/x/telem"
	. "github.com/synnaxlabs/x/testutil"
)

func write(db *ranger.DB, rng telem.TimeRange, data []byte) {
	w := MustSucceed(db.NewWriter(ranger.WriterConfig{
		Start: rng.Start,
		End:   rng.End,
	}))
	n := MustSucceed(w.Write(data))
	Expect(n).To(Equal(len(data)))
	Expect(w.Commit(rng.End)).To(Succeed())
	Expect(w.Close()).To(Succeed())
}

var _ = Describe("Iterator Behavior", func() {
	var db *ranger.DB
	BeforeEach(func() {
		db = MustSucceed(ranger.Open(ranger.Config{FS: fs.NewMem()}))
	})
	AfterEach(func() { Expect(db.Close()).To(Succeed()) })
	Describe("Valid", func() {
		It("Should return false on an iterator with zero span bounds", func() {
			r := db.NewIterator(ranger.IteratorConfig{
				Bounds: (10 * telem.SecondTS).SpanRange(0),
			})
			Expect(r.Valid()).To(BeFalse())
		})
	})
	Describe("SeekFirst + SeekLast", func() {
		BeforeEach(func() {
			write(db, (10 * telem.SecondTS).SpanRange(10*telem.Second), []byte{1, 2, 3, 4, 5, 6})
			write(db, (30 * telem.SecondTS).SpanRange(10*telem.Second), []byte{1, 2, 3, 4, 5, 6})
		})
		DescribeTable("SeekFirst",
			func(
				ts telem.TimeStamp,
				expectedResult bool,
				expectedFirst telem.TimeRange,
			) {
				r := db.NewIterator(ranger.IterRange(ts.SpanRange(telem.TimeSpanMax)))
				Expect(r.SeekFirst()).To(Equal(expectedResult))
				if expectedResult {
					Expect(r.Range()).To(Equal(expectedFirst))
				}
			},
			Entry("Bound start equal to range start",
				10*telem.SecondTS,
				true,
				(10*telem.SecondTS).SpanRange(10*telem.Second),
			),
			Entry("Bound end equal to range start",
				20*telem.SecondTS,
				true,
				(30*telem.SecondTS).SpanRange(10*telem.Second),
			),
			Entry(`
				Bound start strictly greater than range start and strictly less than
				range end
			`,
				15*telem.SecondTS,
				true,
				(10*telem.SecondTS).SpanRange(10*telem.Second),
			),
			Entry("Bound start strictly less than start of first defined range",
				5*telem.SecondTS,
				true,
				(10*telem.SecondTS).SpanRange(10*telem.Second),
			),
			Entry("Bound start strictly greater than end of last defined range",
				40*telem.SecondTS,
				false,
				telem.TimeRangeZero,
			),
		)
		DescribeTable("SeekLast",
			func(
				ts telem.TimeStamp,
				expectedResult bool,
				expectedLast telem.TimeRange,
			) {
				tr := telem.TimeRange{Start: 0, End: ts}
				r := db.NewIterator(ranger.IterRange(tr))
				Expect(r.SeekLast()).To(Equal(expectedResult))
				Expect(r.Range()).To(Equal(expectedLast))
			},
			Entry("Bound end equal to range end",
				40*telem.SecondTS,
				true,
				(30*telem.SecondTS).SpanRange(10*telem.Second),
			),
			Entry("Bound end equal to range start",
				30*telem.SecondTS,
				true,
				(10*telem.SecondTS).SpanRange(10*telem.Second),
			),
			Entry(`
					Bound end strictly greater than range start and strictly less than
					range end
			`,
				35*telem.SecondTS,
				true,
				(30*telem.SecondTS).SpanRange(10*telem.Second),
			),
		)
	})

	Context("Exhaustion", func() {
		BeforeEach(func() {
			write(db, (50 * telem.SecondTS).SpanRange(10*telem.Second), []byte{1, 2, 3, 4, 5, 6})
			write(db, (60 * telem.SecondTS).SpanRange(10*telem.Second), []byte{1, 2, 3, 4, 5, 6})
			write(db, (10 * telem.SecondTS).SpanRange(10*telem.Second), []byte{1, 2, 3, 4, 5, 6})
			write(db, (30 * telem.SecondTS).SpanRange(10*telem.Second), []byte{1, 2, 3, 4, 5, 6})
		})
		Context("Next", func() {
			It("Should return false when the iterator is exhausted", func() {
				iter := db.NewIterator(ranger.IteratorConfig{
					Bounds: (15 * telem.SecondTS).SpanRange(45 * telem.Second),
				})
				Expect(iter.SeekFirst()).To(BeTrue())
				Expect(iter.Range()).To(Equal((10 * telem.SecondTS).SpanRange(10 * telem.Second)))
				Expect(iter.Next()).To(BeTrue())
				Expect(iter.Range()).To(Equal((30 * telem.SecondTS).SpanRange(10 * telem.Second)))
				Expect(iter.Next()).To(BeTrue())
				Expect(iter.Range()).To(Equal((50 * telem.SecondTS).SpanRange(10 * telem.Second)))
				Expect(iter.Next()).To(BeFalse())
			})
		})
		Context("Prev", func() {
			It("Should return false when the iterator is exhausted", func() {
				iter := db.NewIterator(ranger.IteratorConfig{
					Bounds: (15 * telem.SecondTS).SpanRange(45 * telem.Second),
				})
				Expect(iter.SeekLast()).To(BeTrue())
				Expect(iter.Range()).To(Equal((50 * telem.SecondTS).SpanRange(10 * telem.Second)))
				Expect(iter.Prev()).To(BeTrue())
				Expect(iter.Range()).To(Equal((30 * telem.SecondTS).SpanRange(10 * telem.Second)))
				Expect(iter.Prev()).To(BeTrue())
				Expect(iter.Range()).To(Equal((10 * telem.SecondTS).SpanRange(10 * telem.Second)))
				Expect(iter.Prev()).To(BeFalse())
			})
		})
	})
})
