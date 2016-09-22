package miccheck

import (
	"strconv"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestCombinationsAndOrderingOfExpectorsAndExpectations(t *testing.T) {
	Convey("Given single expector with single expectation", t, func() {
		Convey("Given there are no matches", func() {
			data1 := map[string]interface{}{"key1": "value1"}
			secondData1 := map[string]interface{}{"key2": "value2"}
			Convey("When it validates", func() {
				exp1 := &expectation{data: data1}
				secondExp1 := &expectation{data: secondData1}
				Convey("It will result in empty matches", func() {
					returnedData := exp1.validate(secondExp1)
					So(returnedData, ShouldResemble, []map[string]interface{}{})
				})
			})
		})

		Convey("Given there is one match", func() {
			data1 := map[string]interface{}{"key1": "value1"}
			secondData1 := map[string]interface{}{"key1": "value1"}
			Convey("When it validates", func() {
				exp1 := &expectation{data: data1}
				secondExp1 := &expectation{data: secondData1}
				Convey("It will result in that one match", func() {
					returnedData := exp1.validate(secondExp1)
					So(returnedData, ShouldResemble, []map[string]interface{}{data1})
				})
			})
		})
	})

	Convey("Given single expector with multiple expectations", t, func() {
		Convey("Given there are no matches", func() {
			data1 := map[string]interface{}{"key9": "value9"}
			secondData1 := map[string]interface{}{"key1": "value1"}
			secondData2 := map[string]interface{}{"key2": "value2"}
			secondData3 := map[string]interface{}{"key3": "value3"}
			secondData4 := map[string]interface{}{"key4": "value4"}
			secondData5 := map[string]interface{}{"key5": "value5"}
			Convey("When it validates", func() {
				exp1 := &expectation{data: data1}
				secondExp1 := &expectation{data: secondData1}
				secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
				secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
				secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
				secondExp5 := &expectation{data: secondData5, nextExpector: secondExp4}
				Convey("It will result in empty matches", func() {
					returnedData := exp1.validate(secondExp5)
					So(returnedData, ShouldResemble, []map[string]interface{}{})
				})
			})
		})

		Convey("Given there is one match", func() {
			data1 := map[string]interface{}{"key1": "value1"}
			Convey("And that match is at first expectation", func() {
				secondData1 := map[string]interface{}{"key1": "value1"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key3": "value3"}
				secondData4 := map[string]interface{}{"key4": "value4"}
				secondData5 := map[string]interface{}{"key5": "value5"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					secondExp5 := &expectation{data: secondData5, nextExpector: secondExp4}
					Convey("It will result in that one match", func() {
						returnedData := exp1.validate(secondExp5)
						So(returnedData, ShouldResemble, []map[string]interface{}{data1})
					})
				})
			})

			Convey("And that match is at last expectation", func() {
				secondData1 := map[string]interface{}{"key0": "value0"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key3": "value3"}
				secondData4 := map[string]interface{}{"key4": "value4"}
				secondData5 := map[string]interface{}{"key1": "value1"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					secondExp5 := &expectation{data: secondData5, nextExpector: secondExp4}
					Convey("It will result in that one match", func() {
						returnedData := exp1.validate(secondExp5)
						So(returnedData, ShouldResemble, []map[string]interface{}{data1})
					})
				})
			})

			Convey("And that match is at middle expectation", func() {
				secondData1 := map[string]interface{}{"key0": "value0"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key1": "value1"}
				secondData4 := map[string]interface{}{"key4": "value4"}
				secondData5 := map[string]interface{}{"key5": "value5"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					secondExp5 := &expectation{data: secondData5, nextExpector: secondExp4}
					Convey("It will result in that one match", func() {
						returnedData := exp1.validate(secondExp5)
						So(returnedData, ShouldResemble, []map[string]interface{}{data1})
					})
				})
			})
		})

		Convey("Given there are multiple matches", func() {
			data1 := map[string]interface{}{"key1": "value1"}
			Convey("And those matches start at first expectation", func() {
				secondData1 := map[string]interface{}{"key1": "value1"}
				secondData2 := map[string]interface{}{"key1": "value1"}
				secondData3 := map[string]interface{}{"key1": "value1"}
				secondData4 := map[string]interface{}{"key4": "value4"}
				secondData5 := map[string]interface{}{"key5": "value5"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					secondExp5 := &expectation{data: secondData5, nextExpector: secondExp4}
					Convey("It will result in that one match", func() {
						returnedData := exp1.validate(secondExp5)
						So(returnedData, ShouldResemble, []map[string]interface{}{data1})
					})
				})
			})

			Convey("And those matches start at middle and end at last expectation", func() {
				secondData1 := map[string]interface{}{"key0": "value0"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key1": "value1"}
				secondData4 := map[string]interface{}{"key1": "value1"}
				secondData5 := map[string]interface{}{"key1": "value1"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					secondExp5 := &expectation{data: secondData5, nextExpector: secondExp4}
					Convey("It will result in that one match", func() {
						returnedData := exp1.validate(secondExp5)
						So(returnedData, ShouldResemble, []map[string]interface{}{data1})
					})
				})
			})

			Convey("And those matches are in the middle of expectation", func() {
				secondData1 := map[string]interface{}{"key0": "value0"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key1": "value1"}
				secondData4 := map[string]interface{}{"key1": "value1"}
				secondData5 := map[string]interface{}{"key1": "value1"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					secondExp5 := &expectation{data: secondData5, nextExpector: secondExp4}
					Convey("It will result in that one match", func() {
						returnedData := exp1.validate(secondExp5)
						So(returnedData, ShouldResemble, []map[string]interface{}{data1})
					})
				})
			})
		})

		Convey("Given there are all matches", func() {
			data1 := map[string]interface{}{"key1": "value1"}
			secondData1 := map[string]interface{}{"key1": "value1"}
			secondData2 := map[string]interface{}{"key1": "value1"}
			secondData3 := map[string]interface{}{"key1": "value1"}
			secondData4 := map[string]interface{}{"key1": "value1"}
			secondData5 := map[string]interface{}{"key1": "value1"}
			Convey("When it validates", func() {
				exp1 := &expectation{data: data1}
				secondExp1 := &expectation{data: secondData1}
				secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
				secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
				secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
				secondExp5 := &expectation{data: secondData5, nextExpector: secondExp4}
				Convey("It will result in that one match", func() {
					returnedData := exp1.validate(secondExp5)
					So(returnedData, ShouldResemble, []map[string]interface{}{data1})
				})
			})
		})
	})

	Convey("Given multiple expectors with single expectation", t, func() {
		Convey("Given a single match", func() {
			data1 := map[string]interface{}{"key1": "value1"}
			data2 := map[string]interface{}{"key2": "value2"}
			data3 := map[string]interface{}{"key3": "value3"}
			data4 := map[string]interface{}{"key4": "value4"}
			secondData1 := map[string]interface{}{"key1": "value1"}
			Convey("When it validates", func() {
				exp1 := &expectation{data: data1}
				exp2 := &expectation{data: data2, nextExpector: exp1}
				exp3 := &expectation{data: data3, nextExpector: exp2}
				exp4 := &expectation{data: data4, nextExpector: exp3}
				secondExp1 := &expectation{data: secondData1}
				Convey("It will result in that one match", func() {
					returnedData := exp4.validate(secondExp1)
					So(returnedData, ShouldResemble, []map[string]interface{}{secondData1})
				})
			})
		})

		Convey("Given multiple matches", func() {
			Convey("And matches start at first expector", func() {
				data1 := map[string]interface{}{"key1": "value1"}
				data2 := map[string]interface{}{"key1": "value1"}
				data3 := map[string]interface{}{"key3": "value3"}
				data4 := map[string]interface{}{"key4": "value4"}
				secondData1 := map[string]interface{}{"key1": "value1"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp1)
						So(returnedData, ShouldResemble, []map[string]interface{}{secondData1})
					})
				})
			})

			Convey("And matches are in middle of expectors", func() {
				data1 := map[string]interface{}{"key0": "value0"}
				data2 := map[string]interface{}{"key1": "value1"}
				data3 := map[string]interface{}{"key1": "value1"}
				data4 := map[string]interface{}{"key4": "value4"}
				secondData1 := map[string]interface{}{"key1": "value1"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp1)
						So(returnedData, ShouldResemble, []map[string]interface{}{secondData1})
					})
				})
			})

			Convey("And matches start at middle and end at last expector", func() {
				data1 := map[string]interface{}{"key0": "value0"}
				data2 := map[string]interface{}{"key2": "value2"}
				data3 := map[string]interface{}{"key1": "value1"}
				data4 := map[string]interface{}{"key1": "value1"}
				secondData1 := map[string]interface{}{"key1": "value1"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp1)
						So(returnedData, ShouldResemble, []map[string]interface{}{secondData1})
					})
				})
			})
		})
	})

	Convey("Given multiple expectors with multiple expectation", t, func() {
		Convey("Given a single match", func() {
			Convey("And match is at first expector and expectation", func() {
				data1 := map[string]interface{}{"key1": "value1"}
				data2 := map[string]interface{}{"key22": "value22"}
				data3 := map[string]interface{}{"key33": "value33"}
				data4 := map[string]interface{}{"key44": "value44"}
				secondData1 := map[string]interface{}{"key1": "value1"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key3": "value3"}
				secondData4 := map[string]interface{}{"key4": "value4"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{data1})
					})
				})
			})

			Convey("And match is at first expector and middle expectation", func() {
				data1 := map[string]interface{}{"key1": "value1"}
				data2 := map[string]interface{}{"key22": "value22"}
				data3 := map[string]interface{}{"key33": "value33"}
				data4 := map[string]interface{}{"key44": "value44"}
				secondData1 := map[string]interface{}{"key0": "value0"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key1": "value1"}
				secondData4 := map[string]interface{}{"key4": "value4"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{data1})
					})
				})
			})

			Convey("And match is at first expector and last expectation", func() {
				data1 := map[string]interface{}{"key1": "value1"}
				data2 := map[string]interface{}{"key22": "value22"}
				data3 := map[string]interface{}{"key33": "value33"}
				data4 := map[string]interface{}{"key44": "value44"}
				secondData1 := map[string]interface{}{"key0": "value0"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key3": "value3"}
				secondData4 := map[string]interface{}{"key1": "value1"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{data1})
					})
				})
			})

			Convey("And match is at middle expector and first expectation", func() {
				data1 := map[string]interface{}{"key0": "value0"}
				data2 := map[string]interface{}{"key22": "value22"}
				data3 := map[string]interface{}{"key1": "value1"}
				data4 := map[string]interface{}{"key44": "value44"}
				secondData1 := map[string]interface{}{"key1": "value1"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key3": "value3"}
				secondData4 := map[string]interface{}{"key4": "value4"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{secondData1})
					})
				})
			})

			Convey("And match is at middle expector and middle expectation", func() {
				data1 := map[string]interface{}{"key0": "value0"}
				data2 := map[string]interface{}{"key22": "value22"}
				data3 := map[string]interface{}{"key1": "value1"}
				data4 := map[string]interface{}{"key44": "value44"}
				secondData1 := map[string]interface{}{"key00": "value00"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key1": "value1"}
				secondData4 := map[string]interface{}{"key4": "value4"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{data3})
					})
				})
			})

			Convey("And match is at middle expector and last expectation", func() {
				data1 := map[string]interface{}{"key0": "value0"}
				data2 := map[string]interface{}{"key22": "value22"}
				data3 := map[string]interface{}{"key1": "value1"}
				data4 := map[string]interface{}{"key44": "value44"}
				secondData1 := map[string]interface{}{"key00": "value00"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key3": "value3"}
				secondData4 := map[string]interface{}{"key1": "value1"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{data3})
					})
				})
			})

			Convey("And match is at last expector and first expectation", func() {
				data1 := map[string]interface{}{"key11": "value11"}
				data2 := map[string]interface{}{"key22": "value22"}
				data3 := map[string]interface{}{"key33": "value33"}
				data4 := map[string]interface{}{"key1": "value1"}
				secondData1 := map[string]interface{}{"key1": "value1"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key3": "value3"}
				secondData4 := map[string]interface{}{"key4": "value4"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{secondData1})
					})
				})
			})

			Convey("And match is at last expector and middle expectation", func() {
				data1 := map[string]interface{}{"key00": "value00"}
				data2 := map[string]interface{}{"key22": "value22"}
				data3 := map[string]interface{}{"key33": "value33"}
				data4 := map[string]interface{}{"key1": "value1"}
				secondData1 := map[string]interface{}{"key11": "value11"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key1": "value1"}
				secondData4 := map[string]interface{}{"key4": "value4"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{secondData3})
					})
				})
			})

			Convey("And match is at last expector and last expectation", func() {
				data1 := map[string]interface{}{"key11": "value11"}
				data2 := map[string]interface{}{"key22": "value22"}
				data3 := map[string]interface{}{"key33": "value33"}
				data4 := map[string]interface{}{"key1": "value1"}
				secondData1 := map[string]interface{}{"key01": "value01"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key3": "value3"}
				secondData4 := map[string]interface{}{"key1": "value1"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{data4})
					})
				})
			})
		})

		Convey("Given some matches", func() {
			Convey("And matches start at first expector and expectation", func() {
				data1 := map[string]interface{}{"key1": "value1"}
				data2 := map[string]interface{}{"key2": "value2"}
				data3 := map[string]interface{}{"key33": "value33"}
				data4 := map[string]interface{}{"key44": "value44"}
				secondData1 := map[string]interface{}{"key1": "value1"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key3": "value3"}
				secondData4 := map[string]interface{}{"key4": "value4"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{data1, data2})
						So(returnedData, ShouldHaveLength, 2)
					})
				})
			})

			Convey("And matches start at first expector and middle of expectation", func() {
				data1 := map[string]interface{}{"key1": "value1"}
				data2 := map[string]interface{}{"key2": "value2"}
				data3 := map[string]interface{}{"key33": "value33"}
				data4 := map[string]interface{}{"key44": "value44"}
				secondData1 := map[string]interface{}{"key11": "value11"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key1": "value1"}
				secondData4 := map[string]interface{}{"key4": "value4"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{data1, data2})
						So(returnedData, ShouldHaveLength, 2)
					})
				})
			})

			Convey("And matches start at first expector and end at last expectation", func() {
				data1 := map[string]interface{}{"key1": "value1"}
				data2 := map[string]interface{}{"key2": "value2"}
				data3 := map[string]interface{}{"key33": "value33"}
				data4 := map[string]interface{}{"key44": "value44"}
				secondData1 := map[string]interface{}{"key11": "value11"}
				secondData2 := map[string]interface{}{"key22": "value22"}
				secondData3 := map[string]interface{}{"key2": "value2"}
				secondData4 := map[string]interface{}{"key1": "value1"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{data1, data2})
						So(returnedData, ShouldHaveLength, 2)
					})
				})
			})

			Convey("And matches start in middle of expectors and start of expectation", func() {
				data1 := map[string]interface{}{"key11": "value11"}
				data2 := map[string]interface{}{"key2": "value2"}
				data3 := map[string]interface{}{"key1": "value1"}
				data4 := map[string]interface{}{"key44": "value44"}
				secondData1 := map[string]interface{}{"key1": "value1"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key3": "value3"}
				secondData4 := map[string]interface{}{"key4": "value4"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{secondData2, secondData1})
						So(returnedData, ShouldHaveLength, 2)
					})
				})
			})

			Convey("And matches start in middle of expectors and middle of expectations", func() {
				data1 := map[string]interface{}{"key11": "value11"}
				data2 := map[string]interface{}{"key2": "value2"}
				data3 := map[string]interface{}{"key1": "value1"}
				data4 := map[string]interface{}{"key44": "value44"}
				secondData1 := map[string]interface{}{"key10": "value10"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key1": "value1"}
				secondData4 := map[string]interface{}{"key4": "value4"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{secondData2, secondData3})
						So(returnedData, ShouldHaveLength, 2)
					})
				})
			})

			Convey("And matches start in middle of expectors and end at last expectations", func() {
				data1 := map[string]interface{}{"key11": "value11"}
				data2 := map[string]interface{}{"key2": "value2"}
				data3 := map[string]interface{}{"key1": "value1"}
				data4 := map[string]interface{}{"key44": "value44"}
				secondData1 := map[string]interface{}{"key10": "value10"}
				secondData2 := map[string]interface{}{"key22": "value22"}
				secondData3 := map[string]interface{}{"key1": "value1"}
				secondData4 := map[string]interface{}{"key2": "value2"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{data2, data3})
						So(returnedData, ShouldHaveLength, 2)
					})
				})
			})

			Convey("And matches end at last expector and start at first expectation", func() {
				data1 := map[string]interface{}{"key11": "value11"}
				data2 := map[string]interface{}{"key22": "value22"}
				data3 := map[string]interface{}{"key1": "value1"}
				data4 := map[string]interface{}{"key2": "value2"}
				secondData1 := map[string]interface{}{"key1": "value1"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key12": "value12"}
				secondData4 := map[string]interface{}{"key42": "value42"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{data3, data4})
						So(returnedData, ShouldHaveLength, 2)
					})
				})
			})

			Convey("And matches end at last expector and start in middle of expectations", func() {
				data1 := map[string]interface{}{"key11": "value11"}
				data2 := map[string]interface{}{"key22": "value22"}
				data3 := map[string]interface{}{"key1": "value1"}
				data4 := map[string]interface{}{"key2": "value2"}
				secondData1 := map[string]interface{}{"key13": "value13"}
				secondData2 := map[string]interface{}{"key2": "value2"}
				secondData3 := map[string]interface{}{"key1": "value1"}
				secondData4 := map[string]interface{}{"key42": "value42"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{data3, data4})
						So(returnedData, ShouldHaveLength, 2)
					})
				})
			})

			Convey("And matches end at last expector and end at last expectation", func() {
				data1 := map[string]interface{}{"key11": "value11"}
				data2 := map[string]interface{}{"key22": "value22"}
				data3 := map[string]interface{}{"key1": "value1"}
				data4 := map[string]interface{}{"key2": "value2"}
				secondData1 := map[string]interface{}{"key13": "value13"}
				secondData2 := map[string]interface{}{"key23": "value23"}
				secondData3 := map[string]interface{}{"key1": "value1"}
				secondData4 := map[string]interface{}{"key2": "value2"}
				Convey("When it validates", func() {
					exp1 := &expectation{data: data1}
					exp2 := &expectation{data: data2, nextExpector: exp1}
					exp3 := &expectation{data: data3, nextExpector: exp2}
					exp4 := &expectation{data: data4, nextExpector: exp3}
					secondExp1 := &expectation{data: secondData1}
					secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
					secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
					secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
					Convey("It will result in that one match", func() {
						returnedData := exp4.validate(secondExp4)
						So(returnedData, ShouldResemble, []map[string]interface{}{data3, data4})
						So(returnedData, ShouldHaveLength, 2)
					})
				})
			})
		})
	})
}

func TestOutliers(t *testing.T) {
	Convey("Given duplicate matches in expectations", t, func() {
		data1 := map[string]interface{}{"key1": "value1"}
		secondData1 := map[string]interface{}{"key1": "value1"}
		secondData2 := map[string]interface{}{"key2": "value2"}
		secondData3 := map[string]interface{}{"key1": "value1"}
		secondData4 := map[string]interface{}{"key2": "value2"}
		Convey("When it validates", func() {
			exp1 := &expectation{data: data1}
			secondExp1 := &expectation{data: secondData1}
			secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
			secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
			secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
			Convey("It will result in only one of the matches", func() {
				returnedData := exp1.validate(secondExp4)
				So(returnedData, ShouldResemble, []map[string]interface{}{data1})
			})
		})
	})

	Convey("Given duplicate matches in expector and expectations", t, func() {
		data1 := map[string]interface{}{"key1": "value1"}
		data2 := map[string]interface{}{"key1": "value1"}
		secondData1 := map[string]interface{}{"key1": "value1"}
		secondData2 := map[string]interface{}{"key2": "value2"}
		secondData3 := map[string]interface{}{"key1": "value1"}
		secondData4 := map[string]interface{}{"key2": "value2"}
		Convey("When it validates", func() {
			exp1 := &expectation{data: data1}
			exp2 := &expectation{data: data2, nextExpector: exp1}
			secondExp1 := &expectation{data: secondData1}
			secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
			secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
			secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
			Convey("It will result in only one of the matches", func() {
				returnedData := exp2.validate(secondExp4)
				So(returnedData, ShouldResemble, []map[string]interface{}{data1, data1})
			})
		})
	})

	Convey("Given deeply nested matches of hash maps", t, func() {
		nest1 := map[string]interface{}{"key1": "value1"}
		nest2 := map[string]interface{}{"nest": nest1}
		nest3 := map[string]interface{}{"nest1": nest1, "nest2": nest2}
		nest4 := map[string]interface{}{"nest": nest3}
		data1 := map[string]interface{}{"nest": nest4}
		data2 := map[string]interface{}{"key1": "value1"}
		secondData1 := map[string]interface{}{"key1": "value1"}
		secondData2 := map[string]interface{}{"key2": "value2"}
		secondData3 := map[string]interface{}{"nest": nest4}
		secondData4 := map[string]interface{}{"key2": "value2"}
		Convey("When it validates", func() {
			exp1 := &expectation{data: data1}
			exp2 := &expectation{data: data2, nextExpector: exp1}
			secondExp1 := &expectation{data: secondData1}
			secondExp2 := &expectation{data: secondData2, nextExpector: secondExp1}
			secondExp3 := &expectation{data: secondData3, nextExpector: secondExp2}
			secondExp4 := &expectation{data: secondData4, nextExpector: secondExp3}
			Convey("It will result in nested matches", func() {
				returnedData := exp2.validate(secondExp4)
				So(returnedData, ShouldResemble, []map[string]interface{}{data1, data2})
			})
		})
	})
}

func BenchmarkValidation(t *testing.B) {
	terminal1 := &expectation{data: map[string]interface{}{"key": "value"}}
	exp1 := makeExpectations(terminal1, 400, 0, 2)
	terminal2 := &expectation{data: map[string]interface{}{"key": "value"}}
	exp2 := makeExpectations(terminal2, 100, 0, 2)

	now := time.Now()
	exp1.validate(exp2)

	assert.True(t, time.Since(now).Seconds() < 10)
}

func makeExpectations(exp *expectation, times, count, seed int) *expectation {
	if times == count {
		return exp
	}

	count++

	nextExpectation := makeExpectations(exp, times, count, seed)
	num := times % seed
	return &expectation{nextExpector: nextExpectation, data: map[string]interface{}{"key" + strconv.Itoa(num): "value" + strconv.Itoa(num)}}
}
